package notifier

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/MultiBanker/broker/src/config"
	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/MultiBanker/broker/src/director"
)

const defaultTicker = 500 * time.Millisecond

type notifier struct {
	userRepo     repository.UsersRepository
	verifyRepo   repository.VerificationRepository
	recoveryRepo repository.RecoveryRepository
	httpCli      *http.Client
	notifyConfig config.NotifyConfig
}

func NewNotifier(userRepo repository.UsersRepository,
	verifyRepo repository.VerificationRepository,
	recoveryRepo repository.RecoveryRepository,
	notifyConfig config.NotifyConfig) *notifier {
	transport := &http.Transport{
		DialContext:         (&net.Dialer{Timeout: 40 * time.Second}).DialContext,
		TLSHandshakeTimeout: 15 * time.Second,
		MaxIdleConns:        100,
		MaxConnsPerHost:     100,
		MaxIdleConnsPerHost: 5,
	}
	cli := &http.Client{
		Timeout:   time.Second * 60,
		Transport: transport,
	}

	return &notifier{
		userRepo:     userRepo,
		verifyRepo:   verifyRepo,
		recoveryRepo: recoveryRepo,
		httpCli:      cli,
		notifyConfig: notifyConfig,
	}
}

func (n notifier) Name() string {
	return "notifier-worker"
}

func (n notifier) Start(ctx context.Context, cancelFunc context.CancelFunc) error {
	defer cancelFunc()

	notify := director.NewWorker(n.Name(), defaultTicker)

	go notify.Run(ctx, n.Verifier)
	notify.Run(ctx, n.Recovery)

	return nil
}

func (n notifier) Stop(_ context.Context) error {
	n.httpCli.CloseIdleConnections()
	return nil
}

func (n notifier) Verifier(ctx context.Context) error {
	verify, err := n.verifyRepo.GetNewVerification(ctx)
	if err != nil {
		return fmt.Errorf("[ERROR] worker verify: %w", err)
	}

	_, err = n.sendOTP(ctx, verify.Destination, verify.OTP)
	if err != nil {
		log.Printf("[ERROR] verifier send otp: %v", err)
		if vErr := n.verifyRepo.RollbackVerification(ctx, verify.ID); vErr != nil {
			return fmt.Errorf("rollback verifier: %w", vErr)
		}
		return nil
	}

	err = n.verifyRepo.FinishVerification(ctx, verify.ID)
	if err != nil {
		return fmt.Errorf("finish verification: %w", err)
	}
	return nil
}

func (n notifier) Recovery(ctx context.Context) error {
	recovery, err := n.recoveryRepo.GetNewRecovery(ctx)
	if err != nil {
		return fmt.Errorf("[ERROR] worker recovery: %w", err)
	}

	_, err = n.sendOTP(ctx, recovery.Destination, recovery.OTP)
	if err != nil {
		log.Printf("[ERROR] recovery send otp: %v", err)
		if rErr := n.recoveryRepo.RollbackRecovery(ctx, recovery.ID); rErr != nil {
			return fmt.Errorf("rollback recovery: %w", rErr)
		}
		return nil
	}

	err = n.recoveryRepo.FinishRecovery(ctx, recovery.ID)
	if err != nil {
		return fmt.Errorf("finish recovery: %w", err)
	}
	return nil
}

func (n notifier) sendOTP(ctx context.Context, destination, otp string) ([]byte, error) {
	notify := NewKazInfoTehReq(n.notifyConfig.User, n.notifyConfig.Pass, destination, otp)
	b, err := xml.Marshal(&notify)
	if err != nil {
		return nil, fmt.Errorf("send otp: %w", err)
	}
	responseBody, err := n.request(ctx, 3, nil, bytes.NewBuffer(b))
	if err != nil {
		return nil, fmt.Errorf("verifier: %w", err)
	}
	return responseBody, nil
}

func (n notifier) request(ctx context.Context, count int, err error, payload io.Reader) ([]byte, error) {
	const op = "request"
	if count == 0 {
		return nil, fmt.Errorf("%v:%w", op, err)
	}

	req, err := http.NewRequest(http.MethodGet, n.notifyConfig.URL, payload)
	if err != nil {
		return nil, fmt.Errorf("%v:%w", op, err)
	}

	resp, err := n.httpCli.Do(req.WithContext(ctx))
	if err != nil {
		time.Sleep(3 * time.Second)
		count--
		return n.request(ctx, count, err, payload)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code received: %d", resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%v:%w", op, err)
	}
	return b, nil
}

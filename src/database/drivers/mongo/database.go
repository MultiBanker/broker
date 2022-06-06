package mongo

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/go-multierror"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/MultiBanker/broker/src/database/drivers"
)

const (
	connectionTimeout = 3 * time.Second
	ensureIdxTimeout  = 10 * time.Second
	retries           = 1
)

type Mongo struct {
	MongoURL string
	Client   *mongo.Client
	DBName   string
	DB       *mongo.Database
	retries  int

	connectionTimeout time.Duration
	ensureIdxTimeout  time.Duration
}

func (m *Mongo) Name() string { return "mongo" }

func New(conf drivers.DataStoreConfig) (drivers.Datastore, error) {
	return &Mongo{
		MongoURL:          conf.URL,
		DBName:            conf.DBName,
		retries:           retries,
		connectionTimeout: connectionTimeout,
		ensureIdxTimeout:  ensureIdxTimeout,
	}, nil
}

func (m Mongo) Database() interface{} {
	return m.DB
}

func (m *Mongo) Connect() error {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), m.connectionTimeout)
	defer cancel()

	monitor := &event.PoolMonitor{
		Event: m.HandlePoolMonitor,
	}

	log.Printf("[INFO] Connecting to: %s", m.DBName)
	m.Client, err = mongo.Connect(ctx,
		options.Client().
			ApplyURI(m.MongoURL).
			SetMinPoolSize(100).
			SetMaxPoolSize(3000).
			SetHeartbeatInterval(3*time.Second).
			SetPoolMonitor(monitor),
	)
	if err != nil {
		return err
	}

	if err := m.Ping(); err != nil {
		return err
	}

	m.DB = m.Client.Database(m.DBName)

	// убеждаемся что созданы все необходимые индексы
	return m.ensureIndexes()
}

func (m *Mongo) HandlePoolMonitor(evt *event.PoolEvent) {
	switch evt.Type {
	case event.PoolClosedEvent:
		log.Println("[ERROR] DB connection closed.")
		m.reconnect()
	}
}

func (m *Mongo) IsConnecting() (bool, error) {
	if err := m.Connect(); err != nil {
		return false, err
	}
	return true, nil
}

func (m *Mongo) reconnect() {
	for {
		isConn, _ := m.IsConnecting()
		if isConn {
			break
		}
		// Reconnect interval
		time.Sleep(1 * time.Second)
	}
}

func (m *Mongo) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), m.connectionTimeout)
	defer cancel()

	return m.Client.Ping(ctx, readpref.Primary())
}

func (m *Mongo) Close(ctx context.Context) error {
	return m.Client.Disconnect(ctx)
}

// убеждается что все индексы построены
func (m *Mongo) ensureIndexes() error {
	ctx, cancel := context.WithTimeout(context.Background(), m.connectionTimeout)
	defer cancel()

	err := m.LoanIndexes(ctx)
	if err != nil {
		return err
	}
	err = m.MarketIndexes(ctx)
	if err != nil {
		return err
	}
	err = m.OrderIndexes(ctx)
	if err != nil {
		return err
	}
	err = m.PartnerOrderIndexes(ctx)
	if err != nil {
		return err
	}

	err = m.PartnerIndexes(ctx)
	if err != nil {
		return err
	}

	return m.SequenceIndexes(ctx)
}

// indexExistsByName проверяет существование индекса с именем name.
func (m *Mongo) indexExistsByName(ctx context.Context, collection *mongo.Collection, name string) (bool, error) {
	cur, err := collection.Indexes().List(ctx)
	if err != nil {
		return false, err
	}

	for cur.Next(ctx) {
		if name == cur.Current.Lookup("name").StringValue() {
			return true, nil
		}
	}

	return false, nil
}

func (m *Mongo) WithTransaction() func (ctx context.Context, sessionFunc func(session mongo.SessionContext) error) error {
	return func(ctx context.Context, sessionFunc func(session mongo.SessionContext) error) error {
		wc := writeconcern.New(writeconcern.WMajority())
		rc := readconcern.Snapshot()

		tnxOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)

		session, err := m.Client.StartSession()
		if err != nil {
			return err
		}

		defer session.EndSession(ctx)
		mongo.NewSessionContext(ctx, session)
		err = mongo.WithSession(ctx, session, func(sCtx mongo.SessionContext) error {
			if err = session.StartTransaction(tnxOpts); err != nil {
				return err
			}

			if err = sessionFunc(sCtx); err != nil {
				return err
			}

			if err = session.CommitTransaction(sCtx); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			if abortErr := session.AbortTransaction(context.Background()); abortErr != nil {
				return abortErr
			}
			return err
		}

		return nil
	}
}

func (m *Mongo) StartSession(ctx context.Context) (context.Context, drivers.TxCallback, error) {
	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	txOpts := options.Transaction().
		SetWriteConcern(wc).
		SetReadConcern(rc)

	session, err := m.Client.StartSession()
	if err != nil {
		return nil, nil, err
	}

	if err = session.StartTransaction(txOpts); err != nil {
		return nil, nil, err
	}

	return mongo.NewSessionContext(ctx, session), callback(session), nil
}

func callback(session mongo.Session) func(err error) error {
	return func(err error) error {
		var resultErr *multierror.Error

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		defer session.EndSession(ctx)

		if err == nil {
			err = session.CommitTransaction(ctx)
		}
		if err != nil {
			resultErr = multierror.Append(resultErr, err)
			if abortErr := session.AbortTransaction(ctx); abortErr != nil {
				resultErr = multierror.Append(resultErr, abortErr)
			}
		}
		return err
	}
}

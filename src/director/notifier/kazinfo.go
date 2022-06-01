package notifier

import "encoding/xml"

const (
	DefaultSendSMSID = "1"
	KazProvider      = "INFO_KAZ"
)

// http://docs.kazinfoteh.kz/ - претензии к документации сюда

// RESPONSE REQUEST KAZINFOTEH
type KazInfoTehReq struct {
	XMLName  xml.Name       `xml:"package"`
	Text     string         `xml:",chardata"`
	Login    string         `xml:"login,attr"`
	Password string         `xml:"password,attr"`
	Message  MessageInfoReq `xml:"message"`
}

func NewKazInfoTehReq(login string, password string, address string, text string) *KazInfoTehReq {
	return &KazInfoTehReq{
		Login:    login,
		Password: password,
		Message: MessageInfoReq{
			Msg: MessageInfoReqMsg{
				ID:        DefaultSendSMSID,
				Recipient: address,
				Text:      text,
				Sender:    KazProvider,
			},
		}}
}

type MessageInfoReq struct {
	Text string            `xml:",chardata"`
	Def  Default           `xml:"default"`
	Msg  MessageInfoReqMsg `xml:"msg"`
}
type Default struct {
	Text   string `xml:",chardata"`
	Sender string `xml:"sender,attr"`
}

type MessageInfoReqMsg struct {
	Text      string `xml:",chardata"`
	ID        string `xml:"id,attr"`
	Recipient string `xml:"recipient,attr"`
	Sender    string `xml:"sender,attr"`
	Type      int64  `xml:"type,attr"`
}

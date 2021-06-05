package sms

type SMS interface {
	Send(mobile, signName, templateCode, templateParam string) error
	SendVerCode(accessKey, accessKeySecret, endPoint, signName, templateCode, mobile, code string) error
}

func NewSms() SMS {
	return &AliSms{}
}

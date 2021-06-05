package sms

import (
	"errors"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/client"
	"github.com/alibabacloud-go/tea-rpc/client"
)

type AliSms struct{}

func (*AliSms) Send(mobile, signName, templateCode, templateParam string) error {
	var accessKey = ""
	var accessKeySecret = ""
	var endPoint = "dysmsapi.aliyuncs.com"

	config := &client.Config{
		AccessKeyId:     &accessKey,
		AccessKeySecret: &accessKeySecret,
		Endpoint:        &endPoint,
	}
	_result := &dysmsapi20170525.Client{}
	_result, _ = dysmsapi20170525.NewClient(config)

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  &mobile,
		SignName:      &signName,
		TemplateCode:  &templateCode,
		TemplateParam: &templateParam,
	}
	_, _err := _result.SendSms(sendSmsRequest)
	return _err
}

func (*AliSms) SendVerCode(accessKey, accessKeySecret, endPoint, signName, templateCode, mobile, code string) error {
	if len(code) > 6 {
		return errors.New("code len must > 6")
	}
	var templateValue = "{\"code\":\"" + code + "\"}"

	config := &client.Config{
		AccessKeyId:     &accessKey,
		AccessKeySecret: &accessKeySecret,
		Endpoint:        &endPoint,
	}
	_result := &dysmsapi20170525.Client{}
	_result, _ = dysmsapi20170525.NewClient(config)

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  &mobile,
		SignName:      &signName,
		TemplateCode:  &templateCode,
		TemplateParam: &templateValue,
	}
	_, _err := _result.SendSms(sendSmsRequest)
	return _err
}

package test

import (
	"github.com/sendcloud2013/sendcloud-sdk-go/sms"
	"testing"
)

func TestSendTemplateSms(t *testing.T) {
	client, err := sendcloud.NewSendCloudSms("**", "**")
	if err != nil {
		t.Error(err)
	}
	result, err := client.SendTemplateSms(&sendcloud.TemplateSms{
		TemplateId: 1,
		LabelId:    1,
		Phone:      "13800138000,13800138001",
		MsgType:    sendcloud.SMS,
		Vars: map[string]string{
			"name": "sendcloud",
		},
		Tag: map[string]string{
			"key": "value",
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestSendVoiceSms(t *testing.T) {
	client, err := sendcloud.NewSendCloudSms("**", "**")
	if err != nil {
		t.Error(err)
	}
	result, err := client.SendVoiceSms(&sendcloud.VoiceSms{
		Code:    "123456",
		LabelId: 1,
		Phone:   "13800138000",
		Tag: map[string]string{
			"key": "value",
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestSendCodeSms(t *testing.T) {
	client, err := sendcloud.NewSendCloudSms("**", "**")
	if err != nil {
		t.Error(err)
	}
	result, err := client.SendCodeSms(&sendcloud.CodeSms{
		Code:    "123456",
		LabelId: 1,
		Phone:   "13800138000",
		Tag: map[string]string{
			"key": "value",
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

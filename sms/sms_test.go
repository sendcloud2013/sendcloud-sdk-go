package sendcloud

import "testing"

func TestSendTemplateSms(t *testing.T) {
	client, err := NewSendCloudSms("**", "**")
	if err != nil {
		t.Error(err)
	}
	result, err := client.SendTemplateSms(&TemplateSms{
		TemplateId: 1,
		LabelId:    1,
		Phone:      "13800138000,13800138001",
		MsgType:    SMS,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestSendVoiceSms(t *testing.T) {
	client, err := NewSendCloudSms("**", "**")
	if err != nil {
		t.Error(err)
	}
	result, err := client.SendVoiceSms(&VoiceSms{
		Code:    "123456",
		LabelId: 1,
		Phone:   "13800138000",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestSendCodeSms(t *testing.T) {
	client, err := NewSendCloudSms("**", "**")
	if err != nil {
		t.Error(err)
	}
	result, err := client.SendCodeSms(&CodeSms{
		Code:    "123456",
		LabelId: 1,
		Phone:   "13800138000",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

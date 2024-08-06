package sendcloud

import (
	"net/http"
)

const (
	SMS          = 0
	MMS          = 1
	INTERNAT_SMS = 2
	VOICE        = 3
	QR_CODE      = 4
	YX           = 5
)

const (
	smsBasePath         = "https://api.sendcloud.net/smsapi"
	sendSmsTemplatePath = "/send"
	sendSmsVoicePath    = "/sendVoice"
	sendSmsCodePath     = "/sendCode"
)

type SendCloudSms struct {
	smsUser string
	smsKey  string
	apiBase string
	client  *http.Client
}

type Response struct {
	*http.Response
}

type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:"message"` // error message
}

type TemplateSms struct {
	TemplateId    int
	LabelId       int
	MsgType       int
	Phone         string
	Vars          string
	SendRequestId string
	Tag           string
}

type VoiceSms struct {
	Phone         string
	Code          string
	LabelId       int
	SendRequestId string
	Tag           string
}

type CodeSms struct {
	MsgType       int
	Phone         string
	SignId        int
	SignName      string
	Code          string
	LabelId       int
	SendRequestId string
	Tag           string
}

type SendSmsResult struct {
	Result     bool        `json:"result"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Info       interface{} `json:"info"`
}

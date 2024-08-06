package sendcloud

import (
	"errors"
	"strings"
)

func (client *SendCloudSms) validateConfig() error {
	if len(client.apiBase) == 0 {
		client.apiBase = smsBasePath
	}
	switch {
	case len(client.smsUser) == 0:
		return errors.New("smsUser cannot be empty")
	case len(client.smsKey) == 0:
		return errors.New("smsKey cannot be empty")
	}
	return nil
}

func isValidMsgType(msgType int) bool {
	return msgType == SMS ||
		msgType == MMS ||
		msgType == INTERNAT_SMS ||
		msgType == VOICE ||
		msgType == QR_CODE ||
		msgType == YX
}

func ValidatePhoneNumbers(phone string) error {
	phoneNumbers := strings.Split(phone, ",")

	if len(phoneNumbers) > 2000 {
		return errors.New("the number of mobile phone numbers exceeds the maximum limit of 2,000")
	}

	for _, number := range phoneNumbers {
		trimmedNumber := strings.TrimSpace(number)
		if trimmedNumber == "" {
			return errors.New("phone number can not be empty")
		}
	}
	return nil
}

func (s *TemplateSms) validateTemplateSms() error {
	switch {
	case s.TemplateId == 0:
		return errors.New("templateId value is illegal")
	case !isValidMsgType(s.MsgType):
		return errors.New("msgType value is illegal")
	case len(s.Phone) == 0:
		return errors.New("phone cannot be empty")
	case len(s.SendRequestId) > 128:
		return errors.New("sendRequestId cannot exceed 128 characters")
	}
	if err := ValidatePhoneNumbers(s.Phone); err != nil {
		return err
	}
	return nil
}

func (s *VoiceSms) validateVoiceSms() error {
	switch {
	case len(s.Code) == 0:
		return errors.New("code cannot be empty")
	case len(s.Phone) == 0:
		return errors.New("phone cannot be empty")
	case len(s.SendRequestId) > 128:
		return errors.New("sendRequestId cannot exceed 128 characters")
	}
	return nil
}

func (s *CodeSms) validateCodeSms() error {
	switch {
	case !isValidMsgType(s.MsgType):
		return errors.New("msgType value is illegal")
	case len(s.Phone) == 0:
		return errors.New("phone cannot be empty")
	case len(s.Code) == 0:
		return errors.New("code cannot be empty")
	case len(s.SendRequestId) > 128:
		return errors.New("sendRequestId cannot exceed 128 characters")
	}
	return nil
}

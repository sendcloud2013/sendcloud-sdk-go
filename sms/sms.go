package sendcloud

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func NewSendCloudSms(smsUser string, smsKey string) (*SendCloudSms, error) {
	switch {
	case len(smsUser) == 0:
		return nil, errors.New("NewSendCloudSms: smsUser cannot be empty")
	case len(smsKey) == 0:
		return nil, errors.New("NewSendCloudSms: smsKey cannot be empty")
	}
	return &SendCloudSms{
		smsUser: smsUser,
		smsKey:  smsKey,
		apiBase: smsBasePath,
		client:  http.DefaultClient,
	}, nil
}

func (client *SendCloudSms) SendTemplateSms(args *TemplateSms) (*SendSmsResult, error) {
	if err := client.validateConfig(); err != nil {
		return nil, fmt.Errorf("SendTemplateSms: %w", err)
	}
	if err := args.validateTemplateSms(); err != nil {
		return nil, fmt.Errorf("SendTemplateSms: %w", err)
	}
	params, err := client.prepareSendTemplateSmsParams(args)
	if err != nil {
		return nil, fmt.Errorf("SendTemplateSms: %w", err)
	}
	signature := client.calculateSignature(params)
	params.Set("signature", signature)
	sendSmsTemplateUrl := client.apiBase + sendSmsTemplatePath
	formDataEncoded := params.Encode()
	req, err := http.NewRequest("POST", sendSmsTemplateUrl, bytes.NewBufferString(formDataEncoded))
	if err != nil {
		return nil, fmt.Errorf("SendTemplateSms: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	responseData := new(SendSmsResult)
	err = client.request(req, responseData)
	if err != nil {
		return responseData, err
	}
	return responseData, nil
}

func (client *SendCloudSms) SendVoiceSms(args *VoiceSms) (*SendSmsResult, error) {
	if err := client.validateConfig(); err != nil {
		return nil, fmt.Errorf("SendVoiceSms: %w", err)
	}
	if err := args.validateVoiceSms(); err != nil {
		return nil, fmt.Errorf("SendVoiceSms: %w", err)
	}
	params, err := client.prepareSendVoiceSmsParams(args)
	if err != nil {
		return nil, fmt.Errorf("SendVoiceSms: %w", err)
	}
	signature := client.calculateSignature(params)
	params.Set("signature", signature)
	sendSmsVoiceUrl := client.apiBase + sendSmsVoicePath
	formDataEncoded := params.Encode()
	req, err := http.NewRequest("POST", sendSmsVoiceUrl, bytes.NewBufferString(formDataEncoded))
	if err != nil {
		return nil, fmt.Errorf("SendVoiceSms: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	responseData := new(SendSmsResult)
	err = client.request(req, responseData)
	if err != nil {
		return responseData, err
	}
	return responseData, nil
}

func (client *SendCloudSms) SendCodeSms(args *CodeSms) (*SendSmsResult, error) {
	if err := client.validateConfig(); err != nil {
		return nil, fmt.Errorf("SendCodeSms: %w", err)
	}
	if err := args.validateCodeSms(); err != nil {
		return nil, fmt.Errorf("SendCodeSms: %w", err)
	}
	params, err := client.prepareSendCodeSmsParams(args)
	if err != nil {
		return nil, fmt.Errorf("SendCodeSms: %w", err)
	}
	signature := client.calculateSignature(params)
	params.Set("signature", signature)
	sendSmsCodeUrl := client.apiBase + sendSmsCodePath
	formDataEncoded := params.Encode()
	req, err := http.NewRequest("POST", sendSmsCodeUrl, bytes.NewBufferString(formDataEncoded))
	if err != nil {
		return nil, fmt.Errorf("SendCodeSms: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	responseData := new(SendSmsResult)
	err = client.request(req, responseData)
	if err != nil {
		return responseData, err
	}
	return responseData, nil
}

func (client *SendCloudSms) request(req *http.Request, responseResult *SendSmsResult) error {
	resp, err := client.client.Do(req)
	if err != nil {
		return err
	}
	err = checkResponse(resp)
	if err != nil {
		defer resp.Body.Close()
		return err
	}

	if responseResult != nil {
		err = json.NewDecoder(resp.Body).Decode(responseResult)
		if err != nil {
			return err
		}
		if responseResult.StatusCode != http.StatusOK {
			return errors.New(responseResult.Message)
		}
	}
	return err
}

func checkResponse(r *http.Response) error {
	if r.StatusCode == http.StatusOK {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	if r.StatusCode == http.StatusNotFound {
		errorResponse.Message = "Not Found"
		return errorResponse
	}
	data, err := io.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message)
}

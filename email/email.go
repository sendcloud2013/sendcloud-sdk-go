package sendcloud
import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)


func NewSendCloud(apiUser string, apiKey string) (*SendCloud, error) {
	switch {
	case len(apiUser) == 0:
		return nil, errors.New("NewSendCloud: apiUser cannot be empty")
	case len(apiKey) == 0:
		return nil, errors.New("NewSendCloud: apiKey cannot be empty")
	}

	sc := &SendCloud{
		apiUser: apiUser,
		apiKey:  apiKey,
		apiBase: APIBase,
		client:  http.DefaultClient,
	}
	return sc, nil
}

func (client *SendCloud) request(ctx context.Context, req *http.Request, responseResult *SendEmailResult) error {
	req = req.WithContext(ctx)
	resp, err := client.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
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
		if responseResult.StatusCode!= http.StatusOK {
			return errors.New(responseResult.Message)
		}
	}
	return err
}


func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message)
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


func (client *SendCloud) SendCommonEmail(ctx context.Context, args *CommonMail) (*SendEmailResult, error) {
	if err := client.validateConfig(); err != nil {
		return nil, fmt.Errorf("SendCommonEmail: %w", err)
	}
	if err := args.validateCommonEmail(); err != nil {
		return nil, fmt.Errorf("SendCommonEmail: %w", err)
	}
	var req *http.Request
	var err error
	sendCommonUrl := client.apiBase + sendCommonPath
	if args.Body.Attachments == nil {
		params:= client.PrepareSendCommonEmailParams(args)
		formDataEncoded := params.Encode()
		req, err = http.NewRequest("POST", sendCommonUrl, bytes.NewBufferString(formDataEncoded))
		if err != nil {
			return nil, fmt.Errorf("SendCommonEmail: %w", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		multipartWriter,payload, err := client.MultipartSendCommonMail(args)
		if err != nil {
			return nil, fmt.Errorf("SendCommonEmail: %w", err)
		}
		req, err = http.NewRequest("POST", sendCommonUrl, payload)
		if err != nil {
			return nil, fmt.Errorf("SendCommonEmail: %w", err)
		}
		req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	}
	responseData := new(SendEmailResult)
	err = client.request(ctx, req, responseData)
	if err != nil {
		return responseData, err
	}
	return responseData, nil
}

func (client *SendCloud) SendTemplateEmail(ctx context.Context, args *TemplateMail) (*SendEmailResult, error) {
	if err := client.validateConfig(); err != nil {
		return nil, fmt.Errorf("SendTemplateEmail: %w", err)
	}
	if err := args.validateTemplateMail(); err != nil {
		return nil, fmt.Errorf("SendTemplateEmail: %w", err)
	}
	var req *http.Request
	var err error
	sendTemplateUrl := client.apiBase+sendTemplatePath
	if args.Body.Attachments == nil {
		params:= client.PrepareSendTemplateEmailParams(args)
		formDataEncoded := params.Encode()
		req, err = http.NewRequest("POST", sendTemplateUrl, bytes.NewBufferString(formDataEncoded))
		if err != nil {
			return nil, fmt.Errorf("SendTemplateEmail: %w", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		multipartWriter,payload, err := client.MultipartSendTemplateEmail(args)
		if err != nil {
			return nil, fmt.Errorf("SendTemplateEmail: %w", err)
		}
		req, err = http.NewRequest("POST", sendTemplateUrl, payload)
		if err != nil {
			return nil, fmt.Errorf("SendTemplateEmail: %w", err)
		}
		req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	}
	responseData := new(SendEmailResult)
	err = client.request(ctx, req, responseData)
	if err != nil {
		return responseData, err
	}
	return responseData, nil
}

func (client *SendCloud) SendCalendarMail(ctx context.Context, args *CalendarMail) (*SendEmailResult, error) {
	if err := client.validateConfig(); err != nil {
		return nil, fmt.Errorf("SendCalendarMail: %w", err)
	}
	if err := args.validateSendCalendarMail(); err != nil {
		return nil, fmt.Errorf("SendCalendarMail: %w", err)
	}
	if err := args.Calendar.validateCalendarMail(); err != nil {
		return nil, fmt.Errorf("SendCalendarMail: %w", err)
	}
	var req *http.Request
	var err error
	sendCalendarUrl := client.apiBase + sendCalendarPath
	if args.Body.Attachments == nil {
		params:= client.PrepareSendCalendarMailParams(args)
		formDataEncoded := params.Encode()
		req, err = http.NewRequest("POST", sendCalendarUrl, bytes.NewBufferString(formDataEncoded))
		if err != nil {
			return nil, fmt.Errorf("SendCalendarMail: %w", err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		multipartWriter,payload, err := client.MultipartSendCalendarMail(args)
		if err != nil {
			return nil, fmt.Errorf("SendCalendarMail: %w", err)
		}
		req, err = http.NewRequest("POST", sendCalendarUrl, payload)
		if err != nil {
			return nil, fmt.Errorf("SendCalendarMail: %w", err)
		}
		req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	}
	responseData := new(SendEmailResult)
	err = client.request(ctx, req, responseData)
	if err != nil {
		return responseData, err
	}
	return responseData, nil
}
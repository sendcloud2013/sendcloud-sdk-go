package sendcloud

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/url"
	"strconv"
)


func (client *SendCloud) PrepareReceiverParams(e *MailReceiver) url.Values {
	params := url.Values{}
	params.Set("apiUser", client.apiUser)
	params.Set("apiKey", client.apiKey)
	if e.To!= "" {
		params.Set("to", e.To)
	}
	if e.CC!= "" {
		params.Set("cc", e.CC)
	}
	if e.BCC!= "" {
		params.Set("bcc", e.BCC)
	}
	if e.UseAddressList {
		params.Set("useAddressList", strconv.FormatBool(e.UseAddressList))
	}
	return params
}

func (e *MailBody) PrepareMailBodyParams(params *url.Values){
	params.Set("from", e.From)

	params.Set("subject", e.Subject)
	if e.ContentSummary!= "" {
		params.Set("contentSummary", e.ContentSummary)
	}
	if e.FromName!= "" {
		params.Set("fromName", e.FromName)
	}

	if e.ReplyTo!= "" {
		params.Set("replyTo", e.ReplyTo)
	}
	if e.LabelName!= "" {
		params.Set("labelName", e.LabelName)
	}
	if len(e.Headers) > 0 {
		headers, _ := json.Marshal(e.Headers)
		params.Set("headers", string(headers))
	}
	if !e.Xsmtpapi.IsEmpty() {
		xsmtpapi, _ := json.Marshal(e.Xsmtpapi)
		params.Set("xsmtpapi", string(xsmtpapi))
	}

	if e.SendRequestID!= "" {
		params.Set("sendRequestId", e.SendRequestID)
	}
	if e.RespEmailID {
		params.Set("respEmailId", strconv.FormatBool(e.RespEmailID))
	}
	if e.UseNotification {
		params.Set("useNotification", strconv.FormatBool(e.UseNotification))
	}
}

func (e *MailCalendar) PrepareMailCalendarParams(params *url.Values){
	params.Set("startTime", e.StartTime.Format("2006-01-02 15:04:05"))
	params.Set("endTime", e.EndTime.Format("2006-01-02 15:04:05"))
	params.Set("title", e.Title)
	params.Set("organizerName", e.OrganizerName)
	params.Set("organizerEmail", e.OrganizerEmail)
	params.Set("location", e.Location)
	if e.Description!= "" {
		params.Set("description", e.Description)
	}
	params.Set("participatorNames", e.ParticipatorNames)
	if e.ParticipatorEmails!= "" {
        params.Set("participatorEmails", e.ParticipatorEmails)
    }
    if e.UID!= "" {
        params.Set("uid", e.UID)
    }
    if e.IsCancel {
        params.Set("isCancel", strconv.FormatBool(e.IsCancel))
    }
    if e.IsUpdate {
        params.Set("isUpdate", strconv.FormatBool(e.IsUpdate))
    }
    if e.ValarmTime!= 0 {
        params.Set("valarmTime", strconv.Itoa(e.ValarmTime))
    }
}

func (client *SendCloud) PrepareSendCommonEmailParams(e *CommonMail) url.Values {
	params := client.PrepareReceiverParams(&e.Receiver)
	e.Body.PrepareMailBodyParams(&params)
	if e.Content.Plain!= "" {
		params.Set("plain", e.Content.Plain)
	}
	if e.Content.Html!= "" {
		params.Set("html", e.Content.Html)
	}
	return params
}

func (client *SendCloud) PrepareSendTemplateEmailParams (e *TemplateMail) url.Values {
	params := client.PrepareReceiverParams(&e.Receiver)
	e.Body.PrepareMailBodyParams(&params)
	params.Set("templateInvokeName", e.Content.TemplateInvokeName)
	return params
}

func (client *SendCloud) PrepareSendCalendarMailParams(e *CalendarMail) url.Values {
	params := client.PrepareReceiverParams(&e.Receiver)
	e.Body.PrepareMailBodyParams(&params)
	if e.Content.Plain!= "" {
		params.Set("plain", e.Content.Plain)
	}
	if e.Content.Html!= "" {
		params.Set("html", e.Content.Html)
	}
	e.Calendar.PrepareMailCalendarParams(&params)
	return params
}

func (e *MailReceiver) multipartReceiver(client *SendCloud,multipartWriter *multipart.Writer) error {

	var err error

	if client.apiUser != "" {
		err = multipartWriter.WriteField("apiUser", client.apiUser)
		if err!= nil {
			return err
		}
	}

	if client.apiKey != "" {
		err = multipartWriter.WriteField("apiKey", client.apiKey)
		if err!= nil {
			return err
		}
	}

	if e.To != "" {
		err =   multipartWriter.WriteField("to", e.To)
		if err != nil {
			return err
		}
	}

	if e.CC != "" {
		err = multipartWriter.WriteField("cc", e.CC)
		if err != nil {
			return err
		}
	}

	if e.BCC != "" {
		err = multipartWriter.WriteField("bcc", e.BCC)
		if err != nil {
			return err
		}
	}

	if e.UseAddressList {
		useAddressListStr := strconv.FormatBool(e.UseAddressList)
		err = multipartWriter.WriteField("useAddressList", useAddressListStr)
		if err != nil {
			return err
		}
	}

	return nil
}



func (e *MailBody) multipartMailBody(multipartWriter *multipart.Writer) error {

	var err error

	var partWriter io.Writer

	if e.From != "" {
		err = multipartWriter.WriteField("from", e.From)
		if err != nil {
			return err
		}
	}

	if e.Subject != "" {
		err = multipartWriter.WriteField("subject", e.Subject)
		if err!= nil {
			return err
		}
	}


	if e.ContentSummary!= "" {
		err = multipartWriter.WriteField("contentSummary", e.ContentSummary)
		if err!= nil {
			return err
		}
	}

	if e.FromName != "" {
		err = multipartWriter.WriteField("fromName", e.FromName)
		if err != nil {
			return err
		}
	}

	if e.ReplyTo != "" {
		err = multipartWriter.WriteField("replyTo", e.ReplyTo)
		if err != nil {
			return err
		}
	}

	if e.LabelName != "" {
		err = multipartWriter.WriteField("labelName", e.LabelName)
		if err != nil {
			return err
		}
	}

	if len(e.Headers) > 0 {
		headers, _ := json.Marshal(e.Headers)
		err = multipartWriter.WriteField("headers", string(headers))
		if err != nil {
			return err
		}
	}

	if e.Attachments != nil {
		for _, attachment := range e.Attachments {
			defer attachment.Close()
			partWriter, err = multipartWriter.CreateFormFile("attachments", attachment.Name())
			if err != nil {
				return err
			}
			_, err = io.Copy(partWriter, attachment)
			if err != nil {
				return err
			}
		}
	}

	if !e.Xsmtpapi.IsEmpty() {
		xsmtpapi, err := json.Marshal(e.Xsmtpapi)
		if err != nil {
			return err
		}
		err = multipartWriter.WriteField("xsmtpapi", string(xsmtpapi))
		if err != nil {
			return err
		}
	}

	if e.SendRequestID != "" {
		err = multipartWriter.WriteField("sendRequestId", e.SendRequestID)
		if err != nil {
			return err
		}
	}

	if e.RespEmailID {
		respEmailIDStr := strconv.FormatBool(e.RespEmailID)
		err = multipartWriter.WriteField("respEmailId", respEmailIDStr)
		if err!= nil {
			return err
		}
	}

	if e.UseNotification {
		notificationStr := strconv.FormatBool(e.UseNotification)
		err = multipartWriter.WriteField("useNotification", notificationStr)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *MailCalendar) multipartMailCalendar(multipartWriter *multipart.Writer) error {
	var err error

	if !e.StartTime.IsZero() {
		err = multipartWriter.WriteField("startTime", e.StartTime.Format("2006-01-02 15:04:05"))
		if err != nil {
			return err
		}
	}

	if !e.EndTime.IsZero() {
		err = multipartWriter.WriteField("endTime", e.EndTime.Format("2006-01-02 15:04:05"))
        if err != nil {
            return err
        }
	}

	if e.Title!="" {
		err = multipartWriter.WriteField("title", e.Title)
        if err != nil {
            return err
        }
	}
	if e.OrganizerName!= "" {
		err = multipartWriter.WriteField("organizerName", e.OrganizerName)
		if err != nil {
            return err
        }
	}
	if e.OrganizerEmail!= "" {
		err = multipartWriter.WriteField("organizerEmail", e.OrganizerEmail)
        if err != nil {
            return err
        }
	}
	if e.Location!= "" {
		err = multipartWriter.WriteField("location", e.Location)
        if err != nil {
            return err
        }
	}
	if e.Description!= "" {
		err = multipartWriter.WriteField("description", e.Description)
        if err != nil {
            return err
        }
	}
	if e.ParticipatorNames!= "" {
		err = multipartWriter.WriteField("participatorNames", e.ParticipatorNames)
		if err != nil {
			return err
		}
	}
	if e.ParticipatorEmails!= "" {
		err = multipartWriter.WriteField("participatorEmails", e.ParticipatorEmails)
        if err != nil {
            return err
        }
	}
	if e.UID != "" {
		err = multipartWriter.WriteField("uid", e.UID)
        if err != nil {
            return err
        }
	}
	if e.IsCancel{
		isCancelStr := strconv.FormatBool(e.IsCancel)
        err = multipartWriter.WriteField("isCancel", isCancelStr)
        if err != nil {
            return err
        }
	}
	if e.IsUpdate {
		isUpdateStr := strconv.FormatBool(e.IsUpdate)
		err = multipartWriter.WriteField("isUpdate", isUpdateStr)
		if err != nil {
            return err
        }
	}
	if e.ValarmTime != 0 {
		err = multipartWriter.WriteField("valarmTime", strconv.Itoa(e.ValarmTime))
        if err != nil {
            return err
        }
	}
	return nil
}

func (client *SendCloud) MultipartSendCommonMail(e *CommonMail) (*multipart.Writer,*bytes.Buffer, error) {
	buf := bytes.Buffer{}
	multipartWriter := multipart.NewWriter(&buf)
	var err error

	err = e.Receiver.multipartReceiver(client,multipartWriter)
	if err != nil {
		return multipartWriter,nil, err
	}

	err = e.Body.multipartMailBody(multipartWriter)
	if err != nil {
		return multipartWriter,nil, err
	}

	if e.Content.Html!= "" {
		err = multipartWriter.WriteField("html", e.Content.Html)
        if err!= nil {
            return multipartWriter,nil, err
        }
	}

	if e.Content.Plain!= "" {
		err = multipartWriter.WriteField("plain", e.Content.Plain)
        if err!= nil {
            return multipartWriter,nil, err
        }
	}

	multipartWriter.Close()
	return multipartWriter,&buf, nil
}

func (client *SendCloud) MultipartSendTemplateEmail(e *TemplateMail) (*multipart.Writer,*bytes.Buffer, error) {
	buf := bytes.Buffer{}
	multipartWriter := multipart.NewWriter(&buf)
	var err error

	err = e.Receiver.multipartReceiver(client,multipartWriter)
	if err != nil {
		return multipartWriter,nil, err
	}

	err = e.Body.multipartMailBody(multipartWriter)
	if err != nil {
		return multipartWriter,nil, err
	}

	if e.Content.TemplateInvokeName!= "" {
		err = multipartWriter.WriteField("templateInvokeName", e.Content.TemplateInvokeName)
		if err!= nil {
			return multipartWriter,nil, err
		}
	}
	multipartWriter.Close()
	return multipartWriter,&buf, nil
}

func (client *SendCloud) MultipartSendCalendarMail(e *CalendarMail) (*multipart.Writer,*bytes.Buffer, error) {
	buf := bytes.Buffer{}
	multipartWriter := multipart.NewWriter(&buf)
	var err error

	err = e.Receiver.multipartReceiver(client,multipartWriter)
	if err != nil {
		return multipartWriter,nil, err
	}

	err = e.Body.multipartMailBody(multipartWriter)
	if err != nil {
		return multipartWriter,nil, err
	}

	if e.Content.Html!= "" {
		err = multipartWriter.WriteField("html", e.Content.Html)
		if err!= nil {
			return multipartWriter,nil, err
		}
	}

	if e.Content.Plain!= "" {
		err = multipartWriter.WriteField("plain", e.Content.Plain)
		if err!= nil {
			return multipartWriter,nil, err
		}
	}
	err = e.Calendar.multipartMailCalendar(multipartWriter)
	if err != nil {
		return multipartWriter,nil, err
	}
	multipartWriter.Close()
	return multipartWriter,&buf, nil
}
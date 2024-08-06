package sendcloud

import (
	"errors"
	"fmt"
	"strings"
)

const MAX_RECEIVERS = 100
const MAX_MAILLIST = 5

func (client *SendCloud) validateConfig() error {
	if len(client.apiBase) == 0 {
		client.apiBase = APIBase
	}
	switch {
	case len(client.apiUser) == 0:
		return errors.New("apiUser cannot be empty")
	case len(client.apiKey) == 0:
		return errors.New("apiKey cannot be empty")
	}
	return nil
}

func (e *TemplateMail) validateTemplateMail() error {
	if len(e.Receiver.To) == 0 && len(e.Body.Xsmtpapi.To) == 0 {
		return errors.New("to cannot be empty")
	}
	if len(e.Body.Xsmtpapi.To) == 0 || e.Receiver.UseAddressList {
		if err := e.Receiver.validateReceiver(); err != nil {
			return err
		}
	}
	if !e.Receiver.UseAddressList && !e.Body.Xsmtpapi.IsEmpty() {
		err := e.Body.Xsmtpapi.validateXSMTPAPI()
		if err != nil {
			return err
		}
	}
	if err := e.Body.validateMailBody(); err != nil {
		return err
	}
	if e.Content.TemplateInvokeName == "" {
		return errors.New("templateInvokeName cannot be empty")
	}
	return nil
}

func (e *CommonMail) validateCommonEmail() error {
	if len(e.Receiver.To) == 0 && len(e.Body.Xsmtpapi.To) == 0 {
		return errors.New("to cannot be empty")
	}
	if len(e.Body.Xsmtpapi.To) == 0 || e.Receiver.UseAddressList {
		if err := e.Receiver.validateReceiver(); err != nil {
			return err
		}
	}
	if !e.Receiver.UseAddressList && !e.Body.Xsmtpapi.IsEmpty() {
		err := e.Body.Xsmtpapi.validateXSMTPAPI()
		if err != nil {
			return err
		}
	}
	if err := e.Body.validateMailBody(); err != nil {
		return err
	}
	if len(e.Content.Html) == 0 && len(e.Content.Plain) == 0 {
		return errors.New("html or plain cannot be empty")
	}
	return nil
}

func (e *CalendarMail) validateSendCalendarMail() error {
	if len(e.Receiver.To) == 0 && len(e.Body.Xsmtpapi.To) == 0 {
		return errors.New("to cannot be empty")
	}
	if len(e.Body.Xsmtpapi.To) == 0 || e.Receiver.UseAddressList {
		if err := e.Receiver.validateReceiver(); err != nil {
			return err
		}
	}
	if !e.Receiver.UseAddressList && !e.Body.Xsmtpapi.IsEmpty() {
		err := e.Body.Xsmtpapi.validateXSMTPAPI()
		if err != nil {
			return err
		}
	}
	if err := e.Body.validateMailBody(); err != nil {
		return err
	}
	if len(e.Content.Html) == 0 && len(e.Content.Plain) == 0 {
		return errors.New("html or plain cannot be empty")
	}
	if err := e.Calendar.validateCalendarMail(); err != nil {
		return err
	}
	return nil
}

func (e *MailCalendar) validateCalendarMail() error {
	switch {
	case e.StartTime.IsZero():
		return errors.New("startTime cannot be empty")
	case e.EndTime.IsZero():
		return errors.New("endTime cannot be empty")
	case e.StartTime.After(e.EndTime):
		return errors.New("startTime cannot be after endTime")
	case len(e.Title) == 0:
		return errors.New("title cannot be empty")
	case len(e.OrganizerName) == 0:
		return errors.New("organizerName cannot be empty")
	case len(e.OrganizerEmail) == 0:
		return errors.New("organizerEmail cannot be empty")
	case len(e.Location) == 0:
		return errors.New("location cannot be empty")
	case len(e.ParticipatorNames) == 0:
		return errors.New("participatorNames cannot be empty")
	}
	return nil
}

func (e *MailReceiver) validateReceiver() error {
	if len(e.To) == 0 {
		return errors.New("to cannot be empty")
	}
	if e.UseAddressList {
		to := strings.Split(e.To, ";")
		if len(to) > MAX_MAILLIST {
			return errors.New("address list exceeds limit")
		}
	} else {
		to := strings.Split(e.To, ";")
		cc := strings.Split(e.CC, ";")
		bcc := strings.Split(e.BCC, ";")
		receivers := len(to)
		receivers += len(cc)
		receivers += len(bcc)
		// Check if the total number of receivers exceeds the maximum allowed
		if receivers > MAX_RECEIVERS {
			return errors.New("the total number of receivers exceeds the maximum allowed")
		}
	}
	return nil
}

func (e *MailBody) validateMailBody() error {
	switch {
	case len(e.From) == 0:
		return errors.New("from cannot be empty")
	case len(e.Subject) == 0:
		return errors.New("subject cannot be empty")
	}
	return nil
}

func (x XSMTPAPI) validateXSMTPAPI() error {
	if len(x.To) != 0 {
		if len(x.To) > MAX_RECEIVERS {
			return errors.New("the total number of receivers exceeds the maximum allowed")
		}
		if len(x.Sub) != 0 {
			for key, value := range x.Sub {
				if !(len(key) >= 2 && key[0] == '%' && key[len(key)-1] == '%') {
					return errors.New(fmt.Sprintf("the key needs to be in the format '%%...%%'; [%s] does not satisfy this condition", key))
				}
				if len(value) != len(x.To) {
					return errors.New(fmt.Sprintf("the length of sub[%s] list and the number of to elements are not equal", key))
				}
			}
		}
	}
	if len(x.Pubsub) != 0 {
		for key := range x.Pubsub {
			if !(len(key) >= 2 && key[0] == '%' && key[len(key)-1] == '%') {
				return errors.New(fmt.Sprintf("the key needs to be in the format '%%...%%'; [%s] does not satisfy this condition", key))
			}
        }
	}
	if x.Filters != nil {
		err := x.Filters.ValidateFilter()
		if err != nil {
			return err
		}
	}
	return nil
}

func (f Filter) ValidateFilter() error {
	switch {
	case f.SubscriptionTracking.Settings.Enable != "0" && f.SubscriptionTracking.Settings.Enable != "1":
		return errors.New("subscriptionTracking invalid value for Enable, must be '0' or '1'")
	case f.OpenTracking.Settings.Enable != "0" && f.OpenTracking.Settings.Enable != "1":
		return errors.New("openTracking invalid value for Enable, must be '0' or '1'")
	case f.ClickTracking.Settings.Enable != "0" && f.ClickTracking.Settings.Enable != "1":
		return errors.New("clickTracking invalid value for Enable, must be '0' or '1'")
	}
	return nil
}

package test

import (
	"context"
	"github.com/sendcloud2013/sendcloud-sdk-go/email"
	"os"
	"testing"
	"time"
)

func TestSendCommonEmail(t *testing.T) {
	client, err := sendcloud.NewSendCloud("*", "*")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	args := &sendcloud.CommonMail{
		Receiver: sendcloud.MailReceiver{
			To: "a@ifaxin.com;b@ifaxin.com",
		},
		Body: sendcloud.MailBody{
			From:     "SendCloud@SendCloud.com",
			Subject:  "Email from SendCloud SDK",
			FromName: "SendCloud",
		},
		Content: sendcloud.TextContent{
			Html: "<p>This is an HTML email.</p>",
		},
	}
	result, err := client.SendCommonEmail(ctx, args)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestSendCommonEmailWithVars(t *testing.T) {
	client, err := sendcloud.NewSendCloud("*", "*")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	args := &sendcloud.CommonMail{
		Body: sendcloud.MailBody{
			From:     "SendCloud@SendCloud.com",
			Subject:  "Email from SendCloud SDK",
			FromName: "SendCloud",
		},
		Content: sendcloud.TextContent{
			Html: "<p>This is an HTML email.</p>",
		},
	}
	xsmtpapi := sendcloud.XSMTPAPI{
		To: []string{"a@ifaxin.com", "b@ifaxin.com"},
		Sub: map[string][]interface{}{
			"%name%":  {"jack", "rose"},
			"%money%": {"199", "299"},
		},
		Filters: &sendcloud.Filter{
			SubscriptionTracking: sendcloud.TrackingFilter{Settings: sendcloud.FilterSettings{Enable: "1"}},
			OpenTracking:         sendcloud.TrackingFilter{Settings: sendcloud.FilterSettings{Enable: "1"}},
			ClickTracking:        sendcloud.TrackingFilter{Settings: sendcloud.FilterSettings{Enable: "1"}},
		},
		Settings: &sendcloud.Settings{
			Unsubscribe: sendcloud.UnsubscribeSettings{PageID: []int{1, 2}},
		},
	}
	args.Body.SetXsmtpapi(xsmtpapi)

	result, err := client.SendCommonEmail(ctx, args)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestSendCommonEmailWithAttachment(t *testing.T) {
	client, err := sendcloud.NewSendCloud("*", "*")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	attachment1, err := os.Open("path/to/attachment1.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer attachment1.Close()
	args := &sendcloud.CommonMail{
		Receiver: sendcloud.MailReceiver{
			To: "a@ifaxin.com;b@ifaxin.com",
		},
		Body: sendcloud.MailBody{
			From:     "SendCloud@SendCloud.com",
			Subject:  "Email from SendCloud SDK",
			FromName: "SendCloud",
		},
		Content: sendcloud.TextContent{
			Html: "<p>This is an HTML email.</p>",
		},
	}
	args.Body.AddAttachment(attachment1)
	result, err := client.SendCommonEmail(ctx, args)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestSendTemplateEmail(t *testing.T) {
	client, err := sendcloud.NewSendCloud("*", "*")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	args := &sendcloud.TemplateMail{
		Receiver: sendcloud.MailReceiver{
			To: "a@ifaxin.com;b@ifaxin.com",
		},
		Body: sendcloud.MailBody{
			From:     "SendCloud@SendCloud.com",
			Subject:  "Email from SendCloud SDK",
			FromName: "SendCloud",
		},
		Content: sendcloud.TemplateContent{
			TemplateInvokeName: "test_template_active",
		},
	}
	result, err := client.SendTemplateEmail(ctx, args)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestSendTemplateEmailWithVars(t *testing.T) {
	client, err := sendcloud.NewSendCloud("*", "*")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	args := &sendcloud.TemplateMail{
		Body: sendcloud.MailBody{
			From:     "SendCloud@SendCloud.com",
			Subject:  "Email from SendCloud SDK",
			FromName: "SendCloud",
		},
		Content: sendcloud.TemplateContent{
			TemplateInvokeName: "test_template_active",
		},
	}
	xsmtpapi := sendcloud.XSMTPAPI{
		To: []string{"a@ifaxin.com", "b@ifaxin.com"},
		Sub: map[string][]interface{}{
			"%name%":  {"jack", "rose"},
			"%money%": {"199", "299"},
		},
		Filters: &sendcloud.Filter{
			SubscriptionTracking: sendcloud.TrackingFilter{Settings: sendcloud.FilterSettings{Enable: "1"}},
			OpenTracking:         sendcloud.TrackingFilter{Settings: sendcloud.FilterSettings{Enable: "1"}},
			ClickTracking:        sendcloud.TrackingFilter{Settings: sendcloud.FilterSettings{Enable: "1"}},
		},
		Settings: &sendcloud.Settings{
			Unsubscribe: sendcloud.UnsubscribeSettings{PageID: []int{1, 2}},
		},
	}
	args.Body.SetXsmtpapi(xsmtpapi)
	result, err := client.SendTemplateEmail(ctx, args)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestSendTemplateEmailWithAttachment(t *testing.T) {
	client, err := sendcloud.NewSendCloud("*", "*")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	attachment1, err := os.Open("path/to/attachment1.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer attachment1.Close()
	args := &sendcloud.TemplateMail{
		Receiver: sendcloud.MailReceiver{
			To: "a@ifaxin.com;b@ifaxin.com",
		},
		Body: sendcloud.MailBody{
			From:     "SendCloud@SendCloud.com",
			Subject:  "Email from SendCloud SDK",
			FromName: "SendCloud",
		},
		Content: sendcloud.TemplateContent{
			TemplateInvokeName: "test_template_active",
		},
	}
	args.Body.AddAttachment(attachment1)
	result, err := client.SendTemplateEmail(ctx, args)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestSendEmailCalendar(t *testing.T) {
	client, err := sendcloud.NewSendCloud("*", "*")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	args := &sendcloud.CalendarMail{
		Receiver: sendcloud.MailReceiver{
			To: "a@ifaxin.com;b@ifaxin.com",
		},
		Body: sendcloud.MailBody{
			From:     "SendCloud@SendCloud.com",
			Subject:  "Email from SendCloud SDK",
			FromName: "SendCloud",
		},
		Content: sendcloud.TextContent{
			Html: "<p>This is an HTML email.</p>",
		},
	}
	result, err := client.SendCalendarMail(ctx, args)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

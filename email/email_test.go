package sendcloud

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestSendCommonEmail(t *testing.T) {
	client, err := NewSendCloud("*", "*")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	args :=  &CommonMail{
		Receiver: MailReceiver{
			To: "a@ifaxin.com;b@ifaxin.com",
		},
		Body: MailBody{
			From: "SendCloud@SendCloud.com",
			Subject: "Email from SendCloud SDK",
			FromName: "SendCloud",
		},
		Content: TextContent{
			Html:  "<p>This is an HTML email.</p>",
		},
	}
	result, err := client.SendCommonEmail(ctx, args)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestSendCommonEmailWithVars(t *testing.T) {
	client, err := NewSendCloud("*", "*")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	args :=  &CommonMail{
		Body: MailBody{
			From: "SendCloud@SendCloud.com",
			Subject: "Email from SendCloud SDK",
			FromName: "SendCloud",
		},
		Content: TextContent{
			Html:  "<p>This is an HTML email.</p>",
		},
	}
	xsmtpapi := XSMTPAPI{
		To: []string{"a@ifaxin.com","b@ifaxin.com"},
		Sub: map[string][]interface{}{
			"%name%": {"jack", "rose"},
			"%money%": {"199", "299"},
		},
		Filters: &Filter{
			SubscriptionTracking: TrackingFilter{Settings: FilterSettings{Enable: "1"}},
			OpenTracking: TrackingFilter{Settings: FilterSettings{Enable: "1"}},
			ClickTracking: TrackingFilter{Settings: FilterSettings{Enable: "1"}},
		},
		Settings: &Settings{
			Unsubscribe: UnsubscribeSettings{PageID: []int{1, 2}},
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
	client, err := NewSendCloud("*", "*")
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
	args :=  &CommonMail{
		Receiver: MailReceiver{
			To: "a@ifaxin.com;b@ifaxin.com",
		},
		Body: MailBody{
			From: "SendCloud@SendCloud.com",
			Subject: "Email from SendCloud SDK",
			FromName: "SendCloud",
		},
		Content: TextContent{
			Html:  "<p>This is an HTML email.</p>",
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
	client, err := NewSendCloud("*", "*")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	args :=  &TemplateMail{
		Receiver: MailReceiver{
			To: "a@ifaxin.com;b@ifaxin.com",
		},
		Body: MailBody{
			From: "SendCloud@SendCloud.com",
			Subject: "Email from SendCloud SDK",
			FromName: "SendCloud",
		},
		Content: TemplateContent{
			TemplateInvokeName:  "test_template_active",
		},
	}
	result, err := client.SendTemplateEmail(ctx, args)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestSendTemplateEmailWithVars(t *testing.T) {
	client, err := NewSendCloud("*", "*")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	args :=  &TemplateMail{
		Body: MailBody{
			From: "SendCloud@SendCloud.com",
			Subject: "Email from SendCloud SDK",
			FromName: "SendCloud",
		},
		Content: TemplateContent{
			TemplateInvokeName:  "test_template_active",
		},
	}
	xsmtpapi := XSMTPAPI{
		To: []string{"a@ifaxin.com","b@ifaxin.com"},
		Sub: map[string][]interface{}{
			"%name%": {"jack", "rose"},
			"%money%": {"199", "299"},
		},
		Filters: &Filter{
			SubscriptionTracking: TrackingFilter{Settings: FilterSettings{Enable: "1"}},
			OpenTracking: TrackingFilter{Settings: FilterSettings{Enable: "1"}},
			ClickTracking: TrackingFilter{Settings: FilterSettings{Enable: "1"}},
		},
		Settings: &Settings{
			Unsubscribe: UnsubscribeSettings{PageID: []int{1, 2}},
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
	client, err := NewSendCloud("*", "*")
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
	args :=  &TemplateMail{
		Receiver: MailReceiver{
			To: "a@ifaxin.com;b@ifaxin.com",
		},
		Body: MailBody{
			From: "SendCloud@SendCloud.com",
			Subject: "Email from SendCloud SDK",
			FromName: "SendCloud",
		},
		Content: TemplateContent{
			TemplateInvokeName:  "test_template_active",
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
	client, err := NewSendCloud("*", "*")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	args :=  &CalendarMail{
		Receiver: MailReceiver{
			To: "a@ifaxin.com;b@ifaxin.com",
		},
		Body: MailBody{
			From: "SendCloud@SendCloud.com",
			Subject: "Email from SendCloud SDK",
			FromName: "SendCloud",
		},
		Content: TextContent{
			Html:  "<p>This is an HTML email.</p>",
		},
	}
	result, err := client.SendCalendarMail(ctx, args)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
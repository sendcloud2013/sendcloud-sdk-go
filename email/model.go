package sendcloud

import (
	"net/http"
	"os"
	"reflect"
	"time"
)

const (
	APIBase = "https://api.sendcloud.net/apiv2/mail"
	sendCommonPath   = "/send"
	sendTemplatePath = "/sendtemplate"
	sendCalendarPath = "/sendcalendar"
)

type SendCloud struct {
	apiUser string
	apiKey  string
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

type CommonMail struct {
	Receiver MailReceiver
	Body     MailBody
	Content  TextContent
}

type TemplateMail struct {
	Receiver MailReceiver
	Body     MailBody
	Content  TemplateContent
}

type CalendarMail struct {
	Receiver MailReceiver
	Body     MailBody
	Content  TextContent
	Calendar MailCalendar
}


type MailReceiver struct {
	To                  string
	CC                  string
	BCC                 string
	UseAddressList      bool
}

type MailBody struct {
	From                string
	Subject             string
	ContentSummary      string
	FromName            string
	ReplyTo             string
	LabelName           string
	Headers             map[string]string
	Attachments         []*os.File
	Xsmtpapi           XSMTPAPI
	SendRequestID       string
	RespEmailID         bool
	UseNotification     bool
}


type TextContent struct {
	Html              string
	Plain             string
}

type TemplateContent struct {
	TemplateInvokeName string
}


// SetHTML - Set the html content of the email, required if not using a template.
func (e *TextContent) SetHTML(html string) {
	e.Html = html
}

// SetPlain - Set the plain content of the email, required if not using a template.
func (e *TextContent) SetPlain(plain string) {
	e.Plain = plain
}

// SetTemplateInvokeName - Set the template invoke name.
func (e *TemplateContent) SetTemplateInvokeName(name string) {
	e.TemplateInvokeName = name
}


type XSMTPAPI struct {
	To        []string         `json:"to,omitempty"`
	Sub       map[string][]interface{}  `json:"sub,omitempty"`
	Pubsub    map[string]interface{}    `json:"pubsub,omitempty"`
	Filters      *Filter    `json:"filters,omitempty"`
	Settings      *Settings      `json:"settings,omitempty"`
}

func (x XSMTPAPI) IsEmpty() bool {
	return reflect.DeepEqual(x, XSMTPAPI{})
}

type FilterSettings struct {
	Enable string `json:"enable"`
}

type TrackingFilter struct {
	Settings FilterSettings `json:"settings"`
}

type Filter struct {
	SubscriptionTracking TrackingFilter `json:"subscription_tracking"`
	OpenTracking         TrackingFilter `json:"open_tracking"`
	ClickTracking        TrackingFilter `json:"click_tracking"`
}


// UnsubscribeSettings 表示退订设置的结构体
type UnsubscribeSettings struct {
	PageID []int `json:"page_id"`
}

// Settings 表示设置的结构体
type Settings struct {
	Unsubscribe UnsubscribeSettings `json:"unsubscribe"`
}

// SetFrom - Set the from address.
func (e *MailBody) SetFrom(from string) {
	e.From = from
}

// SetContentSummary - Set the content summary of the email.
func (e *MailBody) SetContentSummary(contentSummary string) {
	e.ContentSummary = contentSummary
}

// SetFromName - Set the from name of the email.
func (e *MailBody) SetFromName(fromName string) {
	e.FromName = fromName
}


// SetReplyTo - Set the reply to address.
func (e *MailBody) SetReplyTo(replyTo string) {
	e.ReplyTo = replyTo
}


// SetLabelName - Set the label name of the email.
func (e *MailBody) SetLabelName(labelName string) {
	e.LabelName = labelName
}

// AddHeaders - Add the headers of the email.
func (e *MailBody) AddHeaders(headers map[string]string) {
	e.Headers = headers
}

// AddAttachment - Add an attachment content.
func (e *MailBody) AddAttachment(attachment *os.File) {
	e.Attachments = append(e.Attachments, attachment)
}

// SetXsmtpapi - Set the xsmtpapi of the email.
func (e *MailBody) SetXsmtpapi(xsmtpapi XSMTPAPI) {
	e.Xsmtpapi = xsmtpapi
}

type MailCalendar struct {
	StartTime         time.Time
	EndTime           time.Time
	Title             string
	OrganizerName     string
	OrganizerEmail    string
	Location          string
	Description       string
	ParticipatorNames string
	ParticipatorEmails string
	UID               string
	IsCancel          bool
	IsUpdate          bool
	ValarmTime 		  int
}

//SetStartTime - Set the start time of the calendar.
func (e *MailCalendar) SetStartTime(startTime time.Time) {
    e.StartTime = startTime
}

//SetEndTime - Set the end time of the calendar.
func (e *MailCalendar) SetEndTime(endTime time.Time) {
    e.EndTime = endTime
}

//SetTitle - Set the title of the calendar.
func (e *MailCalendar) SetTitle(title string) {
    e.Title = title
}

//SetOrganizerName - Set the organizer name of the calendar.
func (e *MailCalendar) SetOrganizerName(organizerName string) {
    e.OrganizerName = organizerName
}

//SetOrganizerEmail - Set the organizer email of the calendar.
func (e *MailCalendar) SetOrganizerEmail(organizerEmail string) {
    e.OrganizerEmail = organizerEmail
}

// SetLocation - Set the location of the calendar.
func (e *MailCalendar) SetLocation(location string) {
    e.Location = location
}

// SetDescription - Set the description of the calendar.
func (e *MailCalendar) SetDescription(description string) {
    e.Description = description
}

// SetParticipatorNames - Set the participator names of the calendar.
func (e *MailCalendar) SetParticipatorNames(participatorNames string) {
    e.ParticipatorNames = participatorNames
}

// SetParticipatorEmails - Set the participator emails of the calendar.
func (e *MailCalendar) SetParticipatorEmails(participatorEmails string) {
    e.ParticipatorEmails = participatorEmails
}

// SetUID - Set the UID of the calendar.
func (e *MailCalendar) SetUID(uid string) {
	e.UID = uid
}

// SetIsCancel - Set the isCancel of the calendar.
func (e *MailCalendar) SetIsCancel(isCancel bool) {
    e.IsCancel = isCancel
}

// SetIsUpdate - Set the isUpdate of the calendar.
func (e *MailCalendar) SetIsUpdate(isUpdate bool) {
    e.IsUpdate = isUpdate
}

// SetValarmTime - Set the valarmTime of the calendar.
func (e *MailCalendar) SetValarmTime(valarmTime int) {
    e.ValarmTime = valarmTime
}

type SendEmailResult struct {
	Result     bool        `json:"result"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Info       interface{} `json:"info"`
}



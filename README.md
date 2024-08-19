# sendcloud-sdk-go

## Supported Go Versions  

This SDK supports Go 1.16 and above.

## Email SDK

The [SendCloud Email Go SDK provides a straightforward interface to SendCloud's email delivery service. With this SDK, you can effortlessly send regular emails using the `SendCommonEmail` method, as well as emails based on predefined templates through the `SendTemplateEmail` method.

Whether you're sending a one-off email to a customer or leveraging templates for consistent branding across your communication channels, the Email SDK makes it easy to integrate SendCloud's robust email delivery capabilities into your Go workflows.

### 1. Import the Package

First, you need to import the Go package that contains the `sendcloud`. Let's assume the package name is `sendcloud` (you would need to replace this with the actual package name):

```go
import (  
    "github.com/sendcloud2013/sendcloud-sdk-go/email"
)
```

### 2. Initialize the sendcloud

Next, you need to initialize the `sendcloud` using credentials provided by your SMS service provider, such as an API key or username/password. Assuming there's a `Newsendcloud` function that takes two string parameters (replaced with `**` placeholders):

```go
client, err := sendcloud.NewSendCloud("API_KEY", "API_SECRET")  
if err != nil {  
    // Handle the error, for example, by printing it or returning  
    log.Fatal(err)  
}
```

### 3. Prepare the Send Parameters

Create an instance of the CommonMail struct from the sendcloud package and set the required parameters for sending an email. This struct should include fields such as the recipient email addresses, sender information, subject, and the email content in HTML format. Here's how you can set up the parameters:

```go
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
```

### 4. Send the SMS Template

Now, you can call the `SendCommonEmail` method of the `sendcloud` to send the Email:

```go
result, err := client.SendCommonEmail(ctx, args)
if err != nil {  
    // Handle the error, for example, by printing it or returning  
    log.Fatal(err)  
}
```

### 5. Handle the Result

Finally, you can perform further actions based on the `result` (whose type and structure depend on the `sendcloud` package's definition), such as printing the result or passing it to other functions.

### 6. Usage example

#### SendCommonEmail  

The `SendCommonEmail` method allows you to send a regular email with a custom subject, body, and recipient list.  

#### Usage

```go  
package main  
  
import (
	"context"
	"fmt"
	"github.com/sendcloud2013/sendcloud-sdk-go/email"
	"log"
	"time"
)

func main() {
	client, err := sendcloud.NewSendCloud("*", "*")
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	args :=  &sendcloud.CommonMail{
		Receiver: sendcloud.MailReceiver{
			To: "a@ifaxin.com;b@ifaxin.com",
		},
		Body: sendcloud.MailBody{
			From: "SendCloud@SendCloud.com",
			Subject: "Email from SendCloud SDK",
			FromName: "SendCloud",
		},
		Content: sendcloud.TextContent{
			Html:  "<p>This is an HTML email.</p>",
		},
	}
	result, err := client.SendCommonEmail(ctx, args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
```

#### SendTemplateEmail

The `SendTemplateEmail` method allows you to send an email using a predefined template. This is useful when you want to send emails with consistent design and layout.

#### Usage

```go
package main  
  
import (
	"context"
	"fmt"
	"github.com/sendcloud2013/sendcloud-sdk-go/email"
	"log"
	"time"
)

func main() {
	client, err := sendcloud.NewSendCloud("*", "*")
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	args :=  &sendcloud.TemplateMail{
		Receiver: sendcloud.MailReceiver{
			To: "a@ifaxin.com;b@ifaxin.com",
		},
		Body: sendcloud.MailBody{
			From: "SendCloud@SendCloud.com",
			Subject: "Email from SendCloud SDK",
			FromName: "SendCloud",
		},
		Content: sendcloud.TemplateContent{
			TemplateInvokeName:  "test_template_active",
		},
	}
	result, err := client.SendTemplateEmail(ctx, args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
```

## 

## SMS SDK

The [SendCloud SMS Go SDK](https://github.com/sendcloud2013/sendcloud-sdk-go/blob/main/sms/README.md) offers a similarly convenient way to integrate SendCloud's SMS service into your Go applications. This SDK simplifies the process of sending SMS messages to your users, allowing you to easily incorporate text-based communication into your customer engagement strategies.

With features designed to support both individual and bulk messaging, the SMS SDK enables you to quickly and efficiently send important notifications, reminders, or promotional messages to your audience.

### 1. Import the Package

First, you need to import the Go package that contains the `sendcloud`. Let's assume the package name is `sendcloud` (you would need to replace this with the actual package name):

```go
import (  
    "github.com/sendcloud2013/sendcloud-sdk-go/sms"
)
```

### 2. Initialize the sendcloud

Next, you need to initialize the `sendcloud` using credentials provided by your SMS service provider, such as an API key or username/password. Assuming there's a `Newsendcloud` function that takes two string parameters (replaced with `**` placeholders):

```go
client, err := sendcloud.NewSendCloudSms("SMS_USER", "SMS_KEY")  
if err != nil {  
    // Handle the error, for example, by printing it or returning  
    log.Fatal(err)  
}
```

### 3. Prepare the Send Parameters

Create an instance of the `SendSmsTemplateArgs` struct and set the required parameters. This struct should be defined by the `sendcloud` package and include fields like template ID, label ID, recipient phone numbers, and message type:

```go
args := &sendcloud.SendSmsTemplateArgs{  
    TemplateId: 1,           // Replace with the actual template ID  
    LabelId:    1,           // Replace with the actual label ID (if applicable)  
    Phone:      "13800138000", // Can be a single number or a comma-separated list of numbers  
    MsgType:    sendcloud.SMS,  // Assuming the sendcloud package defines an SMS constant  
}
```

### 4. Send the SMS Template

Now, you can call the `SendSmsTemplate` method of the `sendcloud` to send the SMS:

```go
result, err := client.SendSmsTemplate(args)  
if err != nil {  
    // Handle the error, for example, by printing it or returning  
    log.Fatal(err)  
}
```

### 5. Handle the Result

Finally, you can perform further actions based on the `result` (whose type and structure depend on the `sendcloud` package's definition), such as printing the result or passing it to other functions.

### 6. Usage example

Combining the steps above, here's a complete example code:

```go
package main  
  
import (  
    "fmt"  
    "log"
    "github.com/sendcloud2013/sendcloud-sdk-go/sms"  
)  
  
func main() {  
    client, err := sendcloud.NewSendCloudSms("API_KEY", "API_SECRET")  
    if err != nil {  
        log.Fatal(err)  
    }  
  
    args := &sendcloud.SendSmsTemplateArgs{  
        TemplateId: 1,  
        LabelId:    1,  
        Phone:      "13800138000",  
        MsgType:    sendcloud.SMS,  
    }  
  
    result, err := client.SendSmsTemplate(args)  
    if err != nil {  
        log.Fatal(err)  
    }  
  
    // Handle or print the result  
    fmt.Println(result)  
}
```

Please note that you need to replace the placeholders (like `API_KEY`, `API_SECRET`, and `sendcloud`) with actual credentials and package names. 

## Handling Errors

Always make sure to handle errors returned by the methods. They may indicate issues such as invalid credentials, API errors, or other problems that need to be addressed.

## Why SendCloud?

SendCloud is a leading provider of email and SMS delivery services, trusted by businesses of all sizes to reliably reach their customers' inboxes and mobile devices. By leveraging SendCloud's SDKs, you can focus on building your application's core functionality while enjoying the benefits of a robust, scalable, and secure communication platform.

## Support

If you have any questions or encounter any issues while using the SendCloud Go SDKs, please don't hesitate to reach out to SendCloud's support team for assistance. Additionally, you can find helpful resources and engage with the community on SendCloud's official documentation, forums, and social media channels.

Thank you for choosing SendCloud as your communication platform, and we look forward to helping you succeed with our Go SDKs!
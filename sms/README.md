# sendcloud-sms-go

### 1. Import the Package

First, you need to import the Go package that contains the `sendcloud`. Let's assume the package name is `sendcloud` (you would need to replace this with the actual package name):

```go
import (  
    "github.com/frank-email-sms/sendcloud-sms-go"
)
```

### 2. Initialize the sendcloud

Next, you need to initialize the `sendcloud` using credentials provided by your SMS service provider, such as an API key or username/password. Assuming there's a `Newsendcloud` function that takes two string parameters (replaced with `**` placeholders):

```go
client, err := sendcloud.NewSendCloudSms("API_KEY", "API_SECRET")  
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

### Complete Example

Combining the steps above, here's a complete example code:

```go
package main  
  
import (  
    "fmt"  
    "log"
    "github.com/frank-email-sms/sendcloud-sms-go"  
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

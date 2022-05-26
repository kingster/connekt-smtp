package connekt

import (
	"testing"
)

func TestEmail(t *testing.T) {

	apiEndpoint = "http://127.0.0.1/v3/send/email/"
	apiKey := "your-api-key"

	rq := CreateEmailRequest()
	rq.ChannelInfo.To = append(rq.ChannelInfo.To, EmailAddress{
		Name:    "John Doe",
		Address: "connekt@labworld.org",
	})

	rq.ChannelInfo.From = EmailAddress{
		Name:    "John Doe",
		Address: "hello@labworld.org",
	}
	rq.ChannelData.Text = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec non sollicitudin mauris. Nullam pharetra ligula ut ante vestibulum, quis scelerisque est consequat. Fusce pharetra tempor metus in placerat. Suspendisse a maximus purus, vulputate viverra sem. Aliquam placerat orci quis enim iaculis tincidunt a at nisi. Fusce facilisis gravida enim non malesuada. Praesent ac lectus bibendum tortor suscipit feugiat in id nulla. Nunc in diam at ante consectetur aliquet in sed nunc. Nullam lacinia nec nisl sed ultrices.	"
	rq.ChannelData.Subject = "Hello World"

	result, err := SendEmail(rq, "seller-vendor", apiKey)
	if err != nil {
		t.Fatalf("Send Email failed with error %v", err)
	}

	if result.Status/100 != 2 {
		t.Fatalf("Send Email failed with status %v", result.Status)
	}

	t.Logf("Send Email success %v", result)
}

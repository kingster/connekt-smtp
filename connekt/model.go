package connekt

import (
	"fmt"
	"strings"

	"github.com/emersion/go-message/mail"
)

//Attachment Defines a attachment for an email
type Attachment struct {
	Base64Data string `json:"base64Data"`
	Name       string `json:"name"`
	Mime       string `json:"mime"`
}

type EmailAddress struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type EmailRequest struct {
	SLA         string `json:"sla"`
	ChannelData struct {
		Type        string       `json:"type"`
		Subject     string       `json:"subject"`
		HTML        string       `json:"html"`
		Text        string       `json:"text"`
		Attachments []Attachment `json:"attachments"`
	} `json:"channelData"`
	ChannelInfo struct {
		Type    string         `json:"type"`
		AppName string         `json:"appName"`
		To      []EmailAddress `json:"to"`
		CC      []EmailAddress `json:"cc"`
		From    EmailAddress   `json:"from"`
	} `json:"channelInfo"`
}

type Response struct {
	Status   int         `json:"status"`
	Request  interface{} `json:"request"`
	Response struct {
		Type    string                 `json:"type"`
		Message string                 `json:"message"`
		Success map[string]interface{} `json:"success"`
		Failure interface{}            `json:"failure"`
	} `json:"response"`
}

type ErrorResponse struct {
	Status   int         `json:"status"`
	Request  interface{} `json:"request"`
	Response struct {
		Type    string      `json:"type"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	} `json:"response"`
}

func CreateEmailRequest() EmailRequest {
	rq := EmailRequest{}
	rq.ChannelInfo.Type = "EMAIL"
	rq.ChannelData.Type = "EMAIL"
	rq.SLA = "H"
	return rq
}

func SMTPEmailAddress(ad *mail.Address) EmailAddress {
	return EmailAddress{
		Name:    ad.Name,
		Address: ad.Address,
	}
}

const HTMLTemplate = `<html>
<head>
</head>
<body>
	<p>%s</p>
</body>
</html>`

func Text2Html(s string) string {
	return fmt.Sprintf(HTMLTemplate, strings.ReplaceAll(s, "\n", "<br>"))
}

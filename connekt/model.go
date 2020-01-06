package connekt

import (
	"github.com/emersion/go-message/mail"
)

type ConnektAttachment struct {
	Base64Data string `json:"base64Data"`
	Name       string `json:"name"`
	Mime       string `json:"mime"`
}

type ConnektEmailAddress struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type ConnektEmailRequest struct {
	SLA string `json:"sla"`
	ChannelData struct {
		Type        string              `json:"type"`
		Subject     string              `json:"subject"`
		HTML        string              `json:"html"`
		Text        string              `json:"text"`
		Attachments []ConnektAttachment `json:"attachments"`
	} `json:"channelData"`
	ChannelInfo struct {
		Type string                `json:"type"`
		To   []ConnektEmailAddress `json:"to"`
		CC   []ConnektEmailAddress `json:"cc"`
		From ConnektEmailAddress   `json:"from"`
	} `json:"channelInfo"`
}

type ConnektResponse struct {
	Status   int         `json:"status"`
	Request  interface{} `json:"request"`
	Response struct {
		Type    string `json:"type"`
		Message string `json:"message"`
		Success map[string]interface{} `json:"success"`
		Failure []interface{} `json:"failure"`
	} `json:"response"`
}

func CreateEmailRequest() ConnektEmailRequest {
	rq := ConnektEmailRequest{}
	rq.ChannelInfo.Type = "EMAIL"
	rq.ChannelData.Type = "EMAIL"
	rq.SLA = "H"
	return rq
}

func SMTPEmailAddress(ad *mail.Address) ConnektEmailAddress  {
	return ConnektEmailAddress{
		Name: ad.Name,
		Address:ad.Address,
	}
}


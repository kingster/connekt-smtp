package smtp

import (
	"encoding/base64"
	"github.com/kingster/connekt-smtp/connekt"
)

func ConnektAttachment(a Attachment) connekt.ConnektAttachment {
	return connekt.ConnektAttachment{
		Base64Data: base64.StdEncoding.EncodeToString(a.Data),
		Name:       a.FileName,
		Mime:       a.ContentType,
	}
}

func ConnektEmailRequest(s *Session) connekt.ConnektEmailRequest {
	rq := connekt.CreateEmailRequest()
	for _, addr := range s.To {
		rq.ChannelInfo.To = append(rq.ChannelInfo.To, connekt.SMTPEmailAddress(addr))
	}
	for _, addr := range s.CC {
		rq.ChannelInfo.CC = append(rq.ChannelInfo.CC, connekt.SMTPEmailAddress(addr))
	}
	rq.ChannelInfo.From = connekt.SMTPEmailAddress(s.From)
	rq.ChannelData.Text = s.Text
	rq.ChannelData.Subject = s.Subject
	if len(s.HTML) > 0 {
		rq.ChannelData.HTML = s.HTML
	} else {
		rq.ChannelData.HTML = connekt.Text2Html(s.Text)
	}

	for _, attachment := range s.Attachments {
		rq.ChannelData.Attachments = append(rq.ChannelData.Attachments, ConnektAttachment(attachment))
	}

	return rq
}

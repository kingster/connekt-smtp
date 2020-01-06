package smtp

import (
	"bytes"
	"github.com/emersion/go-message/mail"
	"github.com/emersion/go-smtp"
	"github.com/kingster/connekt-smtp/connekt"
	"io"
	"io/ioutil"
	"log"
	"time"
)

type Attachment struct {
	ContentType string
	FileName    string
	Data        []byte
}

// A SMTPSession is returned after successful login.
type Session struct {
	From        *mail.Address
	To          []*mail.Address
	CC          []*mail.Address
	Text        string
	HTML        string
	Subject     string
	Date        time.Time
	Attachments []Attachment
}

func (s *Session) Mail(from string, opts smtp.MailOptions) error {
	// log.Println("Mail from:", from)
	s.From = &mail.Address{Address: from}
	return nil
}

func (s *Session) Rcpt(to string) error {
	s.To = append(s.To, &mail.Address{Address: to})
	return nil
}

func (s *Session) Dump() {
	log.Println("------------ Dump Email ---------")

	log.Println("Date:", s.Date)
	log.Println("From:", s.From)
	log.Println("To:", s.To)
	log.Println("CC:", s.CC)

	log.Println("Subject:", s.Subject)

	log.Println("Text", s.Text)
	log.Println("HTML", s.HTML)

	for _, at := range s.Attachments {
		log.Println("Got attachment Type:", at.ContentType, "Name:", at.FileName, "Body:\n", string(at.Data))
	}
	log.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
}

func (s *Session) Data(r io.Reader) error {
	if b, err := ioutil.ReadAll(r); err != nil {
		return err
	} else {
		mr, _ := mail.CreateReader(bytes.NewReader(b))
		header := mr.Header
		if date, err := header.Date(); err == nil {
			s.Date = date
		}
		if from, err := header.AddressList("From"); err == nil {
			s.From = from[0]
		}
		if to, err := header.AddressList("To"); err == nil {
			s.To = to
		}
		if cc, err := header.AddressList("CC"); err == nil {
			s.CC = cc
		}
		if subject, err := header.Subject(); err == nil {
			s.Subject = subject
		}
		// Process each message's part
		for {
			p, err := mr.NextPart()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}

			switch h := p.Header.(type) {
			case *mail.InlineHeader:
				// This is the message's text (can be plain-text or HTML)
				b, _ := ioutil.ReadAll(p.Body)
				t, _, _ := h.ContentType()
				switch t {
				case "text/plain":
					s.Text = string(b)
					break
				case "text/html":
					s.HTML = string(b)
					break
				default:
					log.Println("Unsupported body content-type", t)
					break
				}
			case *mail.AttachmentHeader:
				// This is an attachment
				filename, _ := h.Filename()
				t, _, _ := h.ContentType()
				b, _ := ioutil.ReadAll(p.Body)

				s.Attachments = append(s.Attachments, Attachment{
					ContentType: t,
					FileName:    filename,
					Data:        b,
				})
			}
		}

		//s.Dump() //debug
		connekt.SendEmail(ConnektEmailRequest(s))

	}
	return nil
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}

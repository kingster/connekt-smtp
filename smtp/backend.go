package smtp

import (
	"errors"
	"github.com/emersion/go-smtp"
	"log"
)

// The Backend implements SMTP server methods.
type Backend struct{}

// Login handles a login command with username and password.
func (bkd *Backend) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
	if username != "smtp" {
		return nil, errors.New("Invalid username or password")
	}
	log.Println("Connection from", state.RemoteAddr.String(), state.Hostname)
	return &Session{APIKey: password}, nil
}

// AnonymousLogin requires clients to authenticate using SMTP AUTH before sending emails
func (bkd *Backend) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	return nil, smtp.ErrAuthRequired
}

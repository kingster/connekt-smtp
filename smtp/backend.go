package smtp

import (
	"github.com/emersion/go-smtp"
	"log"
	"os"
)

// The Backend implements SMTP server methods.
type Backend struct{}

var defaultConnektApp = os.Getenv("DEFAULT_APP")

// Login handles a login command with username and password.
func (bkd *Backend) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
	if username == "smtp" {
		username = defaultConnektApp
	}
	log.Println("Connection from", state.RemoteAddr.String(), state.Hostname, "AppName: "+username)
	return &Session{APIKey: password, AppName: username}, nil
}

// AnonymousLogin requires clients to authenticate using SMTP AUTH before sending emails
func (bkd *Backend) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	return nil, smtp.ErrAuthRequired
}

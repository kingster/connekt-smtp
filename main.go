package main

import (
	"crypto/tls"
	"log"
	"os"
	"time"

	"github.com/emersion/go-smtp"
	csmtp "github.com/kingster/connekt-smtp/smtp"
)

var Environment = "Development"

func main() {
	be := &csmtp.Backend{}
	s := smtp.NewServer(be)

	bindAddr := os.Getenv("ADDR")
	if bindAddr == "" {
		bindAddr = "0.0.0.0:25"
	}

	s.Addr = bindAddr
	s.Domain = "connekt.flipkart.net"
	s.ReadTimeout = 10 * time.Second
	s.WriteTimeout = 10 * time.Second
	s.MaxMessageBytes = 10 * 1024 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true

	// Load the certificate and key
	certFile := "resources/cert.pem"
	keyFile := "resources/key.pem"

	if Environment == "Production" {
		certFile = "/var/lib/connekt-smtp/cert.pem"
		keyFile = "/var/lib/connekt-smtp/key.pem"
	}
	cer, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
		return
	}
	// Configure the TLS support
	s.TLSConfig = &tls.Config{Certificates: []tls.Certificate{cer}}

	log.SetFlags(log.LstdFlags)
	log.Println("Starting server at", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}

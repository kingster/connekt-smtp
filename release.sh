#!/bin/bash -x

VERSION=$(date +%s)
mkdir -p deb/usr/sbin
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o csmtp main.go
cp csmtp deb/usr/sbin/csmtp
fpm --force\
  --input-type dir\
  --output-type deb\
  --version "1.$VERSION"\
  --name connekt-smtp\
  --architecture amd64\
  --prefix /\
  --description 'An SMTP interface for Connekt'\
  --url "https://github.com/kingster/connekt-smtp"\
  --no-deb-systemd-restart-after-upgrade\
  --chdir deb
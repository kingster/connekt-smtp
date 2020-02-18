#!/bin/bash -x

VERSION=$(date +%s)
mkdir -p deb/usr/sbin
GOOS=linux GOARCH=386 go build -o csmtp main.go
cp csmtp deb/usr/sbin/csmtp
fpm --force \
  --input-type dir\
  --output-type deb\
  --config-files /etc/default/csmtp \
  --version "1.$VERSION"\
  --name connekt-smtp\
  --architecture amd64\
  --depends "systemd (>= 240-1~)" \
  --prefix /\
  --description 'An SMTP interface for Connekt'\
  --url "https://github.com/kingster/connekt-smtp"\
  --no-deb-systemd-restart-after-upgrade\
  --chdir deb \
  --package connekt-smtp.deb

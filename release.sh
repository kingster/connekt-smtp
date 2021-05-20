#!/bin/bash -ex

VERSION=$(date +%s)
export GOOS=linux 
export GOARCH=386 
go build -ldflags "-X main.Environment=Production" -o csmtp main.go

PWD=`pwd`

fpm --force --verbose \
  --input-type dir --output-type deb\
  --config-files /etc/default/csmtp \
  --version "1.$VERSION"\
  --name connekt-smtp --architecture amd64 \
  --package connekt-smtp.deb \
  --depends "systemd (>= 240-1~)" --depends "logrotate" \
  --prefix / \
  --description 'An SMTP interface for Connekt'\
  --url "https://github.com/kingster/connekt-smtp"\
  --maintainer "Kinshuk <hi@kinsh.uk>" \
  --after-install $PWD/package/DEBIAN/postinst \
  --exclude ".DS_Store" \
  ./package/deb/=/  \
  ./csmtp=/usr/sbin/csmtp ./resources/=/var/lib/connekt-smtp/

dpkg -c connekt-smtp.deb

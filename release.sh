#!/bin/bash -ex

VERSION=$(date +%s)
GOOS=linux GOARCH=386 go build -o csmtp main.go


mkdir -p package/deb/usr/sbin
cp csmtp package/deb/usr/sbin/csmtp

PWD=`pwd`

fpm --force --verbose \
  --input-type dir\
  --output-type deb\
  --config-files /etc/default/csmtp \
  --version "1.$VERSION"\
  --name connekt-smtp\
  --architecture amd64\
  --depends "systemd (>= 240-1~)" \
  --prefix / \
  --description 'An SMTP interface for Connekt'\
  --url "https://github.com/kingster/connekt-smtp"\
  --maintainer "Kinshuk <hi@kinsh.uk>" \
  --chdir package/deb \
  --no-deb-systemd-restart-after-upgrade\
  --after-install $PWD/package/DEBIAN/postinst \
  --exclude ".DS_Store" \
  --package connekt-smtp.deb .

dpkg -c connekt-smtp.deb

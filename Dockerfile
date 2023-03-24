FROM golang:1.20-bullseye as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

RUN go build -ldflags "-X main.Environment=Production" -o csmtp main.go


FROM debian:bullseye-slim
LABEL org.opencontainers.image.authors="hi@kinsh.uk"

ENV TZ Asia/Kolkata
ENV API_ENDPOINT="http://127.0.0.1/v3/send/email/"
ENV DEFAULT_APP="appname"
ENV ADDR="0.0.0.0:8025"

RUN apt-get update && apt-get install -y apt-utils ca-certificates tzdata  && rm -rf /var/lib/apt/lists/*

RUN groupadd -g 1001 connekt-smtp && useradd -ms /bin/bash -d /home/connekt-smtp -g 1001 -u 1000 connekt-smtp
USER connekt-smtp

WORKDIR /

COPY --from=builder /app/csmtp  /csmtp
COPY resources /var/lib/connekt-smtp/

EXPOSE 8025

ENTRYPOINT [ "/csmtp" ]

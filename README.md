# connekt-smtp
SMTP Interface for Connekt

## Installation

This depends on `systemd (>= 240-1~)` so if you are on debian 9 you need to install it via backports.

```
apt install systemd -t stretch-backports
```

After that you can do a normal installation

```
# if installing directly via deb
apt install ./connekt-smtp.deb

# or via any repo
apt install connekt-smtp
```

## Configuration

You need to configure connekt endpoint and few other things

open `/etc/default/csmtp` in your favorite editor and make changes as per your need

```
API_ENDPOINT="http://127.0.0.1/v2/send/email/"
DEFAULT_APP="outbound" 
ADDR="0.0.0.0:25"
```


## Restarting Service

```
service connekt-smtp (stop|start|status)
```

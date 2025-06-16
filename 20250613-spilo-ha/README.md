# Generating SSL Certificate
Production systems shouldn't use unencrypted traffic. Follow
this tutorial to generate your own certificates.

## Create Root CA
You may skip this step if your organization already has their
root ca and request them to generate your certificate instead.

**WARNING!!**

You should put your ca.key file somewhere safe, not as plain
text on your git. This repo is only for demo purpose and do not
use the certificate for your production use. If this gets
leaked, someone else can generate their own cert and your
system will trust it.

```shell
cd config/certs/ca
openssl genrsa -out ca.key 2048
openssl req -x509 -subj "/CN=Iqbal Youtube" -nodes -key ca.key -days 3650 -out ca.crt
```

## Create etcd Server and Peer Certificates
Do this for each server. This works on homebrew's openssl version 3.5 on macos.
If you're using ip address on your advertised peer, you should use "IP:x.x.x.x" instead of DNS:etcd1 on subjectAltName.
Or add both DNS or IP just to be safe.
```shell
cd config/certs/etcd1
# replace etcd1 with your domain
openssl req -newkey rsa:2048 -nodes -subj "/CN=etcd1" -addext "subjectAltName = DNS:etcd1,DNS:localhost,IP:127.0.0.1" -keyout server.key -out server.csr
openssl x509 -req -copy_extensions copy -in server.csr -out server.crt -CAcreateserial -CA ../ca/ca.crt -CAkey ../ca/ca.key -days 3600
rm server.csr ../ca/ca.srl
```

## Create Postgres Certificates
Also do this for each server.
```shell
cd config/certs/postgres1
openssl req -newkey rsa:2048 -nodes -subj "/CN=postgres1" -addext "subjectAltName = DNS:postgres1,DNS:localhost,IP:127.0.0.1" -keyout server.key -out server.csr
openssl x509 -req -copy_extensions copy -in server.csr -out server.crt -CAcreateserial -CA ../ca/ca.crt -CAkey ../ca/ca.key -days 3600
rm server.csr ../ca/ca.srl
```
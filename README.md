
This repository consists of a http and grpc client and server setup. The purpose is to enable mutual tls between client and server. 

-------------------------

1.Create domains.ext file and add the config
``` 
authorityKeyIdentifier=keyid,issuer
basicConstraints=CA:FALSE
keyUsage = digitalSignature, nonRepudiation, keyEncipherment, dataEncipherment
subjectAltName = @alt_names
[alt_names]
DNS.1 = localhost
DNS.2 = test.local
```

Add the entry:
127.0.0.1 test.local to /etc/hosts

2.Generate root ca cert
```
openssl req -x509 -nodes -new -sha256 -days 1024 -newkey rsa:2048 -keyout RootCA.key -out RootCA.pem -subj "/C=US/CN=Example-Root-CA"
openssl x509 -outform pem -in RootCA.pem -out RootCA.crt
``` 

3.Generate certs and key for client and server:

In step 1: The rsa private key and a certificate signing request with CN and other details is generated
In step 2: The csr, root CA and details in domains.ext are used to generate a server certificate 

```
For the server:
openssl req -new -nodes -newkey rsa:2048 -keyout localhost.key -out localhost.csr -subj "/C=US/ST=exampleState/L=exampleCity/O=Example-Certificates/CN=localhost"
openssl x509 -req -sha256 -days 1024 -in localhost.csr -CA RootCA.pem -CAkey RootCA.key -CAcreateserial -extfile domains.ext -out localhost.crt
```


```
For the client:
openssl req -new -nodes -newkey rsa:2048 -keyout tls.key -out tls.csr -subj "/C=US/ST=exampleState/L=exampleCity/O=Example-Certificates/CN=test.local"
openssl x509 -req -sha256 -days 1024 -in tls.csr -CA RootCA.pem -CAkey RootCA.key -CAcreateserial -extfile domains.ext -out tls.crt
```

writeDomainInfo:
	echo 'authorityKeyIdentifier=keyid,issuer \
	basicConstraints=CA:FALSE \
	keyUsage = digitalSignature, nonRepudiation, keyEncipherment, dataEncipherment \
	subjectAltName = @alt_names \
	[alt_names] \
	DNS.1 = localhost \
	DNS.2 = test.local' >> domain_info.ext

# root ca certs
root-ca:
	openssl req -x509 -nodes -new -sha256 -days 1024 -newkey rsa:2048 -keyout RootCA.key -out RootCA.pem -subj "/C=US/CN=Example-Root-CA"
	openssl x509 -outform pem -in RootCA.pem -out RootCA.crt

# For the server:
server-cert:
	openssl req -new -nodes -newkey rsa:2048 -keyout localhost.key -out localhost.csr -subj "/C=US/ST=YourState/L=YourCity/O=Example-Certificates/CN=localhost"
	openssl x509 -req -sha256 -days 1 -in localhost.csr -CA RootCA.pem -CAkey RootCA.key -CAcreateserial -extfile domain_info.ext -out localhost.crt

# For the client:
client-cert:
	openssl req -new -nodes -newkey rsa:2048 -keyout tls.key -out tls.csr -subj "/C=US/ST=YourState/L=YourCity/O=Example-Certificates/CN=test.local"
	openssl x509 -req -sha256 -days 1024 -in tls.csr -CA RootCA.pem -CAkey RootCA.key -CAcreateserial -extfile domain_info.ext -out tls.crt

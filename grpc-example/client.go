package main

import (
	"crypto/tls"
	"crypto/x509"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc-ex/chat"
	"io/ioutil"
	"log"
)

var certs = "/certs/"
var address = "test.local"
var port = ":8443"

func main() {
	rootCAContent, _ := ioutil.ReadFile(certs + "RootCA.crt")
	if rootCAContent == nil {
	}

	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(rootCAContent) {
	}

	cert, err := tls.LoadX509KeyPair(certs+"localhost.crt", certs+"localhost.key")
	if err != nil {
		log.Fatal(err)
	}

	config := &tls.Config{
		//ServerName, whose value must match the common name on the certificate.
		ServerName:         address,
		InsecureSkipVerify: false,
		// RootCAs defines the set of root certificate authorities
		// that clients use when verifying server certificates
		RootCAs:      caCertPool,
		Certificates: []tls.Certificate{cert},
	}

	// dial to the server
	conn, err := grpc.Dial(address+port, grpc.WithTransportCredentials(credentials.NewTLS(config)))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)
	response, err := c.SayHello(context.Background(), &chat.Message{Subject: "Client hello", Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Client error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)
}

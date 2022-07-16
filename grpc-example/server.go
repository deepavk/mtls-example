package main

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc-ex/chat"
	"io/ioutil"
	"log"
	"net"
)

// running server as test.local

var certsPath = "/certs/"

func main() {
	certificatePair, err := tls.LoadX509KeyPair(certsPath+"tls.crt", certsPath+"tls.key")
	certPool := x509.NewCertPool()
	rootCACert, _ := ioutil.ReadFile(certsPath + "RootCA.crt")
	ok := certPool.AppendCertsFromPEM(rootCACert)

	if !ok {
		log.Fatalf("failed to serve: %s", err)
	}

	tlsConfig := &tls.Config{
		// ClientAuth determines the server's policy for
		// TLS Client Authentication.
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certificatePair},
		// ClientCAs defines the set of root certificate authorities
		// that servers use if required to verify a client certificate
		ClientCAs:  certPool,
		MinVersion: tls.VersionTLS12,
	}

	listener, err := net.Listen("tcp", "test.local:8443")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := chat.Server{}
	grpcServer := grpc.NewServer(grpc.Creds(credentials.NewTLS(tlsConfig)))
	chat.RegisterChatServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

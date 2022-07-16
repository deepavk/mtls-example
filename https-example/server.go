package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Write "Hello, world!" to the response body
	fmt.Println("Recieved request")
	io.WriteString(w, "Hello, world!\n")
}

func main() {
	// Set up a /hello resource handler
	http.HandleFunc("/hello", helloHandler)

	ca := " /certs/RootCA.crt"
	crt := "/certs/localhost.crt"
	key := "/certs/localhost.key"
	flag.Parse()

	config, err := createServerConfig(ca, crt, key)
	if err != nil {
		log.Fatal("config failed: %s", err.Error())
	}
	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: config,
	}
	log.Fatal(server.ListenAndServeTLS(crt, key))
}

func createServerConfig(ca, crt, key string) (*tls.Config, error) {
	caCertPEM, err := ioutil.ReadFile(ca)
	if err != nil {
		return nil, err
	}

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(caCertPEM)
	if !ok {
		panic("failed to parse root certificate")
	}

	cert, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		return nil, err
	}
	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    roots,
	}, nil
}

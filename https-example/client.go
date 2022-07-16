package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	rootCAContent, _ := ioutil.ReadFile("/certs/RootCA.crt")
	if rootCAContent == nil {
	}

	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(rootCAContent) {
	}

	cert, err := tls.LoadX509KeyPair("/certs/tls.crt",
		"/certs/tls.key")
	if err != nil {
		log.Fatal(err)
	}

	// Create a HTTPS client and supply the created CA pool
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
				RootCAs:            caCertPool,
				Certificates:       []tls.Certificate{cert},
			},
		},
	}
	// Request /hello over port 8080 via the GET method
	r, err := client.Get("https://localhost:8443/hello")
	if err != nil {
		log.Fatal(err)
	}

	// Read the response body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print the response body to stdout
	fmt.Printf("%s\n", body)
}

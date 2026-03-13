package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

// ===== CREATE CERTIFICATE AND KEY
// > openssl req -x509 -newkey rsa:2048 -nodes -keyout key.pem -out cert.pem -days 365

// req: Use the PKCS#10 certificate request and generation utility.
// -x509: Create a self-signed certificate instead of a request (CSR).
// -newkey rsa:2048: Generate a new RSA key that is 2048 bits strong.
// -nodes: Short for "no DES"; don't encrypt the private key with a password.
// -keyout key.pem: Save the generated private key to this file.
// -out cert.pem: Save the generated public certificate to this file.
// -days 365: Set the certificate to expire in one year.

// cert.pem | key.pem
// PEM (Privacy Enhanced Mail) - base64 encoded
// Same as .crt | .key

func main() {

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handling incoming orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Handling users")
	})

	port := 3000

	// Load the TLS certificate and key
	cert := "cert.pem"
	key := "key.pem"

	// Configure TLS
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// Create a custom server
	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		Handler:   nil,
		TLSConfig: tlsConfig,
	}

	// Enable http2
	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Server is running on port:", port)

	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Could not start the server:", err)
	}

	// HTTP 1.2 Server without TLS
	// err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	// if err != nil {
	// 	log.Fatalln("Could not start the server:", err)
	// }
}

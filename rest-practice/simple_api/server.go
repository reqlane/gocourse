package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

// ===== CREATE CERTIFICATE AND KEY
// openssl req -x509 -newkey rsa:2048 -nodes -keyout key.pem -out cert.pem -days 365

// ===== CREATE CERTIFICATE AND KEY WITH CONFIG (for Postman)
// openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem -config openssl.cnf

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

// curl -v -k https://localhost:3000/orders

// mTLS - load client's certificate
func loadClientCAs() *x509.CertPool {
	clientCAs := x509.NewCertPool()
	caCert, err := os.ReadFile("cert.pem")
	if err != nil {
		log.Fatalln("Could not load client CAs:", err)
	}
	clientCAs.AppendCertsFromPEM(caCert)
	return clientCAs
}

func main() {

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		logRequestDetails(r)
		fmt.Fprintf(w, "Handling incoming orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		logRequestDetails(r)
		fmt.Fprintf(w, "Handling users")
	})

	port := 3000

	// Load the TLS certificate and key
	cert := "cert.pem"
	key := "key.pem"

	// Configure TLS
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		// ClientAuth: tls.RequireAndVerifyClientCert, // Enforce mTLS (mutual TLS)
		// ClientCAs:  loadClientCAs(),                // Enforce mTLS (mutual TLS)
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

func logRequestDetails(r *http.Request) {
	httpVersion := r.Proto
	fmt.Println("Received request with HTTP version:", httpVersion)

	if r.TLS != nil {
		tlsVersion := getTLSVersionName(r.TLS.Version)
		fmt.Println("Received request with TLS version:", tlsVersion)
	} else {
		fmt.Println("Received request without TLS")
	}
}

func getTLSVersionName(version uint16) string {
	switch version {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return "Unknown TLS version"
	}
}

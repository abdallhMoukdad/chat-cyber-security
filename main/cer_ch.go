package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"time"
)

func main() {
	// Create a new CA key pair
	caPrivateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		fmt.Println("Error generating CA private key:", err)
		return
	}

	caTemplate := x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "MyCA"},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0), // Valid for 10 years
		KeyUsage:              x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	caDER, err := x509.CreateCertificate(rand.Reader, &caTemplate, &caTemplate, &caPrivateKey.PublicKey, caPrivateKey)
	if err != nil {
		fmt.Println("Error creating CA certificate:", err)
		return
	}

	// Save CA private key to file
	caPrivateKeyFile, err := os.Create("ca_private_key.pem")
	if err != nil {
		fmt.Println("Error creating CA private key file:", err)
		return
	}
	ecPrivateKeyBytes, err := x509.MarshalECPrivateKey(caPrivateKey)
	if err != nil {
		fmt.Println("Error marshaling CA private key:", err)
		return
	}
	pem.Encode(caPrivateKeyFile, &pem.Block{Type: "EC PRIVATE KEY", Bytes: ecPrivateKeyBytes})

	//pem.Encode(caPrivateKeyFile, &pem.Block{Type: "EC PRIVATE KEY", Bytes: x509.MarshalECPrivateKey(caPrivateKey)})
	caPrivateKeyFile.Close()

	// Save CA certificate to file
	caCertFile, err := os.Create("ca_certificate.pem")
	if err != nil {
		fmt.Println("Error creating CA certificate file:", err)
		return
	}
	pem.Encode(caCertFile, &pem.Block{Type: "CERTIFICATE", Bytes: caDER})
	caCertFile.Close()

	// Read CA private key from file
	caPrivateKeyFile, err = os.Open("ca_private_key.pem")
	if err != nil {
		fmt.Println("Error reading CA private key file:", err)
		return
	}
	pemData, err := ioutil.ReadAll(caPrivateKeyFile)
	if err != nil {
		fmt.Println("Error reading CA private key file:", err)
		return
	}
	block, _ := pem.Decode(pemData)
	caPrivateKeyFile.Close()
	caPrivateKey, err = x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("Error parsing CA private key:", err)
		return
	}

	// Read CA certificate from file
	caCertFile, err = os.Open("ca_certificate.pem")
	if err != nil {
		fmt.Println("Error reading CA certificate file:", err)
		return
	}
	pemData, err = ioutil.ReadAll(caCertFile)
	if err != nil {
		fmt.Println("Error reading CA certificate file:", err)
		return
	}
	block, _ = pem.Decode(pemData)
	caCertFile.Close()
	caCert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Println("Error parsing CA certificate:", err)
		return
	}

	// Create a new key pair for the client
	clientPrivateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		fmt.Println("Error generating client private key:", err)
		return
	}

	clientTemplate := x509.Certificate{
		SerialNumber:          big.NewInt(2),
		Subject:               pkix.Name{CommonName: "Client"},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0), // Valid for 1 year
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: false,
	}

	clientDER, err := x509.CreateCertificate(rand.Reader, &clientTemplate, caCert, &clientPrivateKey.PublicKey, caPrivateKey)
	if err != nil {
		fmt.Println("Error creating client certificate:", err)
		return
	}

	// Save client private key to file
	clientPrivateKeyFile, err := os.Create("client_private_key.pem")
	if err != nil {
		fmt.Println("Error creating client private key file:", err)
		return
	}
	clientPrivateKeyBytes, err := x509.MarshalECPrivateKey(caPrivateKey)
	if err != nil {
		fmt.Println("Error marshaling CA private key:", err)
		return
	}
	pem.Encode(clientPrivateKeyFile, &pem.Block{Type: "EC PRIVATE KEY", Bytes: clientPrivateKeyBytes})

	//pem.Encode(clientPrivateKeyFile, &pem.Block{Type: "EC PRIVATE KEY", Bytes: x509.MarshalECPrivateKey(clientPrivateKey)})
	clientPrivateKeyFile.Close()

	// Save client certificate to file
	clientCertFile, err := os.Create("client_certificate.pem")
	if err != nil {
		fmt.Println("Error creating client certificate file:", err)
		return
	}
	pem.Encode(clientCertFile, &pem.Block{Type: "CERTIFICATE", Bytes: clientDER})
	clientCertFile.Close()
}

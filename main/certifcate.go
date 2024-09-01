package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

func generateCertificateAuthority() (*x509.Certificate, *ecdsa.PrivateKey, error) {
	// Generate private key for CA
	caPrivateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	// Create a template for the CA certificate
	caTemplate := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"My CA"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0), // Valid for 10 years
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageKeyEncipherment,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	// Self-sign the CA certificate
	caDERBytes, err := x509.CreateCertificate(rand.Reader, &caTemplate, &caTemplate, &caPrivateKey.PublicKey, caPrivateKey)
	if err != nil {
		return nil, nil, err
	}

	caCert, err := x509.ParseCertificate(caDERBytes)
	if err != nil {
		return nil, nil, err
	}

	return caCert, caPrivateKey, nil
}

func savePEM(fileName string, data interface{}) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	switch v := data.(type) {
	case *x509.Certificate:
		err = pem.Encode(file, &pem.Block{Type: "CERTIFICATE", Bytes: v.Raw})
	case *ecdsa.PrivateKey:
		bytes, err := x509.MarshalECPrivateKey(v)
		if err != nil {
			return err
		}
		err = pem.Encode(file, &pem.Block{Type: "EC PRIVATE KEY", Bytes: bytes})
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}

	if err != nil {
		return err
	}

	return nil
}

func main() {
	caCert, caPrivateKey, err := generateCertificateAuthority()
	if err != nil {
		fmt.Println("Error generating CA:", err)
		return
	}

	// Save CA certificate and private key to files
	err = savePEM("ca.crt", caCert)
	if err != nil {
		fmt.Println("Error saving CA certificate:", err)
		return
	}

	err = savePEM("ca.key", caPrivateKey)
	if err != nil {
		fmt.Println("Error saving CA private key:", err)
		return
	}

	fmt.Println("CA certificate and private key generated successfully.")
}

// client.go

package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/binary"
	"encoding/pem"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to CA server:", err)
		return
	}
	defer conn.Close()

	// Receive CA certificate
	//caCertificate, err := receiveCertificate(conn)
	_, err = receiveCertificate(conn)

	if err != nil {
		fmt.Println("Error receiving CA certificate:", err)
		return
	}

	// Generate client key pair
	clientPrivateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		fmt.Println("Error generating client private key:", err)
		return
	}

	// Create client certificate signing request (CSR)
	csrDER, err := x509.CreateCertificateRequest(rand.Reader, &x509.CertificateRequest{
		Subject:            pkix.Name{CommonName: "Client"},
		SignatureAlgorithm: x509.ECDSAWithSHA256,
	}, clientPrivateKey)
	if err != nil {
		fmt.Println("Error creating CSR:", err)
		return
	}

	// Send CSR to the CA
	err = sendCertificateRequest(conn, csrDER)
	if err != nil {
		fmt.Println("Error sending CSR:", err)
		return
	}

	// Receive signed certificate from the CA
	clientCertificate, err := receiveCertificate(conn)
	if err != nil {
		fmt.Println("Error receiving client certificate:", err)
		return
	}

	// Save client private key to file
	clientPrivateKeyFile, err := os.Create("client_private_key.pem")
	if err != nil {
		fmt.Println("Error creating client private key file:", err)
		return
	}
	ecPrivateKeyBytes, err := x509.MarshalECPrivateKey(clientPrivateKey)
	if err != nil {
		fmt.Println("Error marshaling CA private key:", err)
		return
	}
	pem.Encode(clientPrivateKeyFile, &pem.Block{Type: "EC PRIVATE KEY", Bytes: ecPrivateKeyBytes})

	//pem.Encode(clientPrivateKeyFile, &pem.Block{Type: "EC PRIVATE KEY", Bytes: x509.MarshalECPrivateKey(clientPrivateKey)})
	clientPrivateKeyFile.Close()

	// Save client certificate to file
	clientCertFile, err := os.Create("client_certificate.pem")
	if err != nil {
		fmt.Println("Error creating client certificate file:", err)
		return
	}
	pem.Encode(clientCertFile, &pem.Block{Type: "CERTIFICATE", Bytes: clientCertificate})
	clientCertFile.Close()

	fmt.Println("Client private key and certificate generated successfully.")
}

func receiveCertificate(conn net.Conn) ([]byte, error) {
	lengthBytes := make([]byte, 4)
	_, err := conn.Read(lengthBytes)
	if err != nil {
		return nil, err
	}

	certLength := int(binary.BigEndian.Uint32(lengthBytes))
	certBytes := make([]byte, certLength)
	_, err = conn.Read(certBytes)
	if err != nil {
		return nil, err
	}

	return certBytes, nil
}

func sendCertificateRequest(conn net.Conn, csr []byte) error {
	return sendLengthAndData(conn, csr)
}

func sendLengthAndData(conn net.Conn, data []byte) error {
	err := sendLength(conn, len(data))
	if err != nil {
		return err
	}

	_, err = conn.Write(data)
	return err
}

func sendLength(conn net.Conn, length int) error {
	lengthBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBytes, uint32(length))
	_, err := conn.Write(lengthBytes)
	return err
}

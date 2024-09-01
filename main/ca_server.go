// ca_server.go

package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/binary"
	"fmt"
	"math/big"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting CA server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("CA server is listening on localhost:8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleClient1(conn)
	}
}
func handleClient1(conn net.Conn) {
	defer conn.Close()

	// Generate CA key pair
	caPrivateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		fmt.Println("Error generating CA private key:", err)
		return
	}

	// Create CA certificate
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

	// Send CA certificate to the client
	err = sendCertificate(conn, caDER)
	if err != nil {
		fmt.Println("Error sending CA certificate:", err)
		return
	}
	clientCertificate, err := receiveCertificate(conn)
	if err != nil {
		fmt.Println("Error receiving client certificate:", err)
		return
	}
	fmt.Println(clientCertificate)
	// Wait for client acknowledgment before closing the connection
	_, err = conn.Read(make([]byte, 1))
	if err != nil {
		fmt.Println("Error reading client acknowledgment:", err)
		return
	}
}
func handleClient(conn net.Conn) {
	defer conn.Close()

	// Generate CA key pair
	caPrivateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		fmt.Println("Error generating CA private key:", err)
		return
	}

	// Create CA certificate
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

	// Send CA certificate to the client
	err = sendCertificate(conn, caDER)
	if err != nil {
		fmt.Println("Error sending CA certificate:", err)
		return
	}
}

func sendCertificate(conn net.Conn, certificate []byte) error {
	err := sendLength(conn, len(certificate))
	if err != nil {
		return err
	}

	_, err = conn.Write(certificate)
	return err
}

func sendLength(conn net.Conn, length int) error {
	lengthBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBytes, uint32(length))
	_, err := conn.Write(lengthBytes)
	return err
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

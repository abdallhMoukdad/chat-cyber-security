//package main
//
//import (
//	"bytes"
//	"crypto"
//	"crypto/aes"
//	"crypto/cipher"
//	"crypto/rand"
//	"crypto/rsa"
//	"crypto/sha256"
//	"crypto/x509"
//	"encoding/base64"
//	"encoding/pem"
//	"fmt"
//	"net"
//)
//
//// -------------------2---------------------------------
//// GetAESDecrypted decrypts given text in AES 256 CBC
//func GetAESDecrypted(encrypted, key /*iv*/ string) ([]byte, error) {
//	//key := "my32digitkey12345678901234567890"
//	iv := "my16digitIvKey12"
//
//	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
//	if err != nil {
//		return nil, err
//	}
//
//	block, err := aes.NewCipher([]byte(key))
//	if err != nil {
//		return nil, err
//	}
//
//	if len(ciphertext)%aes.BlockSize != 0 {
//		return nil, fmt.Errorf("block size cant be zero")
//	}
//
//	mode := cipher.NewCBCDecrypter(block, []byte(iv))
//	mode.CryptBlocks(ciphertext, ciphertext)
//	ciphertext = PKCS5UnPadding(ciphertext)
//
//	return ciphertext, nil
//}
//
//// PKCS5UnPadding  pads a certain blob of data with necessary data to be used in AES block cipher
//func PKCS5UnPadding(src []byte) []byte {
//	length := len(src)
//	unpadding := int(src[length-1])
//
//	return src[:(length - unpadding)]
//}
//
//// GetAESEncrypted encrypts given text in AES 256 CBC
//func GetAESEncrypted(plaintext, key /*, iv */ string) (string, error) {
//	//key := "my32digitkey12345678901234567890"
//	iv := "my16digitIvKey12"
//
//	var plainTextBlock []byte
//	length := len(plaintext)
//
//	if length%16 != 0 {
//		extendBlock := 16 - (length % 16)
//		plainTextBlock = make([]byte, length+extendBlock)
//		copy(plainTextBlock[length:], bytes.Repeat([]byte{uint8(extendBlock)}, extendBlock))
//	} else {
//		plainTextBlock = make([]byte, length)
//	}
//
//	copy(plainTextBlock, plaintext)
//	block, err := aes.NewCipher([]byte(key))
//	if err != nil {
//		return "", err
//	}
//
//	ciphertext := make([]byte, len(plainTextBlock))
//	mode := cipher.NewCBCEncrypter(block, []byte(iv))
//	mode.CryptBlocks(ciphertext, plainTextBlock)
//
//	str := base64.StdEncoding.EncodeToString(ciphertext)
//
//	return str, nil
//}
//
////-------------------------------------3--------------------------------------------
//
//func CreateKeys() []byte {
//	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
//	if err != nil {
//		fmt.Println("Error generating private key:", err.Error())
//		return nil
//	}
//
//	publicKey := &privateKey.PublicKey
//
//	pubASN1, err := x509.MarshalPKIXPublicKey(publicKey)
//	if err != nil {
//		fmt.Println("Error marshalling public key:", err.Error())
//		return nil
//	}
//
//	pubPEM := pem.EncodeToMemory(&pem.Block{
//		Type:  "RSA PUBLIC KEY",
//		Bytes: pubASN1,
//	})
//	return pubPEM
//}
//func DecodePublicKey(pubPEMBytes []byte) *rsa.PublicKey {
//	pubBlock, _ := pem.Decode(pubPEMBytes)
//
//	pubKeyInterface, _ := x509.ParsePKIXPublicKey(pubBlock.Bytes)
//	pubKey := pubKeyInterface.(*rsa.PublicKey)
//	fmt.Println("Received Public Key", pubKey.N)
//	return pubKey
//}
//func ServerAsymmetricEncryption(text string) (string, error) {
//	keys := CreateKeys()
//	block, _ := pem.Decode(keys)
//	print(block.Bytes)
//	return
//}
//func generateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
//	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
//	if err != nil {
//		return nil, nil, err
//	}
//	publicKey := &privateKey.PublicKey
//	return privateKey, publicKey, nil
//}
//
//func encodePublicKey(pubkey *rsa.PublicKey) ([]byte, error) {
//	pubkeyBytes, err := x509.MarshalPKIXPublicKey(pubkey)
//	if err != nil {
//		return nil, err
//	}
//	pubkeyPEM := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: pubkeyBytes})
//	return pubkeyPEM, nil
//}
//
//func decodePublicKey(pubkeyPEM []byte) (*rsa.PublicKey, error) {
//	block, _ := pem.Decode(pubkeyPEM)
//	if block == nil {
//		return nil, fmt.Errorf("failed to decode public key")
//	}
//	pubkeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
//	if err != nil {
//		return nil, err
//	}
//	pubkey, ok := pubkeyInterface.(*rsa.PublicKey)
//	if !ok {
//		return nil, fmt.Errorf("failed to parse public key")
//	}
//	return pubkey, nil
//}
//
//func encrypt(data []byte, pubkey *rsa.PublicKey) ([]byte, error) {
//	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pubkey, data)
//	if err != nil {
//		return nil, err
//	}
//	return ciphertext, nil
//}
//
//func decrypt(ciphertext []byte, privatekey *rsa.PrivateKey) ([]byte, error) {
//	data, err := rsa.DecryptPKCS1v15(rand.Reader, privatekey, ciphertext)
//	if err != nil {
//		return []byte{}, err
//	}
//	return data, nil
//}
//
//func main() {
//	serverPrivateKey, serverPublicKey, err := generateKeyPair()
//	if err != nil {
//		fmt.Println("Error generating key pair for server:", err)
//		return
//	}
//	clientPrivateKey, clientPublicKey, err := generateKeyPair()
//	if err != nil {
//		fmt.Println("Error generating key pair for client:", err)
//		return
//	}
//
//	serverPubPEM, err := encodePublicKey(serverPublicKey)
//	if err != nil {
//		fmt.Println("Error encoding server public key:", err)
//		return
//	}
//	clientPubPEM, err := encodePublicKey(clientPublicKey)
//	if err != nil {
//		fmt.Println("Error encoding client public key:", err)
//		return
//	}
//
//	serverPubDecoded, err := decodePublicKey(serverPubPEM)
//
//	clientPubDecoded, err := decodePublicKey(clientPubPEM)
//
//	conn, err := net.Dial("tcp", "localhost:8080")
//	defer conn.Close()
//
//	cipherText, err := encrypt([]byte("Hello from client"), serverPubDecoded)
//
//	conn.Write(cipherText)
//
//	buffer := make([]byte, 1024)
//
//	n, _ := conn.Read(buffer)
//
//	data, _ := decrypt(buffer, clientPrivateKey)
//
//	fmt.Println(string(data))
//
//}
//
//// -----------------------------------4--------------------------
//func main() {
//	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
//	if err != nil {
//		fmt.Println("Error generating private key:", err)
//		return
//	}
//
//	publicKey := &privateKey.PublicKey
//
//	fmt.Println("Private key:", privateKey)
//	fmt.Println("Public key:", publicKey)
//}
//
//// 2. Sign a message using the private key:
//func signMessage(privateKey *rsa.PrivateKey, message []byte) ([]byte, error) {
//	hashed := sha256.Sum256(message)
//	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
//	if err != nil {
//		return nil, err
//	}
//	return signature, nil
//}
//
//// 3. Verify the signature using the public key:
//func verifySignature(publicKey *rsa.PublicKey, message []byte, signature []byte) error {
//	hashed := sha256.Sum256(message)
//	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
//	if err != nil {
//		return fmt.Errorf("signature verification failed: %s", err)
//	}
//	return nil
//}

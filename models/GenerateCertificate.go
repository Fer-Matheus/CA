package models

import (
	"crypto/rand"
	"crypto/x509"
	"fmt"
	"go-CA/internal"
)

func GenerateCertificate(csr []byte) ([]byte, error) {

	csrRequest, err := x509.ParseCertificateRequest(csr)
	if err != nil {
		return nil, err
	}

	privKey, caCert, err := internal.GetKeyCertCA()
	if err != nil {
		return nil, err
	}
	csrTemplate := x509.Certificate{
		PublicKeyAlgorithm: csrRequest.PublicKeyAlgorithm,
		PublicKey:          csrRequest.PublicKey,

		EmailAddresses: caCert.EmailAddresses,
		SerialNumber: caCert.SerialNumber,
		Subject:      csrRequest.Subject,
		NotBefore:    caCert.NotBefore.Local(),
		NotAfter:     caCert.NotAfter,
		KeyUsage:     x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, &csrTemplate, caCert, csrRequest.PublicKey, privKey)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}

	return certBytes, nil
}

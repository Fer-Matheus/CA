package models

import (
	"crypto/rand"
	"crypto/x509"
	"fmt"
	"go-CA/internal"
)

func GenerateCertificate(csr []byte) (x509.Certificate, error) {

	csrRequest, err := x509.ParseCertificateRequest(csr)
	if err != nil {
		return x509.Certificate{}, err
	}

	privKey, caCert, err := internal.GetKeyCertCA()
	if err != nil {
		return x509.Certificate{}, err
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
		return x509.Certificate{}, err
	}

	cert, err := x509.ParseCertificate(certBytes)
	if err != nil {
		return x509.Certificate{}, err
	}

	return *cert, nil
}

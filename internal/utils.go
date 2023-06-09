package internal

import (
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func GetKeyCertCA() (crypto.PrivateKey, *x509.Certificate, error) {

	caCertBytes, err := os.ReadFile("certs/caCert.pem")
	if err != nil {
		return nil, &x509.Certificate{}, err
	}
	caCertPem, _ := pem.Decode(caCertBytes)
	caCert, err := x509.ParseCertificate(caCertPem.Bytes)
	if err != nil {
		return nil, &x509.Certificate{}, err
	}

	keyBytes, err := os.ReadFile("certs/rsaCAPEM.key")
	if err != nil {
		return nil, &x509.Certificate{}, err
	}
	keyPem, _ := pem.Decode(keyBytes)
	key, err := x509.ParsePKCS8PrivateKey(keyPem.Bytes)
	if err != nil {
		return nil, &x509.Certificate{}, err
	}

	return key, caCert, nil
}

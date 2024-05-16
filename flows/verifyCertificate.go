package flows

import (
	"crypto/x509"
	"go-CA/internal"
)

func VerifyCertificate(crt []byte) (bool, error) {

	cert, err := x509.ParseCertificate(crt)
	if err != nil {
		return false, err
	}
	_, caCert, err := internal.GetKeyCertCA()
	if err != nil {
		return false, err
	}

	pool := x509.NewCertPool()
	pool.AddCert(caCert)

	_, err = cert.Verify(x509.VerifyOptions{
		Roots: pool,
		KeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsage(x509.ExtKeyUsageAny)},
	})
	if err != nil {
		return false, err
	}
	
	return true, nil
}
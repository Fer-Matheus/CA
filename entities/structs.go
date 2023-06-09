package entities

import "crypto/x509"

type CSR_Request struct{
	CSR []byte `json:"csr"`
}
type CRT_Response struct {
	CRT x509.Certificate `json:"crt"`
}
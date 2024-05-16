package entities

import "crypto/x509"

type CSR_Request struct{
	CSR []byte `json:"csr"`
}
type CRT_Response struct {
	CRT x509.Certificate `json:"crt"`
}
type CRT_Verification struct {
	CRT []byte `json:"crt"`
}
type CRT_Verification_Response struct {
	PASS bool `json:"pass"`
}

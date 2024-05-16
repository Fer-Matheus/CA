package handlers

import (
	"encoding/json"
	"go-CA/entities"
	"go-CA/flows"
	"net/http"
)

func CertificateRequest(w http.ResponseWriter, r *http.Request) {
	var csr entities.CSR_Request

	err := json.NewDecoder(r.Body).Decode(&csr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	cert, err := flows.GenerateCertificate(csr.CSR)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content=Type", "application/json")
	json.NewEncoder(w).Encode(cert)
}

func CertificateVerification(w http.ResponseWriter, r *http.Request) {
	var crt entities.CRT_Verification

	if err := json.NewDecoder(r.Body).Decode(&crt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	pass, err := flows.VerifyCertificate(crt.CRT)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entities.CRT_Verification_Response{PASS: pass})

}

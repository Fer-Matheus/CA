package handlers

import (
	"encoding/json"
	"go-CA/entities"
	"go-CA/models"
	"net/http"
)

func CertificateRequest(w http.ResponseWriter, r *http.Request) {
	var csr entities.CSR_Request

	err := json.NewDecoder(r.Body).Decode(&csr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	cert, err := models.GenerateCertificate(csr.CSR)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content=Type", "application/json")
	json.NewEncoder(w).Encode(entities.CRT_Response{CRT: cert})
}

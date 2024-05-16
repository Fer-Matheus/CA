package main

import (
	"fmt"
	"go-CA/handlers"
	"log"
	"net/http"
)

func main() {
	
	http.HandleFunc("/csrRequest", handlers.CertificateRequest)
	http.HandleFunc("/crtVerification", handlers.CertificateVerification)
	
	fmt.Println("Api on port :2302")

	if err := http.ListenAndServe(":2302", nil); err != nil {
		log.Fatalf("error listening to port: %v", err)
	}
}

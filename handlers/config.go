package handlers

import "net/http"

var Server = http.Server{
	Addr:    ":2302",
	Handler: http.NewServeMux(),
}

func Setup(mux *http.ServeMux){
	mux.HandleFunc("/csrRequest", CertificateRequest)
}

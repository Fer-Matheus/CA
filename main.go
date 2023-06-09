package main

import (
	"fmt"
	"go-CA/handlers"
	"log"
	"net/http"
)

func main() {
	handler := http.NewServeMux()
	server := handlers.Server

	server.Handler = handler

	handlers.Setup(handler)
	fmt.Println("Api on port ", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("error listening to port: %v", err)
	}

}

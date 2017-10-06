package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodrigodealer/pdfcreator-go/handlers"
)

func handler(w http.ResponseWriter, r *http.Request) {
	return
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/generate", handlers.PdfHandler).Methods("POST")
	r.HandleFunc("/internal/healthcheck", handlers.HealthcheckHandler).Methods("GET")
	http.Handle("/", r)

	log.Print("Starting server on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Panic("Something is wrong : " + err.Error())
	}
}

package main

import (
	"log"
	"net/http"
	Handler "qubeChallenge/Handler"
	Service "qubeChallenge/Service"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	mux := mux.NewRouter()

	H := Handler.Handlers{Service: Service.NewService()}
	mux.HandleFunc("/api/v1/appliances", H.Appliances).Methods(http.MethodGet)
	mux.HandleFunc("/api/v1/appliance/{appliance-id}/info", H.Appliance).Methods(http.MethodGet)

	router := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Add PUT here
		AllowedHeaders: []string{"*"},
	}).Handler(mux)
	listenAddr := ":8080"

	log.Printf("About to listen on 8080. Go to https://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(listenAddr, router))

}

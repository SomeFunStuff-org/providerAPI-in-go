package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/providers", GetProviders).Methods("GET")
	router.HandleFunc("/providers/{id}", GetProvider).Methods("GET")
	router.HandleFunc("/providers/{id}", CreateProvider).Methods("GET")
	router.HandleFunc("/providers/{id}", DeleteProvider).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":7555", router))
}

func GetProviders(w http.ResponseWriter, r *http.Request) {}
func GetProvider(w http.ResponseWriter, r *http.Request) {}
func CreateProvider(w http.ResponseWriter, r *http.Request) {}
func DeleteProvider(w http.ResponseWriter, r *http.Request) {}

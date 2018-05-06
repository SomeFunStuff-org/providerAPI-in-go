package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)



func GetProviders(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(providers)
}

func GetProvider(w http.ResponseWriter, r *http.Request)    {}
func CreateProvider(w http.ResponseWriter, r *http.Request) {}
func DeleteProvider(w http.ResponseWriter, r *http.Request) {}

type Provider struct {
	ID    string `json:"providerID,omitempty"`
	Name  string `json:"name,omitempty"`
	Tier  string `json:"tier,omitempty"`
	Links *Links `json:"links,omitempty"`
}

type Links struct {
	Href string `json:"href,omitempty"`
	Rel  string `json:"lawnmowers"`
	Type string `json:"GET"`
}

var providers []Provider

func main() {
	providers = append(providers, Provider{ID: "1", Name: "lawncare", Tier: "1", Links: &Links{Href: "/lawnmowers", Rel: "Lawnmowers", Type: "GET"}})
	providers = append(providers, Provider{ID: "2", Name: "window cleaning", Tier: "1", Links: &Links{Href: "/windowcleaners", Rel: "windowcleaners", Type: "GET"}})
	providers = append(providers, Provider{ID: "3", Name: "house cleaning", Tier: "1", Links: &Links{Href: "/housecleaners", Rel: "housecleaners", Type: "GET"}})
	providers = append(providers, Provider{ID: "4", Name: "plumbing", Tier: "2", Links: &Links{Href: "/plumbers", Rel: "plumbers", Type: "GET"}})
	providers = append(providers, Provider{ID: "5", Name: "lawyers", Tier: "5", Links: &Links{Href: "/lawyers", Rel: "lawyers", Type: "GET"}})
	router := mux.NewRouter()
	router.HandleFunc("/providers", GetProviders).Methods("GET")
	router.HandleFunc("/providers/{id}", GetProvider).Methods("GET")
	router.HandleFunc("/providers/{id}", CreateProvider).Methods("GET")
	router.HandleFunc("/providers/{id}", DeleteProvider).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":7555", router))
}

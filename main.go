package main

import {
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/providers", GetProviders).Methods("GET")
    router.HandleFunc("/providers/{id}", GetProvider).Methods("GET")
    router.HandleFunc("/providers/{id}", CreateProvider).Methods("GET")
    router.HandleFunc("/providers/{id}", DeleteProvider).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":7555", router))
}

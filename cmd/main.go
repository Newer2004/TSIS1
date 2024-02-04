package main

import (
    "fmt"
    "net/http"
    "github.com/Newer2004/TSIS1/pkg/animal"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/animals", animal.GetAnimalsHandler).Methods("GET")
    r.HandleFunc("/animals/{name}", animal.GetAnimalHandler).Methods("GET")
    r.HandleFunc("/health", animal.HealthCheckHandler).Methods("GET")

    http.Handle("/", r)

    fmt.Println("Server is running on :8080")
    http.ListenAndServe(":8080", nil)
}

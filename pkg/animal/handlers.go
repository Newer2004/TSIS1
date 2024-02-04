package animal

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

var animals = []Animal{
    {"Lion", "Mammal", "Roar"},
    {"Eagle", "Bird", "Screech"},
    {"Fish", "Aquatic", "Bubbles"},
    {"Elephant", "Mammal", "Trumpet"},
    {"Dolphin", "Aquatic", "Clicks"},
    {"Giraffe", "Mammal", "Grunt"},
    {"Snake", "Reptile", "Hiss"},
    {"Owl", "Bird", "Hoot"},
    {"Kangaroo", "Mammal", "Chirrup"},
    {"Tiger", "Mammal", "Roar"},
}

func GetAnimalsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(animals)
}

func GetAnimalHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    name := params["name"]
    for _, a := range animals {
        if a.Name == name {
            json.NewEncoder(w).Encode(a)
            return
        }
    }
    http.NotFound(w, r)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Animal Web App is healthy.")
}


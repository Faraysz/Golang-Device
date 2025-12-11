package handlers

import (
    "encoding/json"
    "fmt"
    "net/http"

    "golang-api/models"
    "golang-api/repositories"
)

func CreateDevice(w http.ResponseWriter, r *http.Request) {
    var device models.Device

    // Decode JSON body
    err := json.NewDecoder(r.Body).Decode(&device)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    device.Status = "online"

    repositories.AddDevice(device)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(device)
}

func GetAllDevices(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(repositories.GetDevices())
}

func Test(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "API Golang Berjalan!")
}

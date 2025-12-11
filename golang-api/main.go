package main

import (
    "log"
    "net/http"

    "golang-api/handlers"
)

func main() {
    http.HandleFunc("/", handlers.Test)
    http.HandleFunc("/devices", handlers.GetAllDevices)
    http.HandleFunc("/device/create", handlers.CreateDevice)

    log.Println("Server berjalan di http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

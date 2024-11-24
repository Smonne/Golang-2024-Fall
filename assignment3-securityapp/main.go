package main

import (
    "log"
    "net/http"
    "go-security-app/handlers"
)

func main() {
    utils.InitLogger()
    defer utils.Logger.Sync()

    router := handlers.SetupRoutes()
    log.Println("Server starting on https://localhost:8443")
    if err := http.ListenAndServeTLS(":8443", "server.crt", "server.key", router); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}

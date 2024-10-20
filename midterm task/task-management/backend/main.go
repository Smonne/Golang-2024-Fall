package main

import (
    "task-manager/database"
    "task-manager/handlers"
    "github.com/gorilla/mux"
    "net/http"
)

func main() {
    database.ConnectDB()

    router := mux.NewRouter()

    router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
    router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
    router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
    router.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

    
    http.ListenAndServe(":8000", router)
}

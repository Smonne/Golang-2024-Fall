package main

import (
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    // Direct SQL routes
    r.HandleFunc("/users-sql", getUsersSQL).Methods("GET")
    r.HandleFunc("/user-sql", createUserSQL).Methods("POST")
    r.HandleFunc("/user-sql/{id}", updateUserSQL).Methods("PUT")
    r.HandleFunc("/user-sql/{id}", deleteUserSQL).Methods("DELETE")

    // GORM routes
    r.HandleFunc("/users-gorm", getUsersGORM).Methods("GET")
    r.HandleFunc("/user-gorm", createUserGORM).Methods("POST")
    r.HandleFunc("/user-gorm/{id}", updateUserGORM).Methods("PUT")
    r.HandleFunc("/user-gorm/{id}", deleteUserGORM).Methods("DELETE")

    http.ListenAndServe(":8080", r)

		
}

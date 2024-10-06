package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "gorm.io/gorm"
)

// SQL Handlers
func getUsersSQL(w http.ResponseWriter, r *http.Request) { /*...*/ }
func createUserSQL(w http.ResponseWriter, r *http.Request) { /*...*/ }
func updateUserSQL(w http.ResponseWriter, r *http.Request) { /*...*/ }
func deleteUserSQL(w http.ResponseWriter, r *http.Request) { /*...*/ }

// GORM Handlers
func getUsersGORM(w http.ResponseWriter, r *http.Request) { /*...*/ }
func createUserGORM(w http.ResponseWriter, r *http.Request) { /*...*/ }
func updateUserGORM(w http.ResponseWriter, r *http.Request) { /*...*/ }
func deleteUserGORM(w http.ResponseWriter, r *http.Request) { /*...*/ }


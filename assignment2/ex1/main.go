package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq" // This imports the pq driver
)

const (
    host     = "localhost"
    port     = 5432
    user     = "your_username"
    password = "your_password"
    dbname   = "your_dbname"
)


func connectDB() (*sql.DB, error) {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        return nil, err
    }

    fmt.Println("Successfully connected to the database!")
    return db, nil
}

func createTable(db *sql.DB) {
    query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        age INT NOT NULL
    );`

    _, err := db.Exec(query)
    if err != nil {
        panic(err)
    }

    fmt.Println("Table created successfully!")
}

func insertUser(db *sql.DB, name string, age int) {
    query := `INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id`
    var id int
    err := db.QueryRow(query, name, age).Scan(&id)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Inserted user with ID: %d\n", id)
}

func getUsers(db *sql.DB) {
    rows, err := db.Query("SELECT id, name, age FROM users")
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    fmt.Println("Users:")
    for rows.Next() {
        var id int
        var name string
        var age int
        err = rows.Scan(&id, &name, &age)
        if err != nil {
            panic(err)
        }
        fmt.Printf("%d: %s, %d years old\n", id, name, age)
    }
}

func main() {
    
    db, err := connectDB()
    if err != nil {
        panic(err)
    }
    defer db.Close()

  
    createTable(db)

    insertUser(db, "Alice", 30)
    insertUser(db, "Bob", 25)

   
    getUsers(db)
}

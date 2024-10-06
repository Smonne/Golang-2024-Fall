package main

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

type User struct {
    ID   uint   `gorm:"primaryKey"`
    Name string `gorm:"unique;not null"`
    Age  int    `gorm:"not null"`
}

func connectGORM() (*gorm.DB, error) {
    dsn := "host=localhost user=your_username password=your_password dbname=your_dbname port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    fmt.Println("Successfully connected using GORM!")
    return db, nil
}

func autoMigrate(db *gorm.DB) {
    err := db.AutoMigrate(&User{})
    if err != nil {
        panic("Failed to migrate the User model")
    }

    fmt.Println("Database migration successful!")
}


func insertUserGORM(db *gorm.DB, name string, age int) {
    user := User{Name: name, Age: age}
    result := db.Create(&user)
    if result.Error != nil {
        panic(result.Error)
    }

    fmt.Printf("Inserted user with ID: %d\n", user.ID)
}

func getUsersGORM(db *gorm.DB) {
    var users []User
    result := db.Find(&users)
    if result.Error != nil {
        panic(result.Error)
    }

    fmt.Println("Users from GORM:")
    for _, user := range users {
        fmt.Printf("%d: %s, %d years old\n", user.ID, user.Name, user.Age)
    }
}


func main() {
 
    db, err := connectGORM()
    if err != nil {
        panic(err)
    }

    autoMigrate(db)

    insertUserGORM(db, "Alice", 30)
    insertUserGORM(db, "Bob", 25)

    getUsersGORM(db)
}


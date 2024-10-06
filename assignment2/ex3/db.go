package main

import (
    "database/sql"
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func connectDB() (*sql.DB, error) {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
    return sql.Open("postgres", psqlInfo)
}

func connectGORM() (*gorm.DB, error) {
    dsn := "host=localhost user=your_username password=your_password dbname=your_dbname port=5432 sslmode=disable"
    return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

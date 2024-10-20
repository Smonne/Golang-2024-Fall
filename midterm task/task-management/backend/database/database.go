package database

import (
    "gorm.io/gorm"
	"gorm.io/driver/sqlite" 
   // "modernc.org/sqlite"
    //"gorm.io/driver/postgres"
    "log"
    "task-manager/models"
)

var DB *gorm.DB


func ConnectDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{}) 
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    log.Println("Database connected successfully")

    // Auto migrate the Task model
    err =  DB.AutoMigrate(&models.Task{}, &models.User{})
    if err != nil {
        log.Fatal("Failed to migrate database:", err)
    }
    log.Println("Database migrated successfully")
}
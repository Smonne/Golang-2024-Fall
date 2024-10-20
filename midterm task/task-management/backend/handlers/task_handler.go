
package handlers

import (
    "net/http"
    "encoding/json"
    "task-manager/database"
    "task-manager/models"
    "github.com/gorilla/mux"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
    var task models.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := database.DB.Create(&task).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(task)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
    var tasks []models.Task
    if err := database.DB.Find(&tasks).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    taskID := vars["id"]

    var task models.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Find and update the task
    if err := database.DB.Model(&task).Where("id = ?", taskID).Updates(task).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    taskID := vars["id"]

    // Delete the task
    if err := database.DB.Delete(&models.Task{}, taskID).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

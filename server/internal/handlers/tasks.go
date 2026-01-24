package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"

	"ToDoGo/internal/config"
	"ToDoGo/internal/models"
)

type PostTaskRequest struct {
	UserID   string `json:"user_id"`
	TaskName string `json:"task_name"`
}

type PostTaskResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	TaskName  string `json:"task_name"`
	CreatedAt string `json:"created_at"`
}

func PostTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req PostTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	if req.UserID == "" || req.TaskName == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "user_id and task_name are required"})
		return
	}

	task := models.Tasks{
		ID:        uuid.New().String(),
		UserID:    req.UserID,
		TaskName:  req.TaskName,
		Completed: false,
		CreatedAt: time.Now(),
	}

	ctx := context.Background()
	_, err := config.DB.Exec(ctx,
		"INSERT INTO Task_Table (id, user_id, task_name, completed, created_at) VALUES ($1, $2, $3, $4, $5)",
		task.ID,
		task.UserID,
		task.TaskName,
		task.Completed,
		task.CreatedAt,
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	response := PostTaskResponse{
		ID:        task.ID,
		UserID:    task.UserID,
		TaskName:  task.TaskName,
		CreatedAt: task.CreatedAt.Format(time.RFC3339),
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

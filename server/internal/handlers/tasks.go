package handlers



import (
		"context"
	"encoding/json"
	"net/http"
	"time"
		// "fmt"

	//pks to do stuff
	"github.com/google/uuid"

	"ToDoGo/internal/config"

	//struct
	"ToDoGo/internal/models"

)

type PostTaskRequest struct {
Tasks : string `json:"tasks`
Email : string `json:"email"`}


type PostTaskRequest struct {
ID string `json:"id"`
Tasks: string `json:"tasks"`
Email: string `json:"email"`
created_at string `json:"created_at"`
}

function PostTasks(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json");

		var PostTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
	return
	}
	
	tasks := models.Tasks {
		ID : uuid.New().String(),
		UserID
	}


}
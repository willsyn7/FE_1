package models


type Tasks struct {
	ID string `json:"id" db:"id"`
	UserID string `json:"user_id" db:"user_id`
	TaskName string `json:"task_name" db: "task_name"`
	Completed bool `json:"completed" db:"completed"`
	CreatedAt time.Time `json:"created_at" db:"created_at`
}







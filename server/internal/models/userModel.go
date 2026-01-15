package models

import "time"

type User struct {
	ID        string    `json:"id" db:"id"`
	UserName  string    `json:"user_name" db:"user_name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
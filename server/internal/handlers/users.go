package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"fmt"

	"github.com/google/uuid"

	"ToDoGo/internal/config"
	"ToDoGo/internal/models"
)
// type UserDataStruct struct {
// 	ID	string `json :"id"`
// 	UserName string `json:"user_name"`
// 	Email	string `json:"email"`
// 	CreatedAt string `json:"created_at"`
// }

type SignUpRequest struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}


type SignUpResponse struct {
	ID        string `json:"id"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

type DeleteUserRequest struct{
Email string `json:"email"`
}

type DeleteUserResponse struct{
	Message string `json:"message"`
}

type GetUserDataRequest struct{
	Email string `json:"email"`
}

type GetUserDataResponse struct{
	ID	string `json :"id"`
	UserName string `json:"user_name"`
	Email	string `json:"email"`
	CreatedAt string `json:"created_at"`

	
}




func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //set appplcaition ehader type

	var req SignUpRequest // declareing that req will use singup rqeueust sturucut
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { // decodoe sturucutre 
		w.WriteHeader(http.StatusBadRequest) //stanrdard method to rwrite bad rqerueusut header 
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"}) // 
		return
	}

	fmt.Println(req)   
	// Validate required fields
	if req.UserName == "" || req.Email == "" || req.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "user_name, email, and password are required"})
		return
	}

	user := models.User{
		ID:        uuid.New().String(),
		UserName:  req.UserName,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: time.Now(),
	}

	// Insert user into database
	ctx := context.Background()
	_, err := config.DB.Exec(ctx,
		"INSERT INTO User_Table (id, user_name, email, password, created_at) VALUES ($1, $2, $3, $4, $5)",
		user.ID, user.UserName, user.Email, user.Password, user.CreatedAt,
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	response := SignUpResponse{
		ID:        user.ID,
		UserName:  user.UserName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}



func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req DeleteUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	if req.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "email is required"})
		return
	}

	ctx := context.Background()
	_, err := config.DB.Exec(ctx, "DELETE FROM User_Table WHERE email = $1", req.Email)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	response := DeleteUserResponse{
		Message: "User has been deleted",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}



func GetUserData(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var req GetUserDataRequest

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
        return
    }

    if req.Email == "" {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "email is required"})
        return
    }

    ctx := context.Background()
    var userData models.User

    err := config.DB.QueryRow(ctx,
        "SELECT id, email, user_name, created_at FROM User_Table WHERE email = $1",
        req.Email,
    ).Scan(&userData.ID, &userData.Email, &userData.UserName, &userData.CreatedAt)

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }  


    fmt.Println(userData)

    response := GetUserDataResponse{
        ID:        userData.ID,
        Email:     userData.Email,
        UserName:  userData.UserName,
        CreatedAt: userData.CreatedAt.Format(time.RFC3339),
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

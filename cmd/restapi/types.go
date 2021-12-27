package main

// User is a struct that holds the user information
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

// JSONResponse is a struct that holds the response information
type JSONResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

type ProjectJSONResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Project `json:"data"`
}

type Project struct {
	ID        int    `json:"id"`
	Name      string `json:"name,omitempty"`
	UserID    int    `json:"user_id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

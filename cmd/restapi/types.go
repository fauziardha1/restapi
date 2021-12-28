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
	Data    []User `json:"data,omitempty"`
}

type ProjectJSONResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Project `json:"data,omitempty"`
}

type Project struct {
	ID        int    `json:"id"`
	Name      string `json:"name,omitempty"`
	UserID    int    `json:"user_id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type Measurement struct {
	ID        int    `json:"id"`
	ProjectID int    `json:"project_id,omitempty"`
	FileName  string `json:"file_name,omitempty"`
	SRPValue  string `json:"srp_value,omitempty"`
	OCPValue  string `json:"ocp_value,omitempty"`
	LSPValue  string `json:"lsp_value,omitempty"`
	ISPValue  string `json:"isp_value,omitempty"`
	DIPValue  string `json:"dip_value,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type MeasurementJSONResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []Measurement `json:"data,omitempty"`
}

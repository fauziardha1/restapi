package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// Product is a struct that holds the product information
type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

// HandleGetProductList is a function that handles the GET request to /api/products
func HandleGetProductList(w http.ResponseWriter, r *http.Request) {
	// Create a slice of Product structs
	products := []Product{
		{ID: "1", Name: "Product 1", Price: "100"},
		{ID: "2", Name: "Product 2", Price: "200"},
		{ID: "3", Name: "Product 3", Price: "300"},
	}

	// Create a new JSON encoder
	jsonEncoder := json.NewEncoder(w)

	// Encode the products slice into the response
	jsonEncoder.Encode(products)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homePage).Methods("GET")

	// api handling
	r.HandleFunc("/api/products", HandleGetProductList).Methods("Get")

	// get users
	r.HandleFunc("/api/users", GetUsers).Methods("GET")

	// account handling
	r.HandleFunc("/api/user/login", login).Methods("POST")
	r.HandleFunc("/api/user/register", register).Methods("POST")

	// projects handling
	r.HandleFunc("/api/project/list", HandleGetProjectList).Methods("POST")
	r.HandleFunc("/api/project/create", HandleCreateProject).Methods("POST")
	r.HandleFunc("/api/project/update", HandleUpdateProject).Methods("PATCH")
	r.HandleFunc("/api/project/delete", HandleDeleteProject).Methods("DELETE")

	// measurements handling
	r.HandleFunc("/api/measurement/list", HandleGetMeasurementList).Methods("POST")
	r.HandleFunc("/api/measurement/create", HandleCreateMeasurement).Methods("POST")
	r.HandleFunc("/api/measurement/delete/all", HandleDeleteAllMeasurements).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":17001", r))
	fmt.Println("Server started on port 17001")

}

func homePage(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// enable cors
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// GetUsers is a function that handles the GET request to /api/users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	db := setupDB()
	fmt.Println("Endpoint Hit: GetUsers")

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}

	users := make([]User, 0, 1)
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	respons := JSONResponse{Status: 200, Message: "Success", Data: users}

	json.NewEncoder(w).Encode(respons)
}

// HandleCreateMeasurement is a function that handles the POST request to /api/measurement/create
func HandleCreateMeasurement(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	projectID := r.FormValue("project_id")
	fileName := r.FormValue("file_name")
	srpval := r.FormValue("srp_val")
	ocpval := r.FormValue("ocp_val")
	lspval := r.FormValue("lsp_val")
	ispval := r.FormValue("isp_val")
	dipval := r.FormValue("dip_val")

	if projectID == "" || fileName == "" || srpval == "" || ocpval == "" || lspval == "" || ispval == "" || dipval == "" {
		respons := JSONResponse{Status: 400, Message: "Please fill all the fields"}
		json.NewEncoder(w).Encode(respons)
		return
	} else {
		db := setupDB()
		fmt.Println("Endpoint Hit: HandleCreateMeasurement")

		var measurementID int
		err := db.QueryRow(`INSERT INTO measurements (
			project_id, 
			file_name, 
			srp_val, 
			ocp_val, 
			lsp_val, 
			isp_val, 
			dip_val,
			created_at, 
			updated_at) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			returning id;`,
			projectID,
			fileName,
			srpval,
			ocpval,
			lspval,
			ispval,
			dipval,
			time.Now(),
			time.Now(),
		).Scan(&measurementID)
		if err != nil {
			panic(err)
		}

		respons := MeasurementJSONResponse{
			Status:  200,
			Message: "Success",
			Data:    []Measurement{{ID: measurementID}},
		}
		json.NewEncoder(w).Encode(respons)
	}
}

// HandleCreateProject is a function that handles the POST request to /api/project/create
func HandleCreateProject(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	user_id := r.FormValue("user_id")
	name := r.FormValue("project_name")

	if user_id == "" || name == "" {
		respons := JSONResponse{Status: 400, Message: "Please fill all the fields"}
		json.NewEncoder(w).Encode(respons)
		return
	} else {
		db := setupDB()
		fmt.Println("Endpoint Hit: HandleCreateProject")

		var projectID int
		err := db.QueryRow("INSERT INTO projects (id,user_id, name, created_at, updated_at) VALUES ((SELECT MAX(id)+1 FROM public.projects), $1, $2, $3, $4) returning id;", user_id, name, time.Now(), time.Now()).Scan(&projectID)
		if err != nil {
			panic(err)
		}

		project := Project{ID: projectID}
		respons := ProjectJSONResponse{Status: 200, Message: "Success", Data: []Project{project}}
		json.NewEncoder(w).Encode(respons)
	}
}

// HandleDeleteAllMeasurements is a function that handles the DELETE request to /api/measurement/delete/all
// it will delete all measurements for a given project
func HandleDeleteAllMeasurements(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	projectID := r.FormValue("project_id")

	if projectID == "" {
		respons := JSONResponse{Status: 400, Message: "Please fill all the fields"}
		json.NewEncoder(w).Encode(respons)
		return
	} else {
		db := setupDB()
		fmt.Println("Endpoint Hit: HandleDeleteAllMeasurements")

		_, err := db.Exec("DELETE FROM measurements WHERE project_id = $1;", projectID)
		if err != nil {
			panic(err)
		}

		respons := JSONResponse{Status: 200, Message: "Success"}
		json.NewEncoder(w).Encode(respons)
	}
}

// HandleDeleteProject is a function that handles the PATCH request to /api/project/delete
func HandleDeleteProject(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	projectID := r.FormValue("project_id")

	if projectID == "" {
		respons := JSONResponse{Status: 400, Message: "Please fill all the fields"}
		json.NewEncoder(w).Encode(respons)
		return
	} else {
		db := setupDB()
		fmt.Println("Endpoint Hit: HandleDeleteProject")

		_, err := db.Exec("DELETE FROM projects WHERE id = $1", projectID)
		if err != nil {
			panic(err)
		}

		respons := ProjectJSONResponse{Status: 200, Message: "Success", Data: []Project{}}
		json.NewEncoder(w).Encode(respons)
	}
}

// HandleGetProjectList is a function that handles the GET request to /api/projects/list
func HandleGetProjectList(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	// get the user data form the request
	user_id := r.FormValue("user_id")

	if user_id == "" {
		respons := JSONResponse{Status: 400, Message: "Please fill all the fields"}
		json.NewEncoder(w).Encode(respons)
		return
	} else {
		db := setupDB()
		fmt.Println("Endpoint Hit: HandleGetProjectList")

		rows, err := db.Query("SELECT * FROM projects WHERE user_id = $1", user_id)
		if err != nil {
			panic(err)
		}

		projects := make([]Project, 0, 1)
		for rows.Next() {
			var project Project
			err = rows.Scan(&project.ID, &project.UserID, &project.Name, &project.CreatedAt, &project.UpdatedAt)
			if err != nil {
				panic(err)
			}
			projects = append(projects, project)
		}

		if len(projects) > 0 {
			respons := ProjectJSONResponse{Status: 200, Message: "Success", Data: projects}
			json.NewEncoder(w).Encode(respons)
		} else {
			respons := JSONResponse{Status: 400, Message: "No projects found"}
			json.NewEncoder(w).Encode(respons)
		}
	}
}

// HandleGetMeasurementList is a function that handles the GET request to /api/measurement/list
func HandleGetMeasurementList(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	projectID := r.FormValue("project_id")

	if projectID == "" {
		respons := JSONResponse{Status: 400, Message: "Please fill all the fields"}
		json.NewEncoder(w).Encode(respons)
		return
	} else {
		db := setupDB()
		fmt.Println("Endpoint Hit: HandleGetMeasurementList")

		rows, err := db.Query("SELECT * FROM measurements WHERE project_id = $1", projectID)
		if err != nil {
			panic(err)
		}

		measurements := make([]Measurement, 0, 1)
		for rows.Next() {
			var measurement Measurement
			err = rows.Scan(
				&measurement.ID,
				&measurement.ProjectID,
				&measurement.FileName,
				&measurement.SRPValue,
				&measurement.OCPValue,
				&measurement.LSPValue,
				&measurement.ISPValue,
				&measurement.DIPValue,
				&measurement.CreatedAt,
				&measurement.UpdatedAt,
			)
			if err != nil {
				panic(err)
			}
			measurements = append(measurements, measurement)
		}

		if len(measurements) > 0 {
			respons := MeasurementJSONResponse{Status: 200, Message: "Success", Data: measurements}
			json.NewEncoder(w).Encode(respons)
		} else {
			respons := JSONResponse{Status: 400, Message: "No measurements found"}
			json.NewEncoder(w).Encode(respons)
		}
	}
}

// HandleUpdateProject is a function that handles the PATCH request to /api/project/update
func HandleUpdateProject(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	projectID := r.FormValue("project_id")
	newname := r.FormValue("project_name")
	updatedTime := time.Now()

	if projectID == "" || newname == "" {
		respons := JSONResponse{Status: 400, Message: "Please fill all the fields"}
		json.NewEncoder(w).Encode(respons)
		return
	} else {
		db := setupDB()
		fmt.Println("Endpoint Hit: HandleUpdateProject")

		_, err := db.Exec("UPDATE projects SET name = $1, updated_at = $2 WHERE id = $3", newname, updatedTime, projectID)
		if err != nil {
			panic(err)
		}

		id, err := strconv.Atoi(projectID)
		if err != nil {
			panic(err)
		}

		project := Project{ID: id}
		respons := ProjectJSONResponse{Status: 200, Message: "Success", Data: []Project{project}}
		json.NewEncoder(w).Encode(respons)
	}
}

// login is a function that handles the POST request to /api/user/login
// it will return a JSON response with the user information
func login(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	// get the user data form the request
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		respons := JSONResponse{Status: 400, Message: "Please fill all the fields"}
		json.NewEncoder(w).Encode(respons)
		return
	} else {
		db := setupDB()
		fmt.Println("Endpoint Hit: login")

		rows, err := db.Query("SELECT * FROM users WHERE username = $1 AND password = $2", username, password)
		if err != nil {
			panic(err)
		}

		users := make([]User, 0, 1)
		for rows.Next() {
			var user User
			err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
			if err != nil {
				panic(err)
			}
			if user.Password != "" {
				user.Password = ""
			}
			users = append(users, user)
		}

		if len(users) > 0 {
			respons := JSONResponse{Status: 200, Message: "Success", Data: users}
			json.NewEncoder(w).Encode(respons)
		} else {
			respons := JSONResponse{Status: 400, Message: "Invalid username or password"}
			json.NewEncoder(w).Encode(respons)
		}
	}
}

// register is a function that handles the POST request to /api/user/register
// it will return a JSON response with the user information
func register(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	if username == "" || password == "" || email == "" {
		respons := JSONResponse{Status: 400, Message: "Please fill all the fields"}
		json.NewEncoder(w).Encode(respons)
		return
	} else {
		db := setupDB()
		fmt.Println("Endpoint Hit: register")

		var lastInsertID int
		err := db.QueryRow("INSERT INTO users (id,username, password, email) VALUES ((SELECT MAX(id)+1 FROM public.users),$1, $2, $3) returning id;", username, password, email).Scan(&lastInsertID)
		if err != nil {
			// log.Fatal(err)
			panic(err)
		}
		users := make([]User, 0, 1)
		if lastInsertID > 0 {
			res, err := db.Query("SELECT * FROM users WHERE id = $1", lastInsertID)
			if err != nil {
				// log.Fatal(err)
				panic(err)
			}
			for res.Next() {
				var user User
				err = res.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
				if err != nil {
					// log.Fatal(err)
					panic(err)
				}
				if user.Password != "" {
					user.Password = ""
				}
				users = append(users, user)
			}
		}

		respons := JSONResponse{Status: 200, Message: "Success", Data: users}
		json.NewEncoder(w).Encode(respons)
	}
}

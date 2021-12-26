package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

	log.Fatal(http.ListenAndServe(":8000", r))

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

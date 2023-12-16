package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Create customers struct Each customer includes ID, Name, Role, Email, Phone, Contacted (i.e., indication of whether or not the customer has been contacted)
type Customer struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Role      string `json:"role,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Contacted bool   `json:"contacted"`
}

// Customers are stored appropriately in a basic data structure (e.g., slice, map, etc.) that represents a "database."
var database = []Customer{
	// adding three customer includes ID, Name, Role, Email, Phone, Contacted randomly
	{ID: 1, Name: "John", Role: "CEO", Email: "john123@gmail.com", Phone: "123456789", Contacted: true},
	{ID: 2, Name: "Mary", Role: "CTO", Email: "mary123@gmail.com,", Phone: "45367812", Contacted: false},
	{ID: 3, Name: "Peter", Role: "CFO", Email: "peter678@hotmail.com,", Phone: "87654321", Contacted: true},
}

func deleteCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// Get the customer ID from the URL path parameter
	customerID := params["ID"]
	// Implement the logic to delete the customer with the given ID from the database
	for index, customer := range database {
		if customerID == fmt.Sprintf("%v", customer.ID) {
			database = append(database[:index], database[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// Get the customer ID from the URL path parameter
	customerID := params["ID"]
	// Implement the logic to retrieve the customer with the given ID from the database
	for _, customer := range database {
		if customerID == fmt.Sprintf("%v", customer.ID) {
			json.NewEncoder(w).Encode(customer)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Implement the logic to add a new customer to the database
	var newCustomer Customer
	json.NewDecoder(r.Body).Decode(&newCustomer)
	database = append(database, newCustomer)
	w.WriteHeader(http.StatusCreated)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// Get the customer ID from the URL path parameter
	customerID := params["ID"]
	// Implement the logic to update the customer with the given ID in the database
	var updatedCustomer Customer
	json.NewDecoder(r.Body).Decode(&updatedCustomer)
	for index, customer := range database {
		if customerID == fmt.Sprintf("%v", customer.ID) {
			database[index] = updatedCustomer
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Implement the logic to retrieve all customers from the database
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(database)
}

func main() {
	// Creating a new router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	// Creating a customer through a /customers path
	router.HandleFunc("/customers", createCustomer).Methods("POST")
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	// Updating a customer through a /customers/{id} path
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/deleteCustomers/{id}", deleteCustomers).Methods("DELETE")

	fmt.Println("Server is starting on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}

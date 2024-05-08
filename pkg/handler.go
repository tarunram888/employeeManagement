package employee

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ListEmployeesHandler handles the request to list employees
func ListEmployeesHandler(store *EmployeeStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get queryParams fromthe URL
		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil || page < 1 {
			// Default page number
			page = 1
		}
		pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
		if err != nil || pageSize < 1 {
			// Default page size
			pageSize = 10
		}

		// To fetch employees with pagination
		employees := store.ListEmployees(page, pageSize)

		// Write the response in JSON format
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(employees)
	}
}

// GetEmployeeHandler handles the request to retrieve an employee by ID
func GetEmployeeHandler(store *EmployeeStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract employee ID from URL path parameters
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid employee ID", http.StatusBadRequest)
			return
		}

		// To retrieve the employee
		emp, err := store.GetEmployeeByID(id)
		if err != nil {
			http.Error(w, "Employee not found", http.StatusNotFound)
			return
		}

		// Write the response in JSON format
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(emp)
	}
}

// CreateEmployeeHandler handles the request to create a new employee
func CreateEmployeeHandler(store *EmployeeStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var emp Employee

		// Parse the request body into an Employee struct
		if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// To add the new employee
		store.CreateEmployee(emp)

		// Set the response status code to 201 Created
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(emp)
	}
}

// UpdateEmployeeHandler handles the request to update an existing employee
func UpdateEmployeeHandler(store *EmployeeStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid employee ID", http.StatusBadRequest)
			return
		}

		var newEmp Employee
		if err := json.NewDecoder(r.Body).Decode(&newEmp); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// To update existing employee details
		if !store.UpdateEmployee(id, newEmp) {
			http.Error(w, "Employee not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newEmp)
	}
}

// DeleteEmployeeHandler handles the request to delete an employee by ID
func DeleteEmployeeHandler(store *EmployeeStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid employee ID", http.StatusBadRequest)
			return
		}

		// To delete the existing employee
		if !store.DeleteEmployee(id) {
			http.Error(w, "Employee not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

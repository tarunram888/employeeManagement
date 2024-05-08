// handlers_test.go

package employee

import (
	"bytes"
	employee "employeeManagement/pkg"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
)

func TestListEmployeesHandler(t *testing.T) {
	// Instance of employeeStore
	store := employee.NewEmployeeStore()

	// Adding employees
	store.CreateEmployee(employee.Employee{ID: 1, Name: "virat kohli", Position: "Manager", Salary: 150000})
	store.CreateEmployee(employee.Employee{ID: 2, Name: "Tarun Ram", Position: "Developer", Salary: 100000})

	// create request
	req, err := http.NewRequest("GET", "/employees", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a recorder to record the HTTP response and run handler
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(employee.ListEmployeesHandler(store))
	handler.ServeHTTP(rr, req)

	// Check the HTTP status code returned by the handler
	if rr.Code != http.StatusOK {
		t.Errorf("ListEmployeesHandler returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}

	// Check the response body
	var employees []employee.Employee
	if err := json.Unmarshal(rr.Body.Bytes(), &employees); err != nil {
		t.Errorf("Error decoding JSON response: %v", err)
	}

	// Check if the correct number of employees is returned
	expectedNumEmployees := 2
	if len(employees) != expectedNumEmployees {
		t.Errorf("ListEmployeesHandler returned wrong number of employees: got %v, want %v", len(employees), expectedNumEmployees)
	}

}

func TestGetEmployeeHandler(t *testing.T) {
	store := employee.NewEmployeeStore()

	// Adding employees
	store.CreateEmployee(employee.Employee{ID: 1, Name: "virat kohli", Position: "Manager", Salary: 150000})
	store.CreateEmployee(employee.Employee{ID: 2, Name: "Tarun Ram", Position: "Developer", Salary: 100000})

	employeeID := 1

	// HTTP GET request to retrieve the employee
	req, err := http.NewRequest("GET", "/employees/"+strconv.Itoa(employeeID), nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set the URL variables for the request
	req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(employeeID)})

	// Create a recorder to record the HTTP response, createHandler and server the request.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(employee.GetEmployeeHandler(store))
	handler.ServeHTTP(rr, req)

	// Check the HTTP status code returned by the handler
	if rr.Code != http.StatusOK {
		t.Errorf("GetEmployeeHandler returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}

	var emp employee.Employee
	if err := json.Unmarshal(rr.Body.Bytes(), &emp); err != nil {
		t.Errorf("Error decoding JSON response: %v", err)
	}

	// Verify employee
	expectedEmployee, _ := store.GetEmployeeByID(employeeID)
	if emp != expectedEmployee {
		t.Errorf("GetEmployeeHandler returned wrong employee: got %v, want %v", emp, expectedEmployee)
	}
}

func TestCreateEmployeeHandler(t *testing.T) {
	store := employee.NewEmployeeStore()

	// Create a new employee to add
	emp := employee.Employee{ID: 1, Name: "Tarun Ram", Position: "Developer", Salary: 100000}
	empJSON, err := json.Marshal(emp)
	if err != nil {
		t.Fatal(err)
	}

	// HTTP POST request to create the employee
	req, err := http.NewRequest("POST", "/employees", bytes.NewBuffer(empJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(employee.CreateEmployeeHandler(store))
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("CreateEmployeeHandler returned wrong status code: got %v, want %v", rr.Code, http.StatusCreated)
	}

	// Check if the employee is added to the store correctly
	createdEmployee, _ := store.GetEmployeeByID(emp.ID)
	if createdEmployee != emp {
		t.Errorf("CreateEmployeeHandler failed to add employee to store correctly: got %v, want %v", createdEmployee, emp)
	}
	fmt.Println("Response body:", rr.Body.String())
	var responseEmployee employee.Employee
	if err := json.NewDecoder(rr.Body).Decode(&responseEmployee); err != nil {
		t.Errorf("Error decoding JSON response: %v", err)
	}

	// Verify that the response body contains the correct data
	if responseEmployee != emp {
		t.Errorf("CreateEmployeeHandler returned wrong employee in response body: got %v, want %v", responseEmployee, emp)
	}
}

func TestUpdateEmployeeHandler(t *testing.T) {
	store := employee.NewEmployeeStore()

	store.CreateEmployee(employee.Employee{ID: 1, Name: "virat kohli", Position: "Manager", Salary: 150000})
	store.CreateEmployee(employee.Employee{ID: 2, Name: "Tarun Ram", Position: "Developer", Salary: 100000})

	// updated employee to send in the request
	updatedEmployee := employee.Employee{ID: 1, Name: "MS Dhoni", Position: "Manager", Salary: 200000}
	updatedEmployeeJSON, err := json.Marshal(updatedEmployee)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/employees/1", bytes.NewBuffer(updatedEmployeeJSON))
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(employee.UpdateEmployeeHandler(store))
	handler.ServeHTTP(rr, req)

	// Check the HTTP status code returned by the handler
	if rr.Code != http.StatusOK {
		t.Errorf("UpdateEmployeeHandler returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}

	// Check if the employee is updated correctly in the store
	updatedEmployeeFromStore, err := store.GetEmployeeByID(1)
	if err != nil {
		t.Errorf("Error getting updated employee from store: %v", err)
	}
	if updatedEmployeeFromStore != updatedEmployee {
		t.Errorf("UpdateEmployeeHandler failed to update employee correctly in store: got %v, want %v", updatedEmployeeFromStore, updatedEmployee)
	}

	fmt.Println("Response body:", rr.Body.String())
	var responseEmployee employee.Employee
	if err := json.NewDecoder(rr.Body).Decode(&responseEmployee); err != nil {
		t.Errorf("Error decoding JSON response: %v", err)
	}

	// Verify that the response body contains the correct data
	if responseEmployee != updatedEmployee {
		t.Errorf("UpdateEmployeeHandler returned wrong employee in response body: got %v, want %v", responseEmployee, updatedEmployee)
	}
}

func TestDeleteEmployeeHandler(t *testing.T) {
	store := employee.NewEmployeeStore()

	store.CreateEmployee(employee.Employee{ID: 1, Name: "virat kohli", Position: "Manager", Salary: 150000})
	store.CreateEmployee(employee.Employee{ID: 2, Name: "Tarun Ram", Position: "Developer", Salary: 100000})

	// HTTP DELETE request to delete the employee
	req, err := http.NewRequest("DELETE", "/employees/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(employee.DeleteEmployeeHandler(store))
	handler.ServeHTTP(rr, req)

	// Check the HTTP status code returned by the handler
	if rr.Code != http.StatusOK {
		t.Errorf("DeleteEmployeeHandler returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}

	// Check if the employee is removed from the store
	_, err = store.GetEmployeeByID(1)
	if err == nil {
		t.Errorf("DeleteEmployeeHandler failed to delete employee from store: employee still exists")
	}

	// Check the response body (it should be empty for a successful delete operation)
	if rr.Body.String() != "" {
		t.Errorf("DeleteEmployeeHandler returned non-empty response body: got %v, want empty", rr.Body.String())
	}
}

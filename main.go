package main

import (
	employee "employeeManagement/pkg"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// To create a instance of employee
	employeeStore := employee.NewEmployeeStore()

	router := setupRouter(employeeStore)

	log.Println("Running server.., port = :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func setupRouter(employeeStore *employee.EmployeeStore) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/employees", employee.ListEmployeesHandler(employeeStore)).Methods("GET")
	router.HandleFunc("/employees/{id}", employee.GetEmployeeHandler(employeeStore)).Methods("GET")
	router.HandleFunc("/employees", employee.CreateEmployeeHandler(employeeStore)).Methods("POST")
	router.HandleFunc("/employees/{id}", employee.UpdateEmployeeHandler(employeeStore)).Methods("PUT")
	router.HandleFunc("/employees/{id}", employee.DeleteEmployeeHandler(employeeStore)).Methods("DELETE")

	return router
}

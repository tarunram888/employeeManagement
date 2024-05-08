package employee

import (
	"errors"
	"sync"
)

// Employee struct represents an employee entity
type Employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}

// EmployeeStore is an in-memory es for employees
type EmployeeStore struct {
	sync.RWMutex
	employees []Employee
}

// NewEmployeeStore creates a new instance of EmployeeStore
func NewEmployeeStore() *EmployeeStore {
	return &EmployeeStore{
		employees: []Employee{},
	}
}

// Custom error for employee not found
var ErrEmployeeNotFound = errors.New("employee not found")

// CreateEmployee adds a new employee to the es
func (es *EmployeeStore) CreateEmployee(emp Employee) {
	es.Lock()
	defer es.Unlock()
	es.employees = append(es.employees, emp)
}

// ListEmployees returns a list of employees with pagination
func (es *EmployeeStore) ListEmployees(page, pageSize int) []Employee {
	es.RLock()
	defer es.RUnlock()
	startIdx := (page - 1) * pageSize
	endIdx := startIdx + pageSize
	if endIdx > len(es.employees) {
		endIdx = len(es.employees)
	}
	return es.employees[startIdx:endIdx]
}

// GetEmployeeByID retrieves an employee from the es by ID
func (es *EmployeeStore) GetEmployeeByID(id int) (Employee, error) {
	es.RLock()
	defer es.RUnlock()
	for _, emp := range es.employees {
		if emp.ID == id {
			return emp, nil
		}
	}
	return Employee{}, ErrEmployeeNotFound
}

// UpdateEmployee updates the details of an existing employee
func (es *EmployeeStore) UpdateEmployee(id int, newEmp Employee) bool {
	es.Lock()
	defer es.Unlock()
	for i, emp := range es.employees {
		if emp.ID == id {
			es.employees[i] = newEmp
			return true
		}
	}
	return false
}

// DeleteEmployee deletes an employee from the es by ID
func (es *EmployeeStore) DeleteEmployee(id int) bool {
	es.Lock()
	defer es.Unlock()
	for i, emp := range es.employees {
		if emp.ID == id {
			// Remove the employee from the slice
			es.employees = append(es.employees[:i], es.employees[i+1:]...)
			return true
		}
	}
	return false
}

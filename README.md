# Employee Management

Employee Management is a RESTful API for managing employee records. It provides endpoints for CRUD operations on employee data and supports pagination for listing employees.

## Features

- **Create, Read, Update, and Delete** employee records
- **List employees** with support for pagination


## Getting Started

### 1. Clone the Repository:
git clone https://github.com/tarunram888/employeeManagement.git


### 2. Install Dependencies:
- cd employeeManagement
- go mod download


### 3. Run the Application:
go run main.go


## API Endpoints

- **GET /api/v1/employees**: List all employees with pagination support.
- **GET /api/v1/employees/{id}**: Get an employee by ID.
- **POST /api/v1/employees**: Create a new employee.
- **PUT /api/v1/employees/{id}**: Update an existing employee.
- **DELETE /api/v1/employees/{id}**: Delete an employee by ID.





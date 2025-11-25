# Zarish-HRA (Human Resources Administration Module)

HR management module for ZarishSphere Platform.

## Features

- Employee Management
- Department tracking
- Contact information management

## API Endpoints

### Employees
- `POST /api/v1/employees` - Create employee
- `GET /api/v1/employees/:id` - Get employee by ID
- `GET /api/v1/employees` - List employees (paginated)

## Running

```bash
go run cmd/server/main.go
```

Server runs on port **8082**.

## Models

- **Employee**: Employee records with personal and work information

## Technology Stack

- Go 1.21+
- Gin Web Framework
- GORM ORM
- PostgreSQL

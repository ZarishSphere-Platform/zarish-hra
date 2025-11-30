# Zarish-HRA (Human Resources Administration Module)

HR management module for ZarishSphere Platform, managing employee records, contracts, and organizational structure.

## Features

- **Employee Management**: Comprehensive employee profiles.
- **Department Tracking**: Organizational hierarchy and department management.
- **Contract Management**: Employment contracts and terms.
- **Leave Management**: Tracking employee leave and attendance (planned).

## Architecture

The system follows the standard ZarishSphere architecture:

- **API Layer**: RESTful endpoints using Gin.
- **Service Layer**: Business logic for HR processes.
- **Repository Layer**: Data persistence with GORM.
- **Models**: Database schemas for HR data.

## API Endpoints

### Employees

- `POST /api/v1/employees` - Create employee
- `GET /api/v1/employees/:id` - Get employee by ID
- `GET /api/v1/employees` - List employees (paginated)

### Departments

- `GET /api/v1/departments` - List departments
- `POST /api/v1/departments` - Create department

## Running Locally

### Prerequisites

- Go 1.23+
- PostgreSQL 15+

### Steps

1. **Configure Database**: Ensure PostgreSQL is running.
2. **Run Server**:
   ```bash
   go run cmd/server/main.go
   ```
   Server runs on port **8082**.

## Docker Support

Build and run using Docker Compose:

```bash
docker-compose up --build zarish-hra
```

## Technology Stack

- **Language**: Go 1.23
- **Framework**: Gin
- **ORM**: GORM
- **Database**: PostgreSQL

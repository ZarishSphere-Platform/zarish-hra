package repository

import (
	"github.com/zarishsphere/zarish-hra/internal/models"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) CreateEmployee(employee *models.Employee) error {
	return r.db.Create(employee).Error
}

func (r *EmployeeRepository) GetEmployeeByID(id uint) (*models.Employee, error) {
	var employee models.Employee
	err := r.db.First(&employee, id).Error
	return &employee, err
}

func (r *EmployeeRepository) UpdateEmployee(employee *models.Employee) error {
	return r.db.Save(employee).Error
}

func (r *EmployeeRepository) ListEmployees(offset, limit int) ([]models.Employee, error) {
	var employees []models.Employee
	err := r.db.Offset(offset).Limit(limit).Find(&employees).Error
	return employees, err
}

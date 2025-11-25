package service

import (
	"github.com/zarishsphere/zarish-hra/internal/models"
	"github.com/zarishsphere/zarish-hra/internal/repository"
)

type EmployeeService struct {
	repo *repository.EmployeeRepository
}

func NewEmployeeService(repo *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) CreateEmployee(employee *models.Employee) error {
	return s.repo.CreateEmployee(employee)
}

func (s *EmployeeService) GetEmployee(id uint) (*models.Employee, error) {
	return s.repo.GetEmployeeByID(id)
}

func (s *EmployeeService) ListEmployees(page, pageSize int) ([]models.Employee, error) {
	offset := (page - 1) * pageSize
	return s.repo.ListEmployees(offset, pageSize)
}

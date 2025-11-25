package models

import (
	"time"
)

// Employee represents an employee record.
// Adapted from Odoo hr.employee model.
type Employee struct {
	BaseModel

	Name        string `gorm:"size:255;not null;index" json:"name"`
	JobTitle    string `gorm:"size:255" json:"job_title"`
	WorkEmail   string `gorm:"size:255;index" json:"work_email"`
	WorkPhone   string `gorm:"size:50" json:"work_phone"`
	MobilePhone string `gorm:"size:50" json:"mobile_phone"`

	DepartmentID *uint `gorm:"index" json:"department_id,omitempty"`
	CompanyID    uint  `gorm:"index" json:"company_id"`
	ParentID     *uint `gorm:"index" json:"parent_id,omitempty"` // Manager
	UserID       *uint `gorm:"index" json:"user_id,omitempty"`   // Related User

	Active   bool       `gorm:"default:true" json:"active"`
	Gender   string     `gorm:"size:20" json:"gender"`
	Birthday *time.Time `json:"birthday,omitempty"`

	// Address
	PrivateAddress   string `gorm:"type:text" json:"private_address"`
	EmergencyContact string `gorm:"size:255" json:"emergency_contact"`
	EmergencyPhone   string `gorm:"size:50" json:"emergency_phone"`
}

// TableName overrides the table name used by User to `employees`
func (Employee) TableName() string {
	return "employees"
}

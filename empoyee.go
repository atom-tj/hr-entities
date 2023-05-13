package entities

import (
	"time"

	"github.com/google/uuid"
)

type Employee struct {
	ID uuid.UUID `json:"id,omitempty" gorm:"primaryKey"`
	LastName   string  `json:"last_name"`   // Фамилия сотрудника
	FirstName  string  `json:"first_name"`  // Имя сотрудника
	MiddleName *string `json:"middle_name"` // Отчество сотрудника (если имеется)

	DepartmentID uuid.UUID `json:"department_id"`

	StartedAt time.Time `json:"started_at"` // Дата и время начала работы в компании
}

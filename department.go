package entities

import (
	"github.com/google/uuid"
)

type Department struct {
	ID        uuid.UUID `json:"id,omitempty" gorm:"primaryKey"` // ID подразделения
	Title     string    `json:"title"`                          // Полное наименование подразделения
	ShortName string    `json:"short_name,omitempty"`           // Краткое наименование подразделения
}

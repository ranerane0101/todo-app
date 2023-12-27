package entity

import (
	"time"

	"github.com/ranerane0101/domain/valueobject"
)

type Todo struct {
	TodoID    valueobject.TodoID `gorm:"primaryKey"`
	Text      string             `gorm:"not null"`
	Done      bool               `gorm:"default:false"`
	StartDate time.Time
	EndDate   time.Time
}

package schemas

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Model struct {
	ID        uuid.UUID  `json:"id" gorm:"primaryKey" gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt *time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"not null"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}

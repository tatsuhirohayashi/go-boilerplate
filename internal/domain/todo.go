package domain

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID    uuid.UUID `json:"user_id" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	Title     string    `json:"title"`
	Content   *string   `json:"content"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	User      User      `gorm:"foreignKey:UserID"`
}

func (Todo) TableName() string {
	return "todos"
}
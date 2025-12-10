package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
	Todos     []Todo     `gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "users"
}

package models

import "github.com/google/uuid"

type User struct {
	ID         uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt  string    `db:"created_at" json:"created_at"`
	UpdatedAt  string    `db:"updated_at" json:"updated_at"`
	Email      string    `db:"email" json:"email" validate:"required,email,lte=255"`
	Password   string    `db:"password" json:"password_hash,omitempty" validate:"required,lte=255"`
	UserStatus string    `db:"user_status" json:"user_status" validate:"required,len=1"`
	UserRole   string    `db:"user_role" json:"user_role" validate:"required,lte=25"`
}

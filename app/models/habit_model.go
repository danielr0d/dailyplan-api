package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"time"
)

type Habit struct {
	ID         uuid.UUID  `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at" json:"updated_at"`
	UserID     uuid.UUID  `db:"user_id" json:"user_id" validate:"required,uuid"`
	Title      string     `db:"title" json:"title" validate:"required, lte=255"`
	IsDone     bool       `db:"is_done" json:"is_done"`
	StartAt    time.Time  `db:"start_at" json:"start_at" validate:"required"`
	EndAt      time.Time  `db:"end_at" json:"end_at" validate:"required"`
	HabitAttrs HabitAttrs `db:"habit_attrs" json:"habit_attrs" validate:"required,dive"`
}

type HabitAttrs struct {
	Picture     string `json:"picture"`
	Description string `json:"description"`
}

func (b HabitAttrs) Value() (driver.Value, error) {
	return json.Marshal(b)
}

func (b *HabitAttrs) Scan(value interface{}) error {
	j, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(j, &b)
}

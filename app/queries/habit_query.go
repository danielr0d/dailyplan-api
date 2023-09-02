package queries

import (
	"github.com/danielr0d/dailyplan-api/app/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type HabitQueries struct {
	*sqlx.DB
}

func (q *HabitQueries) GetHabits() ([]models.Habit, error) {
	habits := []models.Habit{}

	query := `SELECT * FROM habits`

	err := q.Select(&habits, query)
	if err != nil {
		return habits, err
	}

	return habits, nil
}

func (q *HabitQueries) GetHabitsByTitle(title string) ([]models.Habit, error) {
	habits := []models.Habit{}

	query := `SELECT * FROM habits WHERE title = $1`

	err := q.Get(&habits, query, title)
	if err != nil {
		return habits, err
	}

	return habits, nil
}

func (q *HabitQueries) GetHabit(id uuid.UUID) (models.Habit, error) {
	habit := models.Habit{}

	query := `SELECT * FROM habits WHERE id = $1`

	err := q.Get(&habit, query, id)
	if err != nil {
		return habit, err
	}

	return habit, nil
}

func (q *HabitQueries) CreateHabit(h *models.Habit) error {
	query := `INSERT INTO habits VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := q.Exec(query, h.ID, h.CreatedAt, h.UpdatedAt, h.UserID, h.Title, h.IsDone, h.StartAt, h.EndAt, h.HabitAttrs)

	if err != nil {
		return err
	}

	return nil
}

func (q *HabitQueries) UpdateHabit(id uuid.UUID, h *models.Habit) error {
	query := `UPDATE habits SET title = $1, is_done = $2, start_at = $3, end_at = $4, habit_attrs = $5 WHERE id = $6`

	_, err := q.Exec(query, h.Title, h.IsDone, h.StartAt, h.EndAt, h.HabitAttrs, id)

	if err != nil {
		return err
	}

	return nil
}

func (q *HabitQueries) DeleteHabit(id uuid.UUID) error {
	query := `DELETE FROM habits WHERE id = $1`

	_, err := q.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

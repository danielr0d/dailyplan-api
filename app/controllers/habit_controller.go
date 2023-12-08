package controllers

import (
	"github.com/create-go-app/fiber-go-template/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetHabits func gets all exists Habits.
// @Description Get all exists Habits.
// @Summary get all exists Habits
// @Tags Habits
// @Accept json
// @Produce json
// @Success 200 {array} models.Habit
// @Router /v1/Habits [get]
func GetHabits(c *fiber.Ctx) error {
	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get all Habits.
	habits, err := db.GetHabits()
	if err != nil {
		// Return, if Habits not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  true,
			"msg":    "Habits were not found",
			"count":  0,
			"Habits": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"count":  len(habits),
		"habits": habits,
	})
}

// GetHabit func gets Habit by given ID or 404 error.
// @Description Get Habit by given ID.
// @Summary get Habit by given ID
// @Tags Habit
// @Accept json
// @Produce json
// @Param id path string true "Habit ID"
// @Success 200 {object} models.Habit
// @Router /v1/Habit/{id} [get]
func GetHabit(c *fiber.Ctx) error {
	// Catch Habit ID from URL.
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get Habit by ID.
	habit, err := db.GetHabit(id)
	if err != nil {
		// Return, if Habit not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "Habit with the given ID is not found",
			"habit": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"habit": habit,
	})
}

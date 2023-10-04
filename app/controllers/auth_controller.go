package controllers

import (
	"github.com/danielr0d/dailyplan-api/app/models"
	"github.com/danielr0d/dailyplan-api/pkg/utils"
	"github.com/danielr0d/dailyplan-api/platform/database"
	"github.com/google/uuid"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UserSingUp(c *fiber.Ctx) error {

	singUp := &models.SingUp{}

	if err := c.BodyParser(singUp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	validate := utils.NewValidator()

	if err := validate.Struct(singUp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	role, err := utils.VerifyRole(singUp.UserRole)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		}
	}

	user := &models.User{}

	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.Email = singUp.Email
	user.PasswordHash = singUp.Password
	user.UserStatus = "active"
	user.UserRole = role

	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		}
	}

	if err := db.CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		}
	}

	user.PasswordHash = ""

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "User created successfully",
		"user": user,
	})
}

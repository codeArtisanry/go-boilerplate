package v1

import (
	"database/sql"

	"github.com/codeArtisanry/go-boilerplate/config"
	"github.com/codeArtisanry/go-boilerplate/models"
	"go.uber.org/zap"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	model  *models.Queries
	logger *zap.Logger
	cfg    config.AppConfig
}

func NewUserController(db *sql.DB, logger *zap.Logger, cfg config.AppConfig) (*UserController, error) {
	userModel := models.New(db)
	return &UserController{
		model:  userModel,
		logger: logger,
		cfg:    cfg,
	}, nil
}

func (u *UserController) Create(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	if user.Name == "" || user.Password == "" || user.Email == "" {
		return c.Status(400).JSON(fiber.Map{"error": "name, password and email are required"})
	}

	// Retrieve the last inserted ID
	lastID, err := u.model.GetLastId(c.Context())
	if err != nil {
		// Handle case where no users exist yet
		if err == sql.ErrNoRows {
			lastID = 0
		} else {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
	}

	// Increment the ID for the new user
	newID := lastID + 1

	// Create the new user
	_, err = u.model.CreateUser(c.Context(), models.CreateUserParams{
		ID:       newID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error while creating the user": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "user created successfully", "id": newID})
}

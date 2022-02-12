package api

import (
	"log"

	"github.com/den19980107/go-fiber-gorm-starter/config"
	"github.com/den19980107/go-fiber-gorm-starter/db"
	"github.com/den19980107/go-fiber-gorm-starter/db/entity"
	"github.com/den19980107/go-fiber-gorm-starter/db/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"golang.org/x/crypto/bcrypt"
)

func ValidRigisterRequest(c *fiber.Ctx) error {
	var data entity.UserRegisterDto

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	v := validate.Struct(data)
	if !v.Validate() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": v.Errors,
		})
	}

	repo := repository.New(db.ORM)
	record := repo.GetByUsername(data.Username)

	if record != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": "User already exists",
		})
	}

	var user entity.User

	user.Username = data.Username
	user.PasswordHash = CreatePasswordHash(data.Password)

	c.Locals("user", &user)
	c.Locals("repository", repo)

	return c.Next()
}

func Register(c *fiber.Ctx) error {
	user := c.Locals("user").(*entity.User)
	repo := c.Locals("repository").(repository.UserRepositoryInterface)

	repo.Create(user)

	token := user.CreateJWTToken(config.App.JWT.Secret)

	c.Append("X-Access-Token", token.Hash)

	return c.JSON(fiber.Map{
		"data": user,
	})
}

func ValidLoginRequest(c *fiber.Ctx) error {
	var data entity.UserLoginDto

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	v := validate.Struct(data)

	if !v.Validate() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": v.Errors,
		})
	}

	repo := repository.New(db.ORM)
	user := repo.GetByUsername(data.Username)

	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not found.",
		})
	}

	match := ComparePasswordHashIsMatch(data.Password, user.PasswordHash)

	if !match {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password.",
		})
	}

	c.Locals("user", user)

	return c.Next()
}

func Login(c *fiber.Ctx) error {
	user := c.Locals("user").(*entity.User)
	token := user.CreateJWTToken(config.App.JWT.Secret)

	c.Append("X-Access-Token", token.Hash)

	return c.JSON(fiber.Map{
		"data": user,
	})
}

// Create the password hash.
func CreatePasswordHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	return string(hash)
}

func ComparePasswordHashIsMatch(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

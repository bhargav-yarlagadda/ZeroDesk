package main

import (
	"log"
	"zerodesk/database"
	"zerodesk/middleware"
	"zerodesk/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	// "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Connect DB
	database.ConnectToDB()
	log.Println("✅ Database connected")

	// Init fiber
	app := fiber.New()

	// Root route
	app.Get("/", middleware.ValidateJWT(), func(c *fiber.Ctx) error {
		log.Println("📢 Root route hit")

		// Retrieve claims (map[string]interface{})
		claims := c.Locals("user").(jwt.MapClaims)

		// Extract values from claims
		username := claims["username"].(string)
		userID := claims["user_id"]

		return c.JSON(fiber.Map{
			"message":  "Fiber + Neon + GORM connected 🚀",
			"user_id":  userID,
			"username": username,
		})
	})

	// Auth routes
	authRouter := app.Group("/auth")
	routers.AuthRoutes(authRouter)

	// Start server
	log.Println("🚀 Server starting on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal("❌ Failed to start server: ", err)
	}
}

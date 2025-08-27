package main

import (
	"log"
	"zerodesk/database"
	"zerodesk/middleware"
	"zerodesk/routers"
	signallingserver "zerodesk/signalling_server"
	wsserver "zerodesk/ws_server"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
	"github.com/golang-jwt/jwt/v5"
	// "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Connect DB
	database.ConnectToDB()
	log.Println("✅ Database connected")

	// Init fiber
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*", // frontend origin
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin, Content-Type, Authorization",
		AllowCredentials: false, // important if using cookies
	}))
	// ICE server info
	app.Get("/api/ice-servers", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"iceServers": []map[string]interface{}{
				{"urls": "stun:localhost:3478"}, // replace localhost with your public IP/domain for external access
			},
		})
	})
	// Signalling Server
	app.Get("/ws/signaling", middleware.ValidateJWT(), websocket.New(signallingserver.SignalingHandler))
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

	// websocker server to handle keystrokes and mouse movement transmission.
	sessionRouter := app.Group("/session")
	sessionRouter.Use(middleware.ValidateJWT())
	wsserver.SessionRoutes(sessionRouter)

	// Start server
	log.Println("🚀 Server starting on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal("❌ Failed to start server: ", err)
	}
}

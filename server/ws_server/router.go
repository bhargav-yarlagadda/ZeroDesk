package wsserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)


func SessionRoutes(router fiber.Router) {
	router.Get("/:sessionId/:role",websocket.New(SessionHandler))
}
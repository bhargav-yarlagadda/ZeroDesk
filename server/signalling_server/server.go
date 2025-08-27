package signallingserver

import (
	"encoding/json"
	"log"

	"github.com/gofiber/websocket/v2"
)

// Keep track of connected clients per session
var sessions = make(map[string]map[*websocket.Conn]bool)

type SignalMessage struct {
	Type      string `json:"type"`       // "offer", "answer", "candidate"
	SessionID string `json:"sessionId"`  // session identifier
	Payload   string `json:"payload"`    // SDP or ICE candidate
}

func SignalingHandler(c *websocket.Conn) {
	defer c.Close()

	var sessionID string

	// Listen for messages from client
	for {
		_, msgBytes, err := c.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}

		// Decode message
		msg := SignalMessage{}
		if err := json.Unmarshal(msgBytes, &msg); err != nil {
			log.Println("JSON unmarshal error:", err)
			continue
		}

		sessionID = msg.SessionID

		// Register client in session map
		if _, ok := sessions[sessionID]; !ok {
			sessions[sessionID] = make(map[*websocket.Conn]bool)
		}
		sessions[sessionID][c] = true

		// Broadcast message to all other clients in the same session
		for client := range sessions[sessionID] {
			if client != c {
				if err := client.WriteMessage(websocket.TextMessage, msgBytes); err != nil {
					log.Println("Write error:", err)
					client.Close()
					delete(sessions[sessionID], client)
				}
			}
		}
	}

	// Cleanup on disconnect
	if sessionID != "" {
		delete(sessions[sessionID], c)
		if len(sessions[sessionID]) == 0 {
			delete(sessions, sessionID)
		}
	}
}

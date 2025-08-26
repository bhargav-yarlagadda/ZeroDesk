package wsserver

import (
	"encoding/json"
	"log"
	"time"
	"zerodesk/database"

	"github.com/gofiber/websocket/v2"
	"github.com/golang-jwt/jwt/v5"
)

func SessionHandler(conn *websocket.Conn) {
	sessionId := conn.Params("sessionId")
	role := conn.Params("role")

	claims, ok := conn.Locals("user").(jwt.MapClaims)
	if !ok {
		log.Println("no claims found")
		_ = conn.WriteMessage(websocket.TextMessage, []byte("unauthorized"))
		_ = conn.Close()
		return
	}
	userID := claims["user_id"].(string)
	
	client := &Client{
		UserID: userID,
		Role:   role,
		Conn:   conn,
	}

	// ✅ Lock session map
	sessionsMu.Lock()
	sess, ok := Sessions[sessionId]
	if !ok {
		sess = &Session{}
		Sessions[sessionId] = sess
	}

	// ✅ Ensure max 1 host + 1 viewer
	if role == "host" {
		if sess.Host != nil {
			sessionsMu.Unlock()
			conn.WriteMessage(websocket.TextMessage, []byte("host already exists"))
			conn.Close()
			return
		}
		sess.Host = client
	} else if role == "viewer" {
		if sess.Viewer != nil {
			sessionsMu.Unlock()
			conn.WriteMessage(websocket.TextMessage, []byte("viewer already exists"))
			conn.Close()
			return
		}
		sess.Viewer = client
	}
	sessionsMu.Unlock()

	log.Printf("[%s] joined session %s\n", role, sessionId)

	// ✅ If both joined → log session start
	if sess.Host != nil && sess.Viewer != nil && sess.LogID == "" {
		sessionLog := database.SessionLog{
			ViewerID:  sess.Viewer.UserID,
			HostID:    sess.Host.UserID,
			StartTime: time.Now(),
			Status:    "active",
			IPAddress: conn.RemoteAddr().String(),
		}
		if err := database.DB.Create(&sessionLog).Error; err != nil {
			log.Println("DB insert error:", err)
		} else {
			sess.LogID = sessionLog.ID
			log.Println("Session started:", sess.LogID)
		}
	}

	// ✅ Cleanup when handler exits
	defer func() {
		log.Printf("Client %s (%s) disconnected", userID, role)
		conn.Close()

		sessionsMu.Lock()
		if sess, ok := Sessions[sessionId]; ok {
			if role == "host" && sess.Host == client {
				// Host leaving → mark session end
				if sess.LogID != "" {
					end := time.Now()
					if err := database.DB.Model(&database.SessionLog{}).
						Where("id = ?", sess.LogID).
						Updates(map[string]any{
							"end_time": end,
							"status":   "ended",
						}).Error; err != nil {
						log.Println("DB update error:", err)
					} else {
						log.Println("Session ended:", sess.LogID)
					}
				}
				sess.Host = nil
			}
			if role == "viewer" && sess.Viewer == client {
				sess.Viewer = nil
			}
			// Remove session if empty
			if sess.Host == nil && sess.Viewer == nil {
				delete(Sessions, sessionId)
			}
		}
		sessionsMu.Unlock()
	}()

	// ✅ Listen for messages
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}

		if role == "viewer" {
			sessionsMu.RLock()
			target := sess.Host
			sessionsMu.RUnlock()

			if target != nil {
				var event InputEvent
				if err := json.Unmarshal(msg, &event); err != nil {
					log.Println("bad event:", err)
					continue
				}
				if err := target.Conn.WriteJSON(event); err != nil {
					log.Println("forward error:", err)
				}
			}
		}
	}
}

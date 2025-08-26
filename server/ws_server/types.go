package wsserver

import (
	"sync"

	"github.com/gofiber/websocket/v2"
)
type Client struct {
	UserID string
	Role   string // host or viewer
	Conn   *websocket.Conn //connection
}
type Session struct {
	Host   *Client
	Viewer *Client
	LogID  string

}
type InputEvent struct {
	Type      string `json:"type"`       // "keyboard" | "mouse"
	Timestamp int64  `json:"timestamp"`  // for ordering (unix ms)

	// Keyboard
	Key   string `json:"key,omitempty"`   // e.g. "a", "Enter"
	Ctrl  bool   `json:"ctrl,omitempty"`
	Alt   bool   `json:"alt,omitempty"`
	Shift bool   `json:"shift,omitempty"`
	Meta  bool   `json:"meta,omitempty"`

	// Mouse
	X      int    `json:"x,omitempty"`      // X coordinate (relative to screen/canvas)
	Y      int    `json:"y,omitempty"`      // Y coordinate
	Button string `json:"button,omitempty"` // "left" | "right" | "middle"
	Action string `json:"action,omitempty"` // "move" | "down" | "up" | "click" | "dblclick"
}

var (
	Sessions   = make(map[string]*Session)
	sessionsMu sync.RWMutex
	
)
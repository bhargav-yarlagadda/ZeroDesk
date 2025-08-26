package database 


import (
	"time"
)
type User struct{
	ID string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Username string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:now()"` 
	LastLogin time.Time 
}
// SessionLog = logs of viewer-host connections
type SessionLog struct {
    ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    ViewerID  string    `gorm:"type:uuid;not null"`
    HostID    string    `gorm:"type:uuid;not null"`
    StartTime time.Time `gorm:"default:now()"`
    EndTime   *time.Time
    Status    string    `gorm:"default:'active'"`
    IPAddress string
}
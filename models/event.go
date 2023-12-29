package models

import (
	"time"

	"github.com/PrinceNarteh/go-events-api/mongorm"
)

type Event struct {
	mongorm.Model
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int       `json:"user_id"`
}

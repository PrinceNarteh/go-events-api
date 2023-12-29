package models

import (
	"time"

	"github.com/PrinceNarteh/go-events-api/mongorm"
)

type Event struct {
	mongorm.Model `bson:",inline"`
	Name          string    `json:"name" validate:"required"`
	Description   string    `json:"description" validate:"required"`
	Location      string    `json:"location" validate:"required"`
	DateTime      time.Time `json:"date_time" validate:"required"`
	UserID        int       `json:"user_id"`
}

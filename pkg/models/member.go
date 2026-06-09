package models

import (
	"time"

	"github.com/google/uuid"
)

type Membership struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index"`
	CircleID     uuid.UUID `gorm:"type:uuid;not null;index"`
	Roles     []Role
	CreatedAt time.Time
}
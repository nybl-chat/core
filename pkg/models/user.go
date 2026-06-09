package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name      string
	AvatarURL string
    Email     string 
    
	CreatedAt time.Time
	UpdatedAt time.Time

	Devices     []Device     `gorm:"foreginKey:UserID"`
	Memberships []Membership `gorm:"foreignKey:UserID"`
}

type Device struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
    UserID    uuid.UUID `gorm:"type:uuid;not null;index"`

	Name      string
	PublicKey string

	CreatedAt time.Time
	UpdatedAt time.Time

}

type Circle struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name      string
	Slug      string `gorm:"uniqueIndex"`
	CreatedAt time.Time

	Memberships []Membership `gorm:"foreignKey:OrgID"`
}

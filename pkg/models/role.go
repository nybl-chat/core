package models

import (
    "github.com/google/uuid"
    "time"
)

type Permission string

const (
    // Members
    AddMembers       Permission = "members:add"
    InviteMembers    Permission = "members:invite"
    RemoveMembers    Permission = "members:remove"
    BlacklistMembers Permission = "members:blacklist"
    ViewMembers      Permission = "members:view"

    // Circle
    EditCircle Permission = "circle:edit"
	DeleteCircle Permission = "circle:delete"

    // Messages
    SendMessages     Permission = "messages:send"
    ModerateMessages Permission = "messages:moderate"

    // Roles
    EditRoles   Permission = "roles:edit"
    CreateRoles Permission = "roles:create"
    DeleteRoles Permission = "roles:delete"

    // Semicircle
    EditSemicircle   Permission = "semicircle:edit"
    DeleteSemicircle Permission = "semicircle:delete"
    CreateSemicircle Permission = "semicircle:create"

    // Moderation
    ManageTimeouts Permission = "timeouts:manage"
)

// Applied when circle is created
var RolePresets = map[string][]Permission{
    "owner": {
        AddMembers, InviteMembers, RemoveMembers, BlacklistMembers, ViewMembers,
        EditCircle,
        SendMessages, ModerateMessages,
        EditRoles, CreateRoles, DeleteRoles,
        EditSemicircle, DeleteSemicircle, CreateSemicircle,
        ManageTimeouts,
		DeleteCircle,
    },
    "moderator": {
        ViewMembers, RemoveMembers, BlacklistMembers,
        SendMessages, ModerateMessages,
        ManageTimeouts,
    },
    "member": {
        ViewMembers,
        SendMessages,
    },
}

type Role struct {
    ID          uuid.UUID    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
    Name        string       `gorm:"not null"`
    ColorHex    string
    Rank        int
    Hidden      bool         `gorm:"default:false"`
    Permissions []Permission `gorm:"serializer:json"`
    CreatedAt   time.Time
    UpdatedAt   time.Time

    Members []Membership `gorm:"foreignKey:RoleID"`
}

func (r *Role) Can(p Permission) bool {
    for _, perm := range r.Permissions {
        if perm == p {
            return true
        }
    }
    return false
}

func (r *Role) CanAll(perms ...Permission) bool {
    for _, p := range perms {
        if !r.Can(p) {
            return false
        }
    }
    return true
}
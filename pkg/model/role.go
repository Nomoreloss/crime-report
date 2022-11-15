package model

import "time"

// AccessRole represents access role type
type AccessRole int

const (
	// AdminRole has all permissions
	AdminRole AccessRole = 1

	// StaffRole is a standard user
	StaffRole AccessRole = 2

	// UserRole is a standard user
	UserRole AccessRole = 3

	// GuestRole can edit reseler specific things
	GuestRole AccessRole = 4
)

// Role model
type Role struct {
	Id          string     `json:"id,omitempty" bson:"_id"`
	AccessLevel AccessRole `json:"access_level"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at,omitempty"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty"`
}

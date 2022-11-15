package model

import "time"

type Crime struct {
	Id          string    `json:"id,omitempty" bson:"_id"`
	Title       string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Type        string    `json:"type,omitempty"`
	Media       string    `json:"media,omitempty"`
	Location    string    `json:"location,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

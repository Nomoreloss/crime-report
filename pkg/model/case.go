package model

import "time"

// FilterParams is a class contains the incoming query params for listing
type CaseFilterParams struct {
	Title    string `json:"name,omitempty"`
	Location string `json:"location,omitempty"`
	Reporter string `json:"reporter,omitempty"`
	Handler  string `json:"handler,omitempty"`
	Status   string `json:"status,omitempty"`
	Query    string `json:"query,omitempty"`
	Order    string `json:"order,omitempty"`
	Page     string `json:"page,omitempty"`
	Limit    string `json:"limit,omitempty"`
}

// FilterConfig contains the ordering and paging information for listing
type FilterConfig struct {
	Order     string
	OrderBy   string
	PageLimit int
	PageIndex int
}

type Case struct {
	Id          string    `json:"id,omitempty" bson:"_id"`
	Crime       string    `json:"crime,omitempty"`
	Reporter    string    `json:"reporter,omitempty"`
	Handler     string    `json:"handler,omitempty"`
	Status      string    `json:"status,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type CaseResponse struct {
	Id          string    `json:"id,omitempty" bson:"_id"`
	Crime       Crime     `json:"crime,omitempty"`
	Reporter    User      `json:"reporter,omitempty"`
	Handler     User      `json:"handler,omitempty"`
	Status      string    `json:"status,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

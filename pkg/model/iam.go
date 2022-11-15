package model

type Iam struct {
	Id     string  `json:"id,omitempty" bson:"_id"`
	Name   string  `json:"name,omitempty"`
	Email  string  `json:"email,omitempty"`
	Salary float64 `json:"salary,omitempty"`
}

type LoginRequest struct {
	Id     string  `json:"id,omitempty" bson:"_id"`
	Name   string  `json:"name,omitempty"`
	Email  string  `json:"email,omitempty"`
	Salary float64 `json:"salary,omitempty"`
}

type LoginResponse struct {
	Id     string  `json:"id,omitempty" bson:"_id"`
	Name   string  `json:"name,omitempty"`
	Email  string  `json:"email,omitempty"`
	Salary float64 `json:"salary,omitempty"`
}

type RegisterRequest struct {
	Id     string  `json:"id,omitempty" bson:"_id"`
	Name   string  `json:"name,omitempty"`
	Email  string  `json:"email,omitempty"`
	Salary float64 `json:"salary,omitempty"`
}

type RegisterResponse struct {
	Id     string  `json:"id,omitempty" bson:"_id"`
	Name   string  `json:"name,omitempty"`
	Email  string  `json:"email,omitempty"`
	Salary float64 `json:"salary,omitempty"`
}

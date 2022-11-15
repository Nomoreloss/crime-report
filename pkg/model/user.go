package model

import "time"

type UserFilter struct {
	Id        string    `json:"id,omitempty" bson:"_id"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Salary    float64   `json:"salary,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type User struct {
	Id        string    `json:"id,omitempty" bson:"_id"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	Active    bool      `json:"active,omitempty"`
	Status    string    `json:"status,omitempty"`
	Token     string    `json:"token,omitempty"`
	Mobile    string    `json:"mobile,omitempty"`
	About     string    `json:"about,omitempty"`
	Address   string    `json:"address,omitempty"`
	UserType  string    `json:"user_type,omitempty"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Timestamp time.Time

func (t *Timestamp) UnmarshalParam(src string) error {
	ts, err := time.Parse(time.RFC3339, src)
	*t = Timestamp(ts)
	return err
}

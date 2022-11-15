package model

type Permission struct {
	Id        string `json:"id,omitempty" bson:"_id"`
	Group     string `json:"group"`
	Name      string `json:"name"`
	Read      string `json:"read"`
	Add       string `json:"add"`
	Edit      string `json:"edit"`
	Remove    string `json:"remove"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

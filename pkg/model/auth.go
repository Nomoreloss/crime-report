package model

import (
	"github.com/labstack/echo"
)

// AuthToken holds authentication token details with refresh token
type AuthToken struct {
	Token   string `json:"token"`
	Expires string `json:"expires"`
	// Role         AccessRole `json:"role"`
	Role         string `json:"role"`
	Type         string `json:"type"`
	RefreshToken string `json:"refresh_token"`
	User         *User  `json:"user"`
}

// type AuthUser struct {
// 	ID          string     `json:"id,omitempty" bson:"_id"`
// 	Name        string     `json:"name"`
// 	Username    string     `json:"username"`
// 	Email       string     `json:"email"`
// 	Mobile      string     `json:"mobile,omitempty"`
// 	Address     string     `json:"address,omitempty"`
// 	RoleID      AccessRole `json:"-"`
// 	ResellerID  string     `json:"reseller_id"`
// 	ActivatedAt int64      `json:"activated_at"`
// 	CreatedAt   int64      `json:"created_at"`
// }

// RegisterReq represents data stored in JWT token for user
type RegisterReq struct {
	Username         string `json:"username"`
	Email            string `json:"email"`
	Name             string `json:"name"`
	ActivationMedium string `json:"activationMedium"`
	Mobile           string `json:"mobile"`
	ResendType       string `json:"resendType"`
}

// ActivateAcc represents data stored in JWT token for user
type ActivateAcc struct {
	Username       string `json:"username"`
	ActivationCode string `json:"activationCode"`
	Password       string `json:"password"`
}

// RefreshToken holds authentication token details
type RefreshToken struct {
	Token   string `json:"token"`
	Expires string `json:"expires"`
}

// Check holds response status details
type Check struct {
	Body string `json:"body"`
}

// AuthUser represents data stored in JWT token for user
type AuthUser struct {
	ID         string
	ResellerID string
	Username   string
	Email      string
	Role       AccessRole
}

// RBACService represents role-based access control service interface
type RBACService interface {
	User(echo.Context) *AuthUser
	EnforceRole(echo.Context, AccessRole) error
	EnforceUser(echo.Context, int) error
	EnforceReseller(echo.Context, int) error
	AccountCreate(echo.Context, AccessRole, int, int) error
	IsLowerRole(echo.Context, AccessRole) error
}

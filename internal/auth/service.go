package auth

import (
	"database/sql"

	"github.com/ellipizle/crime-report/internal/auth/platform/postgresql"
	"github.com/ellipizle/crime-report/pkg/model"
	"github.com/labstack/echo/v4"
)

// New creates new iam service
func New(udb UserDB, j TokenGenerator, sec Securer, rbac RBAC) *Auth {
	return &Auth{
		udb:  udb,
		tg:   j,
		sec:  sec,
		rbac: rbac,
	}
}

// Initialize initializes auth application service
func Initialize(db *sql.DB, j TokenGenerator, sec Securer, rbac RBAC) *Auth {
	return New(postgresql.New(db), j, sec, rbac)
}

// Service represents auth service interface
type Service interface {
	Authenticate(echo.Context, string, string) (*model.AuthToken, error)
	Register(echo.Context, *model.RegisterReq) (string, error)
	ResendCode(echo.Context, *model.RegisterReq) error
	ChangePassword(echo.Context, string, string, string) error
	Reset(echo.Context, string, string, string) error
	ForgotPassword(echo.Context, string) (string, error)
	FindByUsername(echo.Context, string) (*model.User, error)
	FindByEmail(echo.Context, string) (*model.User, error)
	Me(echo.Context) (*model.User, error)
}

// Auth represents auth application service
type Auth struct {
	udb  UserDB
	tg   TokenGenerator
	sec  Securer
	rbac RBAC
}

// UserDB represents user repository interface
type UserDB interface {
	View(echo.Context, string) (*model.User, error)
	Register(echo.Context, *model.User) (*model.User, error)
	FindByUsername(echo.Context, string) (*model.User, error)
	FindByEmail(echo.Context, string) (*model.User, error)
	FindByMobile(echo.Context, string) (*model.User, error)
	Update(echo.Context, *model.User) error
	Activate(echo.Context, *model.User) error
}

// TokenGenerator represents token generator (jwt) interface
type TokenGenerator interface {
	GenerateToken(*model.User) (string, string, error)
}

// Securer represents security interface
type Securer interface {
	HashMatchesPassword(string, string) bool
	Token(string) string
	Hash(string) string
	GenerateCode() string
}

// RBAC represents role-based-access-control interface
type RBAC interface {
	User(echo.Context) *model.AuthUser
}

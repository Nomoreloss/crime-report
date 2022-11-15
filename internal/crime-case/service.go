package crime_case

import (
	"database/sql"

	"github.com/ellipizle/crime-report/internal/crime-case/platform/postgresql"
	"github.com/ellipizle/crime-report/pkg/model"
	"github.com/labstack/echo/v4"
)

// Service represents user application interface
type Service interface {
	Create(echo.Context, *model.Case) (*model.Case, error)
	List(echo.Context, *model.CaseFilterParams) ([]*model.CaseResponse, error)
	View(echo.Context, string) (*model.CaseResponse, error)
	Delete(echo.Context, string) error
	Status(echo.Context, string, bool) error
	Update(echo.Context, *model.Case) (*model.Case, error)
	Count(echo.Context, *model.CaseFilterParams) int
}

// New creates new user application service
func New(db *sql.DB, udb UDB, rbac RBAC) *Case {
	return &Case{db: db, udb: udb, rbac: rbac}
}

// Initialize initalizes Case application service with defaults
func Initialize(db *sql.DB, rbac RBAC) *Case {
	return New(db, postgresql.New(db), rbac)
}

// Case represents user application service
type Case struct {
	db   *sql.DB
	udb  UDB
	rbac RBAC
}

// UDB represents user repository interface
type UDB interface {
	Create(echo.Context, *model.Case) (*model.Case, error)
	View(echo.Context, string) (*model.CaseResponse, error)
	List(echo.Context, *model.CaseFilterParams) ([]*model.CaseResponse, error)
	Update(echo.Context, *model.Case) (*model.Case, error)
	Delete(echo.Context, string) error
	Count(echo.Context, *model.CaseFilterParams) int
}

// RBAC represents role-based-access-control interface
type RBAC interface {
	User(echo.Context) *model.AuthUser
	EnforceRole(echo.Context, model.AccessRole) error
	IsLowerRole(echo.Context, model.AccessRole) error
}

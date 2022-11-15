package rbac

import (
	"github.com/ellipizle/crime-report/pkg/model"
	"github.com/labstack/echo/v4"
)

// New creates new RBAC service
func New() *Service {
	return &Service{}
}

// Service is RBAC application service
type Service struct{}

func checkBool(b bool) error {
	if b {
		return nil
	}
	return echo.ErrForbidden
}

// User returns user data stored in jwt token
func (s *Service) User(c echo.Context) *model.AuthUser {
	id := c.Get("id").(string)
	user := c.Get("username").(string)
	email := c.Get("email").(string)
	role := c.Get("role").(model.AccessRole)
	return &model.AuthUser{
		ID:       id,
		Username: user,
		Email:    email,
		Role:     role,
	}
}

// EnforceRole authorizes request by AccessRole
func (s *Service) EnforceRole(c echo.Context, r model.AccessRole) error {
	return checkBool(!(c.Get("role").(model.AccessRole) > r))
}

// EnforceUser checks whether the request to change user data is done by the same user
func (s *Service) EnforceUser(c echo.Context, ID string) error {
	// TODO: Implement querying db and checking the requested user's company_id/location_id
	// to allow company/location admins to view the user
	if s.isAdmin(c) {
		return nil
	}
	return checkBool(c.Get("id").(string) == ID)
}

// EnforceCompany checks whether the request to apply change to company data
// is done by the user belonging to the that company and that the user has role CompanyAdmin.
// If user has admin role, the check for company doesnt need to pass.
func (s *Service) EnforceReseller(c echo.Context, ID string) error {
	if s.isAdmin(c) {
		return nil
	}
	if err := s.EnforceRole(c, model.AdminRole); err != nil {
		return err
	}
	return checkBool(c.Get("reseller_id").(string) == ID)
}

func (s *Service) isAdmin(c echo.Context) bool {
	return !(c.Get("role").(model.AccessRole) > model.AdminRole)
}

// IsLowerRole checks whether the requesting user has higher role than the user it wants to change
// Used for account creation/deletion
func (s *Service) IsLowerRole(c echo.Context, r model.AccessRole) error {
	return checkBool(c.Get("role").(model.AccessRole) < r)
}

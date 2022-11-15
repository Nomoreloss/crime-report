// Package user contains user application services
package crime_case

import (
	"net/http"

	"github.com/ellipizle/crime-report/pkg/model"
	"github.com/ellipizle/crime-report/pkg/structs"
	"github.com/labstack/echo/v4"
)

// Custom errors
var (
	ErrInvalidCredentials  = echo.NewHTTPError(http.StatusNotFound, "Username or password does not exist")
	ErrInvalidEmail        = echo.NewHTTPError(http.StatusBadRequest, "Email address is not valid")
	ErrEmailExist          = echo.NewHTTPError(http.StatusBadRequest, "Email already exist")
	ErrUsernameExist       = echo.NewHTTPError(http.StatusBadRequest, "Username ready exist")
	ErrMobileExist         = echo.NewHTTPError(http.StatusBadRequest, "Mobile ready exist")
	ErrInvalidMobile       = echo.NewHTTPError(http.StatusBadRequest, "Mobile number is not valid")
	ErrActivationCode      = echo.NewHTTPError(http.StatusBadRequest, "Activation Code did not match")
	ErrResetCode           = echo.NewHTTPError(http.StatusBadRequest, "Reset Code did not match")
	ErrInvalidMobileString = echo.NewHTTPError(http.StatusBadRequest, "Mobile number is not valid")
	// ErrInvalidCredentials = echo.NewHTTPError(http.StatusNotFound, "Username or email does not exist")
)

// Create creates a new user account
func (u *Case) Create(ctx echo.Context, usr *model.Case) (*model.Case, error) {
	return u.udb.Create(ctx, usr)
}

// List returns list of users
func (u *Case) List(ctx echo.Context, p *model.CaseFilterParams) ([]*model.CaseResponse, error) {
	return u.udb.List(ctx, p)
}

// Count returns total of transactions
func (u *Case) Count(ctx echo.Context, p *model.CaseFilterParams) int {
	return u.udb.Count(ctx, p)
}

// View returns single user
func (u *Case) View(ctx echo.Context, id string) (*model.CaseResponse, error) {
	// if err := u.rbac.EnforceUser(ctx, id); err != nil {
	// 	return nil, err
	// }
	return u.udb.View(ctx, id)
}

// Delete deletes a user
func (u *Case) Delete(ctx echo.Context, id string) error {
	// if err := u.rbac.IsLowerRole(ctx, user.Role.AccessLevel); err != nil {
	// 	return err
	// }
	return u.udb.Delete(ctx, id)
}

// Update updates user's contact information
func (u *Case) Update(ctx echo.Context, req *model.Case) (*model.Case, error) {
	// if err := u.rbac.EnforceUser(ctx, req.ID); err != nil {
	// 	return nil, err
	// }

	singleCase, err := u.udb.View(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	caseStruct := &model.Case{
		Id:          singleCase.Id,
		Crime:       singleCase.Crime.Id,
		Reporter:    singleCase.Reporter.Id,
		Handler:     singleCase.Handler.Id,
		Status:      singleCase.Status,
		Description: singleCase.Description,
		CreatedAt:   singleCase.CreatedAt,
		UpdatedAt:   singleCase.UpdatedAt,
	}

	structs.Merge(caseStruct, req)
	if _, err := u.udb.Update(ctx, caseStruct); err != nil {
		return nil, err
	}

	return caseStruct, nil
}

// Status change user's status information
func (u *Case) Status(ctx echo.Context, id string, status bool) error {
	// if err := u.rbac.EnforceUser(ctx, req.ID); err != nil {
	// 	return nil, err
	// }

	// user, err := u.udb.View(ctx, id)
	// if err != nil {
	// 	return err
	// }

	// user.Active = status

	// if err := u.udb.Update(ctx, user); err != nil {
	// 	return err
	// }

	return nil
}

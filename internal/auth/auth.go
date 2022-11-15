package auth

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	model "github.com/ellipizle/crime-report/pkg/model"
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
	ErrNotAvaliable        = echo.NewHTTPError(http.StatusNotFound, "Not available at the moment")
)

// Authenticate tries to authenticate the user provided by username and password
func (a *Auth) Authenticate(ctx echo.Context, user, pass string) (*model.AuthToken, error) {
	u, err := a.udb.FindByUsername(ctx, strings.ToLower(user))
	if err != nil {
		ErrInvalidCredentials.Message = "Username or password does not exist"
		return nil, ErrInvalidCredentials
	}

	if !a.sec.HashMatchesPassword(u.Password, pass) {
		return nil, ErrInvalidCredentials
	}

	if !u.Active {
		return nil, model.ErrUnauthorized
	}

	token, expire, err := a.tg.GenerateToken(u)
	if err != nil {
		return nil, model.ErrUnauthorized
	}

	if err := a.udb.Update(ctx, u); err != nil {
		return nil, err
	}

	return &model.AuthToken{Token: token, Expires: expire, RefreshToken: u.Token, Role: u.Role, Type: u.UserType, User: u}, nil
}

// Register  refreshes jwt token and puts new claims inside
func (a *Auth) Register(ctx echo.Context, usr *model.RegisterReq) (string, error) {
	// return " ", ErrNotAvaliable
	usr.Email = strings.ToLower(usr.Email)
	usr.Username = strings.ToLower(usr.Username)
	// check for unique email
	if usr.Email != "" {
		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)+[.]*$")
		if !re.MatchString(usr.Email) {
			return "", ErrInvalidEmail
		}
		user, _ := a.udb.FindByEmail(ctx, usr.Email)
		if user != nil {
			return "", ErrEmailExist
		}
	}
	if _, err := strconv.Atoi(usr.Mobile); err != nil {
		return "", ErrInvalidMobileString
	}

	if usr.Mobile != "" {
		if strings.HasPrefix(usr.Mobile, "0") {
			if len(usr.Mobile) != 11 {
				return "", ErrInvalidMobile
			}
		}
		if strings.HasPrefix(usr.Mobile, "+") {
			if len(usr.Mobile) != 14 {
				return "", ErrInvalidMobile
			}
		}
	}

	// check for unique username
	user, _ := a.udb.FindByUsername(ctx, usr.Username)
	if user != nil {
		return "", ErrUsernameExist
	}

	// check for unique mobile
	user, _ = a.udb.FindByMobile(ctx, usr.Mobile)
	if user != nil {
		return "", ErrMobileExist
	}

	// create for unique mobile
	var userRe = new(model.User)
	userRe.Email = usr.Email
	userRe.Username = usr.Username
	_, err := a.udb.Register(ctx, userRe)
	if err != nil {
		return "", err
	}

	return "success", nil
}

// Reset creates a new user account
func (a *Auth) Reset(ctx echo.Context, username, token, password string) error {
	// username = strings.ToLower(username)
	// user, err := a.udb.FindByCode(ctx, username, token)
	// if err != nil {
	// 	return ErrResetCode
	// }
	// user.Password = a.sec.Hash(password)
	// return a.udb.Update(ctx, user)
	return nil
}

// ChangePassword tries to authenticate the user provided by username and password
func (a *Auth) ChangePassword(ctx echo.Context, user, old, pass string) error {
	u, err := a.udb.FindByUsername(ctx, user)
	if err != nil {
		ErrInvalidCredentials.Message = "Username or password does not exist"
		return ErrInvalidCredentials
	}

	if !a.sec.HashMatchesPassword(u.Password, old) {
		return ErrInvalidCredentials
	}

	u.Password = a.sec.Hash(pass)
	if err := a.udb.Update(ctx, u); err != nil {
		return err
	}

	return nil
}

// ResendCode creates a new user account
func (a *Auth) ResendCode(ctx echo.Context, activeAcc *model.RegisterReq) error {
	_, err := a.udb.FindByUsername(ctx, activeAcc.Username)
	if err != nil {
		return ErrInvalidCredentials
	}
	return nil
}

// Forgot  refreshes jwt token and puts new claims inside
func (a *Auth) ForgotPassword(ctx echo.Context, usr string) (string, error) {
	fmt.Print(usr)
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if re.MatchString(usr) {
		user, err := a.udb.FindByEmail(ctx, usr)
		if err != nil {
			return "", ErrInvalidEmail
		}
		err = a.udb.Update(ctx, user)
		if err != nil {
			return "", err
		}
		return user.Username, nil
	}

	if strings.HasPrefix(usr, "0") {
		if len(usr) != 11 {
			return "", ErrInvalidMobile
		}
	}
	if strings.HasPrefix(usr, "+") {
		if len(usr) != 14 {
			return "", ErrInvalidMobile
		}
	}

	user, err := a.udb.FindByMobile(ctx, usr)
	if err != nil {
		return "", err
	}
	err = a.udb.Update(ctx, user)
	if err != nil {
		return "", err
	}
	return user.Username, nil

}

// FindByUsername returns info about currently logged user
func (a *Auth) FindByUsername(ctx echo.Context, uname string) (*model.User, error) {
	return a.udb.FindByUsername(ctx, uname)
}

// FindByEmail returns info about currently logged user
func (a *Auth) FindByEmail(ctx echo.Context, email string) (*model.User, error) {
	return a.udb.FindByEmail(ctx, email)
}

// Me returns info about currently logged user
func (a *Auth) Me(ctx echo.Context) (*model.User, error) {
	au := a.rbac.User(ctx)
	return a.udb.View(ctx, au.ID)
}

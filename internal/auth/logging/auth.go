package auth

import (
	"time"

	"fmt"

	"github.com/ellipizle/crime-report/internal/auth"
	"github.com/ellipizle/crime-report/pkg/model"
	"github.com/labstack/echo/v4"
)

// New creates new auth logging service
func New(svc auth.Service, logger model.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents auth logging service
type LogService struct {
	auth.Service
	logger model.Logger
}

const name = "auth"

// Authenticate logging
func (ls *LogService) Authenticate(c echo.Context, user, password string) (resp *model.AuthToken, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Authenticate request", err,
			map[string]interface{}{
				"req":  user,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Authenticate(c, user, password)
}

// Register logging
func (ls *LogService) Register(c echo.Context, req *model.RegisterReq) (resp string, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Register request", err,
			map[string]interface{}{
				"req":  fmt.Sprintf("%s", req),
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Register(c, req)
}

// ResendCode logging
func (ls *LogService) ResendCode(c echo.Context, req *model.RegisterReq) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Resend Code request", err,
			map[string]interface{}{
				"req":  fmt.Sprintf("%s", req),
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ResendCode(c, req)
}

// ForgotPassword logging
func (ls *LogService) ForgotPassword(c echo.Context, req string) (resp string, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Forgot password request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ForgotPassword(c, req)
}

// Me logging
func (ls *LogService) Me(c echo.Context) (resp *model.User, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Me request", err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Me(c)
}

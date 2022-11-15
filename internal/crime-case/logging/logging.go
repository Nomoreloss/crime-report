package crime_case

import (
	"time"

	crime_case "github.com/ellipizle/crime-report/internal/crime-case"
	"github.com/ellipizle/crime-report/pkg/model"
	"github.com/labstack/echo/v4"
)

// New creates new crime_case logging service
func New(svc crime_case.Service, logger model.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents crime_case logging service
type LogService struct {
	crime_case.Service
	logger model.Logger
}

const name = "crime_case"

// Create logging
func (ls *LogService) Create(c echo.Context, req *model.Case) (resp *model.Case, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Create crime_case request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Create(c, req)
}

// List logging
func (ls *LogService) List(c echo.Context, req *model.CaseFilterParams) (resp []*model.CaseResponse, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "List crime case request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List(c, req)
}

// View logging
func (ls *LogService) View(c echo.Context, req string) (resp *model.CaseResponse, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "View crime_case request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.View(c, req)
}

// Delete logging
func (ls *LogService) Delete(c echo.Context, req string) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Delete crime case request", err,
			map[string]interface{}{
				"req":  req,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Delete(c, req)
}

// Status logging
func (ls *LogService) Status(c echo.Context, req string, status bool) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Status crime case request", err,
			map[string]interface{}{
				"req":  req,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Status(c, req, status)
}

// Update logging
func (ls *LogService) Update(c echo.Context, req *model.Case) (resp *model.Case, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Update crime case request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Update(c, req)
}

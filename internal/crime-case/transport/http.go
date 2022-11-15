package transport

import (
	"net/http"
	"strconv"

	crime_case "github.com/ellipizle/crime-report/internal/crime-case"
	"github.com/ellipizle/crime-report/pkg/model"
	"github.com/labstack/echo/v4"
)

// HTTP represents crime_case http service
type HTTP struct {
	svc crime_case.Service
}

// NewHTTP creates new crime_case http service
func NewHTTP(svc crime_case.Service, er *echo.Group) {
	h := HTTP{svc}
	ur := er.Group("/case")
	// AddUser func adds new crime case.
	//
	//	@Summary      Add new Case
	//	@Description  Add new Case
	//	@Tags         Crime Case
	//	@Accept       json
	//	@Produce      plain
	//	@Param        message  body      model.Case  true  "Account Info"
	//	@Success      200      {object} model.Case
	//	@Failure      500      {string}  string         "fail"
	//	@Router       /v1/case [post]
	ur.POST("", h.create)

	/// GetAllCase func gets all exists employee.
	//
	//		@Summary      Get all crime case
	//		@Description  Get all crime case
	//		@Tags         Crime Case
	//		@Accept       json
	//		@Produce      plain
	//		@Param   sort     query     string     false  "sort by"     desc
	//		@Param   filter     query     string     false  "filter string"     blue
	//	 	@Param   page     query     int     false  "page"     1
	//		@Param   limit     query     int     false  "limit"     10
	//		@Success     200        {array}    model.CaseResponse
	//		@Failure      400         {string}  string  "ok"
	//		@Failure      404         {string}  string  "ok"
	//		@Failure      500         {string}  string  "ok"
	//		@Router       /v1/case  [get]
	ur.GET("", h.list)

	// GetCrimeCase func gets all exists crime case.
	//
	//	@Summary      Get single crime case
	//	@Description  Get single crime case
	//	@Tags         Crime Case
	//	@Accept       json
	//	@Produce      plain
	//	@Param        case_id    path      int     true  "Group ID"
	//	@Success     200        {object}    model.CaseResponse
	//	@Failure      400         {string}  string  "ok"
	//	@Failure      404         {string}  string  "ok"
	//	@Failure      500         {string}  string  "ok"
	//	@Router       /v1/case/{case_id}  [get]
	ur.GET("/:id", h.view)

	// UpdateCase func gets all exists employee.
	//
	//	@Summary      Update crime case request
	//	@Description  Update crime case request
	//	@Tags         Crime Case
	//	@Accept       json
	//	@Produce      plain
	//	@Param        case_id    path      int     true  "Case ID"
	//	@Param        message  body      model.CaseResponse  true  "Crime Case"
	//	@Success      200      {string}  string         "success"
	//	@Failure      500      {string}  string         "fail"
	//	@Router       /v1/case/{case_id} [put]
	ur.PATCH("/:id", h.update)

	//	@Summary      Delete crime case account
	//	@Description  Delete crime case account
	//	@Tags         Crime Case
	//	@Accept       json
	//	@Produce      plain
	//	@Param        case_id    path      string     true  "User ID"
	//	@Success      200         {string}  string  "answer"
	//	@Failure      400         {string}  string  "ok"
	//	@Failure      404         {string}  string  "ok"
	//	@Failure      500         {string}  string  "ok"
	//	@Router       /v1/case/{case_id} [delete]
	ur.DELETE("/:id", h.delete)

}

// Custom errors
var (
	ErrPasswordsNotMaching = echo.NewHTTPError(http.StatusBadRequest, "passwords do not match")
)

// User create request
// swagger:model caseCreate
type createReq struct {
	Crime       string `json:"crime" validate:"required"`
	Reporter    string `json:"reporter" validate:"required"`
	Handler     string `json:"handler" validate:"required"`
	Status      string `json:"status" validate:"required"`
	Description string `json:"description"`
}

func (h *HTTP) create(c echo.Context) error {
	r := new(createReq)

	if err := c.Bind(r); err != nil {
		return err
	}

	usr, err := h.svc.Create(c, &model.Case{
		Crime:       r.Crime,
		Reporter:    r.Reporter,
		Handler:     r.Handler,
		Status:      r.Status,
		Description: r.Description,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, usr)
}

type listResponse struct {
	Cases []*model.CaseResponse `json:"cases"`
	Count int                   `json:"count"`
}

func (h *HTTP) list(c echo.Context) error {
	p := new(model.CaseFilterParams)
	if err := c.Bind(p); err != nil {
		return err
	}

	result, err := h.svc.List(c, p)
	count := h.svc.Count(c, p)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, listResponse{result, count})
}

func (h *HTTP) view(c echo.Context) error {
	id := string(c.Param("id"))
	if id == "" {
		return model.ErrBadRequest
	}

	result, err := h.svc.View(c, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

// User update request
// swagger:model caseUpdate
type updateReq struct {
	ID          string `json:"-"`
	Crime       string `json:"crime" validate:"required"`
	Reporter    string `json:"reporter" validate:"required"`
	Handler     string `json:"handler" validate:"required"`
	Status      string `json:"status" validate:"required"`
	Description string `json:"description"`
}

func (h *HTTP) update(c echo.Context) error {
	id := string(c.Param("id"))
	if id == "" {
		return model.ErrBadRequest
	}

	r := new(updateReq)
	if err := c.Bind(r); err != nil {
		return err
	}

	usr, err := h.svc.Update(c, &model.Case{
		Id:          id,
		Crime:       r.Crime,
		Reporter:    r.Reporter,
		Handler:     r.Handler,
		Status:      r.Status,
		Description: r.Description,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, usr)
}

func (h *HTTP) delete(c echo.Context) error {
	id := string(c.Param("id"))
	if id == "" {
		return model.ErrBadRequest
	}

	if err := h.svc.Delete(c, id); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (h *HTTP) status(c echo.Context) error {
	id := string(c.Param("id"))
	if id == "" {
		return model.ErrBadRequest
	}
	status, err := strconv.ParseBool(c.QueryParam("status"))
	if err != nil {
		return model.ErrBadRequest
	}

	if err := h.svc.Status(c, id, status); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

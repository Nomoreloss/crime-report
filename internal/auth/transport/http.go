package transport

import (
	"net/http"

	"github.com/ellipizle/crime-report/internal/auth"
	"github.com/ellipizle/crime-report/pkg/model"

	"github.com/labstack/echo/v4"
)

// HTTP represents auth http service
type HTTP struct {
	svc auth.Service
}

// NewHTTP creates new auth http service
func NewHTTP(svc auth.Service, e *echo.Echo, mw echo.MiddlewareFunc) {
	h := HTTP{svc}
	// swagger:route POST /login auth login
	// Logs in user by username and password.
	// responses:
	//  200: loginResp
	//  400: errMsg
	//  401: errMsg
	// 	403: err
	//  404: errMsg
	//  500: err
	e.POST("/login", h.login)
	// swagger:operation GET /check-username/{username} check username
	// ---
	// summary: Check username.
	// description: Check username by checking at database whether username exists.
	// parameters:
	// - name: username
	//   in: path
	//   description: check username
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/checkUsernameResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	e.GET("/check-username/:username", h.checkUserName)

	// swagger:operation GET /check-email/{email} check email
	// ---
	// summary: Check email.
	// description: Check email by checking at database whether email exists.
	// parameters:
	// - name: email
	//   in: path
	//   description: check email
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/checkEmaiResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	e.GET("/check-email/:email", h.checkEmail)

	// swagger:route GET /me auth meReq
	// Gets user's info from session.
	// responses:
	//  200: userResp
	//  500: err
	e.GET("/me", h.me, mw)
	e.POST("/register", h.register)
	e.POST("/resend-code", h.resendCode)
	e.POST("/change-password", h.changePassword)
	e.GET("/forgot-password", h.forgotPassword)
}

type credentials struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type cPassword struct {
	OldPassword string `json:"oldpassword" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Username    string `json:"username" validate:"required"`
}

func (h *HTTP) changePassword(c echo.Context) error {
	cred := new(cPassword)
	if err := c.Bind(cred); err != nil {
		return err
	}
	err := h.svc.ChangePassword(c, cred.Username, cred.OldPassword, cred.Password)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, cred)
}

func (h *HTTP) login(c echo.Context) error {
	cred := new(credentials)
	if err := c.Bind(cred); err != nil {
		return err
	}
	r, err := h.svc.Authenticate(c, cred.Username, cred.Password)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, r)
}

func (h *HTTP) me(c echo.Context) error {
	user, err := h.svc.Me(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (h *HTTP) register(c echo.Context) error {
	cred := new(model.RegisterReq)
	if err := c.Bind(cred); err != nil {
		return err
	}
	user, err := h.svc.Register(c, cred)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (h *HTTP) resendCode(c echo.Context) error {
	cred := new(model.RegisterReq)
	if err := c.Bind(cred); err != nil {
		return err
	}
	if err := h.svc.ResendCode(c, cred); err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func (h *HTTP) forgotPassword(c echo.Context) error {
	r, err := h.svc.ForgotPassword(c, c.QueryParam("username"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, r)
}

func (h *HTTP) checkUserName(c echo.Context) error {
	_, err := h.svc.FindByUsername(c, c.Param("username"))
	if err != nil {
		return c.JSON(http.StatusFound, "{status:false}")
	}
	return c.JSON(http.StatusOK, "{status:true}")
}

func (h *HTTP) checkEmail(c echo.Context) error {
	user, err := h.svc.FindByEmail(c, c.Param("email"))
	if err != nil {
		return c.JSON(http.StatusFound, user)
	}
	return c.JSON(http.StatusOK, user)
}

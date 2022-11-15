package transport

import model "github.com/ellipizle/crime-report/pkg/model"

// Login request
// swagger:parameters login
type swaggLoginReq struct {
	// in:body
	// Body credentials
}

// Login response
// swagger:response loginResp
type swaggLoginResp struct {
	// in:body
	Body struct {
		*model.AuthToken
	}
}

// Token refresh response
// swagger:response refreshResp
type swaggRefreshResp struct {
	// in:body
	Body struct {
		*model.RefreshToken
	}
}

// check username response
// swagger:response checkUsernameResp
type swaggCheckUsernameResp struct {
	// in:body
	Body struct {
		*model.RefreshToken
	}
}

// check email response
// swagger:response checkEmailResp
type swaggCheckEmailResp struct {
	// in:body
	Body struct {
		*model.RefreshToken
	}
}

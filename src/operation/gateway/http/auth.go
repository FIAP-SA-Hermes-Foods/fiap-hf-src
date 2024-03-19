package http

import (
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/interfaces"
)

var _ interfaces.UserAuth = (*userAuthGateway)(nil)

type userAuthGateway struct {
	auth interfaces.UserAuth
}

func NewUserAuthGateway(auth interfaces.UserAuth) *userAuthGateway {
	return &userAuthGateway{auth: auth}
}

func (u *userAuthGateway) Auth(input dto.UserInput) (*dto.UserOutput, error) {
	return u.auth.Auth(input)
}

package businesses

import (
	"github.com/TulioGuaraldoB/lemarchand-password/core/dtos/requests"
)

type IUserBusiness interface {
	VerifyPassword(passwordRequest *requests.PasswordRequest)
}

type userBusiness struct {
}

func NewUserBusiness() IUserBusiness {
	return &userBusiness{}
}

func (b *userBusiness) VerifyPassword(passwordRequest *requests.PasswordRequest) {

}

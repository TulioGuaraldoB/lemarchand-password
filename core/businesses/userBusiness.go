package businesses

import (
	"fmt"
	"log"
	"unicode"

	"github.com/TulioGuaraldoB/lemarchand-password/constants"
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
	password := passwordRequest.Password

	for _, passwordRule := range passwordRequest.Rules {
		switch passwordRule.Rule {
		case constants.MinimumCharacters:
			checkPasswordLength(password, passwordRule.Value)

		case constants.MinimumUppercaseCharacters:
			checkUppercase(password, passwordRule.Value)

		case constants.MinimumLowercaseCharacters:

		case constants.MinimumSpecialCharacters:

		case constants.NoRepeated:
		}
	}
}

func checkPasswordLength(password string, minimumSize int64) bool {
	if int64(len(password)) != minimumSize {
		return false
	}

	return true
}

func checkUppercase(password string, minimumSize int64) {
	characters := []rune{}
	for _, character := range password {
		if unicode.IsUpper(character) {
			characters = append(characters, character)
		}
	}

	charactersLength := len(characters)
	if charactersLength != int(minimumSize) {
		errMessage := fmt.Sprintf("Wrong size of characters! password: %v. characters: %v", passwordLength, len(characters))
		log.Fatal(errMessage)
	}
}

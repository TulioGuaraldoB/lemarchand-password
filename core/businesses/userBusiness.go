package businesses

import (
	"unicode"

	"github.com/TulioGuaraldoB/lemarchand-password/constants"
	"github.com/TulioGuaraldoB/lemarchand-password/core/dtos/requests"
	"github.com/TulioGuaraldoB/lemarchand-password/core/dtos/responses"
	"github.com/TulioGuaraldoB/lemarchand-password/util/formatter"
)

type IUserBusiness interface {
	VerifyPassword(passwordRequest *requests.PasswordRequest) *responses.PasswordResponse
}

type userBusiness struct {
}

func NewUserBusiness() IUserBusiness {
	return &userBusiness{}
}

func (b *userBusiness) VerifyPassword(passwordRequest *requests.PasswordRequest) *responses.PasswordResponse {
	password := passwordRequest.Password
	verificationRes := new(responses.PasswordResponse)
	noMatched := []interface{}{}

	for _, passwordRule := range passwordRequest.Rules {
		switch passwordRule.Rule {
		case constants.MinimumCharacters:
			verificationRes.Verification = checkPasswordLength(password, passwordRule.Value)
			verificationRes.NoMatched = passwordRule.Rule
			noMatched = append(noMatched, verificationRes)

		case constants.MinimumUppercaseCharacters:
			verificationRes.Verification = checkUppercase(password, passwordRule.Value)
			verificationRes.NoMatched = passwordRule.Rule
			noMatched = append(noMatched, verificationRes)

		case constants.MinimumLowercaseCharacters:
			verificationRes.Verification = checkLowercase(password, passwordRule.Value)
			verificationRes.NoMatched = passwordRule.Rule
			noMatched = append(noMatched, verificationRes)

		case constants.MinimumSpecialCharacters:
			verificationRes.Verification = checkSpecialCharacters(password, passwordRule.Value)
			verificationRes.NoMatched = passwordRule.Rule
			noMatched = append(noMatched, verificationRes)

		case constants.MinimumDigits:
			verificationRes.Verification = checkDigits(password, passwordRule.Value)
			verificationRes.NoMatched = passwordRule.Rule
			noMatched = append(noMatched, verificationRes)

		case constants.NoRepeated:
		}
	}

	return verificationRes
}

func checkPasswordLength(password string, minimumSize int64) bool {
	return len(password) != int(minimumSize)
}

func checkUppercase(password string, minimumSize int64) bool {
	characters := []rune{}
	for _, character := range password {
		if unicode.IsUpper(character) {
			characters = append(characters, character)
		}
	}

	return compareCharactersToSize(characters, minimumSize)
}

func checkLowercase(password string, minimumSize int64) bool {
	characters := []rune{}
	for _, character := range password {
		if unicode.IsLower(character) {
			characters = append(characters, character)
		}
	}

	return compareCharactersToSize(characters, minimumSize)
}

func checkSpecialCharacters(password string, minimumSize int64) bool {
	newPassword := password
	newPassword = formatter.OnlySpecialCharacters(newPassword)
	passwordLength := len(newPassword)

	return passwordLength == int(minimumSize)
}

func compareCharactersToSize(characters []rune, minimumSize int64) bool {
	charactersLength := len(characters)
	return charactersLength == int(minimumSize)
}

func checkDigits(password string, minimumSize int64) bool {
	characters := []rune{}
	for _, character := range password {
		if unicode.IsDigit(character) {
			characters = append(characters, character)
		}
	}

	return compareCharactersToSize(characters, minimumSize)
}

package businesses

import (
	"fmt"
	"log"
	"regexp"
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
	verifiers := false
	fmt.Printf("%v", verifiers)

	for _, passwordRule := range passwordRequest.Rules {
		switch passwordRule.Rule {
		case constants.MinimumCharacters:
			verifiers = checkPasswordLength(password, passwordRule.Value)

		case constants.MinimumUppercaseCharacters:
			checkUppercase(password, passwordRule.Value)

		case constants.MinimumLowercaseCharacters:
			checkLowercase(password, passwordRule.Value)

		case constants.MinimumSpecialCharacters:
			checkSpecialCharacters(password, passwordRule.Value)

		case constants.NoRepeated:
		}
	}
}

func checkPasswordLength(password string, minimumSize int64) bool {
	return len(password) != int(minimumSize)
}

func checkUppercase(password string, minimumSize int64) {
	characters := []rune{}
	for _, character := range password {
		if unicode.IsUpper(character) {
			characters = append(characters, character)
		}
	}

	compareCharactersToSize(characters, minimumSize)
}

func checkLowercase(password string, minimumSize int64) {
	characters := []rune{}
	for _, character := range password {
		if unicode.IsLower(character) {
			characters = append(characters, character)
		}
	}

	compareCharactersToSize(characters, minimumSize)
}

func checkSpecialCharacters(password string, minimumSize int64) {
	newPassword := password
	newPassword = regexp.MustCompile(`[^!#@$%^&*()-+\\/{}]`).ReplaceAllString(password, "")
	passwordLength := len(newPassword)

	if passwordLength != int(minimumSize) {
		errMessage := fmt.Sprintf("Wrong size of characters! required: %v. characters: %v", minimumSize, passwordLength)
		log.Fatal(errMessage)
	}
}

func compareCharactersToSize(characters []rune, minimumSize int64) {
	charactersLength := len(characters)
	if charactersLength != int(minimumSize) {
		errMessage := fmt.Sprintf("Wrong size of characters! required: %v. characters: %v", minimumSize, len(characters))
		log.Fatal(errMessage)
	}
}

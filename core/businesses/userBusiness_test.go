package businesses_test

import (
	"testing"

	"github.com/TulioGuaraldoB/lemarchand-password/core/businesses"
	"github.com/TulioGuaraldoB/lemarchand-password/core/dtos/requests"
	faker "github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
)

type userBusinessTest struct {
	description              string
	expectedLower            bool
	expectedUpper            bool
	expectedNumeric          bool
	expectedSpecialCharacter bool
	expectedSpace            bool
	expectedNumber           int
	expectedRule             string
	expectedValue            int64
}

func mockPassword(lower, upper, numeric, special, space bool, num int) string {
	return faker.Password(lower, upper, numeric, special, space, num)
}

func TestVerifyPasswordBusiness(t *testing.T) {
	tests := []userBusinessTest{
		{
			description:              "Should return no error in validation",
			expectedLower:            true,
			expectedUpper:            true,
			expectedNumeric:          true,
			expectedSpecialCharacter: true,
			expectedSpace:            true,
			expectedNumber:           1,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Arrange
			mockPassword := faker.Password(
				testCase.expectedLower,
				testCase.expectedUpper,
				testCase.expectedNumeric,
				testCase.expectedSpecialCharacter,
				testCase.expectedSpace,
				testCase.expectedNumber,
			)

			mockPasswordRequest := new(requests.PasswordRequest)
			mockPasswordRequest.Password = mockPassword

			mockPasswordRule := requests.RuleRequest{}
			mockPasswordRule.Rule = testCase.expectedRule
			mockPasswordRule.Value = testCase.expectedValue

			mockPasswordRequest.Rules = append(mockPasswordRequest.Rules, mockPasswordRule)

			// Act
			userBusiness := businesses.NewUserBusiness()
			passwordRes := userBusiness.VerifyPassword(mockPasswordRequest)

			// Assert
			assert.NotNil(t, passwordRes)
		})
	}
}

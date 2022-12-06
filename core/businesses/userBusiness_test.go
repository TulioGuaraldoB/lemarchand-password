package businesses_test

import (
	"testing"

	"github.com/TulioGuaraldoB/lemarchand-password/constants"
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

func TestVerifyPasswordBusiness(t *testing.T) {
	tests := []userBusinessTest{
		{
			description:              "Should return success in character length validation",
			expectedLower:            true,
			expectedUpper:            true,
			expectedNumeric:          true,
			expectedSpecialCharacter: true,
			expectedSpace:            true,
			expectedNumber:           1,
			expectedRule:             constants.MinimumCharacters,
			expectedValue:            10,
		},
		{
			description:              "Should return success in UPPERCASE validation",
			expectedLower:            false,
			expectedUpper:            true,
			expectedNumeric:          false,
			expectedSpecialCharacter: false,
			expectedSpace:            false,
			expectedNumber:           0,
			expectedRule:             constants.MinimumUppercaseCharacters,
			expectedValue:            10,
		},
		{
			description:              "Should return success in lowercase validation",
			expectedLower:            true,
			expectedUpper:            false,
			expectedNumeric:          false,
			expectedSpecialCharacter: false,
			expectedSpace:            false,
			expectedNumber:           0,
			expectedRule:             constants.MinimumLowercaseCharacters,
			expectedValue:            10,
		},
		{
			description:              "Should return success in special characters validation",
			expectedLower:            true,
			expectedUpper:            false,
			expectedNumeric:          false,
			expectedSpecialCharacter: true,
			expectedSpace:            false,
			expectedNumber:           0,
			expectedRule:             constants.MinimumSpecialCharacters,
			expectedValue:            5,
		},
		{
			description:              "Should return success in digits validation",
			expectedLower:            true,
			expectedUpper:            false,
			expectedNumeric:          true,
			expectedSpecialCharacter: false,
			expectedSpace:            false,
			expectedNumber:           0,
			expectedRule:             constants.MinimumDigits,
			expectedValue:            2,
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

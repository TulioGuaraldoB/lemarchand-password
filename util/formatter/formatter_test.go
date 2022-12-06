package formatter_test

import (
	"testing"

	"github.com/TulioGuaraldoB/lemarchand-password/util/formatter"
	"github.com/stretchr/testify/assert"
)

type formatterTest struct {
	description string
	expectedStr string
}

func TestSpecialCharactersFormatter(t *testing.T) {
	tests := []formatterTest{
		{
			description: "should return success string on special characters formatter",
			expectedStr: "senha@!#$%",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			formattedStr := formatter.OnlySpecialCharacters(testCase.expectedStr)
			assert.NotEqual(t, testCase.expectedStr, formattedStr)
		})
	}
}

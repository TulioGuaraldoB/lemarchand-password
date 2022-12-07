package controllers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TulioGuaraldoB/lemarchand-password/constants"
	"github.com/TulioGuaraldoB/lemarchand-password/core/businesses/mock"
	"github.com/TulioGuaraldoB/lemarchand-password/core/controllers"
	"github.com/TulioGuaraldoB/lemarchand-password/core/dtos/requests"
	faker "github.com/brianvoe/gofakeit"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type userControllerTest struct {
	description             string
	expectedRule            string
	expectedPasswordRequest *requests.PasswordRequest
	setMocks                func(*mock.MockIUserBusiness)
	expectedStatusCode      int
	expectedBodyReader      io.Reader
}

func mockGinContext(httpMethod string, bodyReader io.Reader) (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(httpMethod, "/", bodyReader)
	ctx.Request = req
	return ctx, w
}

func mockPassword() *requests.PasswordRequest {
	mockPassword := new(requests.PasswordRequest)
	faker.Struct(&mockPassword)

	return mockPassword
}

func TestVerifyPasswordController(t *testing.T) {
	mockPassword := mockPassword()

	mockRule := new(requests.RuleRequest)
	mockRule.Rule = constants.MinimumUppercaseCharacters
	mockPassword.Rules = append(mockPassword.Rules, *mockRule)

	mockPasswordJSON, _ := json.Marshal(mockPassword)
	mockPasswordBuffer := bytes.NewBuffer(mockPasswordJSON)

	tests := []userControllerTest{
		{
			description:             "Should return StatusCode 200 (OK)",
			expectedRule:            constants.MinimumUppercaseCharacters,
			expectedPasswordRequest: mockPassword,
			setMocks: func(mib *mock.MockIUserBusiness) {
				mib.EXPECT().
					VerifyPassword(mockPassword).
					Return(nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedBodyReader: mockPasswordBuffer,
		},
		{
			description:             "Should return StatusCode 400 (Bad Request)",
			expectedRule:            "",
			expectedPasswordRequest: &requests.PasswordRequest{},
			setMocks:                func(mib *mock.MockIUserBusiness) {},
			expectedStatusCode:      http.StatusBadRequest,
			expectedBodyReader:      nil,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.description, func(t *testing.T) {
			// Arrange
			controller := gomock.NewController(t)
			defer controller.Finish()

			mmb := mock.NewMockIUserBusiness(controller)
			testCase.setMocks(mmb)

			ctx, recorder := mockGinContext(http.MethodPost, testCase.expectedBodyReader)

			// Act
			userController := controllers.NewUserController(mmb)
			userController.VerifyPassword(ctx)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, recorder.Result().StatusCode)
		})
	}
}

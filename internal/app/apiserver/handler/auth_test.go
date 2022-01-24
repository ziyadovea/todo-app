package handler

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/ziyadovea/todo-app/internal/app/service"
	"github.com/ziyadovea/todo-app/internal/app/service/mocks"
	"github.com/ziyadovea/todo-app/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_signUp(t *testing.T) {

	testCases := []struct {
		name                 string
		inputBody            string
		inputUser            *models.User
		mockBehaviour        func(s *mock_service.MockAuthorization, user *models.User)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Valid",
			inputBody: `{"name": "user", "username": "user", "password": "user"}`,
			inputUser: &models.User{
				Name:     "user",
				Username: "user",
				Password: "user",
			},
			mockBehaviour: func(s *mock_service.MockAuthorization, user *models.User) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   http.StatusCreated,
			expectedResponseBody: `{"id": 1}`,
		},
		{
			name:                 "Incorrect input body",
			inputBody:            `{"valid": "no"}`,
			inputUser:            &models.User{},
			mockBehaviour:        func(s *mock_service.MockAuthorization, user *models.User) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: fmt.Sprintf(`{"error_message": "%s"}`, errInvalidInputBody.Error()),
		},
		{
			name:      "Error creating user",
			inputBody: `{"name": "user", "username": "user", "password": "user"}`,
			inputUser: &models.User{
				Name:     "user",
				Username: "user",
				Password: "user",
			},
			mockBehaviour: func(s *mock_service.MockAuthorization, user *models.User) {
				s.EXPECT().CreateUser(user).Return(0, errors.New("error creating user"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: fmt.Sprintf(`{"error_message": "%s"}`, "error creating user"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockAuth := mock_service.NewMockAuthorization(mockCtrl)
			tc.mockBehaviour(mockAuth, tc.inputUser)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(tc.inputBody))

			handler := NewHandler(&service.Service{
				Authorization: mockAuth,
			})
			r := gin.New()
			r.POST("/sign-up", handler.signUp)
			r.ServeHTTP(rec, req)

			assert.EqualValues(t, tc.expectedStatusCode, rec.Code)
			assert.JSONEq(t, tc.expectedResponseBody, rec.Body.String())
		})
	}
}

func TestHandler_SignIn(t *testing.T) {

	testCases := []struct {
		name                 string
		inputBody            string
		input                *signInInput
		mockBehaviour        func(s *mock_service.MockAuthorization, input *signInInput)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Valid",
			inputBody: `{"username": "user", "password": "user"}`,
			input: &signInInput{
				Username: "user",
				Password: "user",
			},
			mockBehaviour: func(s *mock_service.MockAuthorization, input *signInInput) {
				s.EXPECT().GenerateToken(input.Username, input.Username).Return("token", nil)
			},
			expectedStatusCode:   http.StatusFound,
			expectedResponseBody: `{"token": "token"}`,
		},
		{
			name:                 "Bad request",
			inputBody:            `{"invalid": "yes"}`,
			input:                &signInInput{},
			mockBehaviour:        func(s *mock_service.MockAuthorization, input *signInInput) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: fmt.Sprintf(`{"error_message": "%s"}`, errInvalidInputBody.Error()),
		},
		{
			name:      "Error generate token",
			inputBody: `{"username": "user", "password": "user"}`,
			input: &signInInput{
				Username: "user",
				Password: "user",
			},
			mockBehaviour: func(s *mock_service.MockAuthorization, input *signInInput) {
				s.EXPECT().GenerateToken(input.Username, input.Username).Return("", errors.New("error generate token"))
			},
			expectedStatusCode:   http.StatusUnauthorized,
			expectedResponseBody: fmt.Sprintf(`{"error_message": "%s"}`, "error generate token"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockAuth := mock_service.NewMockAuthorization(mockCtrl)
			tc.mockBehaviour(mockAuth, tc.input)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-in", bytes.NewBufferString(tc.inputBody))

			handler := NewHandler(&service.Service{
				Authorization: mockAuth,
			})
			r := gin.New()
			r.POST("/sign-in", handler.signIn)
			r.ServeHTTP(rec, req)

			assert.EqualValues(t, tc.expectedStatusCode, rec.Code)
			assert.JSONEq(t, tc.expectedResponseBody, rec.Body.String())
		})
	}
}

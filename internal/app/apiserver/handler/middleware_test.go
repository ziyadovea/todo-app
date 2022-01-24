package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/ziyadovea/todo-app/internal/app/service"
	"github.com/ziyadovea/todo-app/internal/app/service/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_UserIdentity(t *testing.T) {

	testCases := []struct {
		name                 string
		headerName           string
		headerValue          string
		token                string
		mockBehaviour        func(s *mock_service.MockAuthorization, token string)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Valid",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehaviour: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"id": 1}`,
		},
		{
			name:                 "Empty header",
			headerName:           "Authorization",
			headerValue:          "",
			token:                "token",
			mockBehaviour:        func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:   http.StatusUnauthorized,
			expectedResponseBody: fmt.Sprintf(`{"error_message": "%s"}`, errEmptyAuthHeader.Error()),
		},
		{
			name:                 "All invalid header",
			headerName:           "Authorization",
			headerValue:          "Invalid",
			token:                "token",
			mockBehaviour:        func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:   http.StatusUnauthorized,
			expectedResponseBody: fmt.Sprintf(`{"error_message": "%s"}`, errInvalidAuthHeader.Error()),
		},
		{
			name:                 "Empty token",
			headerName:           "Authorization",
			headerValue:          "Bearer ",
			token:                "token",
			mockBehaviour:        func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:   http.StatusUnauthorized,
			expectedResponseBody: fmt.Sprintf(`{"error_message": "%s"}`, errInvalidAuthHeader.Error()),
		},
		{
			name:                 "Incorrect bearer",
			headerName:           "Authorization",
			headerValue:          "Bear token",
			token:                "token",
			mockBehaviour:        func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:   http.StatusUnauthorized,
			expectedResponseBody: fmt.Sprintf(`{"error_message": "%s"}`, errInvalidAuthHeader.Error()),
		},
		{
			name:        "Error parsing token",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehaviour: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(0, errors.New("error parsing token"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: fmt.Sprintf(`{"error_message": "%s"}`, "error parsing token"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockAuth := mock_service.NewMockAuthorization(mockCtrl)
			tc.mockBehaviour(mockAuth, tc.token)

			handler := NewHandler(&service.Service{
				Authorization: mockAuth,
			})
			r := gin.New()

			r.GET("/protected", handler.userIdentity, func(c *gin.Context) {
				id, _ := c.Get(ContextUserKey)
				c.JSON(http.StatusOK, map[string]interface{}{
					"id": id,
				})
			})

			rec := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/protected", nil)
			req.Header.Set(tc.headerName, tc.headerValue)

			r.ServeHTTP(rec, req)

			assert.EqualValues(t, tc.expectedStatusCode, rec.Code)
			assert.JSONEq(t, tc.expectedResponseBody, rec.Body.String())
		})
	}
}

func TestGetUserId(t *testing.T) {

	getContextWithID := func(id interface{}) *gin.Context {
		c := &gin.Context{}
		c.Set(ContextUserKey, id)
		return c
	}

	testCases := []struct {
		name    string
		ctx     *gin.Context
		isError bool
		id      int
	}{
		{
			name:    "Valid",
			ctx:     getContextWithID(1),
			isError: false,
			id:      1,
		},
		{
			name:    "Context without id",
			ctx:     &gin.Context{},
			isError: true,
			id:      0,
		},
		{
			name:    "ID not int",
			ctx:     getContextWithID("id"),
			isError: true,
			id:      0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			id, err := getUserID(tc.ctx)
			if tc.isError {
				assert.Error(t, err)
			} else {
				assert.EqualValues(t, tc.id, id)
			}
		})
	}
}

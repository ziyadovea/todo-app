package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	ContextUserKey = "userID"
)

// middleware for user auth
func (h *Handler) userIdentity(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		newErrorResponse(c, http.StatusUnauthorized, errEmptyAuthHeader.Error())
		return
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, errInvalidAuthHeader.Error())
		return
	}

	if authHeaderParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, errInvalidAuthHeader.Error())
		return
	}

	if authHeaderParts[1] == "" {
		newErrorResponse(c, http.StatusUnauthorized, errInvalidAuthHeader.Error())
		return
	}

	token := authHeaderParts[1]
	userID, err := h.Services.ParseToken(token)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Set(ContextUserKey, userID)
}

func getUserID(c *gin.Context) (int, error) {
	userID, isExist := c.Get(ContextUserKey)
	if !isExist {
		return 0, errors.New(errIdNotFoundInContext.Error())
	}

	userIDInt, ok := userID.(int)
	if !ok {
		return 0, errors.New(errIdType.Error())
	}

	return userIDInt, nil
}

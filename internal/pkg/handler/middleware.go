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
		newErrorResponse(c, http.StatusUnauthorized, "auth header is empty")
		return
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	token := authHeaderParts[1]
	userID, err := h.Services.ParseToken(token)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "invalid token")
		return
	}

	c.Set(ContextUserKey, userID)
}

func (h *Handler) getUserID(c *gin.Context) (int, error) {
	userID, isExist := c.Get(ContextUserKey)
	if !isExist {
		return 0, errors.New("id not found in context")
	}

	userIDInt, ok := userID.(int)
	if !ok {
		return 0, errors.New("incorrect id: id must be integer number")
	}
	return userIDInt, nil
}

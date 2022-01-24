package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ziyadovea/todo-app/models"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	user := &models.User{}
	if err := c.BindJSON(user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, errInvalidInputBody.Error())
		return
	}

	id, err := h.Services.CreateUser(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	input := &signInInput{}
	if err := c.BindJSON(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, errInvalidInputBody.Error())
		return
	}

	token, err := h.Services.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusFound, map[string]interface{}{
		"token": token,
	})
}

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ziyadovea/todo-app/models"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	user := &models.User{}
	if err := c.BindJSON(user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
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

func (h *Handler) signIn(c *gin.Context) {

}

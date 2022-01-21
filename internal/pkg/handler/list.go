package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getAllLists(c *gin.Context) {
	userID, isExist := c.Get(ContextUserKey)
	if !isExist {
		newErrorResponse(c, http.StatusInternalServerError, "context key not exist")
		return
	}
	c.JSON(http.StatusFound, map[string]interface{}{
		"userID": userID,
	})
}

func (h *Handler) createNewList(c *gin.Context) {

}

func (h *Handler) deleteAllLists(c *gin.Context) {

}

func (h *Handler) getListByID(c *gin.Context) {

}

func (h *Handler) updateListByID(c *gin.Context) {

}

func (h *Handler) deleteListByID(c *gin.Context) {

}

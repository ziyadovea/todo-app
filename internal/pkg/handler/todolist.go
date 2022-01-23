package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ziyadovea/todo-app/models"
	"net/http"
	"strconv"
)

func (h *Handler) createNewList(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	todolist := &models.TodoList{}
	if err := c.BindJSON(todolist); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	listID, err := h.Services.CreateNewList(userID, todolist)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"list_id": listID,
	})
}

func (h *Handler) getAllLists(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	lists, err := h.Services.GetAllLists(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusFound, lists)
}

func (h *Handler) deleteAllLists(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	rowsAffected, err := h.Services.DeleteAllLists(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"rows_deleted": rowsAffected,
	})
}

func (h *Handler) getListByID(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	listIDStr := c.Param("list_id")
	listID, err := strconv.Atoi(listIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	list, err := h.Services.GetListByID(userID, listID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusFound, list)
}

func (h *Handler) deleteListByID(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	listIDStr := c.Param("list_id")
	listID, err := strconv.Atoi(listIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	rowsAffected, err := h.Services.DeleteListByID(userID, listID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"rows_deleted": rowsAffected,
	})
}

func (h *Handler) updateListByID(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	listIDStr := c.Param("list_id")
	listID, err := strconv.Atoi(listIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newList := &models.NewTodoList{}
	if err = c.BindJSON(newList); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.Services.UpdateListByID(userID, listID, newList)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

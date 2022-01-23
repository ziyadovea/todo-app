package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ziyadovea/todo-app/models"
	"net/http"
	"strconv"
)

func (h *Handler) createNewItem(c *gin.Context) {
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

	item := &models.TodoItem{}
	if err = c.BindJSON(item); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Services.CreateNewItem(userID, listID, item)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"item_id": id,
	})
}

func (h *Handler) getAllItems(c *gin.Context) {
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

	items, err := h.Services.GetAllItems(userID, listID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusFound, items)
}

func (h *Handler) deleteAllItems(c *gin.Context) {
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

	rowsAffected, err := h.Services.DeleteAllItems(userID, listID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"rows_deleted": rowsAffected,
	})
}

func (h *Handler) getItemByID(c *gin.Context) {
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

	itemIDStr := c.Param("item_id")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	item, err := h.Services.GetItemByID(userID, listID, itemID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusFound, item)
}

func (h *Handler) updateItemByID(c *gin.Context) {
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

	itemIDStr := c.Param("item_id")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newItem := &models.NewTodoItem{}
	if err = c.BindJSON(newItem); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.Services.UpdateItemByID(userID, listID, itemID, newItem); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) deleteItemByID(c *gin.Context) {
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

	itemIDStr := c.Param("item_id")
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	rowsAffected, err := h.Services.DeleteItemByID(userID, listID, itemID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusFound, map[string]interface{}{
		"rows_deleted": rowsAffected,
	})
}

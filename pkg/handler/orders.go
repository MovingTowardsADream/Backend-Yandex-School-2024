package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"yandex-lavka/entity"
)

func (h *Handler) AddOrders(c *gin.Context) {
	// Getting and validation input
	var input entity.Orders

	if err := c.BindJSON(&input); err != nil {
		NewErrorMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.OrdersList.AddOrders(input)

	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) GetOrders(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 1
	}

	offsetStr := c.Query("offset")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	orders, err := h.service.OrdersList.GetOrders(entity.Parameters{offset, limit})

	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *Handler) GetOrdersById(c *gin.Context) {
	orderId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	order, err := h.service.OrdersList.GetOrdersById(orderId)

	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *Handler) CompleteTheOrder(c *gin.Context) {
	var input entity.Histories

	if err := c.BindJSON(&input); err != nil {
		NewErrorMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.service.OrdersList.CompleteTheOrder(input)

	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

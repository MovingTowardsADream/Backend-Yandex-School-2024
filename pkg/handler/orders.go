package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

	c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Successfully added",
	})
}

func (h *Handler) GetOrders(c *gin.Context) {

}

func (h *Handler) GetOrdersById(c *gin.Context) {

}

func (h *Handler) CompleteTheOrder(c *gin.Context) {

}

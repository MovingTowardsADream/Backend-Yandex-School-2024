package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yandex-lavka/entity"
)

func (h *Handler) AddCouriers(c *gin.Context) {
	// Getting and validation input
	var input entity.Couriers

	if err := c.BindJSON(&input); err != nil {
		NewErrorMessageResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.service.CouriersList.AddCouriers(input)

	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Successfully added",
	})
}

func (h *Handler) GetCouriers(c *gin.Context) {

}

func (h *Handler) GetCouriersById(c *gin.Context) {

}

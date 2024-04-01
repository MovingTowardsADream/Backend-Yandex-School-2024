package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) GetCouriers(c *gin.Context) {
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

	couriers, err := h.service.CouriersList.GetCouriers(entity.Parameters{offset, limit})

	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, couriers)
}

func (h *Handler) GetCouriersById(c *gin.Context) {
	courierId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	courier, err := h.service.CouriersList.GetCouriersById(courierId)

	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, courier)
}

func (h *Handler) GetMetaInfoById(c *gin.Context) {
	courierId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	startStr := c.Query("start_date")
	endStr := c.Query("end_date")

	period := entity.Period{startStr, endStr}

	courier_rating, err := h.service.CouriersList.GetMetaInfoById(courierId, period)

	if err != nil {
		NewErrorMessageResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, courier_rating)
}

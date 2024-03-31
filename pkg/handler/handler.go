package handler

import (
	"github.com/gin-gonic/gin"
	"yandex-lavka/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitHandler() *gin.Engine {
	router := gin.New()
	couriers := router.Group("/couriers")
	{
		couriers.POST("/", h.AddCouriers)
		couriers.GET("/", h.GetCouriers)
		couriers.GET("/:id", h.GetCouriersById)
	}
	orders := router.Group("/orders")
	{
		orders.POST("/", h.AddOrders)
		orders.GET("/", h.GetOrders)
		orders.GET("/:id", h.GetCouriersById)
		complete := orders.Group("/complete")
		{
			complete.POST("/", h.CompleteTheOrder)
		}
	}

	return router
}

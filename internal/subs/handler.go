package subs

import (
	"net/http"
	"strconv"
	"sub-aggregator/cfg"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Config  *cfg.Config
	Service *Service
}

type HandlerDeps struct {
	Config  *cfg.Config
	Service *Service
}

func NewHandler(router *gin.Engine, deps *HandlerDeps) {
	handler := &Handler{
		Config:  deps.Config,
		Service: deps.Service,
	}

	apiV1 := router.Group("/api/v1")
	{
		subscriptions := apiV1.Group("/subscriptions")
		{
			subscriptions.POST("", handler.Create)
			subscriptions.DELETE("/:id", handler.Delete)
			subscriptions.GET("/:id", handler.GetById)
			subscriptions.PUT("/:id", handler.Update)
			subscriptions.PATCH("/:id", handler.Patch)
			subscriptions.GET("", handler.GetList)
			subscriptions.GET("/summary", handler.GetSummary)
		}
	}
}

func (h *Handler) Create(c *gin.Context) {
	var requestData SubscriptionRequest

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.Service.Create(requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "subscription successfully created",
	})
}

func (h *Handler) Delete(c *gin.Context) {
	subIdString := c.Param("id")

	subId, err := strconv.Atoi(subIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid ID format",
		})
	}

	if err := h.Service.Delete(subId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "subscription successfully deleted",
	})
}

func (h *Handler) GetById(c *gin.Context) {
	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid ID format",
		})
	}

	foundedSub, err := h.Service.GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, foundedSub)
}

func (h *Handler) Update(c *gin.Context) {
	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid ID format",
		})
	}

	var requestData SubscriptionRequest
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedSubscription, err := h.Service.Update(uint(id), requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedSubscription)
}

func (h *Handler) Patch(c *gin.Context) {

}

func (h *Handler) GetList(c *gin.Context) {

}

func (h *Handler) GetSummary(c *gin.Context) {

}

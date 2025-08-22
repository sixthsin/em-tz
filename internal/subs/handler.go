package subs

import (
	"log"
	"net/http"
	"strconv"
	"sub-aggregator/cfg"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		return
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
		return
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
		return
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
	idString := c.Param("id")

	id, err := strconv.Atoi(idString)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid ID format",
		})
		return
	}

	var requestData *PatchSubscriptionRequest
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedSubscription, err := h.Service.UpdatePartially(uint(id), requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedSubscription)
}

func (h *Handler) GetList(c *gin.Context) {
	limitString := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitString)
	if err != nil || limit < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid limit parameter",
		})
		return
	}

	offsetString := c.DefaultQuery("offset", "0")
	offset, err := strconv.Atoi(offsetString)
	if err != nil || offset < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid offset parameter",
		})
		return
	}

	var filters SearchParams
	if err := c.ShouldBindQuery(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid query parameters",
		})
		return
	}

	log.Println(filters)

	if filters.StartDate == "" || filters.EndDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "start_date and end_date are required",
		})
		return
	}

	foundData, err := h.Service.GetWithParams(limit, offset, filters)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": foundData,
		"metadata": gin.H{
			"limit":        limit,
			"offset":       offset,
			"current_page": (offset / limit) + 1,
		},
	})
}

func (h *Handler) GetSummary(c *gin.Context) {
	var filters SearchParams
	if err := c.ShouldBindQuery(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	if filters.StartDate == "" || filters.EndDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "start_date and end_date are required",
		})
		return
	}

	total, err := h.Service.GetSummary(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var userId uuid.UUID
	var serviceName string
	if filters.UserId != nil {
		userId = *filters.UserId
	}
	if filters.ServiceName != nil {
		serviceName = *filters.ServiceName
	}

	c.JSON(http.StatusOK, gin.H{
		"total_amount": total,
		"period": gin.H{
			"start_date": filters.StartDate,
			"end_date":   filters.EndDate,
		},
		"filters": gin.H{
			"user_id":      userId,
			"service_name": serviceName,
		},
	})
}

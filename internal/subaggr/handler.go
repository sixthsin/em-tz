package subaggr

import (
	"net/http"
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
	router.POST("/api/v1/sub-aggr", handler.Create)
}

func (h *Handler) Create(c *gin.Context) {
	var requestData SubscriptionData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":       "ok",
		"created_data": requestData,
	})
}

package subs

import (
	"github.com/google/uuid"
)

type CreateSubscriptionRequest struct {
	ServiceName string    `json:"service_name" binding:"required"`
	Price       uint      `json:"price" binding:"required,min=0"`
	UserID      uuid.UUID `json:"user_id" binding:"required"`
	StartDate   string    `json:"start_date" binding:"required"`
	EndDate     string    `json:"end_date" binding:"omitempty"`
}

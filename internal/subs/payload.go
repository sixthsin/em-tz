package subs

import (
	"github.com/google/uuid"
)

type SubscriptionRequest struct {
	ServiceName string    `json:"service_name" binding:"required"`
	Price       uint      `json:"price" binding:"required,min=0"`
	UserId      uuid.UUID `json:"user_id" binding:"required"`
	StartDate   string    `json:"start_date" binding:"required"`
	EndDate     string    `json:"end_date" binding:"omitempty"`
}

type PatchSubscriptionRequest struct {
	ServiceName *string    `json:"service_name" binding:"omitempty"`
	Price       *uint      `json:"price" binding:"omitempty,min=0"`
	UserId      *uuid.UUID `json:"user_id" binding:"omitempty"`
	StartDate   *string    `json:"start_date" binding:"omitempty"`
	EndDate     *string    `json:"end_date" binding:"omitempty"`
}

type SearchParams struct {
	UserId      *uuid.UUID `form:"user_id"`
	ServiceName *string    `form:"service_name"`
	StartDate   string     `form:"start_date" binding:"required"`
	EndDate     string     `form:"end_date" binding:"required"`
}

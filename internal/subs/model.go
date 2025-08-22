package subs

import (
	"github.com/google/uuid"
)

type Subscription struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	ServiceName string    `gorm:"not null;index:idx_service_name" json:"service_name"`
	Price       uint      `gorm:"not null" json:"price"`
	UserID      uuid.UUID `gorm:"type:uuid;not null;index:idx_user_id" json:"user_id"`
	StartDate   string    `gorm:"not null;index:idx_start_date" json:"start_date"`
	EndDate     string    `gorm:"index:idx_end_date" json:"end_date,omitempty"`
}

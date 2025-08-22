package subs

import (
	"sub-aggregator/pkg/db"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	Database *db.Db
}

func NewRepository(database *db.Db) *Repository {
	return &Repository{
		Database: database,
	}
}

func (r *Repository) Create(subscription *Subscription) error {
	res := r.Database.DB.Create(subscription)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *Repository) Delete(subId int) error {
	res := r.Database.DB.Where("id = ?", subId).Delete(&Subscription{})
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *Repository) GetById(id uint) (*Subscription, error) {
	var sub Subscription
	err := r.Database.First(&sub, id).Error
	return &sub, err
}

func (r *Repository) Update(id uint, subscriptionUpdate *Subscription) (*Subscription, error) {
	var updatedSubscription Subscription
	err := r.Database.DB.Transaction(func(tx *gorm.DB) error {
		res := tx.Model(&Subscription{}).
			Where("id = ?", id).
			Omit("id").
			Updates(subscriptionUpdate)

		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}

		return tx.First(&updatedSubscription, id).Error
	})

	if err != nil {
		return nil, err
	}

	return &updatedSubscription, nil
}

func (r *Repository) GetWithParams(limit, offset int, filters SearchParams) []Subscription {
	var foundSubs []Subscription
	query := r.Database.DB.Model(&Subscription{})

	if filters.ServiceName != nil {
		query = query.Where("service_name LIKE ?", "%"+*filters.ServiceName+"%")
	}
	if filters.UserId != nil {
		query = query.Where("user_id = ?", filters.UserId)
	}
	if filters.StartDate != "" && filters.EndDate != "" {
		query = query.Where("start_date BETWEEN ? AND ?", filters.StartDate, filters.EndDate)
	}

	query.Limit(limit).Offset(offset).Find(&foundSubs)

	return foundSubs
}

func (r *Repository) GetSummary(startDate, endDate string, userId *uuid.UUID, serviceName *string) (uint, error) {
	var total struct {
		Total uint `gorm:"column:total_amount"`
	}

	query := r.Database.DB.Model(&Subscription{}).
		Select("COALESCE(SUM(price), 0) as total_amount").
		Where("start_date BETWEEN ? AND ?", endDate, startDate)
	if userId != nil {
		query = query.Where("user_id = ?", *userId)
	}
	if serviceName != nil {
		query = query.Where("service_name = ?", *serviceName)
	}

	if err := query.Scan(&total).Error; err != nil {
		return 0, err
	}

	return total.Total, nil
}

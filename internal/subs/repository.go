package subs

import (
	"sub-aggregator/pkg/db"

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

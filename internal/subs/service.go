package subs

import (
	"fmt"
)

type ServiceDeps struct {
	Repository *Repository
}

type Service struct {
	Repository *Repository
}

func NewService(deps *ServiceDeps) *Service {
	return &Service{
		Repository: deps.Repository,
	}
}

func (s *Service) Create(requestData SubscriptionRequest) error {
	subscription := &Subscription{
		ServiceName: requestData.ServiceName,
		Price:       requestData.Price,
		UserId:      requestData.UserId,
		StartDate:   requestData.StartDate,
		EndDate:     requestData.EndDate,
	}

	if err := s.Repository.Create(subscription); err != nil {
		return err
	}

	return nil
}

func (s *Service) Delete(subId int) error {
	if err := s.Repository.Delete(subId); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetById(id uint) (*Subscription, error) {
	foundedSub, err := s.Repository.GetById(id)
	if err != nil {
		return nil, err
	}

	return foundedSub, nil
}

func (s *Service) Update(id uint, requestData SubscriptionRequest) (*Subscription, error) {
	subscription := &Subscription{
		ServiceName: requestData.ServiceName,
		Price:       requestData.Price,
		UserId:      requestData.UserId,
		StartDate:   requestData.StartDate,
		EndDate:     requestData.EndDate,
	}

	updatedSubscription, err := s.Repository.Update(id, subscription)
	if err != nil {
		return nil, err
	}

	return updatedSubscription, nil
}

func (s *Service) UpdatePartially(id uint, requestData *PatchSubscriptionRequest) (*Subscription, error) {
	subscriptionUpdate := &Subscription{}

	if requestData.ServiceName != nil {
		subscriptionUpdate.ServiceName = *requestData.ServiceName
	}
	if requestData.Price != nil {
		subscriptionUpdate.Price = *requestData.Price
	}
	if requestData.UserId != nil {
		subscriptionUpdate.UserId = *requestData.UserId
	}
	if requestData.StartDate != nil {
		subscriptionUpdate.StartDate = *requestData.StartDate
	}
	if requestData.EndDate != nil {
		subscriptionUpdate.EndDate = *requestData.EndDate
	}

	updatedSubscription, err := s.Repository.Update(id, subscriptionUpdate)
	if err != nil {
		return nil, err
	}

	return updatedSubscription, nil
}

func (s *Service) GetWithParams(limit, offset int, filters SearchParams) ([]Subscription, error) {
	return s.Repository.GetWithParams(limit, offset, filters), nil
}

func (s *Service) GetSummary(filters SearchParams) (uint, error) {
	total, err := s.Repository.GetSummary(filters.StartDate, filters.EndDate, filters.UserId, filters.ServiceName)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate total amount: %w", err)
	}

	return total, nil
}

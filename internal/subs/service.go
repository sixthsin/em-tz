package subs

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
		UserID:      requestData.UserID,
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
		UserID:      requestData.UserID,
		StartDate:   requestData.StartDate,
		EndDate:     requestData.EndDate,
	}

	updatedSubscription, err := s.Repository.Update(id, subscription)
	if err != nil {
		return nil, err
	}

	return updatedSubscription, nil
}

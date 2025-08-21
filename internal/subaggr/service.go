package subaggr

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

func (s *Service) Create(requestData SubscriptionData) (SubscriptionData, error) {
	// if s.Repository.Crea
}

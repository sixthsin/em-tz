package subaggr

type SubscriptionData struct {
	ServiceName string `json:"service_name" validate:"required"`
	Price       uint   `json:"price" validate:"required"`
	UserId      string `json:"user_id" validate:"required"`
	StartDate   string `json:"start_date" validate:"required"`
}

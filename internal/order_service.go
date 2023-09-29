package internal

import "context"

type OrderService struct {
	repo *OrderRepository
}

func NewOrderService(repo *OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) SaveOrder(ctx context.Context, order *Order) error {
	return s.repo.Create(ctx, order)
}

func (s *OrderService) GetOrders(ctx context.Context) ([]*Order, error) {
	return s.repo.FindAll(ctx)
}

func (s *OrderService) GetOrder(ctx context.Context, id string) (*Order, error) {
	return s.repo.FindOne(ctx, id)
}

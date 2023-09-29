package internal

import "context"

type OrderService struct {
	repo *OrderRepository
	metricPublisher *MetricPublisher
}

func NewOrderService(repo *OrderRepository, publisher *MetricPublisher) *OrderService {
	return &OrderService{repo: repo, metricPublisher: publisher}
}

func (s *OrderService) SaveOrder(ctx context.Context, order *Order) error {
	if err := s.repo.Create(ctx, order); err != nil {
		return err
	}

	s.metricPublisher.PublishOrderCreated(order)
	return nil
}

func (s *OrderService) GetOrders(ctx context.Context) ([]Order, error) {
	return s.repo.FindAll(ctx)
}

func (s *OrderService) GetOrder(ctx context.Context, id string) (Order, error) {
	order, err := s.repo.FindOne(ctx, id)
	if err != nil {
		return Order{}, err
	}

	s.metricPublisher.PublishOrderRetrieved()
	return order, nil
}

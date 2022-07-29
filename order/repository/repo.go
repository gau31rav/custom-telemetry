package repository

import (
	"context"
	"fmt"

	"github.com/gau31rav/custom-telemetry/order/pkg/domain/model"
)

var orderRepo map[int][]model.Order = make(map[int][]model.Order)

type Repo interface {
	AddOrder(ctx context.Context, order model.Order) error
	GetOrders(ctx context.Context) ([]model.Order, error)
}

type repo struct{}

func (r *repo) AddOrder(ctx context.Context, order model.Order) error {
	ordersForUser, ok := orderRepo[order.User.ID]
	if !ok {
		var orders []model.Order
		orders = append(orders, order)
		orderRepo[order.User.ID] = orders
		return nil
	}
	ordersForUser = append(ordersForUser, order)
	orderRepo[order.User.ID] = ordersForUser
	return nil
}

func (r *repo) GetOrders(ctx context.Context) ([]model.Order, error) {
	var orders []model.Order
	for _, value := range orderRepo {
		orders = append(orders, value...)
	}
	if len(orders) == 0 {
		return nil, fmt.Errorf("no order found")
	}
	return orders, nil
}

func NewOrderRepo() Repo {
	return &repo{}
}

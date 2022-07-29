package service

import (
	"context"

	"github.com/gau31rav/custom-telemetry/order/pkg/domain/model"
)

type Order interface {
	PlaceOrder(ctx context.Context) (string, error)
	GetOrder(ctx context.Context) ([]model.Order, error)
}

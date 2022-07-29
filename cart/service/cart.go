package service

import (
	"github.com/gau31rav/custom-telemetry/cart/pkg/domain/model"
	"github.com/google/uuid"
)

type Cart interface {
	AddItems(items model.Item) error
	DeleteItems(id uuid.UUID) error
	GetItems() ([]model.Item, error)
}

package domain

import (
	"context"

	"github.com/gau31rav/custom-telemetry/cart/pkg/domain/model"
	"github.com/gau31rav/custom-telemetry/cart/service"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/trace"
)

var inMemoryCart map[uuid.UUID]model.Item = make(map[uuid.UUID]model.Item)

type cart struct {
}

//AddItems ...
func (c *cart) AddItems(ctx context.Context, item model.Item) error {
	span := trace.SpanFromContext(ctx)
	bag := baggage.FromContext(ctx)
	uk := attribute.Key("add item in service")
	span.AddEvent("handling add item at service layer...", trace.WithAttributes(uk.String(bag.Member("add item in service").Value())))
	itemExist, ok := inMemoryCart[item.ID]
	if !ok {
		inMemoryCart[item.ID] = item
		return nil
	}
	itemFromCart := inMemoryCart[itemExist.ID]
	itemFromCart.Quantity++
	inMemoryCart[item.ID] = itemFromCart
	return nil
}

//DeleteItems ...
func (c *cart) DeleteItems(id uuid.UUID) error {
	delete(inMemoryCart, id)
	return nil
}

//GetItems ...
func (c *cart) GetItems() ([]model.Item, error) {
	var items []model.Item
	for _, value := range inMemoryCart {
		items = append(items, value)
	}
	return items, nil
}

// NewCartService ...
func NewCartService() service.Cart {
	return &cart{}
}

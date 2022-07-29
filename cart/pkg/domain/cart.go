package domain

import (
	"github.com/gau31rav/custom-telemetry/cart/pkg/domain/model"
	"github.com/gau31rav/custom-telemetry/cart/service"
	"github.com/google/uuid"
)

var inMemoryCart map[uuid.UUID]model.Item = make(map[uuid.UUID]model.Item)

type cart struct {
}

//AddItems ...
func (c *cart) AddItems(item model.Item) error {
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

package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	customLogger "github.com/gau31rav/cloudera-custom-telemetry/logger"
	"github.com/gau31rav/custom-telemetry/order/pkg/domain/model"
	"github.com/gau31rav/custom-telemetry/order/repository"
	"github.com/gau31rav/custom-telemetry/order/service"
	"github.com/google/uuid"
)

type order struct {
	repo   repository.Repo
	logger customLogger.Logger
}

func (o order) PlaceOrder(ctx context.Context) (string, error) {
	orderID := uuid.New()
	items := o.getItemsFromCart()
	if items == nil {
		o.logger.Debug("no items in cart")
		return "", fmt.Errorf("please add items to cart")
	}
	orderDetails := model.Order{
		ID:    orderID,
		Items: items,
		User: model.User{
			ID:    1,
			Name:  "gaurav",
			Email: "dummyemail@dummy.com",
		},
	}
	err := o.repo.AddOrder(ctx, orderDetails)
	if err != nil {
		o.logger.Error("failed to place order. error %s", err.Error())
		return "", err
	}
	return orderID.String(), nil
}

func (o order) GetOrder(ctx context.Context) ([]model.Order, error) {
	return o.repo.GetOrders(ctx)
}

func NewOrderService(repo repository.Repo, logger customLogger.Logger) service.Order {
	return &order{repo: repo, logger: logger}
}

func (o order) getItemsFromCart() []model.Item {
	resp, err := http.Get("http://localhost:7777/cart")
	if err != nil {
		o.logger.Error("failed to get items from cart. error %s", err.Error())
		return nil
	}
	defer resp.Body.Close()
	var items []model.Item
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	_ = json.Unmarshal(body, &items)
	return items
}

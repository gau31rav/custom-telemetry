package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	customLogger "github.com/gau31rav/cloudera-custom-telemetry/logger"
	"github.com/gau31rav/cloudera-custom-telemetry/telemetry"
	"github.com/gau31rav/custom-telemetry/order/pkg/domain"
	"github.com/gau31rav/custom-telemetry/order/repository"
	"github.com/gau31rav/custom-telemetry/order/service"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/trace"
)

var (
	repo     repository.Repo     = repository.NewOrderRepo()
	logger   customLogger.Logger = customLogger.NewAppLogger()
	orderSVC service.Order       = domain.NewOrderService(repo, logger)
)

func main() {
	tp, err := telemetry.InitTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()
	otelAddHandler := otelhttp.NewHandler(http.HandlerFunc(addHandler), "addOrder")
	http.Handle("/order/add", otelAddHandler)
	getItemHandler := otelhttp.NewHandler(http.HandlerFunc(getHandler), "getOrder")
	http.Handle("/order", getItemHandler)
	err = http.ListenAndServe(":3030", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	span := trace.SpanFromContext(ctx)
	bag := baggage.FromContext(ctx)
	uk := attribute.Key("get orders")
	span.AddEvent("get orders...", trace.WithAttributes(uk.String(bag.Member("get order").Value())))
	orders, err := orderSVC.GetOrder(ctx)
	if err != nil {
		logger.Error("failed to get order", err.Error())
		return
	}
	resp, err := json.Marshal(orders)
	if err != nil {
		logger.Error("failed to get order", err.Error())
		return
	}
	_, _ = io.WriteString(w, fmt.Sprintf("Orders: %s", string(resp)))
}

func addHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	span := trace.SpanFromContext(ctx)
	bag := baggage.FromContext(ctx)
	uk := attribute.Key("add item")
	span.AddEvent("place order...", trace.WithAttributes(uk.String(bag.Member("add order").Value())))
	orderID, err := orderSVC.PlaceOrder(ctx)
	if err != nil {
		logger.Error("failed to place order", err.Error())
		return
	}
	_, _ = io.WriteString(w, fmt.Sprintf("Order Added Successfully. Order Id: %s", orderID))
}

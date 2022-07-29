package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gau31rav/custom-telemetry/cart/pkg/domain"
	"github.com/gau31rav/custom-telemetry/cart/pkg/domain/model"
	"github.com/gau31rav/custom-telemetry/cart/service"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

var cartSVC service.Cart = domain.NewCartService()

func main() {
	tp, err := initTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()
	otelAddHandler := otelhttp.NewHandler(http.HandlerFunc(addHandler), "addItem")
	http.Handle("/cart/addItem", otelAddHandler)
	getItemHandler := otelhttp.NewHandler(http.HandlerFunc(getHandler), "getItem")
	http.Handle("/cart", getItemHandler)
	err = http.ListenAndServe(":7777", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func initTracer() (*sdktrace.TracerProvider, error) {
	exporter, err := stdout.New(stdout.WithPrettyPrint())
	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}

func addHandler(w http.ResponseWriter, req *http.Request) {
	var item model.Item
	ctx := req.Context()
	span := trace.SpanFromContext(ctx)
	bag := baggage.FromContext(ctx)
	uk := attribute.Key("add item")
	span.AddEvent("handling this...", trace.WithAttributes(uk.String(bag.Member("add item").Value())))
	err := json.NewDecoder(req.Body).Decode(&item)
	_ = err
	_ = cartSVC.AddItems(ctx, item)
	_, _ = io.WriteString(w, "Item Added Successfully")
}

func getHandler(w http.ResponseWriter, req *http.Request) {
	cartSVC := domain.NewCartService()
	items, _ := cartSVC.GetItems()
	resp, _ := json.Marshal(items)
	_, _ = io.WriteString(w, string(resp))
}

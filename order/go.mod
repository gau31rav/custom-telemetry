module github.com/gau31rav/custom-telemetry/order

go 1.18

replace github.com/gau31rav/cloudera-custom-telemetry => ../cloudera-custom-telemetry

require (
	github.com/gau31rav/cloudera-custom-telemetry v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.3.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.33.0
)

require (
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/getsentry/sentry-go v0.13.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/keepeye/logrus-filename v0.0.0-20190711075016-ce01a4391dd1 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	go.opentelemetry.io/otel v1.8.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.8.0 // indirect
	go.opentelemetry.io/otel/metric v0.31.0 // indirect
	go.opentelemetry.io/otel/sdk v1.8.0 // indirect
	go.opentelemetry.io/otel/trace v1.8.0 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)

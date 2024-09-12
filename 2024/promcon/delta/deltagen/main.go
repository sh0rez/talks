package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"

	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
)

const Interval = 15 * time.Second

func main() {
	ctx := context.Background()

	otlp, _ := otlpmetrichttp.New(ctx, otlpmetrichttp.WithTemporalitySelector(func(_ metric.InstrumentKind) metricdata.Temporality {
		return metricdata.DeltaTemporality
	}))

	provider := metric.NewMeterProvider(metric.WithReader(
		metric.NewPeriodicReader(otlp, metric.WithInterval(Interval)),
	))

	meter := provider.Meter("promcon.io/deltas")

	sum, _ := meter.Float64Counter("random.sum")
	gauge, _ := meter.Float64Gauge("random.gauge")

	log.Println("sending random increase every", Interval)

	for range time.Tick(Interval) {
		v := float64(rand.Intn(10))
		sum.Add(ctx, v)
		gauge.Record(ctx, v)
	}
}

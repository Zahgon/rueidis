package rueidisotel

import (
	"context"
	"crypto/tls"
	"net"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"

	"github.com/redis/rueidis"
)

var (
	defaultHistogramBuckets = []float64{
		.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10,
	}
)

// MetricAttrs set additional attributes to append to each metric.
func MetricAttrs(attrs ...attribute.KeyValue) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Allocate slices once and use many times

// WithMeterProvider sets the MeterProvider for the otelclient.
func WithMeterProvider(provider metric.MeterProvider) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithOperationMetricAttr sets the operation name as an attribute for duration and error metrics.
// This may cause memory usage to increase with the number of commands used.
func WithOperationMetricAttr() Option { _ = "STUB: not implemented"; return *new(Option) }

type HistogramOption struct {
	Buckets []float64
}

type dialMetrics struct {
	attempt    metric.Int64Counter
	success    metric.Int64Counter
	counts     metric.Int64UpDownCounter
	latency    metric.Float64Histogram
	addOpts    []metric.AddOption
	recordOpts []metric.RecordOption
}

type dialTracer struct {
	trace.Tracer
	tAttrs trace.SpanStartEventOption
}

// WithHistogramOption sets the HistogramOption.
// If not set, DefaultHistogramBuckets will be used.
func WithHistogramOption(histogramOption HistogramOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// NewClient creates a new Client.
// The following metrics are recorded:
// - rueidis_dial_attempt: number of dial attempts
// - rueidis_dial_success: number of successful dials
// - rueidis_dial_conns: number of active connections
// - rueidis_dial_latency: dial latency in seconds
func NewClient(clientOption rueidis.ClientOption, opts ...Option) (rueidis.Client, error) {
	_ = "STUB: not implemented"
	return *new(rueidis.Client), nil
}

func newClient(opts ...Option) (*otelclient, error) { _ = "STUB: not implemented"; return nil, nil }

// Default to global MeterProvider

// Default to global TracerProvider

// Now that we have the meterProvider and tracerProvider, get the Meter and Tracer

// Now create the counters using the meter

func trackDialing(m dialMetrics, t dialTracer, dialFn func(context.Context, string, *net.Dialer, *tls.Config) (conn net.Conn, err error)) func(context.Context, string, *net.Dialer, *tls.Config) (conn net.Conn, err error) {
	_ = "STUB: not implemented"
	return nil
}

// Use floating point division for higher precision (instead of Seconds method).

type connTracker struct {
	net.Conn
	counts  metric.Int64UpDownCounter
	addOpts []metric.AddOption
	once    int32
}

func (t *connTracker) Close() error { _ = "STUB: not implemented"; return nil }

func defaultDialFn(ctx context.Context, dst string, dialer *net.Dialer, cfg *tls.Config) (conn net.Conn, err error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func serverAttrs(dst string) trace.SpanStartEventOption {
	_ = "STUB: not implemented"
	return *new(trace.SpanStartEventOption)
}

func defaultSpanNameFormatter(_ context.Context, op string) string {
	_ = "STUB: not implemented"
	return ""
}

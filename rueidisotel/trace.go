package rueidisotel

import (
	"context"
	"sync"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"

	"github.com/redis/rueidis"
)

var (
	name   = "github.com/redis/rueidis"
	kind   = trace.WithSpanKind(trace.SpanKindClient)
	dbattr = attribute.String("db.system", "redis")
	dbstmt = attribute.Key("db.statement")
)

type contextKey struct{}

var labelerContextKey = contextKey{}

// Labeler is used to allow instrumentation to add additional attributes
// to metrics.
type Labeler struct {
	attrs []attribute.KeyValue
}

// Add appends new attributes to the labeler.
func (l *Labeler) Add(attrs ...attribute.KeyValue) { _ = "STUB: not implemented"; return }

// Get returns the attributes added to the labeler.
func (l *Labeler) Get() []attribute.KeyValue {
	_ = "STUB: not implemented"

	// LabelerFromContext retrieves a Labeler instance from the provided context.
	// It returns the labeler and a boolean indicating whether it was found.
	// If no labeler is found, returns a new empty labeler and false.
	return nil
}

func LabelerFromContext(ctx context.Context) (*Labeler, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// ContextWithLabeler returns a new context with the provided Labeler.
// Attributes added to the labeler will be included in cache metrics.
//
// Example:
//
//	// Check if labeler exists in context, create new context only if needed
//	labeler, ok := rueidisotel.LabelerFromContext(ctx)
//	if !ok {
//		ctx = rueidisotel.ContextWithLabeler(ctx, labeler)
//	}
//	labeler.Add(attribute.String("key_pattern", "book"))
//	client.DoCache(ctx, client.B().Get().Key("book:123").Cache(), time.Minute)
func ContextWithLabeler(ctx context.Context, labeler *Labeler) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

var _ rueidis.Client = (*otelclient)(nil)

// WithClient creates a new rueidis.Client with OpenTelemetry tracing enabled.
//
// Deprecated: use NewClient() instead.
func WithClient(client rueidis.Client, opts ...Option) rueidis.Client {
	_ = "STUB: not implemented"
	return *new(rueidis.Client)
}

// Option is the Functional Options interface
type Option func(o *otelclient)

// TraceAttrs set additional attributes to append to each trace.
func TraceAttrs(attrs ...attribute.KeyValue) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTracerProvider sets the TracerProvider for the otelclient.
func WithTracerProvider(provider trace.TracerProvider) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithDBStatement tells the tracing hook to add raw redis commands to the db.statement attribute.
func WithDBStatement(f StatementFunc) Option { _ = "STUB: not implemented"; return *new(Option) }

// StatementFunc is the function that maps a command's tokens to a string to put in the db.statement attribute
type StatementFunc func(cmdTokens []string) string

// WithOpNameResolver sets a custom [OpNameResolver] used to determine
// the operation name for metrics and spans produced by the client.
//
// By default, [DefaultOpNameResolver] is used.
func WithOpNameResolver(resolver OpNameResolver) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSpanNameFormatter sets a custom function used to format span names created by the client.
func WithSpanNameFormatter(fn SpanNameFormatterFunc) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// SpanNameFormatterFunc defines a function that formats the name of a span.
type SpanNameFormatterFunc func(ctx context.Context, op string) string

type commandMetrics struct {
	duration          metric.Float64Histogram
	errors            metric.Int64Counter
	addOpts           []metric.AddOption
	recordOpts        []metric.RecordOption
	opAttr            bool
	opNameResolver    OpNameResolver
	spanNameFormatter SpanNameFormatterFunc
}

var (
	metricRecordOptionPool = &sync.Pool{
		New: func() any { return &[]metric.RecordOption{} },
	}
	metricAddOptionPool = &sync.Pool{
		New: func() any { return &[]metric.AddOption{} },
	}
)

func (c *commandMetrics) recordDuration(ctx context.Context, op string, startTime time.Time) {
	_ = "STUB: not implemented"
	return
}

func (c *commandMetrics) recordError(ctx context.Context, op string, err error) {
	_ = "STUB: not implemented"
	return
}

type otelclient struct {
	client          rueidis.Client
	meterProvider   metric.MeterProvider
	tracerProvider  trace.TracerProvider
	tracer          trace.Tracer
	meter           metric.Meter
	cscMiss         metric.Int64Counter
	cscHits         metric.Int64Counter
	sAttrs          trace.SpanStartEventOption
	tAttrs          trace.SpanStartEventOption
	dbStmtFunc      StatementFunc
	histogramOption HistogramOption
	commandMetrics
}

func (o *otelclient) B() rueidis.Builder { _ = "STUB: not implemented"; return *new(rueidis.Builder) }

func (o *otelclient) Do(ctx context.Context, cmd rueidis.Completed) (resp rueidis.RedisResult) {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisResult)
}

func (o *otelclient) DoMulti(ctx context.Context, multi ...rueidis.Completed) (resp []rueidis.RedisResult) {
	_ = "STUB: not implemented"
	return nil
}

func (o *otelclient) DoStream(ctx context.Context, cmd rueidis.Completed) (resp rueidis.RedisResultStream) {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisResultStream)
}

func (o *otelclient) DoMultiStream(ctx context.Context, multi ...rueidis.Completed) (resp rueidis.MultiRedisResultStream) {
	_ = "STUB: not implemented"
	return *new(rueidis.MultiRedisResultStream)
}

func (o *otelclient) DoCache(ctx context.Context, cmd rueidis.Cacheable, ttl time.Duration) (resp rueidis.RedisResult) {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisResult)
}

func (o *otelclient) DoMultiCache(ctx context.Context, multi ...rueidis.CacheableTTL) (resps []rueidis.RedisResult) {
	_ = "STUB: not implemented"
	return nil
}

func (o *otelclient) Dedicated(fn func(rueidis.DedicatedClient) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (o *otelclient) Dedicate() (rueidis.DedicatedClient, func()) {
	_ = "STUB: not implemented"
	return *new(rueidis.DedicatedClient), nil
}

func (o *otelclient) Receive(ctx context.Context, subscribe rueidis.Completed, fn func(msg rueidis.PubSubMessage)) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (o *otelclient) Nodes() map[string]rueidis.Client { _ = "STUB: not implemented"; return nil }

func (o *otelclient) Mode() rueidis.ClientMode {
	_ = "STUB: not implemented"
	return *new(rueidis.ClientMode)
}

func (o *otelclient) Close() { _ = "STUB: not implemented"; return }

func (o *otelclient) recordCacheHitMiss(ctx context.Context, resp rueidis.RedisResult) {
	_ = "STUB: not implemented"
	return
}

var _ rueidis.DedicatedClient = (*dedicated)(nil)

type dedicated struct {
	client     rueidis.DedicatedClient
	tracer     trace.Tracer
	sAttrs     trace.SpanStartEventOption
	tAttrs     trace.SpanStartEventOption
	dbStmtFunc StatementFunc
	commandMetrics
}

func (d *dedicated) B() rueidis.Builder { _ = "STUB: not implemented"; return *new(rueidis.Builder) }

func (d *dedicated) Do(ctx context.Context, cmd rueidis.Completed) (resp rueidis.RedisResult) {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisResult)
}

func (d *dedicated) DoMulti(ctx context.Context, multi ...rueidis.Completed) (resp []rueidis.RedisResult) {
	_ = "STUB: not implemented"
	return nil
}

func (d *dedicated) Receive(ctx context.Context, subscribe rueidis.Completed, fn func(msg rueidis.PubSubMessage)) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (d *dedicated) SetPubSubHooks(hooks rueidis.PubSubHooks) <-chan error {
	_ = "STUB: not implemented"
	return nil
}

func (d *dedicated) SetOnInvalidations(fn func([]rueidis.RedisMessage)) <-chan error {
	_ = "STUB: not implemented"
	return nil
}

func (d *dedicated) Close() { _ = "STUB: not implemented"; return }

func sum(s []string) (v int) { _ = "STUB: not implemented"; return 0 }

func firstError(s []rueidis.RedisResult) error { _ = "STUB: not implemented"; return nil }

func multiSum(multi rueidis.Commands) (v int) { _ = "STUB: not implemented"; return 0 }

func multiCacheableSum(multi []rueidis.CacheableTTL) (v int) { _ = "STUB: not implemented"; return 0 }

func (o *otelclient) start(ctx context.Context, op string, size int) (context.Context, trace.Span) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(trace.Span)
}

func (o *otelclient) end(span trace.Span, err error) { _ = "STUB: not implemented"; return }

func (d *dedicated) start(ctx context.Context, op string, size int) (context.Context, trace.Span) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(trace.Span)
}

func (d *dedicated) end(span trace.Span, err error) { _ = "STUB: not implemented"; return }

func startSpan(tracer trace.Tracer, ctx context.Context, op string, size int, sAttrs trace.SpanStartEventOption, tAttrs trace.SpanStartEventOption) (context.Context, trace.Span) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(trace.Span)
}

func endSpan(span trace.Span, err error) { _ = "STUB: not implemented"; return }

// do not record the full db.statement to avoid collecting sensitive data
func attr(op string, size int) trace.SpanStartEventOption {
	_ = "STUB: not implemented"
	return *new(trace.SpanStartEventOption)
}

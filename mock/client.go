package mock

import (
	"context"
	"time"

	"github.com/redis/rueidis"
	"go.uber.org/mock/gomock"
)

var _ rueidis.Client = (*Client)(nil)
var _ rueidis.DedicatedClient = (*DedicatedClient)(nil)

// ClientOption is an optional function parameter for NewClient
type ClientOption func(c any)

// WithSlotCheck enables the command builder of Client to check if the command built across multiple slots and then panic
func WithSlotCheck() ClientOption { _ = "STUB: not implemented"; return *new(ClientOption) }

// Client mocks the Client interface.
type Client struct {
	ctrl     *gomock.Controller
	recorder *ClientMockRecorder
	slot     uint16
}

// ClientMockRecorder is the mock recorder for Client.
type ClientMockRecorder struct {
	mock *Client
}

// NewClient creates a new mock instance.
func NewClient(ctrl *gomock.Controller, options ...ClientOption) *Client {
	_ = "STUB: not implemented"
	return nil
}

// EXPECT returns an object that allows the caller to indicate the expected use.
func (m *Client) EXPECT() *ClientMockRecorder {
	_ = "STUB: not implemented"

	// B mocks base method.
	return nil
}

func (m *Client) B() rueidis.Builder { _ = "STUB: not implemented"; return *new(rueidis.Builder) }

// Mode mocks base method.
func (m *Client) Mode() rueidis.ClientMode {
	_ = "STUB: not implemented"
	return *new(rueidis.ClientMode)
}

// Mode indicates an expected call of Mode.
func (mr *ClientMockRecorder) Mode() *gomock.Call { _ = "STUB: not implemented"; return nil }

// Close mocks base method.
func (m *Client) Close() { _ = "STUB: not implemented"; return }

// Close indicates an expected call of Close.
func (mr *ClientMockRecorder) Close() *gomock.Call { _ = "STUB: not implemented"; return nil }

// Dedicate mocks base method.
func (m *Client) Dedicate() (rueidis.DedicatedClient, func()) {
	_ = "STUB: not implemented"
	return *new(rueidis.DedicatedClient), nil
}

// Dedicate indicates an expected call of Dedicate.
func (mr *ClientMockRecorder) Dedicate() *gomock.Call { _ = "STUB: not implemented"; return nil }

// Dedicated mocks base method.
func (m *Client) Dedicated(arg0 func(rueidis.DedicatedClient) error) error {
	_ = "STUB: not implemented"
	return nil
}

// Dedicated indicates an expected call of Dedicated.
func (mr *ClientMockRecorder) Dedicated(arg0 any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

// Do mocks the base method.
func (m *Client) Do(arg0 context.Context, arg1 rueidis.Completed) rueidis.RedisResult {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisResult)
}

// Do indicates an expected call of Do.
func (mr *ClientMockRecorder) Do(arg0, arg1 any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

// DoStream mocks base method.
func (m *Client) DoStream(arg0 context.Context, arg1 rueidis.Completed) rueidis.RedisResultStream {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisResultStream)
}

// DoStream indicates an expected call of DoStream.
func (mr *ClientMockRecorder) DoStream(arg0, arg1 any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

// DoCache mocks base method.
func (m *Client) DoCache(arg0 context.Context, arg1 rueidis.Cacheable, arg2 time.Duration) rueidis.RedisResult {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisResult)
}

// DoCache indicates an expected call of DoCache.
func (mr *ClientMockRecorder) DoCache(arg0, arg1, arg2 any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

// DoMulti mocks base method.
func (m *Client) DoMulti(arg0 context.Context, arg1 ...rueidis.Completed) []rueidis.RedisResult {
	_ = "STUB: not implemented"
	return nil
}

// DoMulti indicates an expected call of DoMulti.
func (mr *ClientMockRecorder) DoMulti(arg0 any, arg1 ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

// DoMultiStream mocks base method.
func (m *Client) DoMultiStream(arg0 context.Context, arg1 ...rueidis.Completed) rueidis.MultiRedisResultStream {
	_ = "STUB: not implemented"
	return *new(rueidis.MultiRedisResultStream)
}

// DoMultiStream indicates an expected call of DoMultiStream.
func (mr *ClientMockRecorder) DoMultiStream(arg0 any, arg1 ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

// DoMultiCache mocks the base method.
func (m *Client) DoMultiCache(arg0 context.Context, arg1 ...rueidis.CacheableTTL) []rueidis.RedisResult {
	_ = "STUB: not implemented"
	return nil
}

// DoMultiCache indicates an expected call of DoMultiCache.
func (mr *ClientMockRecorder) DoMultiCache(arg0 any, arg1 ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

// Nodes mocks the base method.
func (m *Client) Nodes() map[string]rueidis.Client { _ = "STUB: not implemented"; return nil }

// Nodes indicates an expected call of Nodes.
func (mr *ClientMockRecorder) Nodes() *gomock.Call { _ = "STUB: not implemented"; return nil }

// Receive mocks base method.
func (m *Client) Receive(arg0 context.Context, arg1 rueidis.Completed, arg2 func(rueidis.PubSubMessage)) error {
	_ = "STUB: not implemented"
	return nil
}

// Receive indicates an expected call of Receive.
func (mr *ClientMockRecorder) Receive(arg0, arg1, arg2 any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

// DedicatedClient mocks the DedicatedClient interface.
type DedicatedClient struct {
	ctrl     *gomock.Controller
	recorder *DedicatedClientMockRecorder
	slot     uint16
}

// DedicatedClientMockRecorder is the mock recorder for DedicatedClient.
type DedicatedClientMockRecorder struct {
	mock *DedicatedClient
}

// NewDedicatedClient creates a new mock instance.
func NewDedicatedClient(ctrl *gomock.Controller, options ...ClientOption) *DedicatedClient {
	_ = "STUB: not implemented"
	return nil
}

// EXPECT returns an object that allows the caller to indicate the expected use.
func (m *DedicatedClient) EXPECT() *DedicatedClientMockRecorder {
	_ = "STUB: not implemented"

	// B mocks base method.
	return nil
}

func (m *DedicatedClient) B() rueidis.Builder {
	_ = "STUB: not implemented"
	return *new(rueidis.Builder)
}

// Close mocks base method.
func (m *DedicatedClient) Close() { _ = "STUB: not implemented"; return }

// Close indicates an expected call of Close.
func (mr *DedicatedClientMockRecorder) Close() *gomock.Call { _ = "STUB: not implemented"; return nil }

// Do mocks the base method.
func (m *DedicatedClient) Do(arg0 context.Context, arg1 rueidis.Completed) rueidis.RedisResult {
	_ = "STUB: not implemented"
	return *new(rueidis.RedisResult)
}

// Do indicates an expected call of Do.
func (mr *DedicatedClientMockRecorder) Do(arg0, arg1 any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

// DoMulti mocks base method.
func (m *DedicatedClient) DoMulti(arg0 context.Context, arg1 ...rueidis.Completed) []rueidis.RedisResult {
	_ = "STUB: not implemented"
	return nil
}

// DoMulti indicates an expected call of DoMulti.
func (mr *DedicatedClientMockRecorder) DoMulti(arg0 any, arg1 ...any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

// Receive mocks base method.
func (m *DedicatedClient) Receive(arg0 context.Context, arg1 rueidis.Completed, arg2 func(rueidis.PubSubMessage)) error {
	_ = "STUB: not implemented"
	return nil
}

// Receive indicates an expected call of Receive.
func (mr *DedicatedClientMockRecorder) Receive(arg0, arg1, arg2 any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

// SetPubSubHooks mocks base method.
func (m *DedicatedClient) SetPubSubHooks(arg0 rueidis.PubSubHooks) <-chan error {
	_ = "STUB: not implemented"
	return nil
}

// SetPubSubHooks indicates an expected call of SetPubSubHooks.
func (mr *DedicatedClientMockRecorder) SetPubSubHooks(arg0 any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

// SetOnInvalidations mocks base method.
func (m *DedicatedClient) SetOnInvalidations(arg0 func([]rueidis.RedisMessage)) <-chan error {
	_ = "STUB: not implemented"
	return nil
}

// SetOnInvalidations indicates an expected call of SetOnInvalidations.
func (mr *DedicatedClientMockRecorder) SetOnInvalidations(arg0 any) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

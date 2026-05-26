package rueidisrdma

/*
#cgo LDFLAGS: -libverbs -lrdmacm
#include <errno.h>
#include <stdlib.h>
#include "conn_linux.h"
int rdmaConnect(RdmaContext *ctx, const char *addr, int port, long timeout_msec);
ssize_t rdmaRead(RdmaContext *ctx, char *buf, size_t bufcap, long timeout_msec);
ssize_t rdmaWrite(RdmaContext *ctx, const char *obuf, size_t data_len, long timeout_msec);
void rdmaClose(RdmaContext *ctx);
void rdmaDisconnect(RdmaContext *ctx);
*/
import "C"

import (
	"context"
	"crypto/tls"
	"net"
	"sync"
	"time"
)

var _ net.Conn = (*conn)(nil)

func DialCtxFn(ctx context.Context, dst string, _ *net.Dialer, _ *tls.Config) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

type conn struct {
	ctx   *C.RdmaContext
	mu    sync.RWMutex
	timed int64
	once  int32
}

func (c *conn) timeout() int64 { _ = "STUB: not implemented"; return 0 }

func (c *conn) Read(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (c *conn) Write(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (c *conn) Close() error { _ = "STUB: not implemented"; return nil }

func (c *conn) err() error { _ = "STUB: not implemented"; return nil }

func (c *conn) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// Deadline already passed; use immediate timeout.

func (c *conn) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (c *conn) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (c *conn) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (c *conn) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

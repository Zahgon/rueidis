package om

import (
	"context"
	"errors"

	"github.com/redis/rueidis"
)

var EndOfCursor = errors.New("end of cursor")

func newAggregateCursor(idx string, client rueidis.Client, first []map[string]string, cursor, total int64) *AggregateCursor {
	_ = "STUB: not implemented"
	return nil
}

// AggregateCursor unifies the response of FT.AGGREGATE with or without WITHCURSOR
type AggregateCursor struct {
	client rueidis.Client
	idx    string
	first  []map[string]string
	id     int64
	n      int64
}

// Total return the total numbers of records of the initial FT.AGGREGATE result
func (c *AggregateCursor) Total() int64 {
	_ = "STUB: not implemented"

	// Read return the partial result from the initial FT.AGGREGATE
	// This may invoke FT.CURSOR READ to retrieve a further result
	return 0
}

func (c *AggregateCursor) Read(ctx context.Context) (partial []map[string]string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Del uses FT.CURSOR DEL to destroy the cursor
func (c *AggregateCursor) Del(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

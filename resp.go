package rueidis

import (
	"bufio"
	"errors"
	"io"
	"sync"
)

var errChunked = errors.New("unbounded redis message")
var errOldNull = errors.New("RESP2 null")

const (
	typeBlobString     = byte('$')
	typeSimpleString   = byte('+')
	typeSimpleErr      = byte('-')
	typeInteger        = byte(':')
	typeNull           = byte('_')
	typeEnd            = byte('.')
	typeFloat          = byte(',')
	typeBool           = byte('#')
	typeBlobErr        = byte('!')
	typeVerbatimString = byte('=')
	typeBigNumber      = byte('(')
	typeArray          = byte('*')
	typeMap            = byte('%')
	typeSet            = byte('~')
	typeAttribute      = byte('|')
	typePush           = byte('>')
	typeChunk          = byte(';')
)

var typeNames = make(map[byte]string, 16)

type reader func(i *bufio.Reader) (RedisMessage, error)

var readers = [256]reader{}

func init() {
	readers[typeBlobString] = readBlobString
	readers[typeSimpleString] = readSimpleString
	readers[typeSimpleErr] = readSimpleString
	readers[typeInteger] = readInteger
	readers[typeNull] = readNull
	readers[typeFloat] = readSimpleString
	readers[typeBool] = readBoolean
	readers[typeBlobErr] = readBlobString
	readers[typeVerbatimString] = readBlobString
	readers[typeBigNumber] = readSimpleString
	readers[typeArray] = readArray
	readers[typeMap] = readMap
	readers[typeSet] = readArray
	readers[typeAttribute] = readMap
	readers[typePush] = readArray
	readers[typeEnd] = readNull

	typeNames[typeBlobString] = "blob string"
	typeNames[typeSimpleString] = "simple string"
	typeNames[typeSimpleErr] = "simple error"
	typeNames[typeInteger] = "int64"
	typeNames[typeNull] = "null"
	typeNames[typeFloat] = "float64"
	typeNames[typeBool] = "boolean"
	typeNames[typeBlobErr] = "blob error"
	typeNames[typeVerbatimString] = "verbatim string"
	typeNames[typeBigNumber] = "big number"
	typeNames[typeArray] = "array"
	typeNames[typeMap] = "map"
	typeNames[typeSet] = "set"
	typeNames[typeAttribute] = "attribute"
	typeNames[typePush] = "push"
	typeNames[typeEnd] = "null"
}

func readSimpleString(i *bufio.Reader) (m RedisMessage, err error) {
	_ = "STUB: not implemented"
	return *new(RedisMessage), nil
}

func readBlobString(i *bufio.Reader) (m RedisMessage, err error) {
	_ = "STUB: not implemented"
	return *new(RedisMessage), nil
}

// discard the ';'

func readInteger(i *bufio.Reader) (m RedisMessage, err error) {
	_ = "STUB: not implemented"
	return *new(RedisMessage), nil
}

func readBoolean(i *bufio.Reader) (m RedisMessage, err error) {
	_ = "STUB: not implemented"
	return *new(RedisMessage), nil
}

func readNull(i *bufio.Reader) (m RedisMessage, err error) {
	_ = "STUB: not implemented"
	return *new(RedisMessage), nil
}

func readArray(i *bufio.Reader) (m RedisMessage, err error) {
	_ = "STUB: not implemented"
	return *new(RedisMessage), nil
}

func readMap(i *bufio.Reader) (m RedisMessage, err error) {
	_ = "STUB: not implemented"
	return *new(RedisMessage), nil
}

const ok = "OK"
const okrn = "OK\r\n"

func readS(i *bufio.Reader) (*byte, int64, error) { _ = "STUB: not implemented"; return nil, 0, nil }

func readI(i *bufio.Reader) (v int64, err error) { _ = "STUB: not implemented"; return 0, nil }

func readB(i *bufio.Reader) (*byte, int64, error) { _ = "STUB: not implemented"; return nil, 0, nil }

func readE(i *bufio.Reader) (*RedisMessage, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func readA(i *bufio.Reader, length int64) (*RedisMessage, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func writeB(o *bufio.Writer, id byte, str string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func writeS(o *bufio.Writer, id byte, str string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func writeN(o *bufio.Writer, id byte, n int) (err error) { _ = "STUB: not implemented"; return nil }

func readNextMessage(i *bufio.Reader) (m RedisMessage, err error) {
	_ = "STUB: not implemented"
	return *new(RedisMessage), nil
}

// handle the attributes
// clone the original m first, and then take the address of the clone
// to avoid go compiler allocating the m on heap which causing worse performance.

var lrs = sync.Pool{New: func() any { return &io.LimitedReader{} }}

func streamTo(i *bufio.Reader, w io.Writer) (n int64, err error, clean bool) {
	_ = "STUB: not implemented"
	return 0, nil, false
}

func writeCmd(o *bufio.Writer, cmd []string) (err error) { _ = "STUB: not implemented"; return nil }

// TODO: Can we set cmd[i] = "" here to allow GC to eagerly recycle memory?
// Related: https://github.com/redis/rueidis/issues/364

func flushCmd(o *bufio.Writer, cmd []string) (err error) { _ = "STUB: not implemented"; return nil }

const (
	unexpectedNoCRLF   = "received unexpected simple string message ending without CRLF"
	unexpectedNumByte  = "received unexpected number byte: "
	unknownMessageType = "received unknown message type: "
)

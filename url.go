package rueidis

// ParseURL parses a redis URL into ClientOption.
// https://github.com/redis/redis-specifications/blob/master/uri/redis.txt
// Example:
//
// redis://<user>:<password>@<host>:<port>/<db_number>
// redis://<user>:<password>@<host>:<port>?addr=<host2>:<port2>&addr=<host3>:<port3>
// unix://<user>:<password>@</path/to/redis.sock>?db=<db_number>
func ParseURL(str string) (opt ClientOption, err error) {
	_ = "STUB: not implemented"
	return *new(ClientOption), nil
}

func MustParseURL(str string) ClientOption { _ = "STUB: not implemented"; return *new(ClientOption) }

package main

import (
	"encoding/json"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
)

type goStruct struct {
	Node      *node
	FullName  string
	BuildDef  buildDef
	NextNodes []*node
	Variadic  bool

	MultipleToken bool
}

type buildDef struct {
	MethodName string
	Command    []string
	Parameters []parameter
}

func (d buildDef) hash() uint32 { _ = "STUB: not implemented"; return 0 }

type parameter struct {
	Name string
	Type string
}

type command struct {
	Group     string     `json:"group"`
	Since     string     `json:"since"`
	Arguments []argument `json:"arguments"`
}

type argument struct {
	Name      any        `json:"name"`
	Type      any        `json:"type"`
	Command   string     `json:"command"`
	Token     string     `json:"token"`
	Enum      []string   `json:"enum"`
	Block     []argument `json:"block"`
	Arguments []argument `json:"arguments"`
	Multiple  bool       `json:"multiple"`
	Optional  bool       `json:"optional"`
	Variadic  bool       `json:"variadic"`

	MultipleToken bool `json:"multiple_token"`
}

type node struct {
	Group  string
	Parent *node
	Child  *node
	Next   *node
	Cmd    command
	Arg    argument
	Root   bool
}

func (n *node) FindRoot() (root *node) { _ = "STUB: not implemented"; return nil }

//gocyclo:ignore
func (n *node) GoStructs() (out []goStruct) {
	_ = "STUB: not implemented"

	// fix for XGROUP and XADD
	return nil
}

// fix for TS.MRANGE, TS.MREVRANGE, TS.MGET, TS.QUERYINDEX

// FIXME: this should be handle differently later

// Using time.Duration for EX/PX

// FIXME: this should be handle differently later

// Using time.Time for EXAT/PXAT

// fix for FT.AGGREGATE

// fix for TS.MRANGE, TS.MREVRANGE, TS.MGET

// not change to go type here, change at render

// ignore

func (n *node) MultipleToken() bool { _ = "STUB: not implemented"; return false }

func (n *node) Variadic() bool { _ = "STUB: not implemented"; return false }

func (n *node) FullName() (out string) { _ = "STUB: not implemented"; return "" }

func (n *node) Name() (out string) { _ = "STUB: not implemented"; return "" }

func (n *node) NextNodes() (nodes []*node) { _ = "STUB: not implemented"; return nil }

// block variadic

// don't climb to root

func (n *node) Walk(fn func(node *node)) { _ = "STUB: not implemented"; return }

var (
	inputglob = "*.json"
	outputdir = "../../internal/cmds"
)

// Usages:
// 1) cd hack/cmds && go run gen.go
// 2) go run hack/cmds/gen.go internal/cmds hack/cmds/*.json
func main() {
	var err error
	var defs []string
	var structs = map[string]map[string]goStruct{}

	if len(os.Args) > 1 {
		outputdir = os.Args[1]
	}
	if len(os.Args) > 2 {
		defs = os.Args[2:]
		if len(defs) == 1 {
			defs, err = filepath.Glob(defs[0])
		}
	} else {
		defs, err = filepath.Glob(inputglob)
	}
	if err != nil {
		panic(err)
	}

	for _, p := range defs {
		raw, err := os.ReadFile(p)
		if err != nil {
			panic(err)
		}

		var commands = map[string]command{}
		if err := json.Unmarshal(raw, &commands); err != nil {
			panic(err)
		}

		var roots []string
		nodes := map[string]*node{}
		for k, cmd := range commands {
			if cmd.Group == "" {
				panic(k + " no group")
			}
			root := &node{Group: cmd.Group, Cmd: cmd, Arg: argument{Name: k, Command: k, Type: "command"}, Root: true}
			root.Next = makeChildNodes(root, cmd.Arguments)
			roots = append(roots, k)
			if _, ok := nodes[k]; ok {
				panic(k + " conflict")
			}
			nodes[k] = root
		}
		sort.Strings(roots)

		for _, name := range roots {
			n := nodes[name]
			g := n.Group
			if _, ok := structs[g]; !ok {
				structs[g] = make(map[string]goStruct)
			}
			n.Walk(func(n *node) {
				for _, s := range n.GoStructs() {
					if v, ok := structs[g][s.FullName]; ok {
						if !reflect.DeepEqual(v, s) {
							panic("struct conflict " + v.FullName)
						}
					}
					structs[g][s.FullName] = s
				}
			})
		}
	}

	for g, structs := range structs {
		gfname := filepath.Join(outputdir, "gen_"+g+".go")
		tfname := filepath.Join(outputdir, "gen_"+g+"_test.go")
		gf, err := os.Create(gfname)
		if err != nil {
			panic(err)
		}
		generate(gf, structs)
		if err := gf.Close(); err != nil {
			panic(err)
		}
		if err := exec.Command("gofmt", "-w", gfname).Run(); err != nil {
			panic(err)
		}
		if err := exec.Command("goimports", "-w", gfname).Run(); err != nil {
			panic(err)
		}
		tf, err := os.Create(tfname)
		if err != nil {
			panic(err)
		}
		tests(tf, structs, g)
		if err := tf.Close(); err != nil {
			panic(err)
		}
		if err := exec.Command("gofmt", "-w", tfname).Run(); err != nil {
			panic(err)
		}
		if err := exec.Command("goimports", "-w", tfname).Run(); err != nil {
			panic(err)
		}
	}

	checkAllUsed("noRetCMDs", noRetCMDs)
	checkAllUsed("unsubCMDs", unsubCMDs)
	checkAllUsed("blockingCMDs", blockingCMDs)
	checkAllUsed("cacheableCMDs", cacheableCMDs)
	checkAllUsed("readOnlyCMDs", readOnlyCMDs)
}

func tests(f io.Writer, structs map[string]goStruct, prefix string) {
	_ = "STUB: not implemented"
	return
}

var pathmark = map[string]bool{}
var blockmark = map[*node]bool{}

func makePath(s goStruct, path []goStruct, paths [][]goStruct) [][]goStruct {
	_ = "STUB: not implemented"
	return nil
}

func testParams(defs []parameter) string { _ = "STUB: not implemented"; return "" }

func printPath(f io.Writer, receiver string, path []goStruct, end string) {
	_ = "STUB: not implemented"
	return
}

func generate(f io.Writer, structs map[string]goStruct) { _ = "STUB: not implemented"; return }

func checkAllUsed(name string, tags map[string]bool) { _ = "STUB: not implemented"; return }

func allOptional(s *node, nodes []*node) bool { _ = "STUB: not implemented"; return false }

func toGoType(paramType string) string { _ = "STUB: not implemented"; return "" }

// TODO hack for TS.MRANGE, TS.MREVRANGE, TS.MGET

// TODO hack for FT.CREATE VECTOR

func toGoName(paramName string) string { _ = "STUB: not implemented"; return "" }

func printRootBuilder(w io.Writer, root goStruct) { _ = "STUB: not implemented"; return }

func rootCf(root goStruct) (tag string) { _ = "STUB: not implemented"; return "" }

func printFinalBuilder(w io.Writer, parent goStruct, method, ss string) {
	_ = "STUB: not implemented"
	return
}

//gocyclo:ignore
func printBuilder(w io.Writer, parent, next goStruct) { _ = "STUB: not implemented"; return }

// no parameter

// no parameter

// no parameter

// TODO hack for TS.MRANGE, TS.MREVRANGE, TS.MGET

// TODO hack for FT.CREATE VECTOR

// For seconds

// For seconds

// For milliseconds

func makeChildNodes(parent *node, args []argument) (first *node) {
	_ = "STUB: not implemented"
	return nil
}

func blockEntries(block *node) (nodes []*node) { _ = "STUB: not implemented"; return nil }

func name(s string) (name string) { _ = "STUB: not implemented"; return "" }

func ucFirst(str string) string { _ = "STUB: not implemented"; return "" }

func lcFirst(str string) string { _ = "STUB: not implemented"; return "" }

func within(cmd goStruct, cmds map[string]bool) bool { _ = "STUB: not implemented"; return false }

var noRetCMDs = map[string]bool{
	"subscribe":  false,
	"psubscribe": false,
	"ssubscribe": false,
}

var unsubCMDs = map[string]bool{
	"unsubscribe":  false,
	"punsubscribe": false,
	"sunsubscribe": false,
}

var mtGetCMDs = map[string]bool{
	"mget":     false,
	"jsonmget": false,
}

var scrRoCMDs = map[string]bool{
	"fcallro":   false,
	"evalsharo": false,
	"evalro":    false,
}

var blockingCMDs = map[string]bool{
	"blpop":       false,
	"brpop":       false,
	"brpoplpush":  false,
	"blmove":      false,
	"blmpop":      false,
	"bzpopmin":    false,
	"bzpopmax":    false,
	"bzmpop":      false,
	"clientpause": false,
	"migrate":     false,
	"wait":        false,
	"waitaof":     false,
}

var cacheableCMDs = map[string]bool{
	"bitcount":            false,
	"bitfieldro":          false,
	"bitpos":              false,
	"expiretime":          false,
	"geodist":             false,
	"geohash":             false,
	"geopos":              false,
	"georadiusro":         false,
	"georadiusbymemberro": false,
	"geosearch":           false,
	"get":                 false,
	"mget":                false,
	"getbit":              false,
	"getrange":            false,
	"hexists":             false,
	"hget":                false,
	"hgetall":             false,
	"hkeys":               false,
	"hlen":                false,
	"hmget":               false,
	"hstrlen":             false,
	"hvals":               false,
	"lindex":              false,
	"llen":                false,
	"lpos":                false,
	"lrange":              false,
	"pexpiretime":         false,
	"pttl":                false,
	"scard":               false,
	"sismember":           false,
	"smembers":            false,
	"smismember":          false,
	"sortro":              false,
	"strlen":              false,
	"ttl":                 false,
	"type":                false,
	"zcard":               false,
	"zcount":              false,
	"zlexcount":           false,
	"zmscore":             false,
	"zrange":              false,
	"zrangebylex":         false,
	"zrangebyscore":       false,
	"zrank":               false,
	"zrevrange":           false,
	"zrevrangebylex":      false,
	"zrevrangebyscore":    false,
	"zrevrank":            false,
	"zscore":              false,
	"jsonget":             false,
	"jsonmget":            false,
	"jsonstrlen":          false,
	"jsonarrindex":        false,
	"jsonarrlen":          false,
	"jsonobjkeys":         false,
	"jsonobjlen":          false,
	"jsontype":            false,
	"jsonresp":            false,
	"bfexists":            false,
	"bfinfo":              false,
	"cfexists":            false,
	"cfcount":             false,
	"cfinfo":              false,
	"cmsquery":            false,
	"cmsinfo":             false,
	"topkquery":           false,
	"topklist":            false,
	"topkinfo":            false,
	"fcallro":             false,
	"evalsharo":           false,
	"evalro":              false,
	"graphroquery":        false,
	"aitensorget":         false,
	"aimodelget":          false,
	"aimodelexecute":      false,
	"aiscriptget":         false,
}

var readOnlyCMDs = map[string]bool{
	"bitcount":            false,
	"bitfieldro":          false,
	"bitpos":              false,
	"dbsize":              false,
	"dump":                false,
	"exists":              false,
	"expiretime":          false,
	"geodist":             false,
	"geohash":             false,
	"geopos":              false,
	"georadiusro":         false,
	"georadiusbymemberro": false,
	"geosearch":           false,
	"get":                 false,
	"getbit":              false,
	"getrange":            false,
	"hexists":             false,
	"hget":                false,
	"hgetall":             false,
	"hkeys":               false,
	"hlen":                false,
	"hmget":               false,
	"hrandfield":          false,
	"hscan":               false,
	"hstrlen":             false,
	"hvals":               false,
	"keys":                false,
	"lindex":              false,
	"llen":                false,
	"lolwut":              false,
	"lpos":                false,
	"lrange":              false,
	"memorydoctor":        false,
	"memorystats":         false,
	"memoryusage":         false,
	"memorymallocstats":   false,
	"objectencoding":      false,
	"objectfreq":          false,
	"objecthelp":          false,
	"objectidletime":      false,
	"objectrefcount":      false,
	"pexpiretime":         false,
	"pfcount":             false,
	"pttl":                false,
	"pubsubchannels":      false,
	"pubsubnumpat":        false,
	"pubsubnumsub":        false,
	"pubsubhelp":          false,
	"randomkey":           false,
	"scan":                false,
	"scard":               false,
	"sdiff":               false,
	"sinter":              false,
	"sintercard":          false,
	"sismember":           false,
	"slowlogget":          false,
	"slowloglen":          false,
	"slowloghelp":         false,
	"smembers":            false,
	"smismember":          false,
	"sortro":              false,
	"srandmember":         false,
	"sscan":               false,
	"lcs":                 false,
	"strlen":              false,
	"sunion":              false,
	"touch":               false,
	"ttl":                 false,
	"type":                false,
	"xinfoconsumers":      false,
	"xinfogroups":         false,
	"xinfostream":         false,
	"xinfohelp":           false,
	"xlen":                false,
	"xpending":            false,
	"xrange":              false,
	"xread":               false,
	"xrevrange":           false,
	"zcard":               false,
	"zcount":              false,
	"zdiff":               false,
	"zinter":              false,
	"zlexcount":           false,
	"zmscore":             false,
	"zrandmember":         false,
	"zrange":              false,
	"zrangebylex":         false,
	"zrangebyscore":       false,
	"zrank":               false,
	"zrevrange":           false,
	"zrevrangebylex":      false,
	"zrevrangebyscore":    false,
	"zrevrank":            false,
	"zscan":               false,
	"zscore":              false,
	"zunion":              false,
	"zintercard":          false,
	"jsonget":             false,
	"jsonstrlen":          false,
	"jsonarrindex":        false,
	"jsonarrlen":          false,
	"jsonobjkeys":         false,
	"jsonobjlen":          false,
	"jsontype":            false,
	"jsonresp":            false,
	"bfexists":            false,
	"bfmexists":           false,
	"bfscandump":          false,
	"bfinfo":              false,
	"cfexists":            false,
	"cfcount":             false,
	"cfscandump":          false,
	"cfinfo":              false,
	"cmsquery":            false,
	"cmsinfo":             false,
	"topkquery":           false,
	"topklist":            false,
	"topkinfo":            false,
	"tsrange":             false,
	"tsrevrange":          false,
	"tsget":               false,
	"tsinfo":              false,
	"tsqueryindex":        false,
	"graphroquery":        false,
	"graphexplain":        false,
	"graphslowlog":        false,
	"graphconfigget":      false,
	"graphlist":           false,
	"aitensorget":         false,
	"aimodelget":          false,
	"aimodelexecute":      false,
	"aiscriptget":         false,
	"ftsearch":            false,
	"ftaggregate":         false,
}

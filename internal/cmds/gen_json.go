// Code generated DO NOT EDIT

package cmds

type JsonArrappend Incomplete

func (b Builder) JsonArrappend() (c JsonArrappend) {
	_ = "STUB: not implemented"
	return *new(JsonArrappend)
}

func (c JsonArrappend) Key(key string) JsonArrappendKey {
	_ = "STUB: not implemented"
	return *new(JsonArrappendKey)
}

type JsonArrappendKey Incomplete

func (c JsonArrappendKey) Path(path string) JsonArrappendPath {
	_ = "STUB: not implemented"
	return *new(JsonArrappendPath)
}

func (c JsonArrappendKey) Value(value ...string) JsonArrappendValue {
	_ = "STUB: not implemented"
	return *new(JsonArrappendValue)
}

type JsonArrappendPath Incomplete

func (c JsonArrappendPath) Value(value ...string) JsonArrappendValue {
	_ = "STUB: not implemented"
	return *new(JsonArrappendValue)
}

type JsonArrappendValue Incomplete

func (c JsonArrappendValue) Value(value ...string) JsonArrappendValue {
	_ = "STUB: not implemented"
	return *new(JsonArrappendValue)
}

func (c JsonArrappendValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type JsonArrindex Incomplete

func (b Builder) JsonArrindex() (c JsonArrindex) {
	_ = "STUB: not implemented"
	return *new(JsonArrindex)
}

func (c JsonArrindex) Key(key string) JsonArrindexKey {
	_ = "STUB: not implemented"
	return *new(JsonArrindexKey)
}

type JsonArrindexKey Incomplete

func (c JsonArrindexKey) Path(path string) JsonArrindexPath {
	_ = "STUB: not implemented"
	return *new(JsonArrindexPath)
}

type JsonArrindexPath Incomplete

func (c JsonArrindexPath) Value(value string) JsonArrindexValue {
	_ = "STUB: not implemented"
	return *new(JsonArrindexValue)
}

type JsonArrindexStartStart Incomplete

func (c JsonArrindexStartStart) Stop(stop int64) JsonArrindexStartStop {
	_ = "STUB: not implemented"
	return *new(JsonArrindexStartStop)
}

func (c JsonArrindexStartStart) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c JsonArrindexStartStart) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type JsonArrindexStartStop Incomplete

func (c JsonArrindexStartStop) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c JsonArrindexStartStop) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type JsonArrindexValue Incomplete

func (c JsonArrindexValue) Start(start int64) JsonArrindexStartStart {
	_ = "STUB: not implemented"
	return *new(JsonArrindexStartStart)
}

func (c JsonArrindexValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c JsonArrindexValue) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type JsonArrinsert Incomplete

func (b Builder) JsonArrinsert() (c JsonArrinsert) {
	_ = "STUB: not implemented"
	return *new(JsonArrinsert)
}

func (c JsonArrinsert) Key(key string) JsonArrinsertKey {
	_ = "STUB: not implemented"
	return *new(JsonArrinsertKey)
}

type JsonArrinsertIndex Incomplete

func (c JsonArrinsertIndex) Value(value ...string) JsonArrinsertValue {
	_ = "STUB: not implemented"
	return *new(JsonArrinsertValue)
}

type JsonArrinsertKey Incomplete

func (c JsonArrinsertKey) Path(path string) JsonArrinsertPath {
	_ = "STUB: not implemented"
	return *new(JsonArrinsertPath)
}

type JsonArrinsertPath Incomplete

func (c JsonArrinsertPath) Index(index int64) JsonArrinsertIndex {
	_ = "STUB: not implemented"
	return *new(JsonArrinsertIndex)
}

type JsonArrinsertValue Incomplete

func (c JsonArrinsertValue) Value(value ...string) JsonArrinsertValue {
	_ = "STUB: not implemented"
	return *new(JsonArrinsertValue)
}

func (c JsonArrinsertValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type JsonArrlen Incomplete

func (b Builder) JsonArrlen() (c JsonArrlen) { _ = "STUB: not implemented"; return *new(JsonArrlen) }

func (c JsonArrlen) Key(key string) JsonArrlenKey {
	_ = "STUB: not implemented"
	return *new(JsonArrlenKey)
}

type JsonArrlenKey Incomplete

func (c JsonArrlenKey) Path(path string) JsonArrlenPath {
	_ = "STUB: not implemented"
	return *new(JsonArrlenPath)
}

func (c JsonArrlenKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c JsonArrlenKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type JsonArrlenPath Incomplete

func (c JsonArrlenPath) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c JsonArrlenPath) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type JsonArrpop Incomplete

func (b Builder) JsonArrpop() (c JsonArrpop) { _ = "STUB: not implemented"; return *new(JsonArrpop) }

func (c JsonArrpop) Key(key string) JsonArrpopKey {
	_ = "STUB: not implemented"
	return *new(JsonArrpopKey)
}

type JsonArrpopKey Incomplete

func (c JsonArrpopKey) Path(path string) JsonArrpopPathPath {
	_ = "STUB: not implemented"
	return *new(JsonArrpopPathPath)
}

func (c JsonArrpopKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type JsonArrpopPathIndex Incomplete

func (c JsonArrpopPathIndex) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type JsonArrpopPathPath Incomplete

func (c JsonArrpopPathPath) Index(index int64) JsonArrpopPathIndex {
	_ = "STUB: not implemented"
	return *new(JsonArrpopPathIndex)
}

func (c JsonArrpopPathPath) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type JsonArrtrim Incomplete

func (b Builder) JsonArrtrim() (c JsonArrtrim) { _ = "STUB: not implemented"; return *new(JsonArrtrim) }

func (c JsonArrtrim) Key(key string) JsonArrtrimKey {
	_ = "STUB: not implemented"
	return *new(JsonArrtrimKey)
}

type JsonArrtrimKey Incomplete

func (c JsonArrtrimKey) Path(path string) JsonArrtrimPath {
	_ = "STUB: not implemented"
	return *new(JsonArrtrimPath)
}

type JsonArrtrimPath Incomplete

func (c JsonArrtrimPath) Start(start int64) JsonArrtrimStart {
	_ = "STUB: not implemented"
	return *new(JsonArrtrimStart)
}

type JsonArrtrimStart Incomplete

func (c JsonArrtrimStart) Stop(stop int64) JsonArrtrimStop {
	_ = "STUB: not implemented"
	return *new(JsonArrtrimStop)
}

type JsonArrtrimStop Incomplete

func (c JsonArrtrimStop) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type JsonClear Incomplete

func (b Builder) JsonClear() (c JsonClear) { _ = "STUB: not implemented"; return *new(JsonClear) }

func (c JsonClear) Key(key string) JsonClearKey {
	_ = "STUB: not implemented"
	return *new(JsonClearKey)
}

type JsonClearKey Incomplete

func (c JsonClearKey) Path(path string) JsonClearPath {
	_ = "STUB: not implemented"
	return *new(JsonClearPath)
}

func (c JsonClearKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type JsonClearPath Incomplete

func (c JsonClearPath) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type JsonDebugHelp Incomplete

func (b Builder) JsonDebugHelp() (c JsonDebugHelp) {
	_ = "STUB: not implemented"
	return *new(JsonDebugHelp)
}

func (c JsonDebugHelp) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type JsonDebugMemory Incomplete

func (b Builder) JsonDebugMemory() (c JsonDebugMemory) {
	_ = "STUB: not implemented"
	return *new(JsonDebugMemory)
}

func (c JsonDebugMemory) Key(key string) JsonDebugMemoryKey {
	_ = "STUB: not implemented"
	return *new(JsonDebugMemoryKey)
}

type JsonDebugMemoryKey Incomplete

func (c JsonDebugMemoryKey) Path(path string) JsonDebugMemoryPath {
	_ = "STUB: not implemented"
	return *new(JsonDebugMemoryPath)
}

func (c JsonDebugMemoryKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type JsonDebugMemoryPath Incomplete

func (c JsonDebugMemoryPath) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type JsonDel Incomplete

func (b Builder) JsonDel() (c JsonDel) { _ = "STUB: not implemented"; return *new(JsonDel) }

func (c JsonDel) Key(key string) JsonDelKey { _ = "STUB: not implemented"; return *new(JsonDelKey) }

type JsonDelKey Incomplete

func (c JsonDelKey) Path(path string) JsonDelPath {
	_ = "STUB: not implemented"
	return *new(JsonDelPath)
}

func (c JsonDelKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type JsonDelPath Incomplete

func (c JsonDelPath) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type JsonForget Incomplete

func (b Builder) JsonForget() (c JsonForget) { _ = "STUB: not implemented"; return *new(JsonForget) }

func (c JsonForget) Key(key string) JsonForgetKey {
	_ = "STUB: not implemented"
	return *new(JsonForgetKey)
}

type JsonForgetKey Incomplete

func (c JsonForgetKey) Path(path string) JsonForgetPath {
	_ = "STUB: not implemented"
	return *new(JsonForgetPath)
}

func (c JsonForgetKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type JsonForgetPath Incomplete

func (c JsonForgetPath) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type JsonGet Incomplete

func (b Builder) JsonGet() (c JsonGet) { _ = "STUB: not implemented"; return *new(JsonGet) }

func (c JsonGet) Key(key string) JsonGetKey { _ = "STUB: not implemented"; return *new(JsonGetKey) }

type JsonGetIndent Incomplete

func (c JsonGetIndent) Newline(newline string) JsonGetNewline {
	_ = "STUB: not implemented"
	return *new(JsonGetNewline)
}

func (c JsonGetIndent) Space(space string) JsonGetSpace {
	_ = "STUB: not implemented"
	return *new(JsonGetSpace)
}

func (c JsonGetIndent) Path(path ...string) JsonGetPath {
	_ = "STUB: not implemented"
	return *new(JsonGetPath)
}

func (c JsonGetIndent) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c JsonGetIndent) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type JsonGetKey Incomplete

func (c JsonGetKey) Indent(indent string) JsonGetIndent {
	_ = "STUB: not implemented"
	return *new(JsonGetIndent)
}

func (c JsonGetKey) Newline(newline string) JsonGetNewline {
	_ = "STUB: not implemented"
	return *new(JsonGetNewline)
}

func (c JsonGetKey) Space(space string) JsonGetSpace {
	_ = "STUB: not implemented"
	return *new(JsonGetSpace)
}

func (c JsonGetKey) Path(path ...string) JsonGetPath {
	_ = "STUB: not implemented"
	return *new(JsonGetPath)
}

func (c JsonGetKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c JsonGetKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type JsonGetNewline Incomplete

func (c JsonGetNewline) Space(space string) JsonGetSpace {
	_ = "STUB: not implemented"
	return *new(JsonGetSpace)
}

func (c JsonGetNewline) Path(path ...string) JsonGetPath {
	_ = "STUB: not implemented"
	return *new(JsonGetPath)
}

func (c JsonGetNewline) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c JsonGetNewline) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type JsonGetPath Incomplete

func (c JsonGetPath) Path(path ...string) JsonGetPath {
	_ = "STUB: not implemented"
	return *new(JsonGetPath)
}

func (c JsonGetPath) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c JsonGetPath) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type JsonGetSpace Incomplete

func (c JsonGetSpace) Path(path ...string) JsonGetPath {
	_ = "STUB: not implemented"
	return *new(JsonGetPath)
}

func (c JsonGetSpace) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c JsonGetSpace) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type JsonMerge Incomplete

func (b Builder) JsonMerge() (c JsonMerge) { _ = "STUB: not implemented"; return *new(JsonMerge) }

func (c JsonMerge) Key(key string) JsonMergeKey {
	_ = "STUB: not implemented"
	return *new(JsonMergeKey)
}

type JsonMergeKey Incomplete

func (c JsonMergeKey) Path(path string) JsonMergePath {
	_ = "STUB: not implemented"
	return *new(JsonMergePath)
}

type JsonMergePath Incomplete

func (c JsonMergePath) Value(value string) JsonMergeValue {
	_ = "STUB: not implemented"
	return *new(JsonMergeValue)
}

type JsonMergeValue Incomplete

func (c JsonMergeValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type JsonMget Incomplete

func (b Builder) JsonMget() (c JsonMget) { _ = "STUB: not implemented"; return *new(JsonMget) }

func (c JsonMget) Key(key ...string) JsonMgetKey {
	_ = "STUB: not implemented"
	return *new(JsonMgetKey)
}

type JsonMgetKey Incomplete

func (c JsonMgetKey) Key(key ...string) JsonMgetKey {
	_ = "STUB: not implemented"
	return *new(JsonMgetKey)
}

func (c JsonMgetKey) Path(path string) JsonMgetPath {
	_ = "STUB: not implemented"
	return *new(JsonMgetPath)
}

type JsonMgetPath Incomplete

func (c JsonMgetPath) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c JsonMgetPath) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type JsonMset Incomplete

func (b Builder) JsonMset() (c JsonMset) { _ = "STUB: not implemented"; return *new(JsonMset) }

func (c JsonMset) Key(key string) JsonMsetTripletKey {
	_ = "STUB: not implemented"
	return *new(JsonMsetTripletKey)
}

type JsonMsetTripletKey Incomplete

func (c JsonMsetTripletKey) Path(path string) JsonMsetTripletPath {
	_ = "STUB: not implemented"
	return *new(JsonMsetTripletPath)
}

type JsonMsetTripletPath Incomplete

func (c JsonMsetTripletPath) Value(value string) JsonMsetTripletValue {
	_ = "STUB: not implemented"
	return *new(JsonMsetTripletValue)
}

type JsonMsetTripletValue Incomplete

func (c JsonMsetTripletValue) Key(key string) JsonMsetTripletKey {
	_ = "STUB: not implemented"
	return *new(JsonMsetTripletKey)
}

func (c JsonMsetTripletValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type JsonNumincrby Incomplete

func (b Builder) JsonNumincrby() (c JsonNumincrby) {
	_ = "STUB: not implemented"
	return *new(JsonNumincrby)
}

func (c JsonNumincrby) Key(key string) JsonNumincrbyKey {
	_ = "STUB: not implemented"
	return *new(JsonNumincrbyKey)
}

type JsonNumincrbyKey Incomplete

func (c JsonNumincrbyKey) Path(path string) JsonNumincrbyPath {
	_ = "STUB: not implemented"
	return *new(JsonNumincrbyPath)
}

type JsonNumincrbyPath Incomplete

func (c JsonNumincrbyPath) Value(value float64) JsonNumincrbyValue {
	_ = "STUB: not implemented"
	return *new(JsonNumincrbyValue)
}

type JsonNumincrbyValue Incomplete

func (c JsonNumincrbyValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type JsonNummultby Incomplete

func (b Builder) JsonNummultby() (c JsonNummultby) {
	_ = "STUB: not implemented"
	return *new(JsonNummultby)
}

func (c JsonNummultby) Key(key string) JsonNummultbyKey {
	_ = "STUB: not implemented"
	return *new(JsonNummultbyKey)
}

type JsonNummultbyKey Incomplete

func (c JsonNummultbyKey) Path(path string) JsonNummultbyPath {
	_ = "STUB: not implemented"
	return *new(JsonNummultbyPath)
}

type JsonNummultbyPath Incomplete

func (c JsonNummultbyPath) Value(value float64) JsonNummultbyValue {
	_ = "STUB: not implemented"
	return *new(JsonNummultbyValue)
}

type JsonNummultbyValue Incomplete

func (c JsonNummultbyValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type JsonObjkeys Incomplete

func (b Builder) JsonObjkeys() (c JsonObjkeys) { _ = "STUB: not implemented"; return *new(JsonObjkeys) }

func (c JsonObjkeys) Key(key string) JsonObjkeysKey {
	_ = "STUB: not implemented"
	return *new(JsonObjkeysKey)
}

type JsonObjkeysKey Incomplete

func (c JsonObjkeysKey) Path(path string) JsonObjkeysPath {
	_ = "STUB: not implemented"
	return *new(JsonObjkeysPath)
}

func (c JsonObjkeysKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c JsonObjkeysKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type JsonObjkeysPath Incomplete

func (c JsonObjkeysPath) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c JsonObjkeysPath) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type JsonObjlen Incomplete

func (b Builder) JsonObjlen() (c JsonObjlen) { _ = "STUB: not implemented"; return *new(JsonObjlen) }

func (c JsonObjlen) Key(key string) JsonObjlenKey {
	_ = "STUB: not implemented"
	return *new(JsonObjlenKey)
}

type JsonObjlenKey Incomplete

func (c JsonObjlenKey) Path(path string) JsonObjlenPath {
	_ = "STUB: not implemented"
	return *new(JsonObjlenPath)
}

func (c JsonObjlenKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c JsonObjlenKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type JsonObjlenPath Incomplete

func (c JsonObjlenPath) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c JsonObjlenPath) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type JsonResp Incomplete

func (b Builder) JsonResp() (c JsonResp) { _ = "STUB: not implemented"; return *new(JsonResp) }

func (c JsonResp) Key(key string) JsonRespKey { _ = "STUB: not implemented"; return *new(JsonRespKey) }

type JsonRespKey Incomplete

func (c JsonRespKey) Path(path string) JsonRespPath {
	_ = "STUB: not implemented"
	return *new(JsonRespPath)
}

func (c JsonRespKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c JsonRespKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type JsonRespPath Incomplete

func (c JsonRespPath) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c JsonRespPath) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type JsonSet Incomplete

func (b Builder) JsonSet() (c JsonSet) { _ = "STUB: not implemented"; return *new(JsonSet) }

func (c JsonSet) Key(key string) JsonSetKey { _ = "STUB: not implemented"; return *new(JsonSetKey) }

type JsonSetConditionNx Incomplete

func (c JsonSetConditionNx) Fpha() JsonSetFphaFpha {
	_ = "STUB: not implemented"
	return *new(JsonSetFphaFpha)
}

func (c JsonSetConditionNx) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type JsonSetConditionXx Incomplete

func (c JsonSetConditionXx) Fpha() JsonSetFphaFpha {
	_ = "STUB: not implemented"
	return *new(JsonSetFphaFpha)
}

func (c JsonSetConditionXx) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type JsonSetFphaFpha Incomplete

func (c JsonSetFphaFpha) Fp16() JsonSetFphaFphaTypeFp16 {
	_ = "STUB: not implemented"
	return *new(JsonSetFphaFphaTypeFp16)
}

func (c JsonSetFphaFpha) Bf16() JsonSetFphaFphaTypeBf16 {
	_ = "STUB: not implemented"
	return *new(JsonSetFphaFphaTypeBf16)
}

func (c JsonSetFphaFpha) Fp32() JsonSetFphaFphaTypeFp32 {
	_ = "STUB: not implemented"
	return *new(JsonSetFphaFphaTypeFp32)
}

func (c JsonSetFphaFpha) Fp64() JsonSetFphaFphaTypeFp64 {
	_ = "STUB: not implemented"
	return *new(JsonSetFphaFphaTypeFp64)
}

type JsonSetFphaFphaTypeBf16 Incomplete

func (c JsonSetFphaFphaTypeBf16) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type JsonSetFphaFphaTypeFp16 Incomplete

func (c JsonSetFphaFphaTypeFp16) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type JsonSetFphaFphaTypeFp32 Incomplete

func (c JsonSetFphaFphaTypeFp32) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type JsonSetFphaFphaTypeFp64 Incomplete

func (c JsonSetFphaFphaTypeFp64) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type JsonSetKey Incomplete

func (c JsonSetKey) Path(path string) JsonSetPath {
	_ = "STUB: not implemented"
	return *new(JsonSetPath)
}

type JsonSetPath Incomplete

func (c JsonSetPath) Value(value string) JsonSetValue {
	_ = "STUB: not implemented"
	return *new(JsonSetValue)
}

type JsonSetValue Incomplete

func (c JsonSetValue) Nx() JsonSetConditionNx {
	_ = "STUB: not implemented"
	return *new(JsonSetConditionNx)
}

func (c JsonSetValue) Xx() JsonSetConditionXx {
	_ = "STUB: not implemented"
	return *new(JsonSetConditionXx)
}

func (c JsonSetValue) Fpha() JsonSetFphaFpha {
	_ = "STUB: not implemented"
	return *new(JsonSetFphaFpha)
}

func (c JsonSetValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type JsonStrappend Incomplete

func (b Builder) JsonStrappend() (c JsonStrappend) {
	_ = "STUB: not implemented"
	return *new(JsonStrappend)
}

func (c JsonStrappend) Key(key string) JsonStrappendKey {
	_ = "STUB: not implemented"
	return *new(JsonStrappendKey)
}

type JsonStrappendKey Incomplete

func (c JsonStrappendKey) Path(path string) JsonStrappendPath {
	_ = "STUB: not implemented"
	return *new(JsonStrappendPath)
}

func (c JsonStrappendKey) Value(value string) JsonStrappendValue {
	_ = "STUB: not implemented"
	return *new(JsonStrappendValue)
}

type JsonStrappendPath Incomplete

func (c JsonStrappendPath) Value(value string) JsonStrappendValue {
	_ = "STUB: not implemented"
	return *new(JsonStrappendValue)
}

type JsonStrappendValue Incomplete

func (c JsonStrappendValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type JsonStrlen Incomplete

func (b Builder) JsonStrlen() (c JsonStrlen) { _ = "STUB: not implemented"; return *new(JsonStrlen) }

func (c JsonStrlen) Key(key string) JsonStrlenKey {
	_ = "STUB: not implemented"
	return *new(JsonStrlenKey)
}

type JsonStrlenKey Incomplete

func (c JsonStrlenKey) Path(path string) JsonStrlenPath {
	_ = "STUB: not implemented"
	return *new(JsonStrlenPath)
}

func (c JsonStrlenKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c JsonStrlenKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type JsonStrlenPath Incomplete

func (c JsonStrlenPath) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c JsonStrlenPath) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type JsonToggle Incomplete

func (b Builder) JsonToggle() (c JsonToggle) { _ = "STUB: not implemented"; return *new(JsonToggle) }

func (c JsonToggle) Key(key string) JsonToggleKey {
	_ = "STUB: not implemented"
	return *new(JsonToggleKey)
}

type JsonToggleKey Incomplete

func (c JsonToggleKey) Path(path string) JsonTogglePath {
	_ = "STUB: not implemented"
	return *new(JsonTogglePath)
}

type JsonTogglePath Incomplete

func (c JsonTogglePath) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type JsonType Incomplete

func (b Builder) JsonType() (c JsonType) { _ = "STUB: not implemented"; return *new(JsonType) }

func (c JsonType) Key(key string) JsonTypeKey { _ = "STUB: not implemented"; return *new(JsonTypeKey) }

type JsonTypeKey Incomplete

func (c JsonTypeKey) Path(path string) JsonTypePath {
	_ = "STUB: not implemented"
	return *new(JsonTypePath)
}

func (c JsonTypeKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c JsonTypeKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type JsonTypePath Incomplete

func (c JsonTypePath) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c JsonTypePath) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

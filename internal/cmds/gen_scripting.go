// Code generated DO NOT EDIT

package cmds

type Eval Incomplete

func (b Builder) Eval() (c Eval) { _ = "STUB: not implemented"; return *new(Eval) }

func (c Eval) Script(script string) EvalScript { _ = "STUB: not implemented"; return *new(EvalScript) }

type EvalArg Incomplete

func (c EvalArg) Arg(arg ...string) EvalArg { _ = "STUB: not implemented"; return *new(EvalArg) }

func (c EvalArg) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type EvalKey Incomplete

func (c EvalKey) Key(key ...string) EvalKey { _ = "STUB: not implemented"; return *new(EvalKey) }

func (c EvalKey) Arg(arg ...string) EvalArg { _ = "STUB: not implemented"; return *new(EvalArg) }

func (c EvalKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type EvalNumkeys Incomplete

func (c EvalNumkeys) Key(key ...string) EvalKey { _ = "STUB: not implemented"; return *new(EvalKey) }

func (c EvalNumkeys) Arg(arg ...string) EvalArg { _ = "STUB: not implemented"; return *new(EvalArg) }

func (c EvalNumkeys) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type EvalRo Incomplete

func (b Builder) EvalRo() (c EvalRo) { _ = "STUB: not implemented"; return *new(EvalRo) }

func (c EvalRo) Script(script string) EvalRoScript {
	_ = "STUB: not implemented"
	return *new(EvalRoScript)
}

type EvalRoArg Incomplete

func (c EvalRoArg) Arg(arg ...string) EvalRoArg { _ = "STUB: not implemented"; return *new(EvalRoArg) }

func (c EvalRoArg) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c EvalRoArg) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type EvalRoKey Incomplete

func (c EvalRoKey) Key(key ...string) EvalRoKey { _ = "STUB: not implemented"; return *new(EvalRoKey) }

func (c EvalRoKey) Arg(arg ...string) EvalRoArg { _ = "STUB: not implemented"; return *new(EvalRoArg) }

func (c EvalRoKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c EvalRoKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type EvalRoNumkeys Incomplete

func (c EvalRoNumkeys) Key(key ...string) EvalRoKey {
	_ = "STUB: not implemented"
	return *new(EvalRoKey)
}

func (c EvalRoNumkeys) Arg(arg ...string) EvalRoArg {
	_ = "STUB: not implemented"
	return *new(EvalRoArg)
}

func (c EvalRoNumkeys) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c EvalRoNumkeys) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type EvalRoScript Incomplete

func (c EvalRoScript) Numkeys(numkeys int64) EvalRoNumkeys {
	_ = "STUB: not implemented"
	return *new(EvalRoNumkeys)
}

type EvalScript Incomplete

func (c EvalScript) Numkeys(numkeys int64) EvalNumkeys {
	_ = "STUB: not implemented"
	return *new(EvalNumkeys)
}

type Evalsha Incomplete

func (b Builder) Evalsha() (c Evalsha) { _ = "STUB: not implemented"; return *new(Evalsha) }

func (c Evalsha) Sha1(sha1 string) EvalshaSha1 { _ = "STUB: not implemented"; return *new(EvalshaSha1) }

type EvalshaArg Incomplete

func (c EvalshaArg) Arg(arg ...string) EvalshaArg {
	_ = "STUB: not implemented"
	return *new(EvalshaArg)
}

func (c EvalshaArg) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type EvalshaKey Incomplete

func (c EvalshaKey) Key(key ...string) EvalshaKey {
	_ = "STUB: not implemented"
	return *new(EvalshaKey)
}

func (c EvalshaKey) Arg(arg ...string) EvalshaArg {
	_ = "STUB: not implemented"
	return *new(EvalshaArg)
}

func (c EvalshaKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type EvalshaNumkeys Incomplete

func (c EvalshaNumkeys) Key(key ...string) EvalshaKey {
	_ = "STUB: not implemented"
	return *new(EvalshaKey)
}

func (c EvalshaNumkeys) Arg(arg ...string) EvalshaArg {
	_ = "STUB: not implemented"
	return *new(EvalshaArg)
}

func (c EvalshaNumkeys) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type EvalshaRo Incomplete

func (b Builder) EvalshaRo() (c EvalshaRo) { _ = "STUB: not implemented"; return *new(EvalshaRo) }

func (c EvalshaRo) Sha1(sha1 string) EvalshaRoSha1 {
	_ = "STUB: not implemented"
	return *new(EvalshaRoSha1)
}

type EvalshaRoArg Incomplete

func (c EvalshaRoArg) Arg(arg ...string) EvalshaRoArg {
	_ = "STUB: not implemented"
	return *new(EvalshaRoArg)
}

func (c EvalshaRoArg) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c EvalshaRoArg) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type EvalshaRoKey Incomplete

func (c EvalshaRoKey) Key(key ...string) EvalshaRoKey {
	_ = "STUB: not implemented"
	return *new(EvalshaRoKey)
}

func (c EvalshaRoKey) Arg(arg ...string) EvalshaRoArg {
	_ = "STUB: not implemented"
	return *new(EvalshaRoArg)
}

func (c EvalshaRoKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c EvalshaRoKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type EvalshaRoNumkeys Incomplete

func (c EvalshaRoNumkeys) Key(key ...string) EvalshaRoKey {
	_ = "STUB: not implemented"
	return *new(EvalshaRoKey)
}

func (c EvalshaRoNumkeys) Arg(arg ...string) EvalshaRoArg {
	_ = "STUB: not implemented"
	return *new(EvalshaRoArg)
}

func (c EvalshaRoNumkeys) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c EvalshaRoNumkeys) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type EvalshaRoSha1 Incomplete

func (c EvalshaRoSha1) Numkeys(numkeys int64) EvalshaRoNumkeys {
	_ = "STUB: not implemented"
	return *new(EvalshaRoNumkeys)
}

type EvalshaSha1 Incomplete

func (c EvalshaSha1) Numkeys(numkeys int64) EvalshaNumkeys {
	_ = "STUB: not implemented"
	return *new(EvalshaNumkeys)
}

type Fcall Incomplete

func (b Builder) Fcall() (c Fcall) { _ = "STUB: not implemented"; return *new(Fcall) }

func (c Fcall) Function(function string) FcallFunction {
	_ = "STUB: not implemented"
	return *new(FcallFunction)
}

type FcallArg Incomplete

func (c FcallArg) Arg(arg ...string) FcallArg { _ = "STUB: not implemented"; return *new(FcallArg) }

func (c FcallArg) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FcallFunction Incomplete

func (c FcallFunction) Numkeys(numkeys int64) FcallNumkeys {
	_ = "STUB: not implemented"
	return *new(FcallNumkeys)
}

type FcallKey Incomplete

func (c FcallKey) Key(key ...string) FcallKey { _ = "STUB: not implemented"; return *new(FcallKey) }

func (c FcallKey) Arg(arg ...string) FcallArg { _ = "STUB: not implemented"; return *new(FcallArg) }

func (c FcallKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FcallNumkeys Incomplete

func (c FcallNumkeys) Key(key ...string) FcallKey { _ = "STUB: not implemented"; return *new(FcallKey) }

func (c FcallNumkeys) Arg(arg ...string) FcallArg { _ = "STUB: not implemented"; return *new(FcallArg) }

func (c FcallNumkeys) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FcallRo Incomplete

func (b Builder) FcallRo() (c FcallRo) { _ = "STUB: not implemented"; return *new(FcallRo) }

func (c FcallRo) Function(function string) FcallRoFunction {
	_ = "STUB: not implemented"
	return *new(FcallRoFunction)
}

type FcallRoArg Incomplete

func (c FcallRoArg) Arg(arg ...string) FcallRoArg {
	_ = "STUB: not implemented"
	return *new(FcallRoArg)
}

func (c FcallRoArg) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c FcallRoArg) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type FcallRoFunction Incomplete

func (c FcallRoFunction) Numkeys(numkeys int64) FcallRoNumkeys {
	_ = "STUB: not implemented"
	return *new(FcallRoNumkeys)
}

type FcallRoKey Incomplete

func (c FcallRoKey) Key(key ...string) FcallRoKey {
	_ = "STUB: not implemented"
	return *new(FcallRoKey)
}

func (c FcallRoKey) Arg(arg ...string) FcallRoArg {
	_ = "STUB: not implemented"
	return *new(FcallRoArg)
}

func (c FcallRoKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c FcallRoKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type FcallRoNumkeys Incomplete

func (c FcallRoNumkeys) Key(key ...string) FcallRoKey {
	_ = "STUB: not implemented"
	return *new(FcallRoKey)
}

func (c FcallRoNumkeys) Arg(arg ...string) FcallRoArg {
	_ = "STUB: not implemented"
	return *new(FcallRoArg)
}

func (c FcallRoNumkeys) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c FcallRoNumkeys) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type FunctionDelete Incomplete

func (b Builder) FunctionDelete() (c FunctionDelete) {
	_ = "STUB: not implemented"
	return *new(FunctionDelete)
}

func (c FunctionDelete) LibraryName(libraryName string) FunctionDeleteLibraryName {
	_ = "STUB: not implemented"
	return *new(FunctionDeleteLibraryName)
}

type FunctionDeleteLibraryName Incomplete

func (c FunctionDeleteLibraryName) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FunctionDump Incomplete

func (b Builder) FunctionDump() (c FunctionDump) {
	_ = "STUB: not implemented"
	return *new(FunctionDump)
}

func (c FunctionDump) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FunctionFlush Incomplete

func (b Builder) FunctionFlush() (c FunctionFlush) {
	_ = "STUB: not implemented"
	return *new(FunctionFlush)
}

func (c FunctionFlush) Async() FunctionFlushAsync {
	_ = "STUB: not implemented"
	return *new(FunctionFlushAsync)
}

func (c FunctionFlush) Sync() FunctionFlushAsyncSync {
	_ = "STUB: not implemented"
	return *new(FunctionFlushAsyncSync)
}

func (c FunctionFlush) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FunctionFlushAsync Incomplete

func (c FunctionFlushAsync) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FunctionFlushAsyncSync Incomplete

func (c FunctionFlushAsyncSync) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FunctionHelp Incomplete

func (b Builder) FunctionHelp() (c FunctionHelp) {
	_ = "STUB: not implemented"
	return *new(FunctionHelp)
}

func (c FunctionHelp) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FunctionKill Incomplete

func (b Builder) FunctionKill() (c FunctionKill) {
	_ = "STUB: not implemented"
	return *new(FunctionKill)
}

func (c FunctionKill) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FunctionList Incomplete

func (b Builder) FunctionList() (c FunctionList) {
	_ = "STUB: not implemented"
	return *new(FunctionList)
}

func (c FunctionList) Libraryname(libraryNamePattern string) FunctionListLibraryname {
	_ = "STUB: not implemented"
	return *new(FunctionListLibraryname)
}

func (c FunctionList) Withcode() FunctionListWithcode {
	_ = "STUB: not implemented"
	return *new(FunctionListWithcode)
}

func (c FunctionList) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FunctionListLibraryname Incomplete

func (c FunctionListLibraryname) Withcode() FunctionListWithcode {
	_ = "STUB: not implemented"
	return *new(FunctionListWithcode)
}

func (c FunctionListLibraryname) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FunctionListWithcode Incomplete

func (c FunctionListWithcode) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FunctionLoad Incomplete

func (b Builder) FunctionLoad() (c FunctionLoad) {
	_ = "STUB: not implemented"
	return *new(FunctionLoad)
}

func (c FunctionLoad) Replace() FunctionLoadReplace {
	_ = "STUB: not implemented"
	return *new(FunctionLoadReplace)
}

func (c FunctionLoad) FunctionCode(functionCode string) FunctionLoadFunctionCode {
	_ = "STUB: not implemented"
	return *new(FunctionLoadFunctionCode)
}

type FunctionLoadFunctionCode Incomplete

func (c FunctionLoadFunctionCode) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FunctionLoadReplace Incomplete

func (c FunctionLoadReplace) FunctionCode(functionCode string) FunctionLoadFunctionCode {
	_ = "STUB: not implemented"
	return *new(FunctionLoadFunctionCode)
}

type FunctionRestore Incomplete

func (b Builder) FunctionRestore() (c FunctionRestore) {
	_ = "STUB: not implemented"
	return *new(FunctionRestore)
}

func (c FunctionRestore) SerializedValue(serializedValue string) FunctionRestoreSerializedValue {
	_ = "STUB: not implemented"
	return *new(FunctionRestoreSerializedValue)
}

type FunctionRestorePolicyAppend Incomplete

func (c FunctionRestorePolicyAppend) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FunctionRestorePolicyFlush Incomplete

func (c FunctionRestorePolicyFlush) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FunctionRestorePolicyReplace Incomplete

func (c FunctionRestorePolicyReplace) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FunctionRestoreSerializedValue Incomplete

func (c FunctionRestoreSerializedValue) Flush() FunctionRestorePolicyFlush {
	_ = "STUB: not implemented"
	return *new(FunctionRestorePolicyFlush)
}

func (c FunctionRestoreSerializedValue) Append() FunctionRestorePolicyAppend {
	_ = "STUB: not implemented"
	return *new(FunctionRestorePolicyAppend)
}

func (c FunctionRestoreSerializedValue) Replace() FunctionRestorePolicyReplace {
	_ = "STUB: not implemented"
	return *new(FunctionRestorePolicyReplace)
}

func (c FunctionRestoreSerializedValue) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type FunctionStats Incomplete

func (b Builder) FunctionStats() (c FunctionStats) {
	_ = "STUB: not implemented"
	return *new(FunctionStats)
}

func (c FunctionStats) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ScriptDebug Incomplete

func (b Builder) ScriptDebug() (c ScriptDebug) { _ = "STUB: not implemented"; return *new(ScriptDebug) }

func (c ScriptDebug) Yes() ScriptDebugModeYes {
	_ = "STUB: not implemented"
	return *new(ScriptDebugModeYes)
}

func (c ScriptDebug) Sync() ScriptDebugModeSync {
	_ = "STUB: not implemented"
	return *new(ScriptDebugModeSync)
}

func (c ScriptDebug) No() ScriptDebugModeNo {
	_ = "STUB: not implemented"
	return *new(ScriptDebugModeNo)
}

type ScriptDebugModeNo Incomplete

func (c ScriptDebugModeNo) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ScriptDebugModeSync Incomplete

func (c ScriptDebugModeSync) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ScriptDebugModeYes Incomplete

func (c ScriptDebugModeYes) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ScriptExists Incomplete

func (b Builder) ScriptExists() (c ScriptExists) {
	_ = "STUB: not implemented"
	return *new(ScriptExists)
}

func (c ScriptExists) Sha1(sha1 ...string) ScriptExistsSha1 {
	_ = "STUB: not implemented"
	return *new(ScriptExistsSha1)
}

type ScriptExistsSha1 Incomplete

func (c ScriptExistsSha1) Sha1(sha1 ...string) ScriptExistsSha1 {
	_ = "STUB: not implemented"
	return *new(ScriptExistsSha1)
}

func (c ScriptExistsSha1) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ScriptFlush Incomplete

func (b Builder) ScriptFlush() (c ScriptFlush) { _ = "STUB: not implemented"; return *new(ScriptFlush) }

func (c ScriptFlush) Async() ScriptFlushAsync {
	_ = "STUB: not implemented"
	return *new(ScriptFlushAsync)
}

func (c ScriptFlush) Sync() ScriptFlushAsyncSync {
	_ = "STUB: not implemented"
	return *new(ScriptFlushAsyncSync)
}

func (c ScriptFlush) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ScriptFlushAsync Incomplete

func (c ScriptFlushAsync) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ScriptFlushAsyncSync Incomplete

func (c ScriptFlushAsyncSync) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ScriptKill Incomplete

func (b Builder) ScriptKill() (c ScriptKill) { _ = "STUB: not implemented"; return *new(ScriptKill) }

func (c ScriptKill) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ScriptLoad Incomplete

func (b Builder) ScriptLoad() (c ScriptLoad) { _ = "STUB: not implemented"; return *new(ScriptLoad) }

func (c ScriptLoad) Script(script string) ScriptLoadScript {
	_ = "STUB: not implemented"
	return *new(ScriptLoadScript)
}

type ScriptLoadScript Incomplete

func (c ScriptLoadScript) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ScriptShow Incomplete

func (b Builder) ScriptShow() (c ScriptShow) { _ = "STUB: not implemented"; return *new(ScriptShow) }

func (c ScriptShow) Sha1(sha1 string) ScriptShowSha1 {
	_ = "STUB: not implemented"
	return *new(ScriptShowSha1)
}

type ScriptShowSha1 Incomplete

func (c ScriptShowSha1) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

// Code generated DO NOT EDIT

package cmds

type Tfcall Incomplete

func (b Builder) Tfcall() (c Tfcall) { _ = "STUB: not implemented"; return *new(Tfcall) }

func (c Tfcall) LibraryFunction(libraryFunction string) TfcallLibraryFunction {
	_ = "STUB: not implemented"
	return *new(TfcallLibraryFunction)
}

type TfcallArg Incomplete

func (c TfcallArg) Arg(arg ...string) TfcallArg { _ = "STUB: not implemented"; return *new(TfcallArg) }

func (c TfcallArg) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TfcallKey Incomplete

func (c TfcallKey) Key(key ...string) TfcallKey { _ = "STUB: not implemented"; return *new(TfcallKey) }

func (c TfcallKey) Arg(arg ...string) TfcallArg { _ = "STUB: not implemented"; return *new(TfcallArg) }

func (c TfcallKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TfcallLibraryFunction Incomplete

func (c TfcallLibraryFunction) Numkeys(numkeys int64) TfcallNumkeys {
	_ = "STUB: not implemented"
	return *new(TfcallNumkeys)
}

type TfcallNumkeys Incomplete

func (c TfcallNumkeys) Key(key ...string) TfcallKey {
	_ = "STUB: not implemented"
	return *new(TfcallKey)
}

func (c TfcallNumkeys) Arg(arg ...string) TfcallArg {
	_ = "STUB: not implemented"
	return *new(TfcallArg)
}

func (c TfcallNumkeys) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Tfcallasync Incomplete

func (b Builder) Tfcallasync() (c Tfcallasync) { _ = "STUB: not implemented"; return *new(Tfcallasync) }

func (c Tfcallasync) LibraryFunction(libraryFunction string) TfcallasyncLibraryFunction {
	_ = "STUB: not implemented"
	return *new(TfcallasyncLibraryFunction)
}

type TfcallasyncArg Incomplete

func (c TfcallasyncArg) Arg(arg ...string) TfcallasyncArg {
	_ = "STUB: not implemented"
	return *new(TfcallasyncArg)
}

func (c TfcallasyncArg) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TfcallasyncKey Incomplete

func (c TfcallasyncKey) Key(key ...string) TfcallasyncKey {
	_ = "STUB: not implemented"
	return *new(TfcallasyncKey)
}

func (c TfcallasyncKey) Arg(arg ...string) TfcallasyncArg {
	_ = "STUB: not implemented"
	return *new(TfcallasyncArg)
}

func (c TfcallasyncKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TfcallasyncLibraryFunction Incomplete

func (c TfcallasyncLibraryFunction) Numkeys(numkeys int64) TfcallasyncNumkeys {
	_ = "STUB: not implemented"
	return *new(TfcallasyncNumkeys)
}

type TfcallasyncNumkeys Incomplete

func (c TfcallasyncNumkeys) Key(key ...string) TfcallasyncKey {
	_ = "STUB: not implemented"
	return *new(TfcallasyncKey)
}

func (c TfcallasyncNumkeys) Arg(arg ...string) TfcallasyncArg {
	_ = "STUB: not implemented"
	return *new(TfcallasyncArg)
}

func (c TfcallasyncNumkeys) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TfunctionDelete Incomplete

func (b Builder) TfunctionDelete() (c TfunctionDelete) {
	_ = "STUB: not implemented"
	return *new(TfunctionDelete)
}

func (c TfunctionDelete) LibraryName(libraryName string) TfunctionDeleteLibraryName {
	_ = "STUB: not implemented"
	return *new(TfunctionDeleteLibraryName)
}

type TfunctionDeleteLibraryName Incomplete

func (c TfunctionDeleteLibraryName) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TfunctionList Incomplete

func (b Builder) TfunctionList() (c TfunctionList) {
	_ = "STUB: not implemented"
	return *new(TfunctionList)
}

func (c TfunctionList) LibraryName(libraryName string) TfunctionListLibraryName {
	_ = "STUB: not implemented"
	return *new(TfunctionListLibraryName)
}

func (c TfunctionList) Withcode() TfunctionListWithcode {
	_ = "STUB: not implemented"
	return *new(TfunctionListWithcode)
}

func (c TfunctionList) Verbose() TfunctionListVerbose {
	_ = "STUB: not implemented"
	return *new(TfunctionListVerbose)
}

func (c TfunctionList) V() TfunctionListV { _ = "STUB: not implemented"; return *new(TfunctionListV) }

func (c TfunctionList) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TfunctionListLibraryName Incomplete

func (c TfunctionListLibraryName) Withcode() TfunctionListWithcode {
	_ = "STUB: not implemented"
	return *new(TfunctionListWithcode)
}

func (c TfunctionListLibraryName) Verbose() TfunctionListVerbose {
	_ = "STUB: not implemented"
	return *new(TfunctionListVerbose)
}

func (c TfunctionListLibraryName) V() TfunctionListV {
	_ = "STUB: not implemented"
	return *new(TfunctionListV)
}

func (c TfunctionListLibraryName) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TfunctionListV Incomplete

func (c TfunctionListV) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TfunctionListVerbose Incomplete

func (c TfunctionListVerbose) V() TfunctionListV {
	_ = "STUB: not implemented"
	return *new(TfunctionListV)
}

func (c TfunctionListVerbose) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TfunctionListWithcode Incomplete

func (c TfunctionListWithcode) Verbose() TfunctionListVerbose {
	_ = "STUB: not implemented"
	return *new(TfunctionListVerbose)
}

func (c TfunctionListWithcode) V() TfunctionListV {
	_ = "STUB: not implemented"
	return *new(TfunctionListV)
}

func (c TfunctionListWithcode) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type TfunctionLoad Incomplete

func (b Builder) TfunctionLoad() (c TfunctionLoad) {
	_ = "STUB: not implemented"
	return *new(TfunctionLoad)
}

func (c TfunctionLoad) Replace() TfunctionLoadReplace {
	_ = "STUB: not implemented"
	return *new(TfunctionLoadReplace)
}

func (c TfunctionLoad) Config(config string) TfunctionLoadConfig {
	_ = "STUB: not implemented"
	return *new(TfunctionLoadConfig)
}

func (c TfunctionLoad) LibraryCode(libraryCode string) TfunctionLoadLibraryCode {
	_ = "STUB: not implemented"
	return *new(TfunctionLoadLibraryCode)
}

type TfunctionLoadConfig Incomplete

func (c TfunctionLoadConfig) LibraryCode(libraryCode string) TfunctionLoadLibraryCode {
	_ = "STUB: not implemented"
	return *new(TfunctionLoadLibraryCode)
}

type TfunctionLoadLibraryCode Incomplete

func (c TfunctionLoadLibraryCode) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type TfunctionLoadReplace Incomplete

func (c TfunctionLoadReplace) Config(config string) TfunctionLoadConfig {
	_ = "STUB: not implemented"
	return *new(TfunctionLoadConfig)
}

func (c TfunctionLoadReplace) LibraryCode(libraryCode string) TfunctionLoadLibraryCode {
	_ = "STUB: not implemented"
	return *new(TfunctionLoadLibraryCode)
}

// Code generated DO NOT EDIT

package cmds

type AiScriptdel Incomplete

func (b Builder) AiScriptdel() (c AiScriptdel) { _ = "STUB: not implemented"; return *new(AiScriptdel) }

func (c AiScriptdel) Key(key string) AiScriptdelKey {
	_ = "STUB: not implemented"
	return *new(AiScriptdelKey)
}

type AiScriptdelKey Incomplete

func (c AiScriptdelKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type AiScriptget Incomplete

func (b Builder) AiScriptget() (c AiScriptget) { _ = "STUB: not implemented"; return *new(AiScriptget) }

func (c AiScriptget) Key(key string) AiScriptgetKey {
	_ = "STUB: not implemented"
	return *new(AiScriptgetKey)
}

type AiScriptgetKey Incomplete

func (c AiScriptgetKey) Meta() AiScriptgetMeta {
	_ = "STUB: not implemented"
	return *new(AiScriptgetMeta)
}

func (c AiScriptgetKey) Source() AiScriptgetSource {
	_ = "STUB: not implemented"
	return *new(AiScriptgetSource)
}

func (c AiScriptgetKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c AiScriptgetKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type AiScriptgetMeta Incomplete

func (c AiScriptgetMeta) Source() AiScriptgetSource {
	_ = "STUB: not implemented"
	return *new(AiScriptgetSource)
}

func (c AiScriptgetMeta) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c AiScriptgetMeta) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type AiScriptgetSource Incomplete

func (c AiScriptgetSource) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c AiScriptgetSource) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type AiScriptstore Incomplete

func (b Builder) AiScriptstore() (c AiScriptstore) {
	_ = "STUB: not implemented"
	return *new(AiScriptstore)
}

func (c AiScriptstore) Key(key string) AiScriptstoreKey {
	_ = "STUB: not implemented"
	return *new(AiScriptstoreKey)
}

type AiScriptstoreDeviceCpu Incomplete

func (c AiScriptstoreDeviceCpu) Tag(tag string) AiScriptstoreTag {
	_ = "STUB: not implemented"
	return *new(AiScriptstoreTag)
}

func (c AiScriptstoreDeviceCpu) EntryPoints(entryPointCount int64) AiScriptstoreEntryPointsEntryPoints {
	_ = "STUB: not implemented"
	return *new(AiScriptstoreEntryPointsEntryPoints)
}

type AiScriptstoreDeviceGpu Incomplete

func (c AiScriptstoreDeviceGpu) Tag(tag string) AiScriptstoreTag {
	_ = "STUB: not implemented"
	return *new(AiScriptstoreTag)
}

func (c AiScriptstoreDeviceGpu) EntryPoints(entryPointCount int64) AiScriptstoreEntryPointsEntryPoints {
	_ = "STUB: not implemented"
	return *new(AiScriptstoreEntryPointsEntryPoints)
}

type AiScriptstoreEntryPointsEntryPoint Incomplete

func (c AiScriptstoreEntryPointsEntryPoint) EntryPoint(entryPoint ...string) AiScriptstoreEntryPointsEntryPoint {
	_ = "STUB: not implemented"
	return *new(AiScriptstoreEntryPointsEntryPoint)
}

func (c AiScriptstoreEntryPointsEntryPoint) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type AiScriptstoreEntryPointsEntryPoints Incomplete

func (c AiScriptstoreEntryPointsEntryPoints) EntryPoint(entryPoint ...string) AiScriptstoreEntryPointsEntryPoint {
	_ = "STUB: not implemented"
	return *new(AiScriptstoreEntryPointsEntryPoint)
}

type AiScriptstoreKey Incomplete

func (c AiScriptstoreKey) Cpu() AiScriptstoreDeviceCpu {
	_ = "STUB: not implemented"
	return *new(AiScriptstoreDeviceCpu)
}

func (c AiScriptstoreKey) Gpu() AiScriptstoreDeviceGpu {
	_ = "STUB: not implemented"
	return *new(AiScriptstoreDeviceGpu)
}

type AiScriptstoreTag Incomplete

func (c AiScriptstoreTag) EntryPoints(entryPointCount int64) AiScriptstoreEntryPointsEntryPoints {
	_ = "STUB: not implemented"
	return *new(AiScriptstoreEntryPointsEntryPoints)
}

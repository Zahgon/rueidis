// Code generated DO NOT EDIT

package cmds

type AiModeldel Incomplete

func (b Builder) AiModeldel() (c AiModeldel) { _ = "STUB: not implemented"; return *new(AiModeldel) }

func (c AiModeldel) Key(key string) AiModeldelKey {
	_ = "STUB: not implemented"
	return *new(AiModeldelKey)
}

type AiModeldelKey Incomplete

func (c AiModeldelKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type AiModelget Incomplete

func (b Builder) AiModelget() (c AiModelget) { _ = "STUB: not implemented"; return *new(AiModelget) }

func (c AiModelget) Key(key string) AiModelgetKey {
	_ = "STUB: not implemented"
	return *new(AiModelgetKey)
}

type AiModelgetBlob Incomplete

func (c AiModelgetBlob) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c AiModelgetBlob) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type AiModelgetKey Incomplete

func (c AiModelgetKey) Meta() AiModelgetMeta {
	_ = "STUB: not implemented"
	return *new(AiModelgetMeta)
}

func (c AiModelgetKey) Blob() AiModelgetBlob {
	_ = "STUB: not implemented"
	return *new(AiModelgetBlob)
}

func (c AiModelgetKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c AiModelgetKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type AiModelgetMeta Incomplete

func (c AiModelgetMeta) Blob() AiModelgetBlob {
	_ = "STUB: not implemented"
	return *new(AiModelgetBlob)
}

func (c AiModelgetMeta) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c AiModelgetMeta) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type AiModelstore Incomplete

func (b Builder) AiModelstore() (c AiModelstore) {
	_ = "STUB: not implemented"
	return *new(AiModelstore)
}

func (c AiModelstore) Key(key string) AiModelstoreKey {
	_ = "STUB: not implemented"
	return *new(AiModelstoreKey)
}

type AiModelstoreBackendOnnx Incomplete

func (c AiModelstoreBackendOnnx) Cpu() AiModelstoreDeviceCpu {
	_ = "STUB: not implemented"
	return *new(AiModelstoreDeviceCpu)
}

func (c AiModelstoreBackendOnnx) Gpu() AiModelstoreDeviceGpu {
	_ = "STUB: not implemented"
	return *new(AiModelstoreDeviceGpu)
}

type AiModelstoreBackendTf Incomplete

func (c AiModelstoreBackendTf) Cpu() AiModelstoreDeviceCpu {
	_ = "STUB: not implemented"
	return *new(AiModelstoreDeviceCpu)
}

func (c AiModelstoreBackendTf) Gpu() AiModelstoreDeviceGpu {
	_ = "STUB: not implemented"
	return *new(AiModelstoreDeviceGpu)
}

type AiModelstoreBackendTorch Incomplete

func (c AiModelstoreBackendTorch) Cpu() AiModelstoreDeviceCpu {
	_ = "STUB: not implemented"
	return *new(AiModelstoreDeviceCpu)
}

func (c AiModelstoreBackendTorch) Gpu() AiModelstoreDeviceGpu {
	_ = "STUB: not implemented"
	return *new(AiModelstoreDeviceGpu)
}

type AiModelstoreBatchsize Incomplete

func (c AiModelstoreBatchsize) Minbatchsize(minbatchsize int64) AiModelstoreMinbatchsize {
	_ = "STUB: not implemented"
	return *new(AiModelstoreMinbatchsize)
}

func (c AiModelstoreBatchsize) Minbatchtimeout(minbatchtimeout int64) AiModelstoreMinbatchtimeout {
	_ = "STUB: not implemented"
	return *new(AiModelstoreMinbatchtimeout)
}

func (c AiModelstoreBatchsize) Inputs(inputCount int64) AiModelstoreInputsInputs {
	_ = "STUB: not implemented"
	return *new(AiModelstoreInputsInputs)
}

func (c AiModelstoreBatchsize) Outputs(outputCount int64) AiModelstoreOutputsOutputs {
	_ = "STUB: not implemented"
	return *new(AiModelstoreOutputsOutputs)
}

func (c AiModelstoreBatchsize) Blob(blob string) AiModelstoreBlob {
	_ = "STUB: not implemented"
	return *new(AiModelstoreBlob)
}

func (c AiModelstoreBatchsize) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type AiModelstoreBlob Incomplete

func (c AiModelstoreBlob) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type AiModelstoreDeviceCpu Incomplete

func (c AiModelstoreDeviceCpu) Tag(tag string) AiModelstoreTag {
	_ = "STUB: not implemented"
	return *new(AiModelstoreTag)
}

func (c AiModelstoreDeviceCpu) Batchsize(batchsize int64) AiModelstoreBatchsize {
	_ = "STUB: not implemented"
	return *new(AiModelstoreBatchsize)
}

func (c AiModelstoreDeviceCpu) Minbatchsize(minbatchsize int64) AiModelstoreMinbatchsize {
	_ = "STUB: not implemented"
	return *new(AiModelstoreMinbatchsize)
}

func (c AiModelstoreDeviceCpu) Minbatchtimeout(minbatchtimeout int64) AiModelstoreMinbatchtimeout {
	_ = "STUB: not implemented"
	return *new(AiModelstoreMinbatchtimeout)
}

func (c AiModelstoreDeviceCpu) Inputs(inputCount int64) AiModelstoreInputsInputs {
	_ = "STUB: not implemented"
	return *new(AiModelstoreInputsInputs)
}

func (c AiModelstoreDeviceCpu) Outputs(outputCount int64) AiModelstoreOutputsOutputs {
	_ = "STUB: not implemented"
	return *new(AiModelstoreOutputsOutputs)
}

func (c AiModelstoreDeviceCpu) Blob(blob string) AiModelstoreBlob {
	_ = "STUB: not implemented"
	return *new(AiModelstoreBlob)
}

func (c AiModelstoreDeviceCpu) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type AiModelstoreDeviceGpu Incomplete

func (c AiModelstoreDeviceGpu) Tag(tag string) AiModelstoreTag {
	_ = "STUB: not implemented"
	return *new(AiModelstoreTag)
}

func (c AiModelstoreDeviceGpu) Batchsize(batchsize int64) AiModelstoreBatchsize {
	_ = "STUB: not implemented"
	return *new(AiModelstoreBatchsize)
}

func (c AiModelstoreDeviceGpu) Minbatchsize(minbatchsize int64) AiModelstoreMinbatchsize {
	_ = "STUB: not implemented"
	return *new(AiModelstoreMinbatchsize)
}

func (c AiModelstoreDeviceGpu) Minbatchtimeout(minbatchtimeout int64) AiModelstoreMinbatchtimeout {
	_ = "STUB: not implemented"
	return *new(AiModelstoreMinbatchtimeout)
}

func (c AiModelstoreDeviceGpu) Inputs(inputCount int64) AiModelstoreInputsInputs {
	_ = "STUB: not implemented"
	return *new(AiModelstoreInputsInputs)
}

func (c AiModelstoreDeviceGpu) Outputs(outputCount int64) AiModelstoreOutputsOutputs {
	_ = "STUB: not implemented"
	return *new(AiModelstoreOutputsOutputs)
}

func (c AiModelstoreDeviceGpu) Blob(blob string) AiModelstoreBlob {
	_ = "STUB: not implemented"
	return *new(AiModelstoreBlob)
}

func (c AiModelstoreDeviceGpu) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type AiModelstoreInputsInput Incomplete

func (c AiModelstoreInputsInput) Input(input ...string) AiModelstoreInputsInput {
	_ = "STUB: not implemented"
	return *new(AiModelstoreInputsInput)
}

func (c AiModelstoreInputsInput) Outputs(outputCount int64) AiModelstoreOutputsOutputs {
	_ = "STUB: not implemented"
	return *new(AiModelstoreOutputsOutputs)
}

func (c AiModelstoreInputsInput) Blob(blob string) AiModelstoreBlob {
	_ = "STUB: not implemented"
	return *new(AiModelstoreBlob)
}

func (c AiModelstoreInputsInput) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type AiModelstoreInputsInputs Incomplete

func (c AiModelstoreInputsInputs) Input(input ...string) AiModelstoreInputsInput {
	_ = "STUB: not implemented"
	return *new(AiModelstoreInputsInput)
}

type AiModelstoreKey Incomplete

func (c AiModelstoreKey) Tf() AiModelstoreBackendTf {
	_ = "STUB: not implemented"
	return *new(AiModelstoreBackendTf)
}

func (c AiModelstoreKey) Torch() AiModelstoreBackendTorch {
	_ = "STUB: not implemented"
	return *new(AiModelstoreBackendTorch)
}

func (c AiModelstoreKey) Onnx() AiModelstoreBackendOnnx {
	_ = "STUB: not implemented"
	return *new(AiModelstoreBackendOnnx)
}

type AiModelstoreMinbatchsize Incomplete

func (c AiModelstoreMinbatchsize) Minbatchtimeout(minbatchtimeout int64) AiModelstoreMinbatchtimeout {
	_ = "STUB: not implemented"
	return *new(AiModelstoreMinbatchtimeout)
}

func (c AiModelstoreMinbatchsize) Inputs(inputCount int64) AiModelstoreInputsInputs {
	_ = "STUB: not implemented"
	return *new(AiModelstoreInputsInputs)
}

func (c AiModelstoreMinbatchsize) Outputs(outputCount int64) AiModelstoreOutputsOutputs {
	_ = "STUB: not implemented"
	return *new(AiModelstoreOutputsOutputs)
}

func (c AiModelstoreMinbatchsize) Blob(blob string) AiModelstoreBlob {
	_ = "STUB: not implemented"
	return *new(AiModelstoreBlob)
}

func (c AiModelstoreMinbatchsize) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type AiModelstoreMinbatchtimeout Incomplete

func (c AiModelstoreMinbatchtimeout) Inputs(inputCount int64) AiModelstoreInputsInputs {
	_ = "STUB: not implemented"
	return *new(AiModelstoreInputsInputs)
}

func (c AiModelstoreMinbatchtimeout) Outputs(outputCount int64) AiModelstoreOutputsOutputs {
	_ = "STUB: not implemented"
	return *new(AiModelstoreOutputsOutputs)
}

func (c AiModelstoreMinbatchtimeout) Blob(blob string) AiModelstoreBlob {
	_ = "STUB: not implemented"
	return *new(AiModelstoreBlob)
}

func (c AiModelstoreMinbatchtimeout) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type AiModelstoreOutputsOutput Incomplete

func (c AiModelstoreOutputsOutput) Output(output ...string) AiModelstoreOutputsOutput {
	_ = "STUB: not implemented"
	return *new(AiModelstoreOutputsOutput)
}

func (c AiModelstoreOutputsOutput) Blob(blob string) AiModelstoreBlob {
	_ = "STUB: not implemented"
	return *new(AiModelstoreBlob)
}

func (c AiModelstoreOutputsOutput) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type AiModelstoreOutputsOutputs Incomplete

func (c AiModelstoreOutputsOutputs) Output(output ...string) AiModelstoreOutputsOutput {
	_ = "STUB: not implemented"
	return *new(AiModelstoreOutputsOutput)
}

type AiModelstoreTag Incomplete

func (c AiModelstoreTag) Batchsize(batchsize int64) AiModelstoreBatchsize {
	_ = "STUB: not implemented"
	return *new(AiModelstoreBatchsize)
}

func (c AiModelstoreTag) Minbatchsize(minbatchsize int64) AiModelstoreMinbatchsize {
	_ = "STUB: not implemented"
	return *new(AiModelstoreMinbatchsize)
}

func (c AiModelstoreTag) Minbatchtimeout(minbatchtimeout int64) AiModelstoreMinbatchtimeout {
	_ = "STUB: not implemented"
	return *new(AiModelstoreMinbatchtimeout)
}

func (c AiModelstoreTag) Inputs(inputCount int64) AiModelstoreInputsInputs {
	_ = "STUB: not implemented"
	return *new(AiModelstoreInputsInputs)
}

func (c AiModelstoreTag) Outputs(outputCount int64) AiModelstoreOutputsOutputs {
	_ = "STUB: not implemented"
	return *new(AiModelstoreOutputsOutputs)
}

func (c AiModelstoreTag) Blob(blob string) AiModelstoreBlob {
	_ = "STUB: not implemented"
	return *new(AiModelstoreBlob)
}

func (c AiModelstoreTag) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

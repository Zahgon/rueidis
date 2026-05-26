// Code generated DO NOT EDIT

package cmds

type RgAbortexecution Incomplete

func (b Builder) RgAbortexecution() (c RgAbortexecution) {
	_ = "STUB: not implemented"
	return *new(RgAbortexecution)
}

func (c RgAbortexecution) Id(id string) RgAbortexecutionId {
	_ = "STUB: not implemented"
	return *new(RgAbortexecutionId)
}

type RgAbortexecutionId Incomplete

func (c RgAbortexecutionId) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type RgConfigget Incomplete

func (b Builder) RgConfigget() (c RgConfigget) { _ = "STUB: not implemented"; return *new(RgConfigget) }

func (c RgConfigget) Key(key ...string) RgConfiggetKey {
	_ = "STUB: not implemented"
	return *new(RgConfiggetKey)
}

type RgConfiggetKey Incomplete

func (c RgConfiggetKey) Key(key ...string) RgConfiggetKey {
	_ = "STUB: not implemented"
	return *new(RgConfiggetKey)
}

func (c RgConfiggetKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type RgConfigset Incomplete

func (b Builder) RgConfigset() (c RgConfigset) { _ = "STUB: not implemented"; return *new(RgConfigset) }

func (c RgConfigset) KeyValue() RgConfigsetKeyValue {
	_ = "STUB: not implemented"
	return *new(RgConfigsetKeyValue)
}

type RgConfigsetKeyValue Incomplete

func (c RgConfigsetKeyValue) KeyValue(key string, value string) RgConfigsetKeyValue {
	_ = "STUB: not implemented"
	return *new(RgConfigsetKeyValue)
}

func (c RgConfigsetKeyValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type RgDropexecution Incomplete

func (b Builder) RgDropexecution() (c RgDropexecution) {
	_ = "STUB: not implemented"
	return *new(RgDropexecution)
}

func (c RgDropexecution) Id(id string) RgDropexecutionId {
	_ = "STUB: not implemented"
	return *new(RgDropexecutionId)
}

type RgDropexecutionId Incomplete

func (c RgDropexecutionId) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type RgDumpexecutions Incomplete

func (b Builder) RgDumpexecutions() (c RgDumpexecutions) {
	_ = "STUB: not implemented"
	return *new(RgDumpexecutions)
}

func (c RgDumpexecutions) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type RgDumpregistrations Incomplete

func (b Builder) RgDumpregistrations() (c RgDumpregistrations) {
	_ = "STUB: not implemented"
	return *new(RgDumpregistrations)
}

func (c RgDumpregistrations) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type RgGetexecution Incomplete

func (b Builder) RgGetexecution() (c RgGetexecution) {
	_ = "STUB: not implemented"
	return *new(RgGetexecution)
}

func (c RgGetexecution) Id(id string) RgGetexecutionId {
	_ = "STUB: not implemented"
	return *new(RgGetexecutionId)
}

type RgGetexecutionId Incomplete

func (c RgGetexecutionId) Shard() RgGetexecutionModeShard {
	_ = "STUB: not implemented"
	return *new(RgGetexecutionModeShard)
}

func (c RgGetexecutionId) Cluster() RgGetexecutionModeCluster {
	_ = "STUB: not implemented"
	return *new(RgGetexecutionModeCluster)
}

func (c RgGetexecutionId) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type RgGetexecutionModeCluster Incomplete

func (c RgGetexecutionModeCluster) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type RgGetexecutionModeShard Incomplete

func (c RgGetexecutionModeShard) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type RgGetresults Incomplete

func (b Builder) RgGetresults() (c RgGetresults) {
	_ = "STUB: not implemented"
	return *new(RgGetresults)
}

func (c RgGetresults) Id(id string) RgGetresultsId {
	_ = "STUB: not implemented"
	return *new(RgGetresultsId)
}

type RgGetresultsId Incomplete

func (c RgGetresultsId) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type RgGetresultsblocking Incomplete

func (b Builder) RgGetresultsblocking() (c RgGetresultsblocking) {
	_ = "STUB: not implemented"
	return *new(RgGetresultsblocking)
}

func (c RgGetresultsblocking) Id(id string) RgGetresultsblockingId {
	_ = "STUB: not implemented"
	return *new(RgGetresultsblockingId)
}

type RgGetresultsblockingId Incomplete

func (c RgGetresultsblockingId) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type RgInfocluster Incomplete

func (b Builder) RgInfocluster() (c RgInfocluster) {
	_ = "STUB: not implemented"
	return *new(RgInfocluster)
}

func (c RgInfocluster) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type RgPydumpreqs Incomplete

func (b Builder) RgPydumpreqs() (c RgPydumpreqs) {
	_ = "STUB: not implemented"
	return *new(RgPydumpreqs)
}

func (c RgPydumpreqs) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type RgPyexecute Incomplete

func (b Builder) RgPyexecute() (c RgPyexecute) { _ = "STUB: not implemented"; return *new(RgPyexecute) }

func (c RgPyexecute) Function(function string) RgPyexecuteFunction {
	_ = "STUB: not implemented"
	return *new(RgPyexecuteFunction)
}

type RgPyexecuteDescription Incomplete

func (c RgPyexecuteDescription) Upgrade() RgPyexecuteUpgrade {
	_ = "STUB: not implemented"
	return *new(RgPyexecuteUpgrade)
}

func (c RgPyexecuteDescription) ReplaceWith(replaceWith string) RgPyexecuteReplaceWith {
	_ = "STUB: not implemented"
	return *new(RgPyexecuteReplaceWith)
}

func (c RgPyexecuteDescription) Requirements(requirement ...string) RgPyexecuteRequirementsRequirements {
	_ = "STUB: not implemented"
	return *new(RgPyexecuteRequirementsRequirements)
}

func (c RgPyexecuteDescription) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type RgPyexecuteFunction Incomplete

func (c RgPyexecuteFunction) Unblocking() RgPyexecuteUnblocking {
	_ = "STUB: not implemented"
	return *new(RgPyexecuteUnblocking)
}

func (c RgPyexecuteFunction) Id(id string) RgPyexecuteId {
	_ = "STUB: not implemented"
	return *new(RgPyexecuteId)
}

func (c RgPyexecuteFunction) Description(description string) RgPyexecuteDescription {
	_ = "STUB: not implemented"
	return *new(RgPyexecuteDescription)
}

func (c RgPyexecuteFunction) Upgrade() RgPyexecuteUpgrade {
	_ = "STUB: not implemented"
	return *new(RgPyexecuteUpgrade)
}

func (c RgPyexecuteFunction) ReplaceWith(replaceWith string) RgPyexecuteReplaceWith {
	_ = "STUB: not implemented"
	return *new(RgPyexecuteReplaceWith)
}

func (c RgPyexecuteFunction) Requirements(requirement ...string) RgPyexecuteRequirementsRequirements {
	_ = "STUB: not implemented"
	return *new(RgPyexecuteRequirementsRequirements)
}

func (c RgPyexecuteFunction) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type RgPyexecuteId Incomplete

func (c RgPyexecuteId) Description(description string) RgPyexecuteDescription {
	_ = "STUB: not implemented"
	return *new(RgPyexecuteDescription)
}

func (c RgPyexecuteId) Upgrade() RgPyexecuteUpgrade {
	_ = "STUB: not implemented"
	return *new(RgPyexecuteUpgrade)
}

func (c RgPyexecuteId) ReplaceWith(replaceWith string) RgPyexecuteReplaceWith {
	_ = "STUB: not implemented"
	return *new(RgPyexecuteReplaceWith)
}

func (c RgPyexecuteId) Requirements(requirement ...string) RgPyexecuteRequirementsRequirements {
	_ = "STUB: not implemented"
	return *new(RgPyexecuteRequirementsRequirements)
}

func (c RgPyexecuteId) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type RgPyexecuteReplaceWith Incomplete

func (c RgPyexecuteReplaceWith) Requirements(requirement ...string) RgPyexecuteRequirementsRequirements {
	_ = "STUB: not implemented"
	return *new(RgPyexecuteRequirementsRequirements)
}

func (c RgPyexecuteReplaceWith) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type RgPyexecuteRequirementsRequirements Incomplete

func (c RgPyexecuteRequirementsRequirements) Requirements(requirement ...string) RgPyexecuteRequirementsRequirements {
	_ = "STUB: not implemented"
	return *new(RgPyexecuteRequirementsRequirements)
}

func (c RgPyexecuteRequirementsRequirements) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type RgPyexecuteUnblocking Incomplete

func (c RgPyexecuteUnblocking) Id(id string) RgPyexecuteId {
	_ = "STUB: not implemented"
	return *new(RgPyexecuteId)
}

func (c RgPyexecuteUnblocking) Description(description string) RgPyexecuteDescription {
	_ = "STUB: not implemented"
	return *new(RgPyexecuteDescription)
}

func (c RgPyexecuteUnblocking) Upgrade() RgPyexecuteUpgrade {
	_ = "STUB: not implemented"
	return *new(RgPyexecuteUpgrade)
}

func (c RgPyexecuteUnblocking) ReplaceWith(replaceWith string) RgPyexecuteReplaceWith {
	_ = "STUB: not implemented"
	return *new(RgPyexecuteReplaceWith)
}

func (c RgPyexecuteUnblocking) Requirements(requirement ...string) RgPyexecuteRequirementsRequirements {
	_ = "STUB: not implemented"
	return *new(RgPyexecuteRequirementsRequirements)
}

func (c RgPyexecuteUnblocking) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type RgPyexecuteUpgrade Incomplete

func (c RgPyexecuteUpgrade) ReplaceWith(replaceWith string) RgPyexecuteReplaceWith {
	_ = "STUB: not implemented"
	return *new(RgPyexecuteReplaceWith)
}

func (c RgPyexecuteUpgrade) Requirements(requirement ...string) RgPyexecuteRequirementsRequirements {
	_ = "STUB: not implemented"
	return *new(RgPyexecuteRequirementsRequirements)
}

func (c RgPyexecuteUpgrade) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type RgPystats Incomplete

func (b Builder) RgPystats() (c RgPystats) { _ = "STUB: not implemented"; return *new(RgPystats) }

func (c RgPystats) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type RgRefreshcluster Incomplete

func (b Builder) RgRefreshcluster() (c RgRefreshcluster) {
	_ = "STUB: not implemented"
	return *new(RgRefreshcluster)
}

func (c RgRefreshcluster) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type RgTrigger Incomplete

func (b Builder) RgTrigger() (c RgTrigger) { _ = "STUB: not implemented"; return *new(RgTrigger) }

func (c RgTrigger) Trigger(trigger string) RgTriggerTrigger {
	_ = "STUB: not implemented"
	return *new(RgTriggerTrigger)
}

type RgTriggerArgument Incomplete

func (c RgTriggerArgument) Argument(argument ...string) RgTriggerArgument {
	_ = "STUB: not implemented"
	return *new(RgTriggerArgument)
}

func (c RgTriggerArgument) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type RgTriggerTrigger Incomplete

func (c RgTriggerTrigger) Argument(argument ...string) RgTriggerArgument {
	_ = "STUB: not implemented"
	return *new(RgTriggerArgument)
}

type RgUnregister Incomplete

func (b Builder) RgUnregister() (c RgUnregister) {
	_ = "STUB: not implemented"
	return *new(RgUnregister)
}

func (c RgUnregister) Id(id string) RgUnregisterId {
	_ = "STUB: not implemented"
	return *new(RgUnregisterId)
}

type RgUnregisterId Incomplete

func (c RgUnregisterId) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

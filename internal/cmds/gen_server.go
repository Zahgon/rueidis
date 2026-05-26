// Code generated DO NOT EDIT

package cmds

type AclCat Incomplete

func (b Builder) AclCat() (c AclCat) { _ = "STUB: not implemented"; return *new(AclCat) }

func (c AclCat) Categoryname(categoryname string) AclCatCategoryname {
	_ = "STUB: not implemented"
	return *new(AclCatCategoryname)
}

func (c AclCat) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type AclCatCategoryname Incomplete

func (c AclCatCategoryname) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type AclDeluser Incomplete

func (b Builder) AclDeluser() (c AclDeluser) { _ = "STUB: not implemented"; return *new(AclDeluser) }

func (c AclDeluser) Username(username ...string) AclDeluserUsername {
	_ = "STUB: not implemented"
	return *new(AclDeluserUsername)
}

type AclDeluserUsername Incomplete

func (c AclDeluserUsername) Username(username ...string) AclDeluserUsername {
	_ = "STUB: not implemented"
	return *new(AclDeluserUsername)
}

func (c AclDeluserUsername) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type AclDryrun Incomplete

func (b Builder) AclDryrun() (c AclDryrun) { _ = "STUB: not implemented"; return *new(AclDryrun) }

func (c AclDryrun) Username(username string) AclDryrunUsername {
	_ = "STUB: not implemented"
	return *new(AclDryrunUsername)
}

type AclDryrunArg Incomplete

func (c AclDryrunArg) Arg(arg ...string) AclDryrunArg {
	_ = "STUB: not implemented"
	return *new(AclDryrunArg)
}

func (c AclDryrunArg) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type AclDryrunCommand Incomplete

func (c AclDryrunCommand) Arg(arg ...string) AclDryrunArg {
	_ = "STUB: not implemented"
	return *new(AclDryrunArg)
}

func (c AclDryrunCommand) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type AclDryrunUsername Incomplete

func (c AclDryrunUsername) Command(command string) AclDryrunCommand {
	_ = "STUB: not implemented"
	return *new(AclDryrunCommand)
}

type AclGenpass Incomplete

func (b Builder) AclGenpass() (c AclGenpass) { _ = "STUB: not implemented"; return *new(AclGenpass) }

func (c AclGenpass) Bits(bits int64) AclGenpassBits {
	_ = "STUB: not implemented"
	return *new(AclGenpassBits)
}

func (c AclGenpass) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type AclGenpassBits Incomplete

func (c AclGenpassBits) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type AclGetuser Incomplete

func (b Builder) AclGetuser() (c AclGetuser) { _ = "STUB: not implemented"; return *new(AclGetuser) }

func (c AclGetuser) Username(username string) AclGetuserUsername {
	_ = "STUB: not implemented"
	return *new(AclGetuserUsername)
}

type AclGetuserUsername Incomplete

func (c AclGetuserUsername) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type AclHelp Incomplete

func (b Builder) AclHelp() (c AclHelp) { _ = "STUB: not implemented"; return *new(AclHelp) }

func (c AclHelp) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type AclList Incomplete

func (b Builder) AclList() (c AclList) { _ = "STUB: not implemented"; return *new(AclList) }

func (c AclList) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type AclLoad Incomplete

func (b Builder) AclLoad() (c AclLoad) { _ = "STUB: not implemented"; return *new(AclLoad) }

func (c AclLoad) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type AclLog Incomplete

func (b Builder) AclLog() (c AclLog) { _ = "STUB: not implemented"; return *new(AclLog) }

func (c AclLog) Count(count int64) AclLogCountCount {
	_ = "STUB: not implemented"
	return *new(AclLogCountCount)
}

func (c AclLog) Reset() AclLogCountReset { _ = "STUB: not implemented"; return *new(AclLogCountReset) }

type AclLogCountCount Incomplete

func (c AclLogCountCount) Reset() AclLogCountReset {
	_ = "STUB: not implemented"
	return *new(AclLogCountReset)
}

func (c AclLogCountCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type AclLogCountReset Incomplete

func (c AclLogCountReset) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type AclSave Incomplete

func (b Builder) AclSave() (c AclSave) { _ = "STUB: not implemented"; return *new(AclSave) }

func (c AclSave) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type AclSetuser Incomplete

func (b Builder) AclSetuser() (c AclSetuser) { _ = "STUB: not implemented"; return *new(AclSetuser) }

func (c AclSetuser) Username(username string) AclSetuserUsername {
	_ = "STUB: not implemented"
	return *new(AclSetuserUsername)
}

type AclSetuserRule Incomplete

func (c AclSetuserRule) Rule(rule ...string) AclSetuserRule {
	_ = "STUB: not implemented"
	return *new(AclSetuserRule)
}

func (c AclSetuserRule) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type AclSetuserUsername Incomplete

func (c AclSetuserUsername) Rule(rule ...string) AclSetuserRule {
	_ = "STUB: not implemented"
	return *new(AclSetuserRule)
}

func (c AclSetuserUsername) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type AclUsers Incomplete

func (b Builder) AclUsers() (c AclUsers) { _ = "STUB: not implemented"; return *new(AclUsers) }

func (c AclUsers) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type AclWhoami Incomplete

func (b Builder) AclWhoami() (c AclWhoami) { _ = "STUB: not implemented"; return *new(AclWhoami) }

func (c AclWhoami) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Bgrewriteaof Incomplete

func (b Builder) Bgrewriteaof() (c Bgrewriteaof) {
	_ = "STUB: not implemented"
	return *new(Bgrewriteaof)
}

func (c Bgrewriteaof) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Bgsave Incomplete

func (b Builder) Bgsave() (c Bgsave) { _ = "STUB: not implemented"; return *new(Bgsave) }

func (c Bgsave) Schedule() BgsaveSchedule { _ = "STUB: not implemented"; return *new(BgsaveSchedule) }

func (c Bgsave) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type BgsaveSchedule Incomplete

func (c BgsaveSchedule) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Command Incomplete

func (b Builder) Command() (c Command) { _ = "STUB: not implemented"; return *new(Command) }

func (c Command) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type CommandCount Incomplete

func (b Builder) CommandCount() (c CommandCount) {
	_ = "STUB: not implemented"
	return *new(CommandCount)
}

func (c CommandCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type CommandDocs Incomplete

func (b Builder) CommandDocs() (c CommandDocs) { _ = "STUB: not implemented"; return *new(CommandDocs) }

func (c CommandDocs) CommandName(commandName ...string) CommandDocsCommandName {
	_ = "STUB: not implemented"
	return *new(CommandDocsCommandName)
}

func (c CommandDocs) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type CommandDocsCommandName Incomplete

func (c CommandDocsCommandName) CommandName(commandName ...string) CommandDocsCommandName {
	_ = "STUB: not implemented"
	return *new(CommandDocsCommandName)
}

func (c CommandDocsCommandName) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type CommandGetkeys Incomplete

func (b Builder) CommandGetkeys() (c CommandGetkeys) {
	_ = "STUB: not implemented"
	return *new(CommandGetkeys)
}

func (c CommandGetkeys) Command(command string) CommandGetkeysCommand {
	_ = "STUB: not implemented"
	return *new(CommandGetkeysCommand)
}

type CommandGetkeysArg Incomplete

func (c CommandGetkeysArg) Arg(arg ...string) CommandGetkeysArg {
	_ = "STUB: not implemented"
	return *new(CommandGetkeysArg)
}

func (c CommandGetkeysArg) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type CommandGetkeysCommand Incomplete

func (c CommandGetkeysCommand) Arg(arg ...string) CommandGetkeysArg {
	_ = "STUB: not implemented"
	return *new(CommandGetkeysArg)
}

func (c CommandGetkeysCommand) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type CommandGetkeysandflags Incomplete

func (b Builder) CommandGetkeysandflags() (c CommandGetkeysandflags) {
	_ = "STUB: not implemented"
	return *new(CommandGetkeysandflags)
}

func (c CommandGetkeysandflags) Command(command string) CommandGetkeysandflagsCommand {
	_ = "STUB: not implemented"
	return *new(CommandGetkeysandflagsCommand)
}

type CommandGetkeysandflagsArg Incomplete

func (c CommandGetkeysandflagsArg) Arg(arg ...string) CommandGetkeysandflagsArg {
	_ = "STUB: not implemented"
	return *new(CommandGetkeysandflagsArg)
}

func (c CommandGetkeysandflagsArg) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type CommandGetkeysandflagsCommand Incomplete

func (c CommandGetkeysandflagsCommand) Arg(arg ...string) CommandGetkeysandflagsArg {
	_ = "STUB: not implemented"
	return *new(CommandGetkeysandflagsArg)
}

func (c CommandGetkeysandflagsCommand) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type CommandInfo Incomplete

func (b Builder) CommandInfo() (c CommandInfo) { _ = "STUB: not implemented"; return *new(CommandInfo) }

func (c CommandInfo) CommandName(commandName ...string) CommandInfoCommandName {
	_ = "STUB: not implemented"
	return *new(CommandInfoCommandName)
}

func (c CommandInfo) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type CommandInfoCommandName Incomplete

func (c CommandInfoCommandName) CommandName(commandName ...string) CommandInfoCommandName {
	_ = "STUB: not implemented"
	return *new(CommandInfoCommandName)
}

func (c CommandInfoCommandName) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type CommandList Incomplete

func (b Builder) CommandList() (c CommandList) { _ = "STUB: not implemented"; return *new(CommandList) }

func (c CommandList) FilterbyModuleName(name string) CommandListFilterbyModuleName {
	_ = "STUB: not implemented"
	return *new(CommandListFilterbyModuleName)
}

func (c CommandList) FilterbyAclcatCategory(category string) CommandListFilterbyAclcatCategory {
	_ = "STUB: not implemented"
	return *new(CommandListFilterbyAclcatCategory)
}

func (c CommandList) FilterbyPatternPattern(pattern string) CommandListFilterbyPatternPattern {
	_ = "STUB: not implemented"
	return *new(CommandListFilterbyPatternPattern)
}

func (c CommandList) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type CommandListFilterbyAclcatCategory Incomplete

func (c CommandListFilterbyAclcatCategory) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type CommandListFilterbyModuleName Incomplete

func (c CommandListFilterbyModuleName) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type CommandListFilterbyPatternPattern Incomplete

func (c CommandListFilterbyPatternPattern) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ConfigGet Incomplete

func (b Builder) ConfigGet() (c ConfigGet) { _ = "STUB: not implemented"; return *new(ConfigGet) }

func (c ConfigGet) Parameter(parameter ...string) ConfigGetParameter {
	_ = "STUB: not implemented"
	return *new(ConfigGetParameter)
}

type ConfigGetParameter Incomplete

func (c ConfigGetParameter) Parameter(parameter ...string) ConfigGetParameter {
	_ = "STUB: not implemented"
	return *new(ConfigGetParameter)
}

func (c ConfigGetParameter) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ConfigResetstat Incomplete

func (b Builder) ConfigResetstat() (c ConfigResetstat) {
	_ = "STUB: not implemented"
	return *new(ConfigResetstat)
}

func (c ConfigResetstat) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ConfigRewrite Incomplete

func (b Builder) ConfigRewrite() (c ConfigRewrite) {
	_ = "STUB: not implemented"
	return *new(ConfigRewrite)
}

func (c ConfigRewrite) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ConfigSet Incomplete

func (b Builder) ConfigSet() (c ConfigSet) { _ = "STUB: not implemented"; return *new(ConfigSet) }

func (c ConfigSet) ParameterValue() ConfigSetParameterValue {
	_ = "STUB: not implemented"
	return *new(ConfigSetParameterValue)
}

type ConfigSetParameterValue Incomplete

func (c ConfigSetParameterValue) ParameterValue(parameter string, value string) ConfigSetParameterValue {
	_ = "STUB: not implemented"
	return *new(ConfigSetParameterValue)
}

func (c ConfigSetParameterValue) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type Dbsize Incomplete

func (b Builder) Dbsize() (c Dbsize) { _ = "STUB: not implemented"; return *new(Dbsize) }

func (c Dbsize) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type DebugObject Incomplete

func (b Builder) DebugObject() (c DebugObject) { _ = "STUB: not implemented"; return *new(DebugObject) }

func (c DebugObject) Key(key string) DebugObjectKey {
	_ = "STUB: not implemented"
	return *new(DebugObjectKey)
}

type DebugObjectKey Incomplete

func (c DebugObjectKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type DebugSegfault Incomplete

func (b Builder) DebugSegfault() (c DebugSegfault) {
	_ = "STUB: not implemented"
	return *new(DebugSegfault)
}

func (c DebugSegfault) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Failover Incomplete

func (b Builder) Failover() (c Failover) { _ = "STUB: not implemented"; return *new(Failover) }

func (c Failover) To() FailoverTargetTo { _ = "STUB: not implemented"; return *new(FailoverTargetTo) }

func (c Failover) Abort() FailoverAbort { _ = "STUB: not implemented"; return *new(FailoverAbort) }

func (c Failover) Timeout(milliseconds int64) FailoverTimeout {
	_ = "STUB: not implemented"
	return *new(FailoverTimeout)
}

func (c Failover) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FailoverAbort Incomplete

func (c FailoverAbort) Timeout(milliseconds int64) FailoverTimeout {
	_ = "STUB: not implemented"
	return *new(FailoverTimeout)
}

func (c FailoverAbort) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FailoverTargetForce Incomplete

func (c FailoverTargetForce) Abort() FailoverAbort {
	_ = "STUB: not implemented"
	return *new(FailoverAbort)
}

func (c FailoverTargetForce) Timeout(milliseconds int64) FailoverTimeout {
	_ = "STUB: not implemented"
	return *new(FailoverTimeout)
}

func (c FailoverTargetForce) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FailoverTargetHost Incomplete

func (c FailoverTargetHost) Port(port int64) FailoverTargetPort {
	_ = "STUB: not implemented"
	return *new(FailoverTargetPort)
}

type FailoverTargetPort Incomplete

func (c FailoverTargetPort) Force() FailoverTargetForce {
	_ = "STUB: not implemented"
	return *new(FailoverTargetForce)
}

func (c FailoverTargetPort) Abort() FailoverAbort {
	_ = "STUB: not implemented"
	return *new(FailoverAbort)
}

func (c FailoverTargetPort) Timeout(milliseconds int64) FailoverTimeout {
	_ = "STUB: not implemented"
	return *new(FailoverTimeout)
}

func (c FailoverTargetPort) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FailoverTargetTo Incomplete

func (c FailoverTargetTo) Host(host string) FailoverTargetHost {
	_ = "STUB: not implemented"
	return *new(FailoverTargetHost)
}

type FailoverTimeout Incomplete

func (c FailoverTimeout) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Flushall Incomplete

func (b Builder) Flushall() (c Flushall) { _ = "STUB: not implemented"; return *new(Flushall) }

func (c Flushall) Async() FlushallAsync { _ = "STUB: not implemented"; return *new(FlushallAsync) }

func (c Flushall) Sync() FlushallAsyncSync {
	_ = "STUB: not implemented"
	return *new(FlushallAsyncSync)
}

func (c Flushall) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FlushallAsync Incomplete

func (c FlushallAsync) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FlushallAsyncSync Incomplete

func (c FlushallAsyncSync) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Flushdb Incomplete

func (b Builder) Flushdb() (c Flushdb) { _ = "STUB: not implemented"; return *new(Flushdb) }

func (c Flushdb) Async() FlushdbAsync { _ = "STUB: not implemented"; return *new(FlushdbAsync) }

func (c Flushdb) Sync() FlushdbAsyncSync { _ = "STUB: not implemented"; return *new(FlushdbAsyncSync) }

func (c Flushdb) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FlushdbAsync Incomplete

func (c FlushdbAsync) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type FlushdbAsyncSync Incomplete

func (c FlushdbAsyncSync) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HotkeysGet Incomplete

func (b Builder) HotkeysGet() (c HotkeysGet) { _ = "STUB: not implemented"; return *new(HotkeysGet) }

func (c HotkeysGet) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HotkeysReset Incomplete

func (b Builder) HotkeysReset() (c HotkeysReset) {
	_ = "STUB: not implemented"
	return *new(HotkeysReset)
}

func (c HotkeysReset) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HotkeysStart Incomplete

func (b Builder) HotkeysStart() (c HotkeysStart) {
	_ = "STUB: not implemented"
	return *new(HotkeysStart)
}

func (c HotkeysStart) Metrics() HotkeysStartMetricsMetrics {
	_ = "STUB: not implemented"
	return *new(HotkeysStartMetricsMetrics)
}

type HotkeysStartCount Incomplete

func (c HotkeysStartCount) Duration(seconds int64) HotkeysStartDuration {
	_ = "STUB: not implemented"
	return *new(HotkeysStartDuration)
}

func (c HotkeysStartCount) Sample(ratio int64) HotkeysStartSample {
	_ = "STUB: not implemented"
	return *new(HotkeysStartSample)
}

func (c HotkeysStartCount) Slots() HotkeysStartSlotsSlots {
	_ = "STUB: not implemented"
	return *new(HotkeysStartSlotsSlots)
}

func (c HotkeysStartCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HotkeysStartDuration Incomplete

func (c HotkeysStartDuration) Sample(ratio int64) HotkeysStartSample {
	_ = "STUB: not implemented"
	return *new(HotkeysStartSample)
}

func (c HotkeysStartDuration) Slots() HotkeysStartSlotsSlots {
	_ = "STUB: not implemented"
	return *new(HotkeysStartSlotsSlots)
}

func (c HotkeysStartDuration) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HotkeysStartMetricsCount Incomplete

func (c HotkeysStartMetricsCount) Cpu() HotkeysStartMetricsCpu {
	_ = "STUB: not implemented"
	return *new(HotkeysStartMetricsCpu)
}

func (c HotkeysStartMetricsCount) Net() HotkeysStartMetricsNet {
	_ = "STUB: not implemented"
	return *new(HotkeysStartMetricsNet)
}

func (c HotkeysStartMetricsCount) Count(k int64) HotkeysStartCount {
	_ = "STUB: not implemented"
	return *new(HotkeysStartCount)
}

func (c HotkeysStartMetricsCount) Duration(seconds int64) HotkeysStartDuration {
	_ = "STUB: not implemented"
	return *new(HotkeysStartDuration)
}

func (c HotkeysStartMetricsCount) Sample(ratio int64) HotkeysStartSample {
	_ = "STUB: not implemented"
	return *new(HotkeysStartSample)
}

func (c HotkeysStartMetricsCount) Slots() HotkeysStartSlotsSlots {
	_ = "STUB: not implemented"
	return *new(HotkeysStartSlotsSlots)
}

func (c HotkeysStartMetricsCount) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type HotkeysStartMetricsCpu Incomplete

func (c HotkeysStartMetricsCpu) Net() HotkeysStartMetricsNet {
	_ = "STUB: not implemented"
	return *new(HotkeysStartMetricsNet)
}

func (c HotkeysStartMetricsCpu) Count(k int64) HotkeysStartCount {
	_ = "STUB: not implemented"
	return *new(HotkeysStartCount)
}

func (c HotkeysStartMetricsCpu) Duration(seconds int64) HotkeysStartDuration {
	_ = "STUB: not implemented"
	return *new(HotkeysStartDuration)
}

func (c HotkeysStartMetricsCpu) Sample(ratio int64) HotkeysStartSample {
	_ = "STUB: not implemented"
	return *new(HotkeysStartSample)
}

func (c HotkeysStartMetricsCpu) Slots() HotkeysStartSlotsSlots {
	_ = "STUB: not implemented"
	return *new(HotkeysStartSlotsSlots)
}

func (c HotkeysStartMetricsCpu) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type HotkeysStartMetricsMetrics Incomplete

func (c HotkeysStartMetricsMetrics) Count(count int64) HotkeysStartMetricsCount {
	_ = "STUB: not implemented"
	return *new(HotkeysStartMetricsCount)
}

type HotkeysStartMetricsNet Incomplete

func (c HotkeysStartMetricsNet) Count(k int64) HotkeysStartCount {
	_ = "STUB: not implemented"
	return *new(HotkeysStartCount)
}

func (c HotkeysStartMetricsNet) Duration(seconds int64) HotkeysStartDuration {
	_ = "STUB: not implemented"
	return *new(HotkeysStartDuration)
}

func (c HotkeysStartMetricsNet) Sample(ratio int64) HotkeysStartSample {
	_ = "STUB: not implemented"
	return *new(HotkeysStartSample)
}

func (c HotkeysStartMetricsNet) Slots() HotkeysStartSlotsSlots {
	_ = "STUB: not implemented"
	return *new(HotkeysStartSlotsSlots)
}

func (c HotkeysStartMetricsNet) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type HotkeysStartSample Incomplete

func (c HotkeysStartSample) Slots() HotkeysStartSlotsSlots {
	_ = "STUB: not implemented"
	return *new(HotkeysStartSlotsSlots)
}

func (c HotkeysStartSample) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HotkeysStartSlotsCount Incomplete

func (c HotkeysStartSlotsCount) Slot(slot ...int64) HotkeysStartSlotsSlot {
	_ = "STUB: not implemented"
	return *new(HotkeysStartSlotsSlot)
}

type HotkeysStartSlotsSlot Incomplete

func (c HotkeysStartSlotsSlot) Slot(slot ...int64) HotkeysStartSlotsSlot {
	_ = "STUB: not implemented"
	return *new(HotkeysStartSlotsSlot)
}

func (c HotkeysStartSlotsSlot) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HotkeysStartSlotsSlots Incomplete

func (c HotkeysStartSlotsSlots) Count(count int64) HotkeysStartSlotsCount {
	_ = "STUB: not implemented"
	return *new(HotkeysStartSlotsCount)
}

type HotkeysStop Incomplete

func (b Builder) HotkeysStop() (c HotkeysStop) { _ = "STUB: not implemented"; return *new(HotkeysStop) }

func (c HotkeysStop) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Info Incomplete

func (b Builder) Info() (c Info) { _ = "STUB: not implemented"; return *new(Info) }

func (c Info) Section(section ...string) InfoSection {
	_ = "STUB: not implemented"
	return *new(InfoSection)
}

func (c Info) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type InfoSection Incomplete

func (c InfoSection) Section(section ...string) InfoSection {
	_ = "STUB: not implemented"
	return *new(InfoSection)
}

func (c InfoSection) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Lastsave Incomplete

func (b Builder) Lastsave() (c Lastsave) { _ = "STUB: not implemented"; return *new(Lastsave) }

func (c Lastsave) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type LatencyDoctor Incomplete

func (b Builder) LatencyDoctor() (c LatencyDoctor) {
	_ = "STUB: not implemented"
	return *new(LatencyDoctor)
}

func (c LatencyDoctor) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type LatencyGraph Incomplete

func (b Builder) LatencyGraph() (c LatencyGraph) {
	_ = "STUB: not implemented"
	return *new(LatencyGraph)
}

func (c LatencyGraph) Event(event string) LatencyGraphEvent {
	_ = "STUB: not implemented"
	return *new(LatencyGraphEvent)
}

type LatencyGraphEvent Incomplete

func (c LatencyGraphEvent) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type LatencyHelp Incomplete

func (b Builder) LatencyHelp() (c LatencyHelp) { _ = "STUB: not implemented"; return *new(LatencyHelp) }

func (c LatencyHelp) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type LatencyHistogram Incomplete

func (b Builder) LatencyHistogram() (c LatencyHistogram) {
	_ = "STUB: not implemented"
	return *new(LatencyHistogram)
}

func (c LatencyHistogram) Command(command ...string) LatencyHistogramCommand {
	_ = "STUB: not implemented"
	return *new(LatencyHistogramCommand)
}

func (c LatencyHistogram) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type LatencyHistogramCommand Incomplete

func (c LatencyHistogramCommand) Command(command ...string) LatencyHistogramCommand {
	_ = "STUB: not implemented"
	return *new(LatencyHistogramCommand)
}

func (c LatencyHistogramCommand) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type LatencyHistory Incomplete

func (b Builder) LatencyHistory() (c LatencyHistory) {
	_ = "STUB: not implemented"
	return *new(LatencyHistory)
}

func (c LatencyHistory) Event(event string) LatencyHistoryEvent {
	_ = "STUB: not implemented"
	return *new(LatencyHistoryEvent)
}

type LatencyHistoryEvent Incomplete

func (c LatencyHistoryEvent) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type LatencyLatest Incomplete

func (b Builder) LatencyLatest() (c LatencyLatest) {
	_ = "STUB: not implemented"
	return *new(LatencyLatest)
}

func (c LatencyLatest) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type LatencyReset Incomplete

func (b Builder) LatencyReset() (c LatencyReset) {
	_ = "STUB: not implemented"
	return *new(LatencyReset)
}

func (c LatencyReset) Event(event ...string) LatencyResetEvent {
	_ = "STUB: not implemented"
	return *new(LatencyResetEvent)
}

func (c LatencyReset) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type LatencyResetEvent Incomplete

func (c LatencyResetEvent) Event(event ...string) LatencyResetEvent {
	_ = "STUB: not implemented"
	return *new(LatencyResetEvent)
}

func (c LatencyResetEvent) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Lolwut Incomplete

func (b Builder) Lolwut() (c Lolwut) { _ = "STUB: not implemented"; return *new(Lolwut) }

func (c Lolwut) Version(version int64) LolwutVersion {
	_ = "STUB: not implemented"
	return *new(LolwutVersion)
}

func (c Lolwut) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type LolwutVersion Incomplete

func (c LolwutVersion) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type MemoryDoctor Incomplete

func (b Builder) MemoryDoctor() (c MemoryDoctor) {
	_ = "STUB: not implemented"
	return *new(MemoryDoctor)
}

func (c MemoryDoctor) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type MemoryHelp Incomplete

func (b Builder) MemoryHelp() (c MemoryHelp) { _ = "STUB: not implemented"; return *new(MemoryHelp) }

func (c MemoryHelp) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type MemoryMallocStats Incomplete

func (b Builder) MemoryMallocStats() (c MemoryMallocStats) {
	_ = "STUB: not implemented"
	return *new(MemoryMallocStats)
}

func (c MemoryMallocStats) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type MemoryPurge Incomplete

func (b Builder) MemoryPurge() (c MemoryPurge) { _ = "STUB: not implemented"; return *new(MemoryPurge) }

func (c MemoryPurge) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type MemoryStats Incomplete

func (b Builder) MemoryStats() (c MemoryStats) { _ = "STUB: not implemented"; return *new(MemoryStats) }

func (c MemoryStats) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type MemoryUsage Incomplete

func (b Builder) MemoryUsage() (c MemoryUsage) { _ = "STUB: not implemented"; return *new(MemoryUsage) }

func (c MemoryUsage) Key(key string) MemoryUsageKey {
	_ = "STUB: not implemented"
	return *new(MemoryUsageKey)
}

type MemoryUsageKey Incomplete

func (c MemoryUsageKey) Samples(count int64) MemoryUsageSamples {
	_ = "STUB: not implemented"
	return *new(MemoryUsageSamples)
}

func (c MemoryUsageKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type MemoryUsageSamples Incomplete

func (c MemoryUsageSamples) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ModuleList Incomplete

func (b Builder) ModuleList() (c ModuleList) { _ = "STUB: not implemented"; return *new(ModuleList) }

func (c ModuleList) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ModuleLoad Incomplete

func (b Builder) ModuleLoad() (c ModuleLoad) { _ = "STUB: not implemented"; return *new(ModuleLoad) }

func (c ModuleLoad) Path(path string) ModuleLoadPath {
	_ = "STUB: not implemented"
	return *new(ModuleLoadPath)
}

type ModuleLoadArg Incomplete

func (c ModuleLoadArg) Arg(arg ...string) ModuleLoadArg {
	_ = "STUB: not implemented"
	return *new(ModuleLoadArg)
}

func (c ModuleLoadArg) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ModuleLoadPath Incomplete

func (c ModuleLoadPath) Arg(arg ...string) ModuleLoadArg {
	_ = "STUB: not implemented"
	return *new(ModuleLoadArg)
}

func (c ModuleLoadPath) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ModuleLoadex Incomplete

func (b Builder) ModuleLoadex() (c ModuleLoadex) {
	_ = "STUB: not implemented"
	return *new(ModuleLoadex)
}

func (c ModuleLoadex) Path(path string) ModuleLoadexPath {
	_ = "STUB: not implemented"
	return *new(ModuleLoadexPath)
}

type ModuleLoadexArgs Incomplete

func (c ModuleLoadexArgs) Args(args ...string) ModuleLoadexArgs {
	_ = "STUB: not implemented"
	return *new(ModuleLoadexArgs)
}

func (c ModuleLoadexArgs) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ModuleLoadexConfig Incomplete

func (c ModuleLoadexConfig) Config(name string, value string) ModuleLoadexConfig {
	_ = "STUB: not implemented"
	return *new(ModuleLoadexConfig)
}

func (c ModuleLoadexConfig) Args(args ...string) ModuleLoadexArgs {
	_ = "STUB: not implemented"
	return *new(ModuleLoadexArgs)
}

func (c ModuleLoadexConfig) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ModuleLoadexPath Incomplete

func (c ModuleLoadexPath) Config() ModuleLoadexConfig {
	_ = "STUB: not implemented"
	return *new(ModuleLoadexConfig)
}

func (c ModuleLoadexPath) Args(args ...string) ModuleLoadexArgs {
	_ = "STUB: not implemented"
	return *new(ModuleLoadexArgs)
}

func (c ModuleLoadexPath) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ModuleUnload Incomplete

func (b Builder) ModuleUnload() (c ModuleUnload) {
	_ = "STUB: not implemented"
	return *new(ModuleUnload)
}

func (c ModuleUnload) Name(name string) ModuleUnloadName {
	_ = "STUB: not implemented"
	return *new(ModuleUnloadName)
}

type ModuleUnloadName Incomplete

func (c ModuleUnloadName) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Monitor Incomplete

func (b Builder) Monitor() (c Monitor) { _ = "STUB: not implemented"; return *new(Monitor) }

func (c Monitor) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Psync Incomplete

func (b Builder) Psync() (c Psync) { _ = "STUB: not implemented"; return *new(Psync) }

func (c Psync) Replicationid(replicationid string) PsyncReplicationid {
	_ = "STUB: not implemented"
	return *new(PsyncReplicationid)
}

type PsyncOffset Incomplete

func (c PsyncOffset) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type PsyncReplicationid Incomplete

func (c PsyncReplicationid) Offset(offset int64) PsyncOffset {
	_ = "STUB: not implemented"
	return *new(PsyncOffset)
}

type Replicaof Incomplete

func (b Builder) Replicaof() (c Replicaof) { _ = "STUB: not implemented"; return *new(Replicaof) }

func (c Replicaof) Host(host string) ReplicaofArgsHostPortHost {
	_ = "STUB: not implemented"
	return *new(ReplicaofArgsHostPortHost)
}

func (c Replicaof) No() ReplicaofArgsNoOneNo {
	_ = "STUB: not implemented"
	return *new(ReplicaofArgsNoOneNo)
}

type ReplicaofArgsHostPortHost Incomplete

func (c ReplicaofArgsHostPortHost) Port(port int64) ReplicaofArgsHostPortPort {
	_ = "STUB: not implemented"
	return *new(ReplicaofArgsHostPortPort)
}

type ReplicaofArgsHostPortPort Incomplete

func (c ReplicaofArgsHostPortPort) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ReplicaofArgsNoOneNo Incomplete

func (c ReplicaofArgsNoOneNo) One() ReplicaofArgsNoOneOne {
	_ = "STUB: not implemented"
	return *new(ReplicaofArgsNoOneOne)
}

type ReplicaofArgsNoOneOne Incomplete

func (c ReplicaofArgsNoOneOne) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Role Incomplete

func (b Builder) Role() (c Role) { _ = "STUB: not implemented"; return *new(Role) }

func (c Role) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Save Incomplete

func (b Builder) Save() (c Save) { _ = "STUB: not implemented"; return *new(Save) }

func (c Save) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Shutdown Incomplete

func (b Builder) Shutdown() (c Shutdown) { _ = "STUB: not implemented"; return *new(Shutdown) }

func (c Shutdown) Nosave() ShutdownSaveModeNosave {
	_ = "STUB: not implemented"
	return *new(ShutdownSaveModeNosave)
}

func (c Shutdown) Save() ShutdownSaveModeSave {
	_ = "STUB: not implemented"
	return *new(ShutdownSaveModeSave)
}

func (c Shutdown) Now() ShutdownNow { _ = "STUB: not implemented"; return *new(ShutdownNow) }

func (c Shutdown) Force() ShutdownForce { _ = "STUB: not implemented"; return *new(ShutdownForce) }

func (c Shutdown) Safe() ShutdownSafe { _ = "STUB: not implemented"; return *new(ShutdownSafe) }

func (c Shutdown) Abort() ShutdownAbort { _ = "STUB: not implemented"; return *new(ShutdownAbort) }

func (c Shutdown) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ShutdownAbort Incomplete

func (c ShutdownAbort) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ShutdownForce Incomplete

func (c ShutdownForce) Safe() ShutdownSafe { _ = "STUB: not implemented"; return *new(ShutdownSafe) }

func (c ShutdownForce) Abort() ShutdownAbort { _ = "STUB: not implemented"; return *new(ShutdownAbort) }

func (c ShutdownForce) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ShutdownNow Incomplete

func (c ShutdownNow) Force() ShutdownForce { _ = "STUB: not implemented"; return *new(ShutdownForce) }

func (c ShutdownNow) Safe() ShutdownSafe { _ = "STUB: not implemented"; return *new(ShutdownSafe) }

func (c ShutdownNow) Abort() ShutdownAbort { _ = "STUB: not implemented"; return *new(ShutdownAbort) }

func (c ShutdownNow) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ShutdownSafe Incomplete

func (c ShutdownSafe) Abort() ShutdownAbort { _ = "STUB: not implemented"; return *new(ShutdownAbort) }

func (c ShutdownSafe) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ShutdownSaveModeNosave Incomplete

func (c ShutdownSaveModeNosave) Now() ShutdownNow {
	_ = "STUB: not implemented"
	return *new(ShutdownNow)
}

func (c ShutdownSaveModeNosave) Force() ShutdownForce {
	_ = "STUB: not implemented"
	return *new(ShutdownForce)
}

func (c ShutdownSaveModeNosave) Safe() ShutdownSafe {
	_ = "STUB: not implemented"
	return *new(ShutdownSafe)
}

func (c ShutdownSaveModeNosave) Abort() ShutdownAbort {
	_ = "STUB: not implemented"
	return *new(ShutdownAbort)
}

func (c ShutdownSaveModeNosave) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ShutdownSaveModeSave Incomplete

func (c ShutdownSaveModeSave) Now() ShutdownNow {
	_ = "STUB: not implemented"
	return *new(ShutdownNow)
}

func (c ShutdownSaveModeSave) Force() ShutdownForce {
	_ = "STUB: not implemented"
	return *new(ShutdownForce)
}

func (c ShutdownSaveModeSave) Safe() ShutdownSafe {
	_ = "STUB: not implemented"
	return *new(ShutdownSafe)
}

func (c ShutdownSaveModeSave) Abort() ShutdownAbort {
	_ = "STUB: not implemented"
	return *new(ShutdownAbort)
}

func (c ShutdownSaveModeSave) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Slaveof Incomplete

func (b Builder) Slaveof() (c Slaveof) { _ = "STUB: not implemented"; return *new(Slaveof) }

func (c Slaveof) Host(host string) SlaveofArgsHostPortHost {
	_ = "STUB: not implemented"
	return *new(SlaveofArgsHostPortHost)
}

func (c Slaveof) No() SlaveofArgsNoOneNo {
	_ = "STUB: not implemented"
	return *new(SlaveofArgsNoOneNo)
}

type SlaveofArgsHostPortHost Incomplete

func (c SlaveofArgsHostPortHost) Port(port int64) SlaveofArgsHostPortPort {
	_ = "STUB: not implemented"
	return *new(SlaveofArgsHostPortPort)
}

type SlaveofArgsHostPortPort Incomplete

func (c SlaveofArgsHostPortPort) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type SlaveofArgsNoOneNo Incomplete

func (c SlaveofArgsNoOneNo) One() SlaveofArgsNoOneOne {
	_ = "STUB: not implemented"
	return *new(SlaveofArgsNoOneOne)
}

type SlaveofArgsNoOneOne Incomplete

func (c SlaveofArgsNoOneOne) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SlowlogGet Incomplete

func (b Builder) SlowlogGet() (c SlowlogGet) { _ = "STUB: not implemented"; return *new(SlowlogGet) }

func (c SlowlogGet) Count(count int64) SlowlogGetCount {
	_ = "STUB: not implemented"
	return *new(SlowlogGetCount)
}

func (c SlowlogGet) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SlowlogGetCount Incomplete

func (c SlowlogGetCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SlowlogHelp Incomplete

func (b Builder) SlowlogHelp() (c SlowlogHelp) { _ = "STUB: not implemented"; return *new(SlowlogHelp) }

func (c SlowlogHelp) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SlowlogLen Incomplete

func (b Builder) SlowlogLen() (c SlowlogLen) { _ = "STUB: not implemented"; return *new(SlowlogLen) }

func (c SlowlogLen) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SlowlogReset Incomplete

func (b Builder) SlowlogReset() (c SlowlogReset) {
	_ = "STUB: not implemented"
	return *new(SlowlogReset)
}

func (c SlowlogReset) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Swapdb Incomplete

func (b Builder) Swapdb() (c Swapdb) { _ = "STUB: not implemented"; return *new(Swapdb) }

func (c Swapdb) Index1(index1 int64) SwapdbIndex1 {
	_ = "STUB: not implemented"
	return *new(SwapdbIndex1)
}

type SwapdbIndex1 Incomplete

func (c SwapdbIndex1) Index2(index2 int64) SwapdbIndex2 {
	_ = "STUB: not implemented"
	return *new(SwapdbIndex2)
}

type SwapdbIndex2 Incomplete

func (c SwapdbIndex2) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Sync Incomplete

func (b Builder) Sync() (c Sync) { _ = "STUB: not implemented"; return *new(Sync) }

func (c Sync) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Time Incomplete

func (b Builder) Time() (c Time) { _ = "STUB: not implemented"; return *new(Time) }

func (c Time) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

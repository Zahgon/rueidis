// Code generated DO NOT EDIT

package cmds

type Asking Incomplete

func (b Builder) Asking() (c Asking) { _ = "STUB: not implemented"; return *new(Asking) }

func (c Asking) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClusterAddslots Incomplete

func (b Builder) ClusterAddslots() (c ClusterAddslots) {
	_ = "STUB: not implemented"
	return *new(ClusterAddslots)
}

func (c ClusterAddslots) Slot(slot ...int64) ClusterAddslotsSlot {
	_ = "STUB: not implemented"
	return *new(ClusterAddslotsSlot)
}

type ClusterAddslotsSlot Incomplete

func (c ClusterAddslotsSlot) Slot(slot ...int64) ClusterAddslotsSlot {
	_ = "STUB: not implemented"
	return *new(ClusterAddslotsSlot)
}

func (c ClusterAddslotsSlot) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClusterAddslotsrange Incomplete

func (b Builder) ClusterAddslotsrange() (c ClusterAddslotsrange) {
	_ = "STUB: not implemented"
	return *new(ClusterAddslotsrange)
}

func (c ClusterAddslotsrange) StartSlotEndSlot() ClusterAddslotsrangeStartSlotEndSlot {
	_ = "STUB: not implemented"
	return *new(ClusterAddslotsrangeStartSlotEndSlot)
}

type ClusterAddslotsrangeStartSlotEndSlot Incomplete

func (c ClusterAddslotsrangeStartSlotEndSlot) StartSlotEndSlot(startSlot int64, endSlot int64) ClusterAddslotsrangeStartSlotEndSlot {
	_ = "STUB: not implemented"
	return *new(ClusterAddslotsrangeStartSlotEndSlot)
}

func (c ClusterAddslotsrangeStartSlotEndSlot) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClusterBumpepoch Incomplete

func (b Builder) ClusterBumpepoch() (c ClusterBumpepoch) {
	_ = "STUB: not implemented"
	return *new(ClusterBumpepoch)
}

func (c ClusterBumpepoch) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClusterCancelslotmigrations Incomplete

func (b Builder) ClusterCancelslotmigrations() (c ClusterCancelslotmigrations) {
	_ = "STUB: not implemented"
	return *new(ClusterCancelslotmigrations)
}

func (c ClusterCancelslotmigrations) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClusterCountFailureReports Incomplete

func (b Builder) ClusterCountFailureReports() (c ClusterCountFailureReports) {
	_ = "STUB: not implemented"
	return *new(ClusterCountFailureReports)
}

func (c ClusterCountFailureReports) NodeId(nodeId string) ClusterCountFailureReportsNodeId {
	_ = "STUB: not implemented"
	return *new(ClusterCountFailureReportsNodeId)
}

type ClusterCountFailureReportsNodeId Incomplete

func (c ClusterCountFailureReportsNodeId) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClusterCountkeysinslot Incomplete

func (b Builder) ClusterCountkeysinslot() (c ClusterCountkeysinslot) {
	_ = "STUB: not implemented"
	return *new(ClusterCountkeysinslot)
}

func (c ClusterCountkeysinslot) Slot(slot int64) ClusterCountkeysinslotSlot {
	_ = "STUB: not implemented"
	return *new(ClusterCountkeysinslotSlot)
}

type ClusterCountkeysinslotSlot Incomplete

func (c ClusterCountkeysinslotSlot) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClusterDelslots Incomplete

func (b Builder) ClusterDelslots() (c ClusterDelslots) {
	_ = "STUB: not implemented"
	return *new(ClusterDelslots)
}

func (c ClusterDelslots) Slot(slot ...int64) ClusterDelslotsSlot {
	_ = "STUB: not implemented"
	return *new(ClusterDelslotsSlot)
}

type ClusterDelslotsSlot Incomplete

func (c ClusterDelslotsSlot) Slot(slot ...int64) ClusterDelslotsSlot {
	_ = "STUB: not implemented"
	return *new(ClusterDelslotsSlot)
}

func (c ClusterDelslotsSlot) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClusterDelslotsrange Incomplete

func (b Builder) ClusterDelslotsrange() (c ClusterDelslotsrange) {
	_ = "STUB: not implemented"
	return *new(ClusterDelslotsrange)
}

func (c ClusterDelslotsrange) StartSlotEndSlot() ClusterDelslotsrangeStartSlotEndSlot {
	_ = "STUB: not implemented"
	return *new(ClusterDelslotsrangeStartSlotEndSlot)
}

type ClusterDelslotsrangeStartSlotEndSlot Incomplete

func (c ClusterDelslotsrangeStartSlotEndSlot) StartSlotEndSlot(startSlot int64, endSlot int64) ClusterDelslotsrangeStartSlotEndSlot {
	_ = "STUB: not implemented"
	return *new(ClusterDelslotsrangeStartSlotEndSlot)
}

func (c ClusterDelslotsrangeStartSlotEndSlot) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClusterFailover Incomplete

func (b Builder) ClusterFailover() (c ClusterFailover) {
	_ = "STUB: not implemented"
	return *new(ClusterFailover)
}

func (c ClusterFailover) Force() ClusterFailoverOptionsForce {
	_ = "STUB: not implemented"
	return *new(ClusterFailoverOptionsForce)
}

func (c ClusterFailover) Takeover() ClusterFailoverOptionsTakeover {
	_ = "STUB: not implemented"
	return *new(ClusterFailoverOptionsTakeover)
}

func (c ClusterFailover) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClusterFailoverOptionsForce Incomplete

func (c ClusterFailoverOptionsForce) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClusterFailoverOptionsTakeover Incomplete

func (c ClusterFailoverOptionsTakeover) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClusterFlushslots Incomplete

func (b Builder) ClusterFlushslots() (c ClusterFlushslots) {
	_ = "STUB: not implemented"
	return *new(ClusterFlushslots)
}

func (c ClusterFlushslots) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClusterForget Incomplete

func (b Builder) ClusterForget() (c ClusterForget) {
	_ = "STUB: not implemented"
	return *new(ClusterForget)
}

func (c ClusterForget) NodeId(nodeId string) ClusterForgetNodeId {
	_ = "STUB: not implemented"
	return *new(ClusterForgetNodeId)
}

type ClusterForgetNodeId Incomplete

func (c ClusterForgetNodeId) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClusterGetkeysinslot Incomplete

func (b Builder) ClusterGetkeysinslot() (c ClusterGetkeysinslot) {
	_ = "STUB: not implemented"
	return *new(ClusterGetkeysinslot)
}

func (c ClusterGetkeysinslot) Slot(slot int64) ClusterGetkeysinslotSlot {
	_ = "STUB: not implemented"
	return *new(ClusterGetkeysinslotSlot)
}

type ClusterGetkeysinslotCount Incomplete

func (c ClusterGetkeysinslotCount) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClusterGetkeysinslotSlot Incomplete

func (c ClusterGetkeysinslotSlot) Count(count int64) ClusterGetkeysinslotCount {
	_ = "STUB: not implemented"
	return *new(ClusterGetkeysinslotCount)
}

type ClusterGetslotmigrations Incomplete

func (b Builder) ClusterGetslotmigrations() (c ClusterGetslotmigrations) {
	_ = "STUB: not implemented"
	return *new(ClusterGetslotmigrations)
}

func (c ClusterGetslotmigrations) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClusterInfo Incomplete

func (b Builder) ClusterInfo() (c ClusterInfo) { _ = "STUB: not implemented"; return *new(ClusterInfo) }

func (c ClusterInfo) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClusterKeyslot Incomplete

func (b Builder) ClusterKeyslot() (c ClusterKeyslot) {
	_ = "STUB: not implemented"
	return *new(ClusterKeyslot)
}

func (c ClusterKeyslot) Key(key string) ClusterKeyslotKey {
	_ = "STUB: not implemented"
	return *new(ClusterKeyslotKey)
}

type ClusterKeyslotKey Incomplete

func (c ClusterKeyslotKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClusterLinks Incomplete

func (b Builder) ClusterLinks() (c ClusterLinks) {
	_ = "STUB: not implemented"
	return *new(ClusterLinks)
}

func (c ClusterLinks) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClusterMeet Incomplete

func (b Builder) ClusterMeet() (c ClusterMeet) { _ = "STUB: not implemented"; return *new(ClusterMeet) }

func (c ClusterMeet) Ip(ip string) ClusterMeetIp {
	_ = "STUB: not implemented"
	return *new(ClusterMeetIp)
}

type ClusterMeetClusterBusPort Incomplete

func (c ClusterMeetClusterBusPort) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClusterMeetIp Incomplete

func (c ClusterMeetIp) Port(port int64) ClusterMeetPort {
	_ = "STUB: not implemented"
	return *new(ClusterMeetPort)
}

type ClusterMeetPort Incomplete

func (c ClusterMeetPort) ClusterBusPort(clusterBusPort int64) ClusterMeetClusterBusPort {
	_ = "STUB: not implemented"
	return *new(ClusterMeetClusterBusPort)
}

func (c ClusterMeetPort) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClusterMigrateslots Incomplete

func (b Builder) ClusterMigrateslots() (c ClusterMigrateslots) {
	_ = "STUB: not implemented"
	return *new(ClusterMigrateslots)
}

func (c ClusterMigrateslots) Slotsrange() ClusterMigrateslotsSlotMigrationSpecSlotsrange {
	_ = "STUB: not implemented"
	return *new(ClusterMigrateslotsSlotMigrationSpecSlotsrange)
}

type ClusterMigrateslotsSlotMigrationSpecNode Incomplete

func (c ClusterMigrateslotsSlotMigrationSpecNode) Slotsrange() ClusterMigrateslotsSlotMigrationSpecSlotsrange {
	_ = "STUB: not implemented"
	return *new(ClusterMigrateslotsSlotMigrationSpecSlotsrange)
}

func (c ClusterMigrateslotsSlotMigrationSpecNode) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClusterMigrateslotsSlotMigrationSpecSlotsrange Incomplete

func (c ClusterMigrateslotsSlotMigrationSpecSlotsrange) Slotsrange(startSlot int64, endSlot int64) ClusterMigrateslotsSlotMigrationSpecSlotsrange {
	_ = "STUB: not implemented"
	return *new(ClusterMigrateslotsSlotMigrationSpecSlotsrange)
}

func (c ClusterMigrateslotsSlotMigrationSpecSlotsrange) Node(nodeId string) ClusterMigrateslotsSlotMigrationSpecNode {
	_ = "STUB: not implemented"
	return *new(ClusterMigrateslotsSlotMigrationSpecNode)
}

type ClusterMyid Incomplete

func (b Builder) ClusterMyid() (c ClusterMyid) { _ = "STUB: not implemented"; return *new(ClusterMyid) }

func (c ClusterMyid) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClusterMyshardid Incomplete

func (b Builder) ClusterMyshardid() (c ClusterMyshardid) {
	_ = "STUB: not implemented"
	return *new(ClusterMyshardid)
}

func (c ClusterMyshardid) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClusterNodes Incomplete

func (b Builder) ClusterNodes() (c ClusterNodes) {
	_ = "STUB: not implemented"
	return *new(ClusterNodes)
}

func (c ClusterNodes) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClusterReplicas Incomplete

func (b Builder) ClusterReplicas() (c ClusterReplicas) {
	_ = "STUB: not implemented"
	return *new(ClusterReplicas)
}

func (c ClusterReplicas) NodeId(nodeId string) ClusterReplicasNodeId {
	_ = "STUB: not implemented"
	return *new(ClusterReplicasNodeId)
}

type ClusterReplicasNodeId Incomplete

func (c ClusterReplicasNodeId) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClusterReplicate Incomplete

func (b Builder) ClusterReplicate() (c ClusterReplicate) {
	_ = "STUB: not implemented"
	return *new(ClusterReplicate)
}

func (c ClusterReplicate) NodeId(nodeId string) ClusterReplicateNodeId {
	_ = "STUB: not implemented"
	return *new(ClusterReplicateNodeId)
}

type ClusterReplicateNodeId Incomplete

func (c ClusterReplicateNodeId) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClusterReset Incomplete

func (b Builder) ClusterReset() (c ClusterReset) {
	_ = "STUB: not implemented"
	return *new(ClusterReset)
}

func (c ClusterReset) Hard() ClusterResetResetTypeHard {
	_ = "STUB: not implemented"
	return *new(ClusterResetResetTypeHard)
}

func (c ClusterReset) Soft() ClusterResetResetTypeSoft {
	_ = "STUB: not implemented"
	return *new(ClusterResetResetTypeSoft)
}

func (c ClusterReset) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClusterResetResetTypeHard Incomplete

func (c ClusterResetResetTypeHard) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClusterResetResetTypeSoft Incomplete

func (c ClusterResetResetTypeSoft) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClusterSaveconfig Incomplete

func (b Builder) ClusterSaveconfig() (c ClusterSaveconfig) {
	_ = "STUB: not implemented"
	return *new(ClusterSaveconfig)
}

func (c ClusterSaveconfig) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClusterSetConfigEpoch Incomplete

func (b Builder) ClusterSetConfigEpoch() (c ClusterSetConfigEpoch) {
	_ = "STUB: not implemented"
	return *new(ClusterSetConfigEpoch)
}

func (c ClusterSetConfigEpoch) ConfigEpoch(configEpoch int64) ClusterSetConfigEpochConfigEpoch {
	_ = "STUB: not implemented"
	return *new(ClusterSetConfigEpochConfigEpoch)
}

type ClusterSetConfigEpochConfigEpoch Incomplete

func (c ClusterSetConfigEpochConfigEpoch) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClusterSetslot Incomplete

func (b Builder) ClusterSetslot() (c ClusterSetslot) {
	_ = "STUB: not implemented"
	return *new(ClusterSetslot)
}

func (c ClusterSetslot) Slot(slot int64) ClusterSetslotSlot {
	_ = "STUB: not implemented"
	return *new(ClusterSetslotSlot)
}

type ClusterSetslotNodeId Incomplete

func (c ClusterSetslotNodeId) Timeout(timeout int64) ClusterSetslotTimeout {
	_ = "STUB: not implemented"
	return *new(ClusterSetslotTimeout)
}

func (c ClusterSetslotNodeId) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClusterSetslotSlot Incomplete

func (c ClusterSetslotSlot) Importing() ClusterSetslotSubcommandImporting {
	_ = "STUB: not implemented"
	return *new(ClusterSetslotSubcommandImporting)
}

func (c ClusterSetslotSlot) Migrating() ClusterSetslotSubcommandMigrating {
	_ = "STUB: not implemented"
	return *new(ClusterSetslotSubcommandMigrating)
}

func (c ClusterSetslotSlot) Stable() ClusterSetslotSubcommandStable {
	_ = "STUB: not implemented"
	return *new(ClusterSetslotSubcommandStable)
}

func (c ClusterSetslotSlot) Node() ClusterSetslotSubcommandNode {
	_ = "STUB: not implemented"
	return *new(ClusterSetslotSubcommandNode)
}

type ClusterSetslotSubcommandImporting Incomplete

func (c ClusterSetslotSubcommandImporting) NodeId(nodeId string) ClusterSetslotNodeId {
	_ = "STUB: not implemented"
	return *new(ClusterSetslotNodeId)
}

func (c ClusterSetslotSubcommandImporting) Timeout(timeout int64) ClusterSetslotTimeout {
	_ = "STUB: not implemented"
	return *new(ClusterSetslotTimeout)
}

func (c ClusterSetslotSubcommandImporting) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClusterSetslotSubcommandMigrating Incomplete

func (c ClusterSetslotSubcommandMigrating) NodeId(nodeId string) ClusterSetslotNodeId {
	_ = "STUB: not implemented"
	return *new(ClusterSetslotNodeId)
}

func (c ClusterSetslotSubcommandMigrating) Timeout(timeout int64) ClusterSetslotTimeout {
	_ = "STUB: not implemented"
	return *new(ClusterSetslotTimeout)
}

func (c ClusterSetslotSubcommandMigrating) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClusterSetslotSubcommandNode Incomplete

func (c ClusterSetslotSubcommandNode) NodeId(nodeId string) ClusterSetslotNodeId {
	_ = "STUB: not implemented"
	return *new(ClusterSetslotNodeId)
}

func (c ClusterSetslotSubcommandNode) Timeout(timeout int64) ClusterSetslotTimeout {
	_ = "STUB: not implemented"
	return *new(ClusterSetslotTimeout)
}

func (c ClusterSetslotSubcommandNode) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClusterSetslotSubcommandStable Incomplete

func (c ClusterSetslotSubcommandStable) NodeId(nodeId string) ClusterSetslotNodeId {
	_ = "STUB: not implemented"
	return *new(ClusterSetslotNodeId)
}

func (c ClusterSetslotSubcommandStable) Timeout(timeout int64) ClusterSetslotTimeout {
	_ = "STUB: not implemented"
	return *new(ClusterSetslotTimeout)
}

func (c ClusterSetslotSubcommandStable) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClusterSetslotTimeout Incomplete

func (c ClusterSetslotTimeout) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClusterShards Incomplete

func (b Builder) ClusterShards() (c ClusterShards) {
	_ = "STUB: not implemented"
	return *new(ClusterShards)
}

func (c ClusterShards) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClusterSlaves Incomplete

func (b Builder) ClusterSlaves() (c ClusterSlaves) {
	_ = "STUB: not implemented"
	return *new(ClusterSlaves)
}

func (c ClusterSlaves) NodeId(nodeId string) ClusterSlavesNodeId {
	_ = "STUB: not implemented"
	return *new(ClusterSlavesNodeId)
}

type ClusterSlavesNodeId Incomplete

func (c ClusterSlavesNodeId) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClusterSlotStats Incomplete

func (b Builder) ClusterSlotStats() (c ClusterSlotStats) {
	_ = "STUB: not implemented"
	return *new(ClusterSlotStats)
}

func (c ClusterSlotStats) Slotsrange() ClusterSlotStatsFilterSlotsrangeSlotsrange {
	_ = "STUB: not implemented"
	return *new(ClusterSlotStatsFilterSlotsrangeSlotsrange)
}

func (c ClusterSlotStats) Orderby() ClusterSlotStatsFilterOrderbyOrderby {
	_ = "STUB: not implemented"
	return *new(ClusterSlotStatsFilterOrderbyOrderby)
}

type ClusterSlotStatsFilterOrderbyLimit Incomplete

func (c ClusterSlotStatsFilterOrderbyLimit) Asc() ClusterSlotStatsFilterOrderbyOrderAsc {
	_ = "STUB: not implemented"
	return *new(ClusterSlotStatsFilterOrderbyOrderAsc)
}

func (c ClusterSlotStatsFilterOrderbyLimit) Desc() ClusterSlotStatsFilterOrderbyOrderDesc {
	_ = "STUB: not implemented"
	return *new(ClusterSlotStatsFilterOrderbyOrderDesc)
}

func (c ClusterSlotStatsFilterOrderbyLimit) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClusterSlotStatsFilterOrderbyMetric Incomplete

func (c ClusterSlotStatsFilterOrderbyMetric) Limit(limit int64) ClusterSlotStatsFilterOrderbyLimit {
	_ = "STUB: not implemented"
	return *new(ClusterSlotStatsFilterOrderbyLimit)
}

func (c ClusterSlotStatsFilterOrderbyMetric) Asc() ClusterSlotStatsFilterOrderbyOrderAsc {
	_ = "STUB: not implemented"
	return *new(ClusterSlotStatsFilterOrderbyOrderAsc)
}

func (c ClusterSlotStatsFilterOrderbyMetric) Desc() ClusterSlotStatsFilterOrderbyOrderDesc {
	_ = "STUB: not implemented"
	return *new(ClusterSlotStatsFilterOrderbyOrderDesc)
}

func (c ClusterSlotStatsFilterOrderbyMetric) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClusterSlotStatsFilterOrderbyOrderAsc Incomplete

func (c ClusterSlotStatsFilterOrderbyOrderAsc) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClusterSlotStatsFilterOrderbyOrderDesc Incomplete

func (c ClusterSlotStatsFilterOrderbyOrderDesc) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClusterSlotStatsFilterOrderbyOrderby Incomplete

func (c ClusterSlotStatsFilterOrderbyOrderby) Metric(metric string) ClusterSlotStatsFilterOrderbyMetric {
	_ = "STUB: not implemented"
	return *new(ClusterSlotStatsFilterOrderbyMetric)
}

type ClusterSlotStatsFilterSlotsrangeEndSlot Incomplete

func (c ClusterSlotStatsFilterSlotsrangeEndSlot) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClusterSlotStatsFilterSlotsrangeSlotsrange Incomplete

func (c ClusterSlotStatsFilterSlotsrangeSlotsrange) StartSlot(startSlot int64) ClusterSlotStatsFilterSlotsrangeStartSlot {
	_ = "STUB: not implemented"
	return *new(ClusterSlotStatsFilterSlotsrangeStartSlot)
}

type ClusterSlotStatsFilterSlotsrangeStartSlot Incomplete

func (c ClusterSlotStatsFilterSlotsrangeStartSlot) EndSlot(endSlot int64) ClusterSlotStatsFilterSlotsrangeEndSlot {
	_ = "STUB: not implemented"
	return *new(ClusterSlotStatsFilterSlotsrangeEndSlot)
}

type ClusterSlots Incomplete

func (b Builder) ClusterSlots() (c ClusterSlots) {
	_ = "STUB: not implemented"
	return *new(ClusterSlots)
}

func (c ClusterSlots) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Readonly Incomplete

func (b Builder) Readonly() (c Readonly) { _ = "STUB: not implemented"; return *new(Readonly) }

func (c Readonly) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Readwrite Incomplete

func (b Builder) Readwrite() (c Readwrite) { _ = "STUB: not implemented"; return *new(Readwrite) }

func (c Readwrite) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

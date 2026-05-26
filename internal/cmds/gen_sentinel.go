// Code generated DO NOT EDIT

package cmds

type SentinelFailover Incomplete

func (b Builder) SentinelFailover() (c SentinelFailover) {
	_ = "STUB: not implemented"
	return *new(SentinelFailover)
}

func (c SentinelFailover) Master(master string) SentinelFailoverMaster {
	_ = "STUB: not implemented"
	return *new(SentinelFailoverMaster)
}

type SentinelFailoverMaster Incomplete

func (c SentinelFailoverMaster) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type SentinelGetMasterAddrByName Incomplete

func (b Builder) SentinelGetMasterAddrByName() (c SentinelGetMasterAddrByName) {
	_ = "STUB: not implemented"
	return *new(SentinelGetMasterAddrByName)
}

func (c SentinelGetMasterAddrByName) Master(master string) SentinelGetMasterAddrByNameMaster {
	_ = "STUB: not implemented"
	return *new(SentinelGetMasterAddrByNameMaster)
}

type SentinelGetMasterAddrByNameMaster Incomplete

func (c SentinelGetMasterAddrByNameMaster) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type SentinelReplicas Incomplete

func (b Builder) SentinelReplicas() (c SentinelReplicas) {
	_ = "STUB: not implemented"
	return *new(SentinelReplicas)
}

func (c SentinelReplicas) Master(master string) SentinelReplicasMaster {
	_ = "STUB: not implemented"
	return *new(SentinelReplicasMaster)
}

type SentinelReplicasMaster Incomplete

func (c SentinelReplicasMaster) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type SentinelSentinels Incomplete

func (b Builder) SentinelSentinels() (c SentinelSentinels) {
	_ = "STUB: not implemented"
	return *new(SentinelSentinels)
}

func (c SentinelSentinels) Master(master string) SentinelSentinelsMaster {
	_ = "STUB: not implemented"
	return *new(SentinelSentinelsMaster)
}

type SentinelSentinelsMaster Incomplete

func (c SentinelSentinelsMaster) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

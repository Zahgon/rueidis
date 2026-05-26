// Code generated DO NOT EDIT

package cmds

type Auth Incomplete

func (b Builder) Auth() (c Auth) { _ = "STUB: not implemented"; return *new(Auth) }

func (c Auth) Username(username string) AuthUsername {
	_ = "STUB: not implemented"
	return *new(AuthUsername)
}

func (c Auth) Password(password string) AuthPassword {
	_ = "STUB: not implemented"
	return *new(AuthPassword)
}

type AuthPassword Incomplete

func (c AuthPassword) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type AuthUsername Incomplete

func (c AuthUsername) Password(password string) AuthPassword {
	_ = "STUB: not implemented"
	return *new(AuthPassword)
}

type ClientCaching Incomplete

func (b Builder) ClientCaching() (c ClientCaching) {
	_ = "STUB: not implemented"
	return *new(ClientCaching)
}

func (c ClientCaching) Yes() ClientCachingModeYes {
	_ = "STUB: not implemented"
	return *new(ClientCachingModeYes)
}

func (c ClientCaching) No() ClientCachingModeNo {
	_ = "STUB: not implemented"
	return *new(ClientCachingModeNo)
}

type ClientCachingModeNo Incomplete

func (c ClientCachingModeNo) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientCachingModeYes Incomplete

func (c ClientCachingModeYes) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientCapa Incomplete

func (b Builder) ClientCapa() (c ClientCapa) { _ = "STUB: not implemented"; return *new(ClientCapa) }

func (c ClientCapa) Capability(capability ...string) ClientCapaCapability {
	_ = "STUB: not implemented"
	return *new(ClientCapaCapability)
}

type ClientCapaCapability Incomplete

func (c ClientCapaCapability) Capability(capability ...string) ClientCapaCapability {
	_ = "STUB: not implemented"
	return *new(ClientCapaCapability)
}

func (c ClientCapaCapability) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientGetname Incomplete

func (b Builder) ClientGetname() (c ClientGetname) {
	_ = "STUB: not implemented"
	return *new(ClientGetname)
}

func (c ClientGetname) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientGetredir Incomplete

func (b Builder) ClientGetredir() (c ClientGetredir) {
	_ = "STUB: not implemented"
	return *new(ClientGetredir)
}

func (c ClientGetredir) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientId Incomplete

func (b Builder) ClientId() (c ClientId) { _ = "STUB: not implemented"; return *new(ClientId) }

func (c ClientId) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientInfo Incomplete

func (b Builder) ClientInfo() (c ClientInfo) { _ = "STUB: not implemented"; return *new(ClientInfo) }

func (c ClientInfo) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKill Incomplete

func (b Builder) ClientKill() (c ClientKill) { _ = "STUB: not implemented"; return *new(ClientKill) }

func (c ClientKill) IpPort(ipPort string) ClientKillIpPort {
	_ = "STUB: not implemented"
	return *new(ClientKillIpPort)
}

func (c ClientKill) Id(clientId int64) ClientKillId {
	_ = "STUB: not implemented"
	return *new(ClientKillId)
}

func (c ClientKill) TypeNormal() ClientKillTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientKillTypeNormal)
}

func (c ClientKill) TypeMaster() ClientKillTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientKillTypeMaster)
}

func (c ClientKill) TypePrimary() ClientKillTypePrimary {
	_ = "STUB: not implemented"
	return *new(ClientKillTypePrimary)
}

func (c ClientKill) TypeSlave() ClientKillTypeSlave {
	_ = "STUB: not implemented"
	return *new(ClientKillTypeSlave)
}

func (c ClientKill) TypeReplica() ClientKillTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientKillTypeReplica)
}

func (c ClientKill) TypePubsub() ClientKillTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientKillTypePubsub)
}

func (c ClientKill) User(username string) ClientKillUser {
	_ = "STUB: not implemented"
	return *new(ClientKillUser)
}

func (c ClientKill) Addr(ipPort string) ClientKillAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillAddr)
}

func (c ClientKill) Laddr(ipPort string) ClientKillLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillLaddr)
}

func (c ClientKill) SkipmeYes() ClientKillSkipmeYes {
	_ = "STUB: not implemented"
	return *new(ClientKillSkipmeYes)
}

func (c ClientKill) SkipmeNo() ClientKillSkipmeNo {
	_ = "STUB: not implemented"
	return *new(ClientKillSkipmeNo)
}

func (c ClientKill) Maxage(maxage int64) ClientKillMaxage {
	_ = "STUB: not implemented"
	return *new(ClientKillMaxage)
}

func (c ClientKill) Name(name string) ClientKillName {
	_ = "STUB: not implemented"
	return *new(ClientKillName)
}

func (c ClientKill) Idle(idle int64) ClientKillIdle {
	_ = "STUB: not implemented"
	return *new(ClientKillIdle)
}

func (c ClientKill) Flags(flags string) ClientKillFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillFlags)
}

func (c ClientKill) LibName(libName string) ClientKillLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillLibName)
}

func (c ClientKill) LibVer(libVer string) ClientKillLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillLibVer)
}

func (c ClientKill) Db(db int64) ClientKillDb { _ = "STUB: not implemented"; return *new(ClientKillDb) }

func (c ClientKill) Capa(capa string) ClientKillCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillCapa)
}

func (c ClientKill) Ip(ip string) ClientKillIp {
	_ = "STUB: not implemented"
	return *new(ClientKillIp)
}

func (c ClientKill) NotTypeNormal() ClientKillNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeNormal)
}

func (c ClientKill) NotTypeMaster() ClientKillNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeMaster)
}

func (c ClientKill) NotTypePrimary() ClientKillNotTypePrimary {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePrimary)
}

func (c ClientKill) NotTypeSlave() ClientKillNotTypeSlave {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeSlave)
}

func (c ClientKill) NotTypeReplica() ClientKillNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeReplica)
}

func (c ClientKill) NotTypePubsub() ClientKillNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePubsub)
}

func (c ClientKill) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKill) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKill) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKill) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKill) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKill) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKill) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKill) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKill) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKill) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKill) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKill) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillAddr Incomplete

func (c ClientKillAddr) Laddr(ipPort string) ClientKillLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillLaddr)
}

func (c ClientKillAddr) SkipmeYes() ClientKillSkipmeYes {
	_ = "STUB: not implemented"
	return *new(ClientKillSkipmeYes)
}

func (c ClientKillAddr) SkipmeNo() ClientKillSkipmeNo {
	_ = "STUB: not implemented"
	return *new(ClientKillSkipmeNo)
}

func (c ClientKillAddr) Maxage(maxage int64) ClientKillMaxage {
	_ = "STUB: not implemented"
	return *new(ClientKillMaxage)
}

func (c ClientKillAddr) Name(name string) ClientKillName {
	_ = "STUB: not implemented"
	return *new(ClientKillName)
}

func (c ClientKillAddr) Idle(idle int64) ClientKillIdle {
	_ = "STUB: not implemented"
	return *new(ClientKillIdle)
}

func (c ClientKillAddr) Flags(flags string) ClientKillFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillFlags)
}

func (c ClientKillAddr) LibName(libName string) ClientKillLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillLibName)
}

func (c ClientKillAddr) LibVer(libVer string) ClientKillLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillLibVer)
}

func (c ClientKillAddr) Db(db int64) ClientKillDb {
	_ = "STUB: not implemented"
	return *new(ClientKillDb)
}

func (c ClientKillAddr) Capa(capa string) ClientKillCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillCapa)
}

func (c ClientKillAddr) Ip(ip string) ClientKillIp {
	_ = "STUB: not implemented"
	return *new(ClientKillIp)
}

func (c ClientKillAddr) NotTypeNormal() ClientKillNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeNormal)
}

func (c ClientKillAddr) NotTypeMaster() ClientKillNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeMaster)
}

func (c ClientKillAddr) NotTypePrimary() ClientKillNotTypePrimary {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePrimary)
}

func (c ClientKillAddr) NotTypeSlave() ClientKillNotTypeSlave {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeSlave)
}

func (c ClientKillAddr) NotTypeReplica() ClientKillNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeReplica)
}

func (c ClientKillAddr) NotTypePubsub() ClientKillNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePubsub)
}

func (c ClientKillAddr) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillAddr) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillAddr) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillAddr) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillAddr) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillAddr) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillAddr) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillAddr) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillAddr) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillAddr) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillAddr) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillAddr) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillCapa Incomplete

func (c ClientKillCapa) Ip(ip string) ClientKillIp {
	_ = "STUB: not implemented"
	return *new(ClientKillIp)
}

func (c ClientKillCapa) NotTypeNormal() ClientKillNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeNormal)
}

func (c ClientKillCapa) NotTypeMaster() ClientKillNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeMaster)
}

func (c ClientKillCapa) NotTypePrimary() ClientKillNotTypePrimary {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePrimary)
}

func (c ClientKillCapa) NotTypeSlave() ClientKillNotTypeSlave {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeSlave)
}

func (c ClientKillCapa) NotTypeReplica() ClientKillNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeReplica)
}

func (c ClientKillCapa) NotTypePubsub() ClientKillNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePubsub)
}

func (c ClientKillCapa) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillCapa) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillCapa) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillCapa) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillCapa) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillCapa) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillCapa) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillCapa) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillCapa) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillCapa) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillCapa) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillCapa) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillDb Incomplete

func (c ClientKillDb) Capa(capa string) ClientKillCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillCapa)
}

func (c ClientKillDb) Ip(ip string) ClientKillIp {
	_ = "STUB: not implemented"
	return *new(ClientKillIp)
}

func (c ClientKillDb) NotTypeNormal() ClientKillNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeNormal)
}

func (c ClientKillDb) NotTypeMaster() ClientKillNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeMaster)
}

func (c ClientKillDb) NotTypePrimary() ClientKillNotTypePrimary {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePrimary)
}

func (c ClientKillDb) NotTypeSlave() ClientKillNotTypeSlave {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeSlave)
}

func (c ClientKillDb) NotTypeReplica() ClientKillNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeReplica)
}

func (c ClientKillDb) NotTypePubsub() ClientKillNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePubsub)
}

func (c ClientKillDb) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillDb) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillDb) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillDb) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillDb) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillDb) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillDb) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillDb) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillDb) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillDb) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillDb) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillDb) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillFlags Incomplete

func (c ClientKillFlags) LibName(libName string) ClientKillLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillLibName)
}

func (c ClientKillFlags) LibVer(libVer string) ClientKillLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillLibVer)
}

func (c ClientKillFlags) Db(db int64) ClientKillDb {
	_ = "STUB: not implemented"
	return *new(ClientKillDb)
}

func (c ClientKillFlags) Capa(capa string) ClientKillCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillCapa)
}

func (c ClientKillFlags) Ip(ip string) ClientKillIp {
	_ = "STUB: not implemented"
	return *new(ClientKillIp)
}

func (c ClientKillFlags) NotTypeNormal() ClientKillNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeNormal)
}

func (c ClientKillFlags) NotTypeMaster() ClientKillNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeMaster)
}

func (c ClientKillFlags) NotTypePrimary() ClientKillNotTypePrimary {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePrimary)
}

func (c ClientKillFlags) NotTypeSlave() ClientKillNotTypeSlave {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeSlave)
}

func (c ClientKillFlags) NotTypeReplica() ClientKillNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeReplica)
}

func (c ClientKillFlags) NotTypePubsub() ClientKillNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePubsub)
}

func (c ClientKillFlags) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillFlags) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillFlags) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillFlags) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillFlags) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillFlags) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillFlags) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillFlags) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillFlags) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillFlags) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillFlags) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillFlags) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillId Incomplete

func (c ClientKillId) TypeNormal() ClientKillTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientKillTypeNormal)
}

func (c ClientKillId) TypeMaster() ClientKillTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientKillTypeMaster)
}

func (c ClientKillId) TypePrimary() ClientKillTypePrimary {
	_ = "STUB: not implemented"
	return *new(ClientKillTypePrimary)
}

func (c ClientKillId) TypeSlave() ClientKillTypeSlave {
	_ = "STUB: not implemented"
	return *new(ClientKillTypeSlave)
}

func (c ClientKillId) TypeReplica() ClientKillTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientKillTypeReplica)
}

func (c ClientKillId) TypePubsub() ClientKillTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientKillTypePubsub)
}

func (c ClientKillId) User(username string) ClientKillUser {
	_ = "STUB: not implemented"
	return *new(ClientKillUser)
}

func (c ClientKillId) Addr(ipPort string) ClientKillAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillAddr)
}

func (c ClientKillId) Laddr(ipPort string) ClientKillLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillLaddr)
}

func (c ClientKillId) SkipmeYes() ClientKillSkipmeYes {
	_ = "STUB: not implemented"
	return *new(ClientKillSkipmeYes)
}

func (c ClientKillId) SkipmeNo() ClientKillSkipmeNo {
	_ = "STUB: not implemented"
	return *new(ClientKillSkipmeNo)
}

func (c ClientKillId) Maxage(maxage int64) ClientKillMaxage {
	_ = "STUB: not implemented"
	return *new(ClientKillMaxage)
}

func (c ClientKillId) Name(name string) ClientKillName {
	_ = "STUB: not implemented"
	return *new(ClientKillName)
}

func (c ClientKillId) Idle(idle int64) ClientKillIdle {
	_ = "STUB: not implemented"
	return *new(ClientKillIdle)
}

func (c ClientKillId) Flags(flags string) ClientKillFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillFlags)
}

func (c ClientKillId) LibName(libName string) ClientKillLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillLibName)
}

func (c ClientKillId) LibVer(libVer string) ClientKillLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillLibVer)
}

func (c ClientKillId) Db(db int64) ClientKillDb {
	_ = "STUB: not implemented"
	return *new(ClientKillDb)
}

func (c ClientKillId) Capa(capa string) ClientKillCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillCapa)
}

func (c ClientKillId) Ip(ip string) ClientKillIp {
	_ = "STUB: not implemented"
	return *new(ClientKillIp)
}

func (c ClientKillId) NotTypeNormal() ClientKillNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeNormal)
}

func (c ClientKillId) NotTypeMaster() ClientKillNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeMaster)
}

func (c ClientKillId) NotTypePrimary() ClientKillNotTypePrimary {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePrimary)
}

func (c ClientKillId) NotTypeSlave() ClientKillNotTypeSlave {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeSlave)
}

func (c ClientKillId) NotTypeReplica() ClientKillNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeReplica)
}

func (c ClientKillId) NotTypePubsub() ClientKillNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePubsub)
}

func (c ClientKillId) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillId) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillId) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillId) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillId) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillId) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillId) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillId) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillId) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillId) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillId) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillId) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillIdle Incomplete

func (c ClientKillIdle) Flags(flags string) ClientKillFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillFlags)
}

func (c ClientKillIdle) LibName(libName string) ClientKillLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillLibName)
}

func (c ClientKillIdle) LibVer(libVer string) ClientKillLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillLibVer)
}

func (c ClientKillIdle) Db(db int64) ClientKillDb {
	_ = "STUB: not implemented"
	return *new(ClientKillDb)
}

func (c ClientKillIdle) Capa(capa string) ClientKillCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillCapa)
}

func (c ClientKillIdle) Ip(ip string) ClientKillIp {
	_ = "STUB: not implemented"
	return *new(ClientKillIp)
}

func (c ClientKillIdle) NotTypeNormal() ClientKillNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeNormal)
}

func (c ClientKillIdle) NotTypeMaster() ClientKillNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeMaster)
}

func (c ClientKillIdle) NotTypePrimary() ClientKillNotTypePrimary {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePrimary)
}

func (c ClientKillIdle) NotTypeSlave() ClientKillNotTypeSlave {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeSlave)
}

func (c ClientKillIdle) NotTypeReplica() ClientKillNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeReplica)
}

func (c ClientKillIdle) NotTypePubsub() ClientKillNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePubsub)
}

func (c ClientKillIdle) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillIdle) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillIdle) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillIdle) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillIdle) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillIdle) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillIdle) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillIdle) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillIdle) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillIdle) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillIdle) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillIdle) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillIp Incomplete

func (c ClientKillIp) NotTypeNormal() ClientKillNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeNormal)
}

func (c ClientKillIp) NotTypeMaster() ClientKillNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeMaster)
}

func (c ClientKillIp) NotTypePrimary() ClientKillNotTypePrimary {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePrimary)
}

func (c ClientKillIp) NotTypeSlave() ClientKillNotTypeSlave {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeSlave)
}

func (c ClientKillIp) NotTypeReplica() ClientKillNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeReplica)
}

func (c ClientKillIp) NotTypePubsub() ClientKillNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePubsub)
}

func (c ClientKillIp) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillIp) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillIp) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillIp) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillIp) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillIp) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillIp) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillIp) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillIp) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillIp) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillIp) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillIp) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillIpPort Incomplete

func (c ClientKillIpPort) Id(clientId int64) ClientKillId {
	_ = "STUB: not implemented"
	return *new(ClientKillId)
}

func (c ClientKillIpPort) TypeNormal() ClientKillTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientKillTypeNormal)
}

func (c ClientKillIpPort) TypeMaster() ClientKillTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientKillTypeMaster)
}

func (c ClientKillIpPort) TypePrimary() ClientKillTypePrimary {
	_ = "STUB: not implemented"
	return *new(ClientKillTypePrimary)
}

func (c ClientKillIpPort) TypeSlave() ClientKillTypeSlave {
	_ = "STUB: not implemented"
	return *new(ClientKillTypeSlave)
}

func (c ClientKillIpPort) TypeReplica() ClientKillTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientKillTypeReplica)
}

func (c ClientKillIpPort) TypePubsub() ClientKillTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientKillTypePubsub)
}

func (c ClientKillIpPort) User(username string) ClientKillUser {
	_ = "STUB: not implemented"
	return *new(ClientKillUser)
}

func (c ClientKillIpPort) Addr(ipPort string) ClientKillAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillAddr)
}

func (c ClientKillIpPort) Laddr(ipPort string) ClientKillLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillLaddr)
}

func (c ClientKillIpPort) SkipmeYes() ClientKillSkipmeYes {
	_ = "STUB: not implemented"
	return *new(ClientKillSkipmeYes)
}

func (c ClientKillIpPort) SkipmeNo() ClientKillSkipmeNo {
	_ = "STUB: not implemented"
	return *new(ClientKillSkipmeNo)
}

func (c ClientKillIpPort) Maxage(maxage int64) ClientKillMaxage {
	_ = "STUB: not implemented"
	return *new(ClientKillMaxage)
}

func (c ClientKillIpPort) Name(name string) ClientKillName {
	_ = "STUB: not implemented"
	return *new(ClientKillName)
}

func (c ClientKillIpPort) Idle(idle int64) ClientKillIdle {
	_ = "STUB: not implemented"
	return *new(ClientKillIdle)
}

func (c ClientKillIpPort) Flags(flags string) ClientKillFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillFlags)
}

func (c ClientKillIpPort) LibName(libName string) ClientKillLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillLibName)
}

func (c ClientKillIpPort) LibVer(libVer string) ClientKillLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillLibVer)
}

func (c ClientKillIpPort) Db(db int64) ClientKillDb {
	_ = "STUB: not implemented"
	return *new(ClientKillDb)
}

func (c ClientKillIpPort) Capa(capa string) ClientKillCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillCapa)
}

func (c ClientKillIpPort) Ip(ip string) ClientKillIp {
	_ = "STUB: not implemented"
	return *new(ClientKillIp)
}

func (c ClientKillIpPort) NotTypeNormal() ClientKillNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeNormal)
}

func (c ClientKillIpPort) NotTypeMaster() ClientKillNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeMaster)
}

func (c ClientKillIpPort) NotTypePrimary() ClientKillNotTypePrimary {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePrimary)
}

func (c ClientKillIpPort) NotTypeSlave() ClientKillNotTypeSlave {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeSlave)
}

func (c ClientKillIpPort) NotTypeReplica() ClientKillNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeReplica)
}

func (c ClientKillIpPort) NotTypePubsub() ClientKillNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePubsub)
}

func (c ClientKillIpPort) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillIpPort) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillIpPort) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillIpPort) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillIpPort) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillIpPort) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillIpPort) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillIpPort) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillIpPort) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillIpPort) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillIpPort) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillIpPort) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillLaddr Incomplete

func (c ClientKillLaddr) SkipmeYes() ClientKillSkipmeYes {
	_ = "STUB: not implemented"
	return *new(ClientKillSkipmeYes)
}

func (c ClientKillLaddr) SkipmeNo() ClientKillSkipmeNo {
	_ = "STUB: not implemented"
	return *new(ClientKillSkipmeNo)
}

func (c ClientKillLaddr) Maxage(maxage int64) ClientKillMaxage {
	_ = "STUB: not implemented"
	return *new(ClientKillMaxage)
}

func (c ClientKillLaddr) Name(name string) ClientKillName {
	_ = "STUB: not implemented"
	return *new(ClientKillName)
}

func (c ClientKillLaddr) Idle(idle int64) ClientKillIdle {
	_ = "STUB: not implemented"
	return *new(ClientKillIdle)
}

func (c ClientKillLaddr) Flags(flags string) ClientKillFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillFlags)
}

func (c ClientKillLaddr) LibName(libName string) ClientKillLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillLibName)
}

func (c ClientKillLaddr) LibVer(libVer string) ClientKillLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillLibVer)
}

func (c ClientKillLaddr) Db(db int64) ClientKillDb {
	_ = "STUB: not implemented"
	return *new(ClientKillDb)
}

func (c ClientKillLaddr) Capa(capa string) ClientKillCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillCapa)
}

func (c ClientKillLaddr) Ip(ip string) ClientKillIp {
	_ = "STUB: not implemented"
	return *new(ClientKillIp)
}

func (c ClientKillLaddr) NotTypeNormal() ClientKillNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeNormal)
}

func (c ClientKillLaddr) NotTypeMaster() ClientKillNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeMaster)
}

func (c ClientKillLaddr) NotTypePrimary() ClientKillNotTypePrimary {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePrimary)
}

func (c ClientKillLaddr) NotTypeSlave() ClientKillNotTypeSlave {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeSlave)
}

func (c ClientKillLaddr) NotTypeReplica() ClientKillNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeReplica)
}

func (c ClientKillLaddr) NotTypePubsub() ClientKillNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePubsub)
}

func (c ClientKillLaddr) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillLaddr) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillLaddr) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillLaddr) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillLaddr) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillLaddr) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillLaddr) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillLaddr) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillLaddr) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillLaddr) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillLaddr) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillLaddr) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillLibName Incomplete

func (c ClientKillLibName) LibVer(libVer string) ClientKillLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillLibVer)
}

func (c ClientKillLibName) Db(db int64) ClientKillDb {
	_ = "STUB: not implemented"
	return *new(ClientKillDb)
}

func (c ClientKillLibName) Capa(capa string) ClientKillCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillCapa)
}

func (c ClientKillLibName) Ip(ip string) ClientKillIp {
	_ = "STUB: not implemented"
	return *new(ClientKillIp)
}

func (c ClientKillLibName) NotTypeNormal() ClientKillNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeNormal)
}

func (c ClientKillLibName) NotTypeMaster() ClientKillNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeMaster)
}

func (c ClientKillLibName) NotTypePrimary() ClientKillNotTypePrimary {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePrimary)
}

func (c ClientKillLibName) NotTypeSlave() ClientKillNotTypeSlave {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeSlave)
}

func (c ClientKillLibName) NotTypeReplica() ClientKillNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeReplica)
}

func (c ClientKillLibName) NotTypePubsub() ClientKillNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePubsub)
}

func (c ClientKillLibName) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillLibName) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillLibName) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillLibName) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillLibName) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillLibName) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillLibName) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillLibName) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillLibName) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillLibName) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillLibName) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillLibName) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillLibVer Incomplete

func (c ClientKillLibVer) Db(db int64) ClientKillDb {
	_ = "STUB: not implemented"
	return *new(ClientKillDb)
}

func (c ClientKillLibVer) Capa(capa string) ClientKillCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillCapa)
}

func (c ClientKillLibVer) Ip(ip string) ClientKillIp {
	_ = "STUB: not implemented"
	return *new(ClientKillIp)
}

func (c ClientKillLibVer) NotTypeNormal() ClientKillNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeNormal)
}

func (c ClientKillLibVer) NotTypeMaster() ClientKillNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeMaster)
}

func (c ClientKillLibVer) NotTypePrimary() ClientKillNotTypePrimary {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePrimary)
}

func (c ClientKillLibVer) NotTypeSlave() ClientKillNotTypeSlave {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeSlave)
}

func (c ClientKillLibVer) NotTypeReplica() ClientKillNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeReplica)
}

func (c ClientKillLibVer) NotTypePubsub() ClientKillNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePubsub)
}

func (c ClientKillLibVer) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillLibVer) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillLibVer) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillLibVer) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillLibVer) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillLibVer) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillLibVer) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillLibVer) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillLibVer) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillLibVer) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillLibVer) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillLibVer) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillMaxage Incomplete

func (c ClientKillMaxage) Name(name string) ClientKillName {
	_ = "STUB: not implemented"
	return *new(ClientKillName)
}

func (c ClientKillMaxage) Idle(idle int64) ClientKillIdle {
	_ = "STUB: not implemented"
	return *new(ClientKillIdle)
}

func (c ClientKillMaxage) Flags(flags string) ClientKillFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillFlags)
}

func (c ClientKillMaxage) LibName(libName string) ClientKillLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillLibName)
}

func (c ClientKillMaxage) LibVer(libVer string) ClientKillLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillLibVer)
}

func (c ClientKillMaxage) Db(db int64) ClientKillDb {
	_ = "STUB: not implemented"
	return *new(ClientKillDb)
}

func (c ClientKillMaxage) Capa(capa string) ClientKillCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillCapa)
}

func (c ClientKillMaxage) Ip(ip string) ClientKillIp {
	_ = "STUB: not implemented"
	return *new(ClientKillIp)
}

func (c ClientKillMaxage) NotTypeNormal() ClientKillNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeNormal)
}

func (c ClientKillMaxage) NotTypeMaster() ClientKillNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeMaster)
}

func (c ClientKillMaxage) NotTypePrimary() ClientKillNotTypePrimary {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePrimary)
}

func (c ClientKillMaxage) NotTypeSlave() ClientKillNotTypeSlave {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeSlave)
}

func (c ClientKillMaxage) NotTypeReplica() ClientKillNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeReplica)
}

func (c ClientKillMaxage) NotTypePubsub() ClientKillNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePubsub)
}

func (c ClientKillMaxage) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillMaxage) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillMaxage) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillMaxage) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillMaxage) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillMaxage) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillMaxage) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillMaxage) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillMaxage) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillMaxage) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillMaxage) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillMaxage) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillName Incomplete

func (c ClientKillName) Idle(idle int64) ClientKillIdle {
	_ = "STUB: not implemented"
	return *new(ClientKillIdle)
}

func (c ClientKillName) Flags(flags string) ClientKillFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillFlags)
}

func (c ClientKillName) LibName(libName string) ClientKillLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillLibName)
}

func (c ClientKillName) LibVer(libVer string) ClientKillLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillLibVer)
}

func (c ClientKillName) Db(db int64) ClientKillDb {
	_ = "STUB: not implemented"
	return *new(ClientKillDb)
}

func (c ClientKillName) Capa(capa string) ClientKillCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillCapa)
}

func (c ClientKillName) Ip(ip string) ClientKillIp {
	_ = "STUB: not implemented"
	return *new(ClientKillIp)
}

func (c ClientKillName) NotTypeNormal() ClientKillNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeNormal)
}

func (c ClientKillName) NotTypeMaster() ClientKillNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeMaster)
}

func (c ClientKillName) NotTypePrimary() ClientKillNotTypePrimary {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePrimary)
}

func (c ClientKillName) NotTypeSlave() ClientKillNotTypeSlave {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeSlave)
}

func (c ClientKillName) NotTypeReplica() ClientKillNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeReplica)
}

func (c ClientKillName) NotTypePubsub() ClientKillNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePubsub)
}

func (c ClientKillName) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillName) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillName) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillName) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillName) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillName) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillName) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillName) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillName) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillName) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillName) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillName) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillNotAddr Incomplete

func (c ClientKillNotAddr) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillNotAddr) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillNotAddr) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillNotAddr) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillNotAddr) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillNotAddr) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillNotAddr) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillNotAddr) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillNotAddr) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillNotCapa Incomplete

func (c ClientKillNotCapa) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillNotCapa) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillNotDb Incomplete

func (c ClientKillNotDb) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillNotDb) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillNotDb) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillNotFlags Incomplete

func (c ClientKillNotFlags) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillNotFlags) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillNotFlags) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillNotFlags) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillNotFlags) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillNotFlags) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillNotIdNotClientId Incomplete

func (c ClientKillNotIdNotClientId) NotClientId(notClientId ...int64) ClientKillNotIdNotClientId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotClientId)
}

func (c ClientKillNotIdNotClientId) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillNotIdNotClientId) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillNotIdNotClientId) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillNotIdNotClientId) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillNotIdNotClientId) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillNotIdNotClientId) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillNotIdNotClientId) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillNotIdNotClientId) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillNotIdNotClientId) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillNotIdNotClientId) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillNotIdNotClientId) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClientKillNotIdNotId Incomplete

func (c ClientKillNotIdNotId) NotClientId(notClientId ...int64) ClientKillNotIdNotClientId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotClientId)
}

type ClientKillNotIp Incomplete

func (c ClientKillNotIp) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillNotLaddr Incomplete

func (c ClientKillNotLaddr) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillNotLaddr) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillNotLaddr) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillNotLaddr) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillNotLaddr) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillNotLaddr) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillNotLaddr) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillNotLaddr) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillNotLibName Incomplete

func (c ClientKillNotLibName) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillNotLibName) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillNotLibName) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillNotLibName) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillNotLibName) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillNotLibVer Incomplete

func (c ClientKillNotLibVer) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillNotLibVer) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillNotLibVer) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillNotLibVer) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillNotName Incomplete

func (c ClientKillNotName) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillNotName) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillNotName) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillNotName) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillNotName) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillNotName) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillNotName) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillNotTypeMaster Incomplete

func (c ClientKillNotTypeMaster) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillNotTypeMaster) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillNotTypeMaster) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillNotTypeMaster) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillNotTypeMaster) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillNotTypeMaster) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillNotTypeMaster) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillNotTypeMaster) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillNotTypeMaster) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillNotTypeMaster) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillNotTypeMaster) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillNotTypeMaster) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClientKillNotTypeNormal Incomplete

func (c ClientKillNotTypeNormal) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillNotTypeNormal) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillNotTypeNormal) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillNotTypeNormal) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillNotTypeNormal) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillNotTypeNormal) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillNotTypeNormal) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillNotTypeNormal) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillNotTypeNormal) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillNotTypeNormal) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillNotTypeNormal) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillNotTypeNormal) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClientKillNotTypePrimary Incomplete

func (c ClientKillNotTypePrimary) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillNotTypePrimary) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillNotTypePrimary) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillNotTypePrimary) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillNotTypePrimary) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillNotTypePrimary) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillNotTypePrimary) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillNotTypePrimary) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillNotTypePrimary) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillNotTypePrimary) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillNotTypePrimary) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillNotTypePrimary) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClientKillNotTypePubsub Incomplete

func (c ClientKillNotTypePubsub) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillNotTypePubsub) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillNotTypePubsub) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillNotTypePubsub) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillNotTypePubsub) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillNotTypePubsub) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillNotTypePubsub) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillNotTypePubsub) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillNotTypePubsub) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillNotTypePubsub) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillNotTypePubsub) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillNotTypePubsub) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClientKillNotTypeReplica Incomplete

func (c ClientKillNotTypeReplica) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillNotTypeReplica) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillNotTypeReplica) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillNotTypeReplica) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillNotTypeReplica) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillNotTypeReplica) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillNotTypeReplica) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillNotTypeReplica) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillNotTypeReplica) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillNotTypeReplica) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillNotTypeReplica) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillNotTypeReplica) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClientKillNotTypeSlave Incomplete

func (c ClientKillNotTypeSlave) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillNotTypeSlave) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillNotTypeSlave) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillNotTypeSlave) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillNotTypeSlave) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillNotTypeSlave) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillNotTypeSlave) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillNotTypeSlave) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillNotTypeSlave) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillNotTypeSlave) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillNotTypeSlave) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillNotTypeSlave) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClientKillNotUser Incomplete

func (c ClientKillNotUser) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillNotUser) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillNotUser) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillNotUser) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillNotUser) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillNotUser) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillNotUser) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillNotUser) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillNotUser) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillNotUser) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillSkipmeNo Incomplete

func (c ClientKillSkipmeNo) Maxage(maxage int64) ClientKillMaxage {
	_ = "STUB: not implemented"
	return *new(ClientKillMaxage)
}

func (c ClientKillSkipmeNo) Name(name string) ClientKillName {
	_ = "STUB: not implemented"
	return *new(ClientKillName)
}

func (c ClientKillSkipmeNo) Idle(idle int64) ClientKillIdle {
	_ = "STUB: not implemented"
	return *new(ClientKillIdle)
}

func (c ClientKillSkipmeNo) Flags(flags string) ClientKillFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillFlags)
}

func (c ClientKillSkipmeNo) LibName(libName string) ClientKillLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillLibName)
}

func (c ClientKillSkipmeNo) LibVer(libVer string) ClientKillLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillLibVer)
}

func (c ClientKillSkipmeNo) Db(db int64) ClientKillDb {
	_ = "STUB: not implemented"
	return *new(ClientKillDb)
}

func (c ClientKillSkipmeNo) Capa(capa string) ClientKillCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillCapa)
}

func (c ClientKillSkipmeNo) Ip(ip string) ClientKillIp {
	_ = "STUB: not implemented"
	return *new(ClientKillIp)
}

func (c ClientKillSkipmeNo) NotTypeNormal() ClientKillNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeNormal)
}

func (c ClientKillSkipmeNo) NotTypeMaster() ClientKillNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeMaster)
}

func (c ClientKillSkipmeNo) NotTypePrimary() ClientKillNotTypePrimary {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePrimary)
}

func (c ClientKillSkipmeNo) NotTypeSlave() ClientKillNotTypeSlave {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeSlave)
}

func (c ClientKillSkipmeNo) NotTypeReplica() ClientKillNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeReplica)
}

func (c ClientKillSkipmeNo) NotTypePubsub() ClientKillNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePubsub)
}

func (c ClientKillSkipmeNo) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillSkipmeNo) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillSkipmeNo) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillSkipmeNo) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillSkipmeNo) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillSkipmeNo) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillSkipmeNo) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillSkipmeNo) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillSkipmeNo) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillSkipmeNo) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillSkipmeNo) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillSkipmeNo) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillSkipmeYes Incomplete

func (c ClientKillSkipmeYes) Maxage(maxage int64) ClientKillMaxage {
	_ = "STUB: not implemented"
	return *new(ClientKillMaxage)
}

func (c ClientKillSkipmeYes) Name(name string) ClientKillName {
	_ = "STUB: not implemented"
	return *new(ClientKillName)
}

func (c ClientKillSkipmeYes) Idle(idle int64) ClientKillIdle {
	_ = "STUB: not implemented"
	return *new(ClientKillIdle)
}

func (c ClientKillSkipmeYes) Flags(flags string) ClientKillFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillFlags)
}

func (c ClientKillSkipmeYes) LibName(libName string) ClientKillLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillLibName)
}

func (c ClientKillSkipmeYes) LibVer(libVer string) ClientKillLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillLibVer)
}

func (c ClientKillSkipmeYes) Db(db int64) ClientKillDb {
	_ = "STUB: not implemented"
	return *new(ClientKillDb)
}

func (c ClientKillSkipmeYes) Capa(capa string) ClientKillCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillCapa)
}

func (c ClientKillSkipmeYes) Ip(ip string) ClientKillIp {
	_ = "STUB: not implemented"
	return *new(ClientKillIp)
}

func (c ClientKillSkipmeYes) NotTypeNormal() ClientKillNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeNormal)
}

func (c ClientKillSkipmeYes) NotTypeMaster() ClientKillNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeMaster)
}

func (c ClientKillSkipmeYes) NotTypePrimary() ClientKillNotTypePrimary {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePrimary)
}

func (c ClientKillSkipmeYes) NotTypeSlave() ClientKillNotTypeSlave {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeSlave)
}

func (c ClientKillSkipmeYes) NotTypeReplica() ClientKillNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeReplica)
}

func (c ClientKillSkipmeYes) NotTypePubsub() ClientKillNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePubsub)
}

func (c ClientKillSkipmeYes) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillSkipmeYes) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillSkipmeYes) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillSkipmeYes) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillSkipmeYes) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillSkipmeYes) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillSkipmeYes) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillSkipmeYes) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillSkipmeYes) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillSkipmeYes) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillSkipmeYes) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillSkipmeYes) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillTypeMaster Incomplete

func (c ClientKillTypeMaster) User(username string) ClientKillUser {
	_ = "STUB: not implemented"
	return *new(ClientKillUser)
}

func (c ClientKillTypeMaster) Addr(ipPort string) ClientKillAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillAddr)
}

func (c ClientKillTypeMaster) Laddr(ipPort string) ClientKillLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillLaddr)
}

func (c ClientKillTypeMaster) SkipmeYes() ClientKillSkipmeYes {
	_ = "STUB: not implemented"
	return *new(ClientKillSkipmeYes)
}

func (c ClientKillTypeMaster) SkipmeNo() ClientKillSkipmeNo {
	_ = "STUB: not implemented"
	return *new(ClientKillSkipmeNo)
}

func (c ClientKillTypeMaster) Maxage(maxage int64) ClientKillMaxage {
	_ = "STUB: not implemented"
	return *new(ClientKillMaxage)
}

func (c ClientKillTypeMaster) Name(name string) ClientKillName {
	_ = "STUB: not implemented"
	return *new(ClientKillName)
}

func (c ClientKillTypeMaster) Idle(idle int64) ClientKillIdle {
	_ = "STUB: not implemented"
	return *new(ClientKillIdle)
}

func (c ClientKillTypeMaster) Flags(flags string) ClientKillFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillFlags)
}

func (c ClientKillTypeMaster) LibName(libName string) ClientKillLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillLibName)
}

func (c ClientKillTypeMaster) LibVer(libVer string) ClientKillLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillLibVer)
}

func (c ClientKillTypeMaster) Db(db int64) ClientKillDb {
	_ = "STUB: not implemented"
	return *new(ClientKillDb)
}

func (c ClientKillTypeMaster) Capa(capa string) ClientKillCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillCapa)
}

func (c ClientKillTypeMaster) Ip(ip string) ClientKillIp {
	_ = "STUB: not implemented"
	return *new(ClientKillIp)
}

func (c ClientKillTypeMaster) NotTypeNormal() ClientKillNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeNormal)
}

func (c ClientKillTypeMaster) NotTypeMaster() ClientKillNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeMaster)
}

func (c ClientKillTypeMaster) NotTypePrimary() ClientKillNotTypePrimary {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePrimary)
}

func (c ClientKillTypeMaster) NotTypeSlave() ClientKillNotTypeSlave {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeSlave)
}

func (c ClientKillTypeMaster) NotTypeReplica() ClientKillNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeReplica)
}

func (c ClientKillTypeMaster) NotTypePubsub() ClientKillNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePubsub)
}

func (c ClientKillTypeMaster) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillTypeMaster) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillTypeMaster) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillTypeMaster) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillTypeMaster) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillTypeMaster) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillTypeMaster) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillTypeMaster) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillTypeMaster) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillTypeMaster) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillTypeMaster) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillTypeMaster) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillTypeNormal Incomplete

func (c ClientKillTypeNormal) User(username string) ClientKillUser {
	_ = "STUB: not implemented"
	return *new(ClientKillUser)
}

func (c ClientKillTypeNormal) Addr(ipPort string) ClientKillAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillAddr)
}

func (c ClientKillTypeNormal) Laddr(ipPort string) ClientKillLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillLaddr)
}

func (c ClientKillTypeNormal) SkipmeYes() ClientKillSkipmeYes {
	_ = "STUB: not implemented"
	return *new(ClientKillSkipmeYes)
}

func (c ClientKillTypeNormal) SkipmeNo() ClientKillSkipmeNo {
	_ = "STUB: not implemented"
	return *new(ClientKillSkipmeNo)
}

func (c ClientKillTypeNormal) Maxage(maxage int64) ClientKillMaxage {
	_ = "STUB: not implemented"
	return *new(ClientKillMaxage)
}

func (c ClientKillTypeNormal) Name(name string) ClientKillName {
	_ = "STUB: not implemented"
	return *new(ClientKillName)
}

func (c ClientKillTypeNormal) Idle(idle int64) ClientKillIdle {
	_ = "STUB: not implemented"
	return *new(ClientKillIdle)
}

func (c ClientKillTypeNormal) Flags(flags string) ClientKillFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillFlags)
}

func (c ClientKillTypeNormal) LibName(libName string) ClientKillLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillLibName)
}

func (c ClientKillTypeNormal) LibVer(libVer string) ClientKillLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillLibVer)
}

func (c ClientKillTypeNormal) Db(db int64) ClientKillDb {
	_ = "STUB: not implemented"
	return *new(ClientKillDb)
}

func (c ClientKillTypeNormal) Capa(capa string) ClientKillCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillCapa)
}

func (c ClientKillTypeNormal) Ip(ip string) ClientKillIp {
	_ = "STUB: not implemented"
	return *new(ClientKillIp)
}

func (c ClientKillTypeNormal) NotTypeNormal() ClientKillNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeNormal)
}

func (c ClientKillTypeNormal) NotTypeMaster() ClientKillNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeMaster)
}

func (c ClientKillTypeNormal) NotTypePrimary() ClientKillNotTypePrimary {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePrimary)
}

func (c ClientKillTypeNormal) NotTypeSlave() ClientKillNotTypeSlave {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeSlave)
}

func (c ClientKillTypeNormal) NotTypeReplica() ClientKillNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeReplica)
}

func (c ClientKillTypeNormal) NotTypePubsub() ClientKillNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePubsub)
}

func (c ClientKillTypeNormal) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillTypeNormal) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillTypeNormal) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillTypeNormal) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillTypeNormal) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillTypeNormal) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillTypeNormal) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillTypeNormal) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillTypeNormal) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillTypeNormal) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillTypeNormal) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillTypeNormal) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillTypePrimary Incomplete

func (c ClientKillTypePrimary) User(username string) ClientKillUser {
	_ = "STUB: not implemented"
	return *new(ClientKillUser)
}

func (c ClientKillTypePrimary) Addr(ipPort string) ClientKillAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillAddr)
}

func (c ClientKillTypePrimary) Laddr(ipPort string) ClientKillLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillLaddr)
}

func (c ClientKillTypePrimary) SkipmeYes() ClientKillSkipmeYes {
	_ = "STUB: not implemented"
	return *new(ClientKillSkipmeYes)
}

func (c ClientKillTypePrimary) SkipmeNo() ClientKillSkipmeNo {
	_ = "STUB: not implemented"
	return *new(ClientKillSkipmeNo)
}

func (c ClientKillTypePrimary) Maxage(maxage int64) ClientKillMaxage {
	_ = "STUB: not implemented"
	return *new(ClientKillMaxage)
}

func (c ClientKillTypePrimary) Name(name string) ClientKillName {
	_ = "STUB: not implemented"
	return *new(ClientKillName)
}

func (c ClientKillTypePrimary) Idle(idle int64) ClientKillIdle {
	_ = "STUB: not implemented"
	return *new(ClientKillIdle)
}

func (c ClientKillTypePrimary) Flags(flags string) ClientKillFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillFlags)
}

func (c ClientKillTypePrimary) LibName(libName string) ClientKillLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillLibName)
}

func (c ClientKillTypePrimary) LibVer(libVer string) ClientKillLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillLibVer)
}

func (c ClientKillTypePrimary) Db(db int64) ClientKillDb {
	_ = "STUB: not implemented"
	return *new(ClientKillDb)
}

func (c ClientKillTypePrimary) Capa(capa string) ClientKillCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillCapa)
}

func (c ClientKillTypePrimary) Ip(ip string) ClientKillIp {
	_ = "STUB: not implemented"
	return *new(ClientKillIp)
}

func (c ClientKillTypePrimary) NotTypeNormal() ClientKillNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeNormal)
}

func (c ClientKillTypePrimary) NotTypeMaster() ClientKillNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeMaster)
}

func (c ClientKillTypePrimary) NotTypePrimary() ClientKillNotTypePrimary {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePrimary)
}

func (c ClientKillTypePrimary) NotTypeSlave() ClientKillNotTypeSlave {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeSlave)
}

func (c ClientKillTypePrimary) NotTypeReplica() ClientKillNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeReplica)
}

func (c ClientKillTypePrimary) NotTypePubsub() ClientKillNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePubsub)
}

func (c ClientKillTypePrimary) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillTypePrimary) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillTypePrimary) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillTypePrimary) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillTypePrimary) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillTypePrimary) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillTypePrimary) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillTypePrimary) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillTypePrimary) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillTypePrimary) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillTypePrimary) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillTypePrimary) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillTypePubsub Incomplete

func (c ClientKillTypePubsub) User(username string) ClientKillUser {
	_ = "STUB: not implemented"
	return *new(ClientKillUser)
}

func (c ClientKillTypePubsub) Addr(ipPort string) ClientKillAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillAddr)
}

func (c ClientKillTypePubsub) Laddr(ipPort string) ClientKillLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillLaddr)
}

func (c ClientKillTypePubsub) SkipmeYes() ClientKillSkipmeYes {
	_ = "STUB: not implemented"
	return *new(ClientKillSkipmeYes)
}

func (c ClientKillTypePubsub) SkipmeNo() ClientKillSkipmeNo {
	_ = "STUB: not implemented"
	return *new(ClientKillSkipmeNo)
}

func (c ClientKillTypePubsub) Maxage(maxage int64) ClientKillMaxage {
	_ = "STUB: not implemented"
	return *new(ClientKillMaxage)
}

func (c ClientKillTypePubsub) Name(name string) ClientKillName {
	_ = "STUB: not implemented"
	return *new(ClientKillName)
}

func (c ClientKillTypePubsub) Idle(idle int64) ClientKillIdle {
	_ = "STUB: not implemented"
	return *new(ClientKillIdle)
}

func (c ClientKillTypePubsub) Flags(flags string) ClientKillFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillFlags)
}

func (c ClientKillTypePubsub) LibName(libName string) ClientKillLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillLibName)
}

func (c ClientKillTypePubsub) LibVer(libVer string) ClientKillLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillLibVer)
}

func (c ClientKillTypePubsub) Db(db int64) ClientKillDb {
	_ = "STUB: not implemented"
	return *new(ClientKillDb)
}

func (c ClientKillTypePubsub) Capa(capa string) ClientKillCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillCapa)
}

func (c ClientKillTypePubsub) Ip(ip string) ClientKillIp {
	_ = "STUB: not implemented"
	return *new(ClientKillIp)
}

func (c ClientKillTypePubsub) NotTypeNormal() ClientKillNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeNormal)
}

func (c ClientKillTypePubsub) NotTypeMaster() ClientKillNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeMaster)
}

func (c ClientKillTypePubsub) NotTypePrimary() ClientKillNotTypePrimary {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePrimary)
}

func (c ClientKillTypePubsub) NotTypeSlave() ClientKillNotTypeSlave {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeSlave)
}

func (c ClientKillTypePubsub) NotTypeReplica() ClientKillNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeReplica)
}

func (c ClientKillTypePubsub) NotTypePubsub() ClientKillNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePubsub)
}

func (c ClientKillTypePubsub) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillTypePubsub) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillTypePubsub) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillTypePubsub) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillTypePubsub) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillTypePubsub) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillTypePubsub) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillTypePubsub) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillTypePubsub) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillTypePubsub) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillTypePubsub) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillTypePubsub) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillTypeReplica Incomplete

func (c ClientKillTypeReplica) User(username string) ClientKillUser {
	_ = "STUB: not implemented"
	return *new(ClientKillUser)
}

func (c ClientKillTypeReplica) Addr(ipPort string) ClientKillAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillAddr)
}

func (c ClientKillTypeReplica) Laddr(ipPort string) ClientKillLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillLaddr)
}

func (c ClientKillTypeReplica) SkipmeYes() ClientKillSkipmeYes {
	_ = "STUB: not implemented"
	return *new(ClientKillSkipmeYes)
}

func (c ClientKillTypeReplica) SkipmeNo() ClientKillSkipmeNo {
	_ = "STUB: not implemented"
	return *new(ClientKillSkipmeNo)
}

func (c ClientKillTypeReplica) Maxage(maxage int64) ClientKillMaxage {
	_ = "STUB: not implemented"
	return *new(ClientKillMaxage)
}

func (c ClientKillTypeReplica) Name(name string) ClientKillName {
	_ = "STUB: not implemented"
	return *new(ClientKillName)
}

func (c ClientKillTypeReplica) Idle(idle int64) ClientKillIdle {
	_ = "STUB: not implemented"
	return *new(ClientKillIdle)
}

func (c ClientKillTypeReplica) Flags(flags string) ClientKillFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillFlags)
}

func (c ClientKillTypeReplica) LibName(libName string) ClientKillLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillLibName)
}

func (c ClientKillTypeReplica) LibVer(libVer string) ClientKillLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillLibVer)
}

func (c ClientKillTypeReplica) Db(db int64) ClientKillDb {
	_ = "STUB: not implemented"
	return *new(ClientKillDb)
}

func (c ClientKillTypeReplica) Capa(capa string) ClientKillCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillCapa)
}

func (c ClientKillTypeReplica) Ip(ip string) ClientKillIp {
	_ = "STUB: not implemented"
	return *new(ClientKillIp)
}

func (c ClientKillTypeReplica) NotTypeNormal() ClientKillNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeNormal)
}

func (c ClientKillTypeReplica) NotTypeMaster() ClientKillNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeMaster)
}

func (c ClientKillTypeReplica) NotTypePrimary() ClientKillNotTypePrimary {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePrimary)
}

func (c ClientKillTypeReplica) NotTypeSlave() ClientKillNotTypeSlave {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeSlave)
}

func (c ClientKillTypeReplica) NotTypeReplica() ClientKillNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeReplica)
}

func (c ClientKillTypeReplica) NotTypePubsub() ClientKillNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePubsub)
}

func (c ClientKillTypeReplica) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillTypeReplica) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillTypeReplica) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillTypeReplica) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillTypeReplica) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillTypeReplica) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillTypeReplica) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillTypeReplica) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillTypeReplica) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillTypeReplica) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillTypeReplica) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillTypeReplica) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillTypeSlave Incomplete

func (c ClientKillTypeSlave) User(username string) ClientKillUser {
	_ = "STUB: not implemented"
	return *new(ClientKillUser)
}

func (c ClientKillTypeSlave) Addr(ipPort string) ClientKillAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillAddr)
}

func (c ClientKillTypeSlave) Laddr(ipPort string) ClientKillLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillLaddr)
}

func (c ClientKillTypeSlave) SkipmeYes() ClientKillSkipmeYes {
	_ = "STUB: not implemented"
	return *new(ClientKillSkipmeYes)
}

func (c ClientKillTypeSlave) SkipmeNo() ClientKillSkipmeNo {
	_ = "STUB: not implemented"
	return *new(ClientKillSkipmeNo)
}

func (c ClientKillTypeSlave) Maxage(maxage int64) ClientKillMaxage {
	_ = "STUB: not implemented"
	return *new(ClientKillMaxage)
}

func (c ClientKillTypeSlave) Name(name string) ClientKillName {
	_ = "STUB: not implemented"
	return *new(ClientKillName)
}

func (c ClientKillTypeSlave) Idle(idle int64) ClientKillIdle {
	_ = "STUB: not implemented"
	return *new(ClientKillIdle)
}

func (c ClientKillTypeSlave) Flags(flags string) ClientKillFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillFlags)
}

func (c ClientKillTypeSlave) LibName(libName string) ClientKillLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillLibName)
}

func (c ClientKillTypeSlave) LibVer(libVer string) ClientKillLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillLibVer)
}

func (c ClientKillTypeSlave) Db(db int64) ClientKillDb {
	_ = "STUB: not implemented"
	return *new(ClientKillDb)
}

func (c ClientKillTypeSlave) Capa(capa string) ClientKillCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillCapa)
}

func (c ClientKillTypeSlave) Ip(ip string) ClientKillIp {
	_ = "STUB: not implemented"
	return *new(ClientKillIp)
}

func (c ClientKillTypeSlave) NotTypeNormal() ClientKillNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeNormal)
}

func (c ClientKillTypeSlave) NotTypeMaster() ClientKillNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeMaster)
}

func (c ClientKillTypeSlave) NotTypePrimary() ClientKillNotTypePrimary {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePrimary)
}

func (c ClientKillTypeSlave) NotTypeSlave() ClientKillNotTypeSlave {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeSlave)
}

func (c ClientKillTypeSlave) NotTypeReplica() ClientKillNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeReplica)
}

func (c ClientKillTypeSlave) NotTypePubsub() ClientKillNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePubsub)
}

func (c ClientKillTypeSlave) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillTypeSlave) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillTypeSlave) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillTypeSlave) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillTypeSlave) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillTypeSlave) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillTypeSlave) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillTypeSlave) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillTypeSlave) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillTypeSlave) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillTypeSlave) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillTypeSlave) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientKillUser Incomplete

func (c ClientKillUser) Addr(ipPort string) ClientKillAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillAddr)
}

func (c ClientKillUser) Laddr(ipPort string) ClientKillLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillLaddr)
}

func (c ClientKillUser) SkipmeYes() ClientKillSkipmeYes {
	_ = "STUB: not implemented"
	return *new(ClientKillSkipmeYes)
}

func (c ClientKillUser) SkipmeNo() ClientKillSkipmeNo {
	_ = "STUB: not implemented"
	return *new(ClientKillSkipmeNo)
}

func (c ClientKillUser) Maxage(maxage int64) ClientKillMaxage {
	_ = "STUB: not implemented"
	return *new(ClientKillMaxage)
}

func (c ClientKillUser) Name(name string) ClientKillName {
	_ = "STUB: not implemented"
	return *new(ClientKillName)
}

func (c ClientKillUser) Idle(idle int64) ClientKillIdle {
	_ = "STUB: not implemented"
	return *new(ClientKillIdle)
}

func (c ClientKillUser) Flags(flags string) ClientKillFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillFlags)
}

func (c ClientKillUser) LibName(libName string) ClientKillLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillLibName)
}

func (c ClientKillUser) LibVer(libVer string) ClientKillLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillLibVer)
}

func (c ClientKillUser) Db(db int64) ClientKillDb {
	_ = "STUB: not implemented"
	return *new(ClientKillDb)
}

func (c ClientKillUser) Capa(capa string) ClientKillCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillCapa)
}

func (c ClientKillUser) Ip(ip string) ClientKillIp {
	_ = "STUB: not implemented"
	return *new(ClientKillIp)
}

func (c ClientKillUser) NotTypeNormal() ClientKillNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeNormal)
}

func (c ClientKillUser) NotTypeMaster() ClientKillNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeMaster)
}

func (c ClientKillUser) NotTypePrimary() ClientKillNotTypePrimary {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePrimary)
}

func (c ClientKillUser) NotTypeSlave() ClientKillNotTypeSlave {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeSlave)
}

func (c ClientKillUser) NotTypeReplica() ClientKillNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypeReplica)
}

func (c ClientKillUser) NotTypePubsub() ClientKillNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientKillNotTypePubsub)
}

func (c ClientKillUser) NotId() ClientKillNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIdNotId)
}

func (c ClientKillUser) NotUser(notUsername string) ClientKillNotUser {
	_ = "STUB: not implemented"
	return *new(ClientKillNotUser)
}

func (c ClientKillUser) NotAddr(notAddr string) ClientKillNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotAddr)
}

func (c ClientKillUser) NotLaddr(notLaddr string) ClientKillNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLaddr)
}

func (c ClientKillUser) NotName(notName string) ClientKillNotName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotName)
}

func (c ClientKillUser) NotFlags(notFlags string) ClientKillNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientKillNotFlags)
}

func (c ClientKillUser) NotLibName(notLibName string) ClientKillNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibName)
}

func (c ClientKillUser) NotLibVer(notLibVer string) ClientKillNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientKillNotLibVer)
}

func (c ClientKillUser) NotDb(notDb int64) ClientKillNotDb {
	_ = "STUB: not implemented"
	return *new(ClientKillNotDb)
}

func (c ClientKillUser) NotCapa(notCapa string) ClientKillNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientKillNotCapa)
}

func (c ClientKillUser) NotIp(notIp string) ClientKillNotIp {
	_ = "STUB: not implemented"
	return *new(ClientKillNotIp)
}

func (c ClientKillUser) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientList Incomplete

func (b Builder) ClientList() (c ClientList) { _ = "STUB: not implemented"; return *new(ClientList) }

func (c ClientList) TypeNormal() ClientListTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientListTypeNormal)
}

func (c ClientList) TypeMaster() ClientListTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientListTypeMaster)
}

func (c ClientList) TypeReplica() ClientListTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientListTypeReplica)
}

func (c ClientList) TypePubsub() ClientListTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientListTypePubsub)
}

func (c ClientList) Id() ClientListIdId { _ = "STUB: not implemented"; return *new(ClientListIdId) }

func (c ClientList) User(username string) ClientListUser {
	_ = "STUB: not implemented"
	return *new(ClientListUser)
}

func (c ClientList) Addr(ipPort string) ClientListAddr {
	_ = "STUB: not implemented"
	return *new(ClientListAddr)
}

func (c ClientList) Laddr(ipPort string) ClientListLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListLaddr)
}

func (c ClientList) SkipmeYes() ClientListSkipmeYes {
	_ = "STUB: not implemented"
	return *new(ClientListSkipmeYes)
}

func (c ClientList) SkipmeNo() ClientListSkipmeNo {
	_ = "STUB: not implemented"
	return *new(ClientListSkipmeNo)
}

func (c ClientList) Maxage(maxage int64) ClientListMaxage {
	_ = "STUB: not implemented"
	return *new(ClientListMaxage)
}

func (c ClientList) Name(name string) ClientListName {
	_ = "STUB: not implemented"
	return *new(ClientListName)
}

func (c ClientList) Idle(idle int64) ClientListIdle {
	_ = "STUB: not implemented"
	return *new(ClientListIdle)
}

func (c ClientList) Flags(flags string) ClientListFlags {
	_ = "STUB: not implemented"
	return *new(ClientListFlags)
}

func (c ClientList) LibName(libName string) ClientListLibName {
	_ = "STUB: not implemented"
	return *new(ClientListLibName)
}

func (c ClientList) LibVer(libVer string) ClientListLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListLibVer)
}

func (c ClientList) Db(db int64) ClientListDb { _ = "STUB: not implemented"; return *new(ClientListDb) }

func (c ClientList) Capa(capa string) ClientListCapa {
	_ = "STUB: not implemented"
	return *new(ClientListCapa)
}

func (c ClientList) Ip(ip string) ClientListIp {
	_ = "STUB: not implemented"
	return *new(ClientListIp)
}

func (c ClientList) NotTypeNormal() ClientListNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeNormal)
}

func (c ClientList) NotTypeMaster() ClientListNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeMaster)
}

func (c ClientList) NotTypeReplica() ClientListNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeReplica)
}

func (c ClientList) NotTypePubsub() ClientListNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypePubsub)
}

func (c ClientList) NotId() ClientListNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientListNotIdNotId)
}

func (c ClientList) NotUser(notUsername string) ClientListNotUser {
	_ = "STUB: not implemented"
	return *new(ClientListNotUser)
}

func (c ClientList) NotAddr(notAddr string) ClientListNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotAddr)
}

func (c ClientList) NotLaddr(notLaddr string) ClientListNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotLaddr)
}

func (c ClientList) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientList) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientList) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientList) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientList) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientList) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientList) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientList) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListAddr Incomplete

func (c ClientListAddr) Laddr(ipPort string) ClientListLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListLaddr)
}

func (c ClientListAddr) SkipmeYes() ClientListSkipmeYes {
	_ = "STUB: not implemented"
	return *new(ClientListSkipmeYes)
}

func (c ClientListAddr) SkipmeNo() ClientListSkipmeNo {
	_ = "STUB: not implemented"
	return *new(ClientListSkipmeNo)
}

func (c ClientListAddr) Maxage(maxage int64) ClientListMaxage {
	_ = "STUB: not implemented"
	return *new(ClientListMaxage)
}

func (c ClientListAddr) Name(name string) ClientListName {
	_ = "STUB: not implemented"
	return *new(ClientListName)
}

func (c ClientListAddr) Idle(idle int64) ClientListIdle {
	_ = "STUB: not implemented"
	return *new(ClientListIdle)
}

func (c ClientListAddr) Flags(flags string) ClientListFlags {
	_ = "STUB: not implemented"
	return *new(ClientListFlags)
}

func (c ClientListAddr) LibName(libName string) ClientListLibName {
	_ = "STUB: not implemented"
	return *new(ClientListLibName)
}

func (c ClientListAddr) LibVer(libVer string) ClientListLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListLibVer)
}

func (c ClientListAddr) Db(db int64) ClientListDb {
	_ = "STUB: not implemented"
	return *new(ClientListDb)
}

func (c ClientListAddr) Capa(capa string) ClientListCapa {
	_ = "STUB: not implemented"
	return *new(ClientListCapa)
}

func (c ClientListAddr) Ip(ip string) ClientListIp {
	_ = "STUB: not implemented"
	return *new(ClientListIp)
}

func (c ClientListAddr) NotTypeNormal() ClientListNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeNormal)
}

func (c ClientListAddr) NotTypeMaster() ClientListNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeMaster)
}

func (c ClientListAddr) NotTypeReplica() ClientListNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeReplica)
}

func (c ClientListAddr) NotTypePubsub() ClientListNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypePubsub)
}

func (c ClientListAddr) NotId() ClientListNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientListNotIdNotId)
}

func (c ClientListAddr) NotUser(notUsername string) ClientListNotUser {
	_ = "STUB: not implemented"
	return *new(ClientListNotUser)
}

func (c ClientListAddr) NotAddr(notAddr string) ClientListNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotAddr)
}

func (c ClientListAddr) NotLaddr(notLaddr string) ClientListNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotLaddr)
}

func (c ClientListAddr) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientListAddr) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListAddr) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListAddr) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListAddr) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListAddr) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListAddr) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListAddr) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListCapa Incomplete

func (c ClientListCapa) Ip(ip string) ClientListIp {
	_ = "STUB: not implemented"
	return *new(ClientListIp)
}

func (c ClientListCapa) NotTypeNormal() ClientListNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeNormal)
}

func (c ClientListCapa) NotTypeMaster() ClientListNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeMaster)
}

func (c ClientListCapa) NotTypeReplica() ClientListNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeReplica)
}

func (c ClientListCapa) NotTypePubsub() ClientListNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypePubsub)
}

func (c ClientListCapa) NotId() ClientListNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientListNotIdNotId)
}

func (c ClientListCapa) NotUser(notUsername string) ClientListNotUser {
	_ = "STUB: not implemented"
	return *new(ClientListNotUser)
}

func (c ClientListCapa) NotAddr(notAddr string) ClientListNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotAddr)
}

func (c ClientListCapa) NotLaddr(notLaddr string) ClientListNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotLaddr)
}

func (c ClientListCapa) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientListCapa) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListCapa) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListCapa) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListCapa) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListCapa) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListCapa) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListCapa) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListDb Incomplete

func (c ClientListDb) Capa(capa string) ClientListCapa {
	_ = "STUB: not implemented"
	return *new(ClientListCapa)
}

func (c ClientListDb) Ip(ip string) ClientListIp {
	_ = "STUB: not implemented"
	return *new(ClientListIp)
}

func (c ClientListDb) NotTypeNormal() ClientListNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeNormal)
}

func (c ClientListDb) NotTypeMaster() ClientListNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeMaster)
}

func (c ClientListDb) NotTypeReplica() ClientListNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeReplica)
}

func (c ClientListDb) NotTypePubsub() ClientListNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypePubsub)
}

func (c ClientListDb) NotId() ClientListNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientListNotIdNotId)
}

func (c ClientListDb) NotUser(notUsername string) ClientListNotUser {
	_ = "STUB: not implemented"
	return *new(ClientListNotUser)
}

func (c ClientListDb) NotAddr(notAddr string) ClientListNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotAddr)
}

func (c ClientListDb) NotLaddr(notLaddr string) ClientListNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotLaddr)
}

func (c ClientListDb) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientListDb) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListDb) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListDb) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListDb) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListDb) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListDb) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListDb) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListFlags Incomplete

func (c ClientListFlags) LibName(libName string) ClientListLibName {
	_ = "STUB: not implemented"
	return *new(ClientListLibName)
}

func (c ClientListFlags) LibVer(libVer string) ClientListLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListLibVer)
}

func (c ClientListFlags) Db(db int64) ClientListDb {
	_ = "STUB: not implemented"
	return *new(ClientListDb)
}

func (c ClientListFlags) Capa(capa string) ClientListCapa {
	_ = "STUB: not implemented"
	return *new(ClientListCapa)
}

func (c ClientListFlags) Ip(ip string) ClientListIp {
	_ = "STUB: not implemented"
	return *new(ClientListIp)
}

func (c ClientListFlags) NotTypeNormal() ClientListNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeNormal)
}

func (c ClientListFlags) NotTypeMaster() ClientListNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeMaster)
}

func (c ClientListFlags) NotTypeReplica() ClientListNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeReplica)
}

func (c ClientListFlags) NotTypePubsub() ClientListNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypePubsub)
}

func (c ClientListFlags) NotId() ClientListNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientListNotIdNotId)
}

func (c ClientListFlags) NotUser(notUsername string) ClientListNotUser {
	_ = "STUB: not implemented"
	return *new(ClientListNotUser)
}

func (c ClientListFlags) NotAddr(notAddr string) ClientListNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotAddr)
}

func (c ClientListFlags) NotLaddr(notLaddr string) ClientListNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotLaddr)
}

func (c ClientListFlags) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientListFlags) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListFlags) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListFlags) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListFlags) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListFlags) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListFlags) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListFlags) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListIdClientId Incomplete

func (c ClientListIdClientId) ClientId(clientId ...int64) ClientListIdClientId {
	_ = "STUB: not implemented"
	return *new(ClientListIdClientId)
}

func (c ClientListIdClientId) User(username string) ClientListUser {
	_ = "STUB: not implemented"
	return *new(ClientListUser)
}

func (c ClientListIdClientId) Addr(ipPort string) ClientListAddr {
	_ = "STUB: not implemented"
	return *new(ClientListAddr)
}

func (c ClientListIdClientId) Laddr(ipPort string) ClientListLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListLaddr)
}

func (c ClientListIdClientId) SkipmeYes() ClientListSkipmeYes {
	_ = "STUB: not implemented"
	return *new(ClientListSkipmeYes)
}

func (c ClientListIdClientId) SkipmeNo() ClientListSkipmeNo {
	_ = "STUB: not implemented"
	return *new(ClientListSkipmeNo)
}

func (c ClientListIdClientId) Maxage(maxage int64) ClientListMaxage {
	_ = "STUB: not implemented"
	return *new(ClientListMaxage)
}

func (c ClientListIdClientId) Name(name string) ClientListName {
	_ = "STUB: not implemented"
	return *new(ClientListName)
}

func (c ClientListIdClientId) Idle(idle int64) ClientListIdle {
	_ = "STUB: not implemented"
	return *new(ClientListIdle)
}

func (c ClientListIdClientId) Flags(flags string) ClientListFlags {
	_ = "STUB: not implemented"
	return *new(ClientListFlags)
}

func (c ClientListIdClientId) LibName(libName string) ClientListLibName {
	_ = "STUB: not implemented"
	return *new(ClientListLibName)
}

func (c ClientListIdClientId) LibVer(libVer string) ClientListLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListLibVer)
}

func (c ClientListIdClientId) Db(db int64) ClientListDb {
	_ = "STUB: not implemented"
	return *new(ClientListDb)
}

func (c ClientListIdClientId) Capa(capa string) ClientListCapa {
	_ = "STUB: not implemented"
	return *new(ClientListCapa)
}

func (c ClientListIdClientId) Ip(ip string) ClientListIp {
	_ = "STUB: not implemented"
	return *new(ClientListIp)
}

func (c ClientListIdClientId) NotTypeNormal() ClientListNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeNormal)
}

func (c ClientListIdClientId) NotTypeMaster() ClientListNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeMaster)
}

func (c ClientListIdClientId) NotTypeReplica() ClientListNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeReplica)
}

func (c ClientListIdClientId) NotTypePubsub() ClientListNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypePubsub)
}

func (c ClientListIdClientId) NotId() ClientListNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientListNotIdNotId)
}

func (c ClientListIdClientId) NotUser(notUsername string) ClientListNotUser {
	_ = "STUB: not implemented"
	return *new(ClientListNotUser)
}

func (c ClientListIdClientId) NotAddr(notAddr string) ClientListNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotAddr)
}

func (c ClientListIdClientId) NotLaddr(notLaddr string) ClientListNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotLaddr)
}

func (c ClientListIdClientId) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientListIdClientId) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListIdClientId) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListIdClientId) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListIdClientId) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListIdClientId) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListIdClientId) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListIdClientId) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListIdId Incomplete

func (c ClientListIdId) ClientId(clientId ...int64) ClientListIdClientId {
	_ = "STUB: not implemented"
	return *new(ClientListIdClientId)
}

type ClientListIdle Incomplete

func (c ClientListIdle) Flags(flags string) ClientListFlags {
	_ = "STUB: not implemented"
	return *new(ClientListFlags)
}

func (c ClientListIdle) LibName(libName string) ClientListLibName {
	_ = "STUB: not implemented"
	return *new(ClientListLibName)
}

func (c ClientListIdle) LibVer(libVer string) ClientListLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListLibVer)
}

func (c ClientListIdle) Db(db int64) ClientListDb {
	_ = "STUB: not implemented"
	return *new(ClientListDb)
}

func (c ClientListIdle) Capa(capa string) ClientListCapa {
	_ = "STUB: not implemented"
	return *new(ClientListCapa)
}

func (c ClientListIdle) Ip(ip string) ClientListIp {
	_ = "STUB: not implemented"
	return *new(ClientListIp)
}

func (c ClientListIdle) NotTypeNormal() ClientListNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeNormal)
}

func (c ClientListIdle) NotTypeMaster() ClientListNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeMaster)
}

func (c ClientListIdle) NotTypeReplica() ClientListNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeReplica)
}

func (c ClientListIdle) NotTypePubsub() ClientListNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypePubsub)
}

func (c ClientListIdle) NotId() ClientListNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientListNotIdNotId)
}

func (c ClientListIdle) NotUser(notUsername string) ClientListNotUser {
	_ = "STUB: not implemented"
	return *new(ClientListNotUser)
}

func (c ClientListIdle) NotAddr(notAddr string) ClientListNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotAddr)
}

func (c ClientListIdle) NotLaddr(notLaddr string) ClientListNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotLaddr)
}

func (c ClientListIdle) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientListIdle) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListIdle) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListIdle) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListIdle) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListIdle) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListIdle) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListIdle) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListIp Incomplete

func (c ClientListIp) NotTypeNormal() ClientListNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeNormal)
}

func (c ClientListIp) NotTypeMaster() ClientListNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeMaster)
}

func (c ClientListIp) NotTypeReplica() ClientListNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeReplica)
}

func (c ClientListIp) NotTypePubsub() ClientListNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypePubsub)
}

func (c ClientListIp) NotId() ClientListNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientListNotIdNotId)
}

func (c ClientListIp) NotUser(notUsername string) ClientListNotUser {
	_ = "STUB: not implemented"
	return *new(ClientListNotUser)
}

func (c ClientListIp) NotAddr(notAddr string) ClientListNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotAddr)
}

func (c ClientListIp) NotLaddr(notLaddr string) ClientListNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotLaddr)
}

func (c ClientListIp) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientListIp) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListIp) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListIp) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListIp) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListIp) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListIp) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListIp) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListLaddr Incomplete

func (c ClientListLaddr) SkipmeYes() ClientListSkipmeYes {
	_ = "STUB: not implemented"
	return *new(ClientListSkipmeYes)
}

func (c ClientListLaddr) SkipmeNo() ClientListSkipmeNo {
	_ = "STUB: not implemented"
	return *new(ClientListSkipmeNo)
}

func (c ClientListLaddr) Maxage(maxage int64) ClientListMaxage {
	_ = "STUB: not implemented"
	return *new(ClientListMaxage)
}

func (c ClientListLaddr) Name(name string) ClientListName {
	_ = "STUB: not implemented"
	return *new(ClientListName)
}

func (c ClientListLaddr) Idle(idle int64) ClientListIdle {
	_ = "STUB: not implemented"
	return *new(ClientListIdle)
}

func (c ClientListLaddr) Flags(flags string) ClientListFlags {
	_ = "STUB: not implemented"
	return *new(ClientListFlags)
}

func (c ClientListLaddr) LibName(libName string) ClientListLibName {
	_ = "STUB: not implemented"
	return *new(ClientListLibName)
}

func (c ClientListLaddr) LibVer(libVer string) ClientListLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListLibVer)
}

func (c ClientListLaddr) Db(db int64) ClientListDb {
	_ = "STUB: not implemented"
	return *new(ClientListDb)
}

func (c ClientListLaddr) Capa(capa string) ClientListCapa {
	_ = "STUB: not implemented"
	return *new(ClientListCapa)
}

func (c ClientListLaddr) Ip(ip string) ClientListIp {
	_ = "STUB: not implemented"
	return *new(ClientListIp)
}

func (c ClientListLaddr) NotTypeNormal() ClientListNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeNormal)
}

func (c ClientListLaddr) NotTypeMaster() ClientListNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeMaster)
}

func (c ClientListLaddr) NotTypeReplica() ClientListNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeReplica)
}

func (c ClientListLaddr) NotTypePubsub() ClientListNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypePubsub)
}

func (c ClientListLaddr) NotId() ClientListNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientListNotIdNotId)
}

func (c ClientListLaddr) NotUser(notUsername string) ClientListNotUser {
	_ = "STUB: not implemented"
	return *new(ClientListNotUser)
}

func (c ClientListLaddr) NotAddr(notAddr string) ClientListNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotAddr)
}

func (c ClientListLaddr) NotLaddr(notLaddr string) ClientListNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotLaddr)
}

func (c ClientListLaddr) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientListLaddr) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListLaddr) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListLaddr) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListLaddr) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListLaddr) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListLaddr) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListLaddr) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListLibName Incomplete

func (c ClientListLibName) LibVer(libVer string) ClientListLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListLibVer)
}

func (c ClientListLibName) Db(db int64) ClientListDb {
	_ = "STUB: not implemented"
	return *new(ClientListDb)
}

func (c ClientListLibName) Capa(capa string) ClientListCapa {
	_ = "STUB: not implemented"
	return *new(ClientListCapa)
}

func (c ClientListLibName) Ip(ip string) ClientListIp {
	_ = "STUB: not implemented"
	return *new(ClientListIp)
}

func (c ClientListLibName) NotTypeNormal() ClientListNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeNormal)
}

func (c ClientListLibName) NotTypeMaster() ClientListNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeMaster)
}

func (c ClientListLibName) NotTypeReplica() ClientListNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeReplica)
}

func (c ClientListLibName) NotTypePubsub() ClientListNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypePubsub)
}

func (c ClientListLibName) NotId() ClientListNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientListNotIdNotId)
}

func (c ClientListLibName) NotUser(notUsername string) ClientListNotUser {
	_ = "STUB: not implemented"
	return *new(ClientListNotUser)
}

func (c ClientListLibName) NotAddr(notAddr string) ClientListNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotAddr)
}

func (c ClientListLibName) NotLaddr(notLaddr string) ClientListNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotLaddr)
}

func (c ClientListLibName) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientListLibName) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListLibName) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListLibName) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListLibName) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListLibName) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListLibName) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListLibName) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListLibVer Incomplete

func (c ClientListLibVer) Db(db int64) ClientListDb {
	_ = "STUB: not implemented"
	return *new(ClientListDb)
}

func (c ClientListLibVer) Capa(capa string) ClientListCapa {
	_ = "STUB: not implemented"
	return *new(ClientListCapa)
}

func (c ClientListLibVer) Ip(ip string) ClientListIp {
	_ = "STUB: not implemented"
	return *new(ClientListIp)
}

func (c ClientListLibVer) NotTypeNormal() ClientListNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeNormal)
}

func (c ClientListLibVer) NotTypeMaster() ClientListNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeMaster)
}

func (c ClientListLibVer) NotTypeReplica() ClientListNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeReplica)
}

func (c ClientListLibVer) NotTypePubsub() ClientListNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypePubsub)
}

func (c ClientListLibVer) NotId() ClientListNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientListNotIdNotId)
}

func (c ClientListLibVer) NotUser(notUsername string) ClientListNotUser {
	_ = "STUB: not implemented"
	return *new(ClientListNotUser)
}

func (c ClientListLibVer) NotAddr(notAddr string) ClientListNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotAddr)
}

func (c ClientListLibVer) NotLaddr(notLaddr string) ClientListNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotLaddr)
}

func (c ClientListLibVer) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientListLibVer) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListLibVer) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListLibVer) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListLibVer) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListLibVer) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListLibVer) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListLibVer) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListMaxage Incomplete

func (c ClientListMaxage) Name(name string) ClientListName {
	_ = "STUB: not implemented"
	return *new(ClientListName)
}

func (c ClientListMaxage) Idle(idle int64) ClientListIdle {
	_ = "STUB: not implemented"
	return *new(ClientListIdle)
}

func (c ClientListMaxage) Flags(flags string) ClientListFlags {
	_ = "STUB: not implemented"
	return *new(ClientListFlags)
}

func (c ClientListMaxage) LibName(libName string) ClientListLibName {
	_ = "STUB: not implemented"
	return *new(ClientListLibName)
}

func (c ClientListMaxage) LibVer(libVer string) ClientListLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListLibVer)
}

func (c ClientListMaxage) Db(db int64) ClientListDb {
	_ = "STUB: not implemented"
	return *new(ClientListDb)
}

func (c ClientListMaxage) Capa(capa string) ClientListCapa {
	_ = "STUB: not implemented"
	return *new(ClientListCapa)
}

func (c ClientListMaxage) Ip(ip string) ClientListIp {
	_ = "STUB: not implemented"
	return *new(ClientListIp)
}

func (c ClientListMaxage) NotTypeNormal() ClientListNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeNormal)
}

func (c ClientListMaxage) NotTypeMaster() ClientListNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeMaster)
}

func (c ClientListMaxage) NotTypeReplica() ClientListNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeReplica)
}

func (c ClientListMaxage) NotTypePubsub() ClientListNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypePubsub)
}

func (c ClientListMaxage) NotId() ClientListNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientListNotIdNotId)
}

func (c ClientListMaxage) NotUser(notUsername string) ClientListNotUser {
	_ = "STUB: not implemented"
	return *new(ClientListNotUser)
}

func (c ClientListMaxage) NotAddr(notAddr string) ClientListNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotAddr)
}

func (c ClientListMaxage) NotLaddr(notLaddr string) ClientListNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotLaddr)
}

func (c ClientListMaxage) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientListMaxage) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListMaxage) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListMaxage) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListMaxage) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListMaxage) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListMaxage) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListMaxage) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListName Incomplete

func (c ClientListName) Idle(idle int64) ClientListIdle {
	_ = "STUB: not implemented"
	return *new(ClientListIdle)
}

func (c ClientListName) Flags(flags string) ClientListFlags {
	_ = "STUB: not implemented"
	return *new(ClientListFlags)
}

func (c ClientListName) LibName(libName string) ClientListLibName {
	_ = "STUB: not implemented"
	return *new(ClientListLibName)
}

func (c ClientListName) LibVer(libVer string) ClientListLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListLibVer)
}

func (c ClientListName) Db(db int64) ClientListDb {
	_ = "STUB: not implemented"
	return *new(ClientListDb)
}

func (c ClientListName) Capa(capa string) ClientListCapa {
	_ = "STUB: not implemented"
	return *new(ClientListCapa)
}

func (c ClientListName) Ip(ip string) ClientListIp {
	_ = "STUB: not implemented"
	return *new(ClientListIp)
}

func (c ClientListName) NotTypeNormal() ClientListNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeNormal)
}

func (c ClientListName) NotTypeMaster() ClientListNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeMaster)
}

func (c ClientListName) NotTypeReplica() ClientListNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeReplica)
}

func (c ClientListName) NotTypePubsub() ClientListNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypePubsub)
}

func (c ClientListName) NotId() ClientListNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientListNotIdNotId)
}

func (c ClientListName) NotUser(notUsername string) ClientListNotUser {
	_ = "STUB: not implemented"
	return *new(ClientListNotUser)
}

func (c ClientListName) NotAddr(notAddr string) ClientListNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotAddr)
}

func (c ClientListName) NotLaddr(notLaddr string) ClientListNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotLaddr)
}

func (c ClientListName) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientListName) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListName) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListName) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListName) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListName) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListName) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListName) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListNotAddr Incomplete

func (c ClientListNotAddr) NotLaddr(notLaddr string) ClientListNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotLaddr)
}

func (c ClientListNotAddr) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientListNotAddr) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListNotAddr) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListNotAddr) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListNotAddr) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListNotAddr) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListNotAddr) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListNotAddr) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListNotCapa Incomplete

func (c ClientListNotCapa) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListNotCapa) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListNotDb Incomplete

func (c ClientListNotDb) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListNotDb) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListNotDb) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListNotFlags Incomplete

func (c ClientListNotFlags) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListNotFlags) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListNotFlags) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListNotFlags) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListNotFlags) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListNotFlags) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListNotIdNotClientId Incomplete

func (c ClientListNotIdNotClientId) NotClientId(notClientId ...int64) ClientListNotIdNotClientId {
	_ = "STUB: not implemented"
	return *new(ClientListNotIdNotClientId)
}

func (c ClientListNotIdNotClientId) NotUser(notUsername string) ClientListNotUser {
	_ = "STUB: not implemented"
	return *new(ClientListNotUser)
}

func (c ClientListNotIdNotClientId) NotAddr(notAddr string) ClientListNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotAddr)
}

func (c ClientListNotIdNotClientId) NotLaddr(notLaddr string) ClientListNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotLaddr)
}

func (c ClientListNotIdNotClientId) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientListNotIdNotClientId) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListNotIdNotClientId) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListNotIdNotClientId) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListNotIdNotClientId) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListNotIdNotClientId) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListNotIdNotClientId) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListNotIdNotClientId) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClientListNotIdNotId Incomplete

func (c ClientListNotIdNotId) NotClientId(notClientId ...int64) ClientListNotIdNotClientId {
	_ = "STUB: not implemented"
	return *new(ClientListNotIdNotClientId)
}

type ClientListNotIp Incomplete

func (c ClientListNotIp) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListNotLaddr Incomplete

func (c ClientListNotLaddr) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientListNotLaddr) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListNotLaddr) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListNotLaddr) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListNotLaddr) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListNotLaddr) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListNotLaddr) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListNotLaddr) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListNotLibName Incomplete

func (c ClientListNotLibName) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListNotLibName) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListNotLibName) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListNotLibName) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListNotLibName) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListNotLibVer Incomplete

func (c ClientListNotLibVer) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListNotLibVer) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListNotLibVer) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListNotLibVer) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListNotName Incomplete

func (c ClientListNotName) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListNotName) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListNotName) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListNotName) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListNotName) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListNotName) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListNotName) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListNotTypeMaster Incomplete

func (c ClientListNotTypeMaster) NotId() ClientListNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientListNotIdNotId)
}

func (c ClientListNotTypeMaster) NotUser(notUsername string) ClientListNotUser {
	_ = "STUB: not implemented"
	return *new(ClientListNotUser)
}

func (c ClientListNotTypeMaster) NotAddr(notAddr string) ClientListNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotAddr)
}

func (c ClientListNotTypeMaster) NotLaddr(notLaddr string) ClientListNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotLaddr)
}

func (c ClientListNotTypeMaster) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientListNotTypeMaster) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListNotTypeMaster) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListNotTypeMaster) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListNotTypeMaster) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListNotTypeMaster) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListNotTypeMaster) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListNotTypeMaster) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClientListNotTypeNormal Incomplete

func (c ClientListNotTypeNormal) NotId() ClientListNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientListNotIdNotId)
}

func (c ClientListNotTypeNormal) NotUser(notUsername string) ClientListNotUser {
	_ = "STUB: not implemented"
	return *new(ClientListNotUser)
}

func (c ClientListNotTypeNormal) NotAddr(notAddr string) ClientListNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotAddr)
}

func (c ClientListNotTypeNormal) NotLaddr(notLaddr string) ClientListNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotLaddr)
}

func (c ClientListNotTypeNormal) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientListNotTypeNormal) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListNotTypeNormal) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListNotTypeNormal) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListNotTypeNormal) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListNotTypeNormal) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListNotTypeNormal) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListNotTypeNormal) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClientListNotTypePubsub Incomplete

func (c ClientListNotTypePubsub) NotId() ClientListNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientListNotIdNotId)
}

func (c ClientListNotTypePubsub) NotUser(notUsername string) ClientListNotUser {
	_ = "STUB: not implemented"
	return *new(ClientListNotUser)
}

func (c ClientListNotTypePubsub) NotAddr(notAddr string) ClientListNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotAddr)
}

func (c ClientListNotTypePubsub) NotLaddr(notLaddr string) ClientListNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotLaddr)
}

func (c ClientListNotTypePubsub) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientListNotTypePubsub) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListNotTypePubsub) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListNotTypePubsub) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListNotTypePubsub) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListNotTypePubsub) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListNotTypePubsub) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListNotTypePubsub) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClientListNotTypeReplica Incomplete

func (c ClientListNotTypeReplica) NotId() ClientListNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientListNotIdNotId)
}

func (c ClientListNotTypeReplica) NotUser(notUsername string) ClientListNotUser {
	_ = "STUB: not implemented"
	return *new(ClientListNotUser)
}

func (c ClientListNotTypeReplica) NotAddr(notAddr string) ClientListNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotAddr)
}

func (c ClientListNotTypeReplica) NotLaddr(notLaddr string) ClientListNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotLaddr)
}

func (c ClientListNotTypeReplica) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientListNotTypeReplica) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListNotTypeReplica) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListNotTypeReplica) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListNotTypeReplica) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListNotTypeReplica) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListNotTypeReplica) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListNotTypeReplica) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClientListNotUser Incomplete

func (c ClientListNotUser) NotAddr(notAddr string) ClientListNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotAddr)
}

func (c ClientListNotUser) NotLaddr(notLaddr string) ClientListNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotLaddr)
}

func (c ClientListNotUser) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientListNotUser) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListNotUser) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListNotUser) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListNotUser) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListNotUser) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListNotUser) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListNotUser) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListSkipmeNo Incomplete

func (c ClientListSkipmeNo) Maxage(maxage int64) ClientListMaxage {
	_ = "STUB: not implemented"
	return *new(ClientListMaxage)
}

func (c ClientListSkipmeNo) Name(name string) ClientListName {
	_ = "STUB: not implemented"
	return *new(ClientListName)
}

func (c ClientListSkipmeNo) Idle(idle int64) ClientListIdle {
	_ = "STUB: not implemented"
	return *new(ClientListIdle)
}

func (c ClientListSkipmeNo) Flags(flags string) ClientListFlags {
	_ = "STUB: not implemented"
	return *new(ClientListFlags)
}

func (c ClientListSkipmeNo) LibName(libName string) ClientListLibName {
	_ = "STUB: not implemented"
	return *new(ClientListLibName)
}

func (c ClientListSkipmeNo) LibVer(libVer string) ClientListLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListLibVer)
}

func (c ClientListSkipmeNo) Db(db int64) ClientListDb {
	_ = "STUB: not implemented"
	return *new(ClientListDb)
}

func (c ClientListSkipmeNo) Capa(capa string) ClientListCapa {
	_ = "STUB: not implemented"
	return *new(ClientListCapa)
}

func (c ClientListSkipmeNo) Ip(ip string) ClientListIp {
	_ = "STUB: not implemented"
	return *new(ClientListIp)
}

func (c ClientListSkipmeNo) NotTypeNormal() ClientListNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeNormal)
}

func (c ClientListSkipmeNo) NotTypeMaster() ClientListNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeMaster)
}

func (c ClientListSkipmeNo) NotTypeReplica() ClientListNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeReplica)
}

func (c ClientListSkipmeNo) NotTypePubsub() ClientListNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypePubsub)
}

func (c ClientListSkipmeNo) NotId() ClientListNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientListNotIdNotId)
}

func (c ClientListSkipmeNo) NotUser(notUsername string) ClientListNotUser {
	_ = "STUB: not implemented"
	return *new(ClientListNotUser)
}

func (c ClientListSkipmeNo) NotAddr(notAddr string) ClientListNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotAddr)
}

func (c ClientListSkipmeNo) NotLaddr(notLaddr string) ClientListNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotLaddr)
}

func (c ClientListSkipmeNo) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientListSkipmeNo) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListSkipmeNo) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListSkipmeNo) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListSkipmeNo) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListSkipmeNo) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListSkipmeNo) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListSkipmeNo) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListSkipmeYes Incomplete

func (c ClientListSkipmeYes) Maxage(maxage int64) ClientListMaxage {
	_ = "STUB: not implemented"
	return *new(ClientListMaxage)
}

func (c ClientListSkipmeYes) Name(name string) ClientListName {
	_ = "STUB: not implemented"
	return *new(ClientListName)
}

func (c ClientListSkipmeYes) Idle(idle int64) ClientListIdle {
	_ = "STUB: not implemented"
	return *new(ClientListIdle)
}

func (c ClientListSkipmeYes) Flags(flags string) ClientListFlags {
	_ = "STUB: not implemented"
	return *new(ClientListFlags)
}

func (c ClientListSkipmeYes) LibName(libName string) ClientListLibName {
	_ = "STUB: not implemented"
	return *new(ClientListLibName)
}

func (c ClientListSkipmeYes) LibVer(libVer string) ClientListLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListLibVer)
}

func (c ClientListSkipmeYes) Db(db int64) ClientListDb {
	_ = "STUB: not implemented"
	return *new(ClientListDb)
}

func (c ClientListSkipmeYes) Capa(capa string) ClientListCapa {
	_ = "STUB: not implemented"
	return *new(ClientListCapa)
}

func (c ClientListSkipmeYes) Ip(ip string) ClientListIp {
	_ = "STUB: not implemented"
	return *new(ClientListIp)
}

func (c ClientListSkipmeYes) NotTypeNormal() ClientListNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeNormal)
}

func (c ClientListSkipmeYes) NotTypeMaster() ClientListNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeMaster)
}

func (c ClientListSkipmeYes) NotTypeReplica() ClientListNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeReplica)
}

func (c ClientListSkipmeYes) NotTypePubsub() ClientListNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypePubsub)
}

func (c ClientListSkipmeYes) NotId() ClientListNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientListNotIdNotId)
}

func (c ClientListSkipmeYes) NotUser(notUsername string) ClientListNotUser {
	_ = "STUB: not implemented"
	return *new(ClientListNotUser)
}

func (c ClientListSkipmeYes) NotAddr(notAddr string) ClientListNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotAddr)
}

func (c ClientListSkipmeYes) NotLaddr(notLaddr string) ClientListNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotLaddr)
}

func (c ClientListSkipmeYes) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientListSkipmeYes) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListSkipmeYes) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListSkipmeYes) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListSkipmeYes) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListSkipmeYes) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListSkipmeYes) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListSkipmeYes) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListTypeMaster Incomplete

func (c ClientListTypeMaster) Id() ClientListIdId {
	_ = "STUB: not implemented"
	return *new(ClientListIdId)
}

func (c ClientListTypeMaster) User(username string) ClientListUser {
	_ = "STUB: not implemented"
	return *new(ClientListUser)
}

func (c ClientListTypeMaster) Addr(ipPort string) ClientListAddr {
	_ = "STUB: not implemented"
	return *new(ClientListAddr)
}

func (c ClientListTypeMaster) Laddr(ipPort string) ClientListLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListLaddr)
}

func (c ClientListTypeMaster) SkipmeYes() ClientListSkipmeYes {
	_ = "STUB: not implemented"
	return *new(ClientListSkipmeYes)
}

func (c ClientListTypeMaster) SkipmeNo() ClientListSkipmeNo {
	_ = "STUB: not implemented"
	return *new(ClientListSkipmeNo)
}

func (c ClientListTypeMaster) Maxage(maxage int64) ClientListMaxage {
	_ = "STUB: not implemented"
	return *new(ClientListMaxage)
}

func (c ClientListTypeMaster) Name(name string) ClientListName {
	_ = "STUB: not implemented"
	return *new(ClientListName)
}

func (c ClientListTypeMaster) Idle(idle int64) ClientListIdle {
	_ = "STUB: not implemented"
	return *new(ClientListIdle)
}

func (c ClientListTypeMaster) Flags(flags string) ClientListFlags {
	_ = "STUB: not implemented"
	return *new(ClientListFlags)
}

func (c ClientListTypeMaster) LibName(libName string) ClientListLibName {
	_ = "STUB: not implemented"
	return *new(ClientListLibName)
}

func (c ClientListTypeMaster) LibVer(libVer string) ClientListLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListLibVer)
}

func (c ClientListTypeMaster) Db(db int64) ClientListDb {
	_ = "STUB: not implemented"
	return *new(ClientListDb)
}

func (c ClientListTypeMaster) Capa(capa string) ClientListCapa {
	_ = "STUB: not implemented"
	return *new(ClientListCapa)
}

func (c ClientListTypeMaster) Ip(ip string) ClientListIp {
	_ = "STUB: not implemented"
	return *new(ClientListIp)
}

func (c ClientListTypeMaster) NotTypeNormal() ClientListNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeNormal)
}

func (c ClientListTypeMaster) NotTypeMaster() ClientListNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeMaster)
}

func (c ClientListTypeMaster) NotTypeReplica() ClientListNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeReplica)
}

func (c ClientListTypeMaster) NotTypePubsub() ClientListNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypePubsub)
}

func (c ClientListTypeMaster) NotId() ClientListNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientListNotIdNotId)
}

func (c ClientListTypeMaster) NotUser(notUsername string) ClientListNotUser {
	_ = "STUB: not implemented"
	return *new(ClientListNotUser)
}

func (c ClientListTypeMaster) NotAddr(notAddr string) ClientListNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotAddr)
}

func (c ClientListTypeMaster) NotLaddr(notLaddr string) ClientListNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotLaddr)
}

func (c ClientListTypeMaster) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientListTypeMaster) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListTypeMaster) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListTypeMaster) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListTypeMaster) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListTypeMaster) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListTypeMaster) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListTypeMaster) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListTypeNormal Incomplete

func (c ClientListTypeNormal) Id() ClientListIdId {
	_ = "STUB: not implemented"
	return *new(ClientListIdId)
}

func (c ClientListTypeNormal) User(username string) ClientListUser {
	_ = "STUB: not implemented"
	return *new(ClientListUser)
}

func (c ClientListTypeNormal) Addr(ipPort string) ClientListAddr {
	_ = "STUB: not implemented"
	return *new(ClientListAddr)
}

func (c ClientListTypeNormal) Laddr(ipPort string) ClientListLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListLaddr)
}

func (c ClientListTypeNormal) SkipmeYes() ClientListSkipmeYes {
	_ = "STUB: not implemented"
	return *new(ClientListSkipmeYes)
}

func (c ClientListTypeNormal) SkipmeNo() ClientListSkipmeNo {
	_ = "STUB: not implemented"
	return *new(ClientListSkipmeNo)
}

func (c ClientListTypeNormal) Maxage(maxage int64) ClientListMaxage {
	_ = "STUB: not implemented"
	return *new(ClientListMaxage)
}

func (c ClientListTypeNormal) Name(name string) ClientListName {
	_ = "STUB: not implemented"
	return *new(ClientListName)
}

func (c ClientListTypeNormal) Idle(idle int64) ClientListIdle {
	_ = "STUB: not implemented"
	return *new(ClientListIdle)
}

func (c ClientListTypeNormal) Flags(flags string) ClientListFlags {
	_ = "STUB: not implemented"
	return *new(ClientListFlags)
}

func (c ClientListTypeNormal) LibName(libName string) ClientListLibName {
	_ = "STUB: not implemented"
	return *new(ClientListLibName)
}

func (c ClientListTypeNormal) LibVer(libVer string) ClientListLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListLibVer)
}

func (c ClientListTypeNormal) Db(db int64) ClientListDb {
	_ = "STUB: not implemented"
	return *new(ClientListDb)
}

func (c ClientListTypeNormal) Capa(capa string) ClientListCapa {
	_ = "STUB: not implemented"
	return *new(ClientListCapa)
}

func (c ClientListTypeNormal) Ip(ip string) ClientListIp {
	_ = "STUB: not implemented"
	return *new(ClientListIp)
}

func (c ClientListTypeNormal) NotTypeNormal() ClientListNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeNormal)
}

func (c ClientListTypeNormal) NotTypeMaster() ClientListNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeMaster)
}

func (c ClientListTypeNormal) NotTypeReplica() ClientListNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeReplica)
}

func (c ClientListTypeNormal) NotTypePubsub() ClientListNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypePubsub)
}

func (c ClientListTypeNormal) NotId() ClientListNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientListNotIdNotId)
}

func (c ClientListTypeNormal) NotUser(notUsername string) ClientListNotUser {
	_ = "STUB: not implemented"
	return *new(ClientListNotUser)
}

func (c ClientListTypeNormal) NotAddr(notAddr string) ClientListNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotAddr)
}

func (c ClientListTypeNormal) NotLaddr(notLaddr string) ClientListNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotLaddr)
}

func (c ClientListTypeNormal) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientListTypeNormal) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListTypeNormal) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListTypeNormal) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListTypeNormal) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListTypeNormal) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListTypeNormal) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListTypeNormal) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListTypePubsub Incomplete

func (c ClientListTypePubsub) Id() ClientListIdId {
	_ = "STUB: not implemented"
	return *new(ClientListIdId)
}

func (c ClientListTypePubsub) User(username string) ClientListUser {
	_ = "STUB: not implemented"
	return *new(ClientListUser)
}

func (c ClientListTypePubsub) Addr(ipPort string) ClientListAddr {
	_ = "STUB: not implemented"
	return *new(ClientListAddr)
}

func (c ClientListTypePubsub) Laddr(ipPort string) ClientListLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListLaddr)
}

func (c ClientListTypePubsub) SkipmeYes() ClientListSkipmeYes {
	_ = "STUB: not implemented"
	return *new(ClientListSkipmeYes)
}

func (c ClientListTypePubsub) SkipmeNo() ClientListSkipmeNo {
	_ = "STUB: not implemented"
	return *new(ClientListSkipmeNo)
}

func (c ClientListTypePubsub) Maxage(maxage int64) ClientListMaxage {
	_ = "STUB: not implemented"
	return *new(ClientListMaxage)
}

func (c ClientListTypePubsub) Name(name string) ClientListName {
	_ = "STUB: not implemented"
	return *new(ClientListName)
}

func (c ClientListTypePubsub) Idle(idle int64) ClientListIdle {
	_ = "STUB: not implemented"
	return *new(ClientListIdle)
}

func (c ClientListTypePubsub) Flags(flags string) ClientListFlags {
	_ = "STUB: not implemented"
	return *new(ClientListFlags)
}

func (c ClientListTypePubsub) LibName(libName string) ClientListLibName {
	_ = "STUB: not implemented"
	return *new(ClientListLibName)
}

func (c ClientListTypePubsub) LibVer(libVer string) ClientListLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListLibVer)
}

func (c ClientListTypePubsub) Db(db int64) ClientListDb {
	_ = "STUB: not implemented"
	return *new(ClientListDb)
}

func (c ClientListTypePubsub) Capa(capa string) ClientListCapa {
	_ = "STUB: not implemented"
	return *new(ClientListCapa)
}

func (c ClientListTypePubsub) Ip(ip string) ClientListIp {
	_ = "STUB: not implemented"
	return *new(ClientListIp)
}

func (c ClientListTypePubsub) NotTypeNormal() ClientListNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeNormal)
}

func (c ClientListTypePubsub) NotTypeMaster() ClientListNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeMaster)
}

func (c ClientListTypePubsub) NotTypeReplica() ClientListNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeReplica)
}

func (c ClientListTypePubsub) NotTypePubsub() ClientListNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypePubsub)
}

func (c ClientListTypePubsub) NotId() ClientListNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientListNotIdNotId)
}

func (c ClientListTypePubsub) NotUser(notUsername string) ClientListNotUser {
	_ = "STUB: not implemented"
	return *new(ClientListNotUser)
}

func (c ClientListTypePubsub) NotAddr(notAddr string) ClientListNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotAddr)
}

func (c ClientListTypePubsub) NotLaddr(notLaddr string) ClientListNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotLaddr)
}

func (c ClientListTypePubsub) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientListTypePubsub) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListTypePubsub) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListTypePubsub) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListTypePubsub) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListTypePubsub) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListTypePubsub) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListTypePubsub) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListTypeReplica Incomplete

func (c ClientListTypeReplica) Id() ClientListIdId {
	_ = "STUB: not implemented"
	return *new(ClientListIdId)
}

func (c ClientListTypeReplica) User(username string) ClientListUser {
	_ = "STUB: not implemented"
	return *new(ClientListUser)
}

func (c ClientListTypeReplica) Addr(ipPort string) ClientListAddr {
	_ = "STUB: not implemented"
	return *new(ClientListAddr)
}

func (c ClientListTypeReplica) Laddr(ipPort string) ClientListLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListLaddr)
}

func (c ClientListTypeReplica) SkipmeYes() ClientListSkipmeYes {
	_ = "STUB: not implemented"
	return *new(ClientListSkipmeYes)
}

func (c ClientListTypeReplica) SkipmeNo() ClientListSkipmeNo {
	_ = "STUB: not implemented"
	return *new(ClientListSkipmeNo)
}

func (c ClientListTypeReplica) Maxage(maxage int64) ClientListMaxage {
	_ = "STUB: not implemented"
	return *new(ClientListMaxage)
}

func (c ClientListTypeReplica) Name(name string) ClientListName {
	_ = "STUB: not implemented"
	return *new(ClientListName)
}

func (c ClientListTypeReplica) Idle(idle int64) ClientListIdle {
	_ = "STUB: not implemented"
	return *new(ClientListIdle)
}

func (c ClientListTypeReplica) Flags(flags string) ClientListFlags {
	_ = "STUB: not implemented"
	return *new(ClientListFlags)
}

func (c ClientListTypeReplica) LibName(libName string) ClientListLibName {
	_ = "STUB: not implemented"
	return *new(ClientListLibName)
}

func (c ClientListTypeReplica) LibVer(libVer string) ClientListLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListLibVer)
}

func (c ClientListTypeReplica) Db(db int64) ClientListDb {
	_ = "STUB: not implemented"
	return *new(ClientListDb)
}

func (c ClientListTypeReplica) Capa(capa string) ClientListCapa {
	_ = "STUB: not implemented"
	return *new(ClientListCapa)
}

func (c ClientListTypeReplica) Ip(ip string) ClientListIp {
	_ = "STUB: not implemented"
	return *new(ClientListIp)
}

func (c ClientListTypeReplica) NotTypeNormal() ClientListNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeNormal)
}

func (c ClientListTypeReplica) NotTypeMaster() ClientListNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeMaster)
}

func (c ClientListTypeReplica) NotTypeReplica() ClientListNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeReplica)
}

func (c ClientListTypeReplica) NotTypePubsub() ClientListNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypePubsub)
}

func (c ClientListTypeReplica) NotId() ClientListNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientListNotIdNotId)
}

func (c ClientListTypeReplica) NotUser(notUsername string) ClientListNotUser {
	_ = "STUB: not implemented"
	return *new(ClientListNotUser)
}

func (c ClientListTypeReplica) NotAddr(notAddr string) ClientListNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotAddr)
}

func (c ClientListTypeReplica) NotLaddr(notLaddr string) ClientListNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotLaddr)
}

func (c ClientListTypeReplica) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientListTypeReplica) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListTypeReplica) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListTypeReplica) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListTypeReplica) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListTypeReplica) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListTypeReplica) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListTypeReplica) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientListUser Incomplete

func (c ClientListUser) Addr(ipPort string) ClientListAddr {
	_ = "STUB: not implemented"
	return *new(ClientListAddr)
}

func (c ClientListUser) Laddr(ipPort string) ClientListLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListLaddr)
}

func (c ClientListUser) SkipmeYes() ClientListSkipmeYes {
	_ = "STUB: not implemented"
	return *new(ClientListSkipmeYes)
}

func (c ClientListUser) SkipmeNo() ClientListSkipmeNo {
	_ = "STUB: not implemented"
	return *new(ClientListSkipmeNo)
}

func (c ClientListUser) Maxage(maxage int64) ClientListMaxage {
	_ = "STUB: not implemented"
	return *new(ClientListMaxage)
}

func (c ClientListUser) Name(name string) ClientListName {
	_ = "STUB: not implemented"
	return *new(ClientListName)
}

func (c ClientListUser) Idle(idle int64) ClientListIdle {
	_ = "STUB: not implemented"
	return *new(ClientListIdle)
}

func (c ClientListUser) Flags(flags string) ClientListFlags {
	_ = "STUB: not implemented"
	return *new(ClientListFlags)
}

func (c ClientListUser) LibName(libName string) ClientListLibName {
	_ = "STUB: not implemented"
	return *new(ClientListLibName)
}

func (c ClientListUser) LibVer(libVer string) ClientListLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListLibVer)
}

func (c ClientListUser) Db(db int64) ClientListDb {
	_ = "STUB: not implemented"
	return *new(ClientListDb)
}

func (c ClientListUser) Capa(capa string) ClientListCapa {
	_ = "STUB: not implemented"
	return *new(ClientListCapa)
}

func (c ClientListUser) Ip(ip string) ClientListIp {
	_ = "STUB: not implemented"
	return *new(ClientListIp)
}

func (c ClientListUser) NotTypeNormal() ClientListNotTypeNormal {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeNormal)
}

func (c ClientListUser) NotTypeMaster() ClientListNotTypeMaster {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeMaster)
}

func (c ClientListUser) NotTypeReplica() ClientListNotTypeReplica {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypeReplica)
}

func (c ClientListUser) NotTypePubsub() ClientListNotTypePubsub {
	_ = "STUB: not implemented"
	return *new(ClientListNotTypePubsub)
}

func (c ClientListUser) NotId() ClientListNotIdNotId {
	_ = "STUB: not implemented"
	return *new(ClientListNotIdNotId)
}

func (c ClientListUser) NotUser(notUsername string) ClientListNotUser {
	_ = "STUB: not implemented"
	return *new(ClientListNotUser)
}

func (c ClientListUser) NotAddr(notAddr string) ClientListNotAddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotAddr)
}

func (c ClientListUser) NotLaddr(notLaddr string) ClientListNotLaddr {
	_ = "STUB: not implemented"
	return *new(ClientListNotLaddr)
}

func (c ClientListUser) NotName(notName string) ClientListNotName {
	_ = "STUB: not implemented"
	return *new(ClientListNotName)
}

func (c ClientListUser) NotFlags(notFlags string) ClientListNotFlags {
	_ = "STUB: not implemented"
	return *new(ClientListNotFlags)
}

func (c ClientListUser) NotLibName(notLibName string) ClientListNotLibName {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibName)
}

func (c ClientListUser) NotLibVer(notLibVer string) ClientListNotLibVer {
	_ = "STUB: not implemented"
	return *new(ClientListNotLibVer)
}

func (c ClientListUser) NotDb(notDb int64) ClientListNotDb {
	_ = "STUB: not implemented"
	return *new(ClientListNotDb)
}

func (c ClientListUser) NotCapa(notCapa string) ClientListNotCapa {
	_ = "STUB: not implemented"
	return *new(ClientListNotCapa)
}

func (c ClientListUser) NotIp(notIp string) ClientListNotIp {
	_ = "STUB: not implemented"
	return *new(ClientListNotIp)
}

func (c ClientListUser) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientNoEvict Incomplete

func (b Builder) ClientNoEvict() (c ClientNoEvict) {
	_ = "STUB: not implemented"
	return *new(ClientNoEvict)
}

func (c ClientNoEvict) On() ClientNoEvictEnabledOn {
	_ = "STUB: not implemented"
	return *new(ClientNoEvictEnabledOn)
}

func (c ClientNoEvict) Off() ClientNoEvictEnabledOff {
	_ = "STUB: not implemented"
	return *new(ClientNoEvictEnabledOff)
}

type ClientNoEvictEnabledOff Incomplete

func (c ClientNoEvictEnabledOff) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClientNoEvictEnabledOn Incomplete

func (c ClientNoEvictEnabledOn) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClientNoTouch Incomplete

func (b Builder) ClientNoTouch() (c ClientNoTouch) {
	_ = "STUB: not implemented"
	return *new(ClientNoTouch)
}

func (c ClientNoTouch) On() ClientNoTouchEnabledOn {
	_ = "STUB: not implemented"
	return *new(ClientNoTouchEnabledOn)
}

func (c ClientNoTouch) Off() ClientNoTouchEnabledOff {
	_ = "STUB: not implemented"
	return *new(ClientNoTouchEnabledOff)
}

type ClientNoTouchEnabledOff Incomplete

func (c ClientNoTouchEnabledOff) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClientNoTouchEnabledOn Incomplete

func (c ClientNoTouchEnabledOn) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClientPause Incomplete

func (b Builder) ClientPause() (c ClientPause) { _ = "STUB: not implemented"; return *new(ClientPause) }

func (c ClientPause) Timeout(timeout int64) ClientPauseTimeout {
	_ = "STUB: not implemented"
	return *new(ClientPauseTimeout)
}

type ClientPauseModeAll Incomplete

func (c ClientPauseModeAll) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientPauseModeWrite Incomplete

func (c ClientPauseModeWrite) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientPauseTimeout Incomplete

func (c ClientPauseTimeout) Write() ClientPauseModeWrite {
	_ = "STUB: not implemented"
	return *new(ClientPauseModeWrite)
}

func (c ClientPauseTimeout) All() ClientPauseModeAll {
	_ = "STUB: not implemented"
	return *new(ClientPauseModeAll)
}

func (c ClientPauseTimeout) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientReply Incomplete

func (b Builder) ClientReply() (c ClientReply) { _ = "STUB: not implemented"; return *new(ClientReply) }

func (c ClientReply) On() ClientReplyReplyModeOn {
	_ = "STUB: not implemented"
	return *new(ClientReplyReplyModeOn)
}

func (c ClientReply) Off() ClientReplyReplyModeOff {
	_ = "STUB: not implemented"
	return *new(ClientReplyReplyModeOff)
}

func (c ClientReply) Skip() ClientReplyReplyModeSkip {
	_ = "STUB: not implemented"
	return *new(ClientReplyReplyModeSkip)
}

type ClientReplyReplyModeOff Incomplete

func (c ClientReplyReplyModeOff) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClientReplyReplyModeOn Incomplete

func (c ClientReplyReplyModeOn) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClientReplyReplyModeSkip Incomplete

func (c ClientReplyReplyModeSkip) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClientSetinfo Incomplete

func (b Builder) ClientSetinfo() (c ClientSetinfo) {
	_ = "STUB: not implemented"
	return *new(ClientSetinfo)
}

func (c ClientSetinfo) Libname(libname string) ClientSetinfoAttrLibname {
	_ = "STUB: not implemented"
	return *new(ClientSetinfoAttrLibname)
}

func (c ClientSetinfo) Libver(libver string) ClientSetinfoAttrLibver {
	_ = "STUB: not implemented"
	return *new(ClientSetinfoAttrLibver)
}

type ClientSetinfoAttrLibname Incomplete

func (c ClientSetinfoAttrLibname) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClientSetinfoAttrLibver Incomplete

func (c ClientSetinfoAttrLibver) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClientSetname Incomplete

func (b Builder) ClientSetname() (c ClientSetname) {
	_ = "STUB: not implemented"
	return *new(ClientSetname)
}

func (c ClientSetname) ConnectionName(connectionName string) ClientSetnameConnectionName {
	_ = "STUB: not implemented"
	return *new(ClientSetnameConnectionName)
}

type ClientSetnameConnectionName Incomplete

func (c ClientSetnameConnectionName) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClientTracking Incomplete

func (b Builder) ClientTracking() (c ClientTracking) {
	_ = "STUB: not implemented"
	return *new(ClientTracking)
}

func (c ClientTracking) On() ClientTrackingStatusOn {
	_ = "STUB: not implemented"
	return *new(ClientTrackingStatusOn)
}

func (c ClientTracking) Off() ClientTrackingStatusOff {
	_ = "STUB: not implemented"
	return *new(ClientTrackingStatusOff)
}

type ClientTrackingBcast Incomplete

func (c ClientTrackingBcast) Optin() ClientTrackingOptin {
	_ = "STUB: not implemented"
	return *new(ClientTrackingOptin)
}

func (c ClientTrackingBcast) Optout() ClientTrackingOptout {
	_ = "STUB: not implemented"
	return *new(ClientTrackingOptout)
}

func (c ClientTrackingBcast) Noloop() ClientTrackingNoloop {
	_ = "STUB: not implemented"
	return *new(ClientTrackingNoloop)
}

func (c ClientTrackingBcast) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientTrackingNoloop Incomplete

func (c ClientTrackingNoloop) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientTrackingOptin Incomplete

func (c ClientTrackingOptin) Optout() ClientTrackingOptout {
	_ = "STUB: not implemented"
	return *new(ClientTrackingOptout)
}

func (c ClientTrackingOptin) Noloop() ClientTrackingNoloop {
	_ = "STUB: not implemented"
	return *new(ClientTrackingNoloop)
}

func (c ClientTrackingOptin) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientTrackingOptout Incomplete

func (c ClientTrackingOptout) Noloop() ClientTrackingNoloop {
	_ = "STUB: not implemented"
	return *new(ClientTrackingNoloop)
}

func (c ClientTrackingOptout) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientTrackingPrefix Incomplete

func (c ClientTrackingPrefix) Prefix(prefix string) ClientTrackingPrefix {
	_ = "STUB: not implemented"
	return *new(ClientTrackingPrefix)
}

func (c ClientTrackingPrefix) Bcast() ClientTrackingBcast {
	_ = "STUB: not implemented"
	return *new(ClientTrackingBcast)
}

func (c ClientTrackingPrefix) Optin() ClientTrackingOptin {
	_ = "STUB: not implemented"
	return *new(ClientTrackingOptin)
}

func (c ClientTrackingPrefix) Optout() ClientTrackingOptout {
	_ = "STUB: not implemented"
	return *new(ClientTrackingOptout)
}

func (c ClientTrackingPrefix) Noloop() ClientTrackingNoloop {
	_ = "STUB: not implemented"
	return *new(ClientTrackingNoloop)
}

func (c ClientTrackingPrefix) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientTrackingRedirect Incomplete

func (c ClientTrackingRedirect) Prefix() ClientTrackingPrefix {
	_ = "STUB: not implemented"
	return *new(ClientTrackingPrefix)
}

func (c ClientTrackingRedirect) Bcast() ClientTrackingBcast {
	_ = "STUB: not implemented"
	return *new(ClientTrackingBcast)
}

func (c ClientTrackingRedirect) Optin() ClientTrackingOptin {
	_ = "STUB: not implemented"
	return *new(ClientTrackingOptin)
}

func (c ClientTrackingRedirect) Optout() ClientTrackingOptout {
	_ = "STUB: not implemented"
	return *new(ClientTrackingOptout)
}

func (c ClientTrackingRedirect) Noloop() ClientTrackingNoloop {
	_ = "STUB: not implemented"
	return *new(ClientTrackingNoloop)
}

func (c ClientTrackingRedirect) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClientTrackingStatusOff Incomplete

func (c ClientTrackingStatusOff) Redirect(clientId int64) ClientTrackingRedirect {
	_ = "STUB: not implemented"
	return *new(ClientTrackingRedirect)
}

func (c ClientTrackingStatusOff) Prefix() ClientTrackingPrefix {
	_ = "STUB: not implemented"
	return *new(ClientTrackingPrefix)
}

func (c ClientTrackingStatusOff) Bcast() ClientTrackingBcast {
	_ = "STUB: not implemented"
	return *new(ClientTrackingBcast)
}

func (c ClientTrackingStatusOff) Optin() ClientTrackingOptin {
	_ = "STUB: not implemented"
	return *new(ClientTrackingOptin)
}

func (c ClientTrackingStatusOff) Optout() ClientTrackingOptout {
	_ = "STUB: not implemented"
	return *new(ClientTrackingOptout)
}

func (c ClientTrackingStatusOff) Noloop() ClientTrackingNoloop {
	_ = "STUB: not implemented"
	return *new(ClientTrackingNoloop)
}

func (c ClientTrackingStatusOff) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClientTrackingStatusOn Incomplete

func (c ClientTrackingStatusOn) Redirect(clientId int64) ClientTrackingRedirect {
	_ = "STUB: not implemented"
	return *new(ClientTrackingRedirect)
}

func (c ClientTrackingStatusOn) Prefix() ClientTrackingPrefix {
	_ = "STUB: not implemented"
	return *new(ClientTrackingPrefix)
}

func (c ClientTrackingStatusOn) Bcast() ClientTrackingBcast {
	_ = "STUB: not implemented"
	return *new(ClientTrackingBcast)
}

func (c ClientTrackingStatusOn) Optin() ClientTrackingOptin {
	_ = "STUB: not implemented"
	return *new(ClientTrackingOptin)
}

func (c ClientTrackingStatusOn) Optout() ClientTrackingOptout {
	_ = "STUB: not implemented"
	return *new(ClientTrackingOptout)
}

func (c ClientTrackingStatusOn) Noloop() ClientTrackingNoloop {
	_ = "STUB: not implemented"
	return *new(ClientTrackingNoloop)
}

func (c ClientTrackingStatusOn) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClientTrackinginfo Incomplete

func (b Builder) ClientTrackinginfo() (c ClientTrackinginfo) {
	_ = "STUB: not implemented"
	return *new(ClientTrackinginfo)
}

func (c ClientTrackinginfo) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientUnblock Incomplete

func (b Builder) ClientUnblock() (c ClientUnblock) {
	_ = "STUB: not implemented"
	return *new(ClientUnblock)
}

func (c ClientUnblock) ClientId(clientId int64) ClientUnblockClientId {
	_ = "STUB: not implemented"
	return *new(ClientUnblockClientId)
}

type ClientUnblockClientId Incomplete

func (c ClientUnblockClientId) Timeout() ClientUnblockUnblockTypeTimeout {
	_ = "STUB: not implemented"
	return *new(ClientUnblockUnblockTypeTimeout)
}

func (c ClientUnblockClientId) Error() ClientUnblockUnblockTypeError {
	_ = "STUB: not implemented"
	return *new(ClientUnblockUnblockTypeError)
}

func (c ClientUnblockClientId) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type ClientUnblockUnblockTypeError Incomplete

func (c ClientUnblockUnblockTypeError) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClientUnblockUnblockTypeTimeout Incomplete

func (c ClientUnblockUnblockTypeTimeout) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type ClientUnpause Incomplete

func (b Builder) ClientUnpause() (c ClientUnpause) {
	_ = "STUB: not implemented"
	return *new(ClientUnpause)
}

func (c ClientUnpause) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Echo Incomplete

func (b Builder) Echo() (c Echo) { _ = "STUB: not implemented"; return *new(Echo) }

func (c Echo) Message(message string) EchoMessage {
	_ = "STUB: not implemented"
	return *new(EchoMessage)
}

type EchoMessage Incomplete

func (c EchoMessage) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Hello Incomplete

func (b Builder) Hello() (c Hello) { _ = "STUB: not implemented"; return *new(Hello) }

func (c Hello) Protover(protover int64) HelloArgumentsProtover {
	_ = "STUB: not implemented"
	return *new(HelloArgumentsProtover)
}

func (c Hello) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HelloArgumentsAuth Incomplete

func (c HelloArgumentsAuth) Setname(clientname string) HelloArgumentsSetname {
	_ = "STUB: not implemented"
	return *new(HelloArgumentsSetname)
}

func (c HelloArgumentsAuth) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type HelloArgumentsProtover Incomplete

func (c HelloArgumentsProtover) Auth(username string, password string) HelloArgumentsAuth {
	_ = "STUB: not implemented"
	return *new(HelloArgumentsAuth)
}

func (c HelloArgumentsProtover) Setname(clientname string) HelloArgumentsSetname {
	_ = "STUB: not implemented"
	return *new(HelloArgumentsSetname)
}

func (c HelloArgumentsProtover) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type HelloArgumentsSetname Incomplete

func (c HelloArgumentsSetname) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Ping Incomplete

func (b Builder) Ping() (c Ping) { _ = "STUB: not implemented"; return *new(Ping) }

func (c Ping) Message(message string) PingMessage {
	_ = "STUB: not implemented"
	return *new(PingMessage)
}

func (c Ping) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type PingMessage Incomplete

func (c PingMessage) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Quit Incomplete

func (b Builder) Quit() (c Quit) { _ = "STUB: not implemented"; return *new(Quit) }

func (c Quit) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Reset Incomplete

func (b Builder) Reset() (c Reset) { _ = "STUB: not implemented"; return *new(Reset) }

func (c Reset) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Select Incomplete

func (b Builder) Select() (c Select) { _ = "STUB: not implemented"; return *new(Select) }

func (c Select) Index(index int64) SelectIndex { _ = "STUB: not implemented"; return *new(SelectIndex) }

type SelectIndex Incomplete

func (c SelectIndex) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

// Code generated DO NOT EDIT

package cmds

type Geoadd Incomplete

func (b Builder) Geoadd() (c Geoadd) { _ = "STUB: not implemented"; return *new(Geoadd) }

func (c Geoadd) Key(key string) GeoaddKey { _ = "STUB: not implemented"; return *new(GeoaddKey) }

type GeoaddChangeCh Incomplete

func (c GeoaddChangeCh) LongitudeLatitudeMember() GeoaddLongitudeLatitudeMember {
	_ = "STUB: not implemented"
	return *new(GeoaddLongitudeLatitudeMember)
}

type GeoaddConditionNx Incomplete

func (c GeoaddConditionNx) Ch() GeoaddChangeCh {
	_ = "STUB: not implemented"
	return *new(GeoaddChangeCh)
}

func (c GeoaddConditionNx) LongitudeLatitudeMember() GeoaddLongitudeLatitudeMember {
	_ = "STUB: not implemented"
	return *new(GeoaddLongitudeLatitudeMember)
}

type GeoaddConditionXx Incomplete

func (c GeoaddConditionXx) Ch() GeoaddChangeCh {
	_ = "STUB: not implemented"
	return *new(GeoaddChangeCh)
}

func (c GeoaddConditionXx) LongitudeLatitudeMember() GeoaddLongitudeLatitudeMember {
	_ = "STUB: not implemented"
	return *new(GeoaddLongitudeLatitudeMember)
}

type GeoaddKey Incomplete

func (c GeoaddKey) Nx() GeoaddConditionNx {
	_ = "STUB: not implemented"
	return *new(GeoaddConditionNx)
}

func (c GeoaddKey) Xx() GeoaddConditionXx {
	_ = "STUB: not implemented"
	return *new(GeoaddConditionXx)
}

func (c GeoaddKey) Ch() GeoaddChangeCh { _ = "STUB: not implemented"; return *new(GeoaddChangeCh) }

func (c GeoaddKey) LongitudeLatitudeMember() GeoaddLongitudeLatitudeMember {
	_ = "STUB: not implemented"
	return *new(GeoaddLongitudeLatitudeMember)
}

type GeoaddLongitudeLatitudeMember Incomplete

func (c GeoaddLongitudeLatitudeMember) LongitudeLatitudeMember(longitude float64, latitude float64, member string) GeoaddLongitudeLatitudeMember {
	_ = "STUB: not implemented"
	return *new(GeoaddLongitudeLatitudeMember)
}

func (c GeoaddLongitudeLatitudeMember) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type Geodist Incomplete

func (b Builder) Geodist() (c Geodist) { _ = "STUB: not implemented"; return *new(Geodist) }

func (c Geodist) Key(key string) GeodistKey { _ = "STUB: not implemented"; return *new(GeodistKey) }

type GeodistKey Incomplete

func (c GeodistKey) Member1(member1 string) GeodistMember1 {
	_ = "STUB: not implemented"
	return *new(GeodistMember1)
}

type GeodistMember1 Incomplete

func (c GeodistMember1) Member2(member2 string) GeodistMember2 {
	_ = "STUB: not implemented"
	return *new(GeodistMember2)
}

type GeodistMember2 Incomplete

func (c GeodistMember2) M() GeodistUnitM { _ = "STUB: not implemented"; return *new(GeodistUnitM) }

func (c GeodistMember2) Km() GeodistUnitKm { _ = "STUB: not implemented"; return *new(GeodistUnitKm) }

func (c GeodistMember2) Ft() GeodistUnitFt { _ = "STUB: not implemented"; return *new(GeodistUnitFt) }

func (c GeodistMember2) Mi() GeodistUnitMi { _ = "STUB: not implemented"; return *new(GeodistUnitMi) }

func (c GeodistMember2) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GeodistMember2) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type GeodistUnitFt Incomplete

func (c GeodistUnitFt) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GeodistUnitFt) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type GeodistUnitKm Incomplete

func (c GeodistUnitKm) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GeodistUnitKm) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type GeodistUnitM Incomplete

func (c GeodistUnitM) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GeodistUnitM) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type GeodistUnitMi Incomplete

func (c GeodistUnitMi) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GeodistUnitMi) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Geohash Incomplete

func (b Builder) Geohash() (c Geohash) { _ = "STUB: not implemented"; return *new(Geohash) }

func (c Geohash) Key(key string) GeohashKey { _ = "STUB: not implemented"; return *new(GeohashKey) }

type GeohashKey Incomplete

func (c GeohashKey) Member(member ...string) GeohashMember {
	_ = "STUB: not implemented"
	return *new(GeohashMember)
}

func (c GeohashKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GeohashKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type GeohashMember Incomplete

func (c GeohashMember) Member(member ...string) GeohashMember {
	_ = "STUB: not implemented"
	return *new(GeohashMember)
}

func (c GeohashMember) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GeohashMember) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Geopos Incomplete

func (b Builder) Geopos() (c Geopos) { _ = "STUB: not implemented"; return *new(Geopos) }

func (c Geopos) Key(key string) GeoposKey { _ = "STUB: not implemented"; return *new(GeoposKey) }

type GeoposKey Incomplete

func (c GeoposKey) Member(member ...string) GeoposMember {
	_ = "STUB: not implemented"
	return *new(GeoposMember)
}

func (c GeoposKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GeoposKey) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type GeoposMember Incomplete

func (c GeoposMember) Member(member ...string) GeoposMember {
	_ = "STUB: not implemented"
	return *new(GeoposMember)
}

func (c GeoposMember) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GeoposMember) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Georadius Incomplete

func (b Builder) Georadius() (c Georadius) { _ = "STUB: not implemented"; return *new(Georadius) }

func (c Georadius) Key(key string) GeoradiusKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusKey)
}

type GeoradiusCountAny Incomplete

func (c GeoradiusCountAny) Asc() GeoradiusOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusOrderAsc)
}

func (c GeoradiusCountAny) Desc() GeoradiusOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusOrderDesc)
}

func (c GeoradiusCountAny) Store(key string) GeoradiusStoreKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusStoreKey)
}

func (c GeoradiusCountAny) Storedist(key string) GeoradiusStoredistKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusStoredistKey)
}

func (c GeoradiusCountAny) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type GeoradiusCountCount Incomplete

func (c GeoradiusCountCount) Any() GeoradiusCountAny {
	_ = "STUB: not implemented"
	return *new(GeoradiusCountAny)
}

func (c GeoradiusCountCount) Asc() GeoradiusOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusOrderAsc)
}

func (c GeoradiusCountCount) Desc() GeoradiusOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusOrderDesc)
}

func (c GeoradiusCountCount) Store(key string) GeoradiusStoreKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusStoreKey)
}

func (c GeoradiusCountCount) Storedist(key string) GeoradiusStoredistKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusStoredistKey)
}

func (c GeoradiusCountCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type GeoradiusKey Incomplete

func (c GeoradiusKey) Longitude(longitude float64) GeoradiusLongitude {
	_ = "STUB: not implemented"
	return *new(GeoradiusLongitude)
}

type GeoradiusLatitude Incomplete

func (c GeoradiusLatitude) Radius(radius float64) GeoradiusRadius {
	_ = "STUB: not implemented"
	return *new(GeoradiusRadius)
}

type GeoradiusLongitude Incomplete

func (c GeoradiusLongitude) Latitude(latitude float64) GeoradiusLatitude {
	_ = "STUB: not implemented"
	return *new(GeoradiusLatitude)
}

type GeoradiusOrderAsc Incomplete

func (c GeoradiusOrderAsc) Store(key string) GeoradiusStoreKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusStoreKey)
}

func (c GeoradiusOrderAsc) Storedist(key string) GeoradiusStoredistKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusStoredistKey)
}

func (c GeoradiusOrderAsc) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type GeoradiusOrderDesc Incomplete

func (c GeoradiusOrderDesc) Store(key string) GeoradiusStoreKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusStoreKey)
}

func (c GeoradiusOrderDesc) Storedist(key string) GeoradiusStoredistKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusStoredistKey)
}

func (c GeoradiusOrderDesc) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type GeoradiusRadius Incomplete

func (c GeoradiusRadius) M() GeoradiusUnitM { _ = "STUB: not implemented"; return *new(GeoradiusUnitM) }

func (c GeoradiusRadius) Km() GeoradiusUnitKm {
	_ = "STUB: not implemented"
	return *new(GeoradiusUnitKm)
}

func (c GeoradiusRadius) Ft() GeoradiusUnitFt {
	_ = "STUB: not implemented"
	return *new(GeoradiusUnitFt)
}

func (c GeoradiusRadius) Mi() GeoradiusUnitMi {
	_ = "STUB: not implemented"
	return *new(GeoradiusUnitMi)
}

type GeoradiusRo Incomplete

func (b Builder) GeoradiusRo() (c GeoradiusRo) { _ = "STUB: not implemented"; return *new(GeoradiusRo) }

func (c GeoradiusRo) Key(key string) GeoradiusRoKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoKey)
}

type GeoradiusRoCountAny Incomplete

func (c GeoradiusRoCountAny) Asc() GeoradiusRoOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoOrderAsc)
}

func (c GeoradiusRoCountAny) Desc() GeoradiusRoOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoOrderDesc)
}

func (c GeoradiusRoCountAny) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GeoradiusRoCountAny) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type GeoradiusRoCountCount Incomplete

func (c GeoradiusRoCountCount) Any() GeoradiusRoCountAny {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoCountAny)
}

func (c GeoradiusRoCountCount) Asc() GeoradiusRoOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoOrderAsc)
}

func (c GeoradiusRoCountCount) Desc() GeoradiusRoOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoOrderDesc)
}

func (c GeoradiusRoCountCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GeoradiusRoCountCount) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type GeoradiusRoKey Incomplete

func (c GeoradiusRoKey) Longitude(longitude float64) GeoradiusRoLongitude {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoLongitude)
}

type GeoradiusRoLatitude Incomplete

func (c GeoradiusRoLatitude) Radius(radius float64) GeoradiusRoRadius {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoRadius)
}

type GeoradiusRoLongitude Incomplete

func (c GeoradiusRoLongitude) Latitude(latitude float64) GeoradiusRoLatitude {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoLatitude)
}

type GeoradiusRoOrderAsc Incomplete

func (c GeoradiusRoOrderAsc) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GeoradiusRoOrderAsc) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type GeoradiusRoOrderDesc Incomplete

func (c GeoradiusRoOrderDesc) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GeoradiusRoOrderDesc) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type GeoradiusRoRadius Incomplete

func (c GeoradiusRoRadius) M() GeoradiusRoUnitM {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoUnitM)
}

func (c GeoradiusRoRadius) Km() GeoradiusRoUnitKm {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoUnitKm)
}

func (c GeoradiusRoRadius) Ft() GeoradiusRoUnitFt {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoUnitFt)
}

func (c GeoradiusRoRadius) Mi() GeoradiusRoUnitMi {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoUnitMi)
}

type GeoradiusRoUnitFt Incomplete

func (c GeoradiusRoUnitFt) Withcoord() GeoradiusRoWithcoord {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoWithcoord)
}

func (c GeoradiusRoUnitFt) Withdist() GeoradiusRoWithdist {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoWithdist)
}

func (c GeoradiusRoUnitFt) Withhash() GeoradiusRoWithhash {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoWithhash)
}

func (c GeoradiusRoUnitFt) Count(count int64) GeoradiusRoCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoCountCount)
}

func (c GeoradiusRoUnitFt) Asc() GeoradiusRoOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoOrderAsc)
}

func (c GeoradiusRoUnitFt) Desc() GeoradiusRoOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoOrderDesc)
}

func (c GeoradiusRoUnitFt) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GeoradiusRoUnitFt) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type GeoradiusRoUnitKm Incomplete

func (c GeoradiusRoUnitKm) Withcoord() GeoradiusRoWithcoord {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoWithcoord)
}

func (c GeoradiusRoUnitKm) Withdist() GeoradiusRoWithdist {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoWithdist)
}

func (c GeoradiusRoUnitKm) Withhash() GeoradiusRoWithhash {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoWithhash)
}

func (c GeoradiusRoUnitKm) Count(count int64) GeoradiusRoCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoCountCount)
}

func (c GeoradiusRoUnitKm) Asc() GeoradiusRoOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoOrderAsc)
}

func (c GeoradiusRoUnitKm) Desc() GeoradiusRoOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoOrderDesc)
}

func (c GeoradiusRoUnitKm) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GeoradiusRoUnitKm) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type GeoradiusRoUnitM Incomplete

func (c GeoradiusRoUnitM) Withcoord() GeoradiusRoWithcoord {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoWithcoord)
}

func (c GeoradiusRoUnitM) Withdist() GeoradiusRoWithdist {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoWithdist)
}

func (c GeoradiusRoUnitM) Withhash() GeoradiusRoWithhash {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoWithhash)
}

func (c GeoradiusRoUnitM) Count(count int64) GeoradiusRoCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoCountCount)
}

func (c GeoradiusRoUnitM) Asc() GeoradiusRoOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoOrderAsc)
}

func (c GeoradiusRoUnitM) Desc() GeoradiusRoOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoOrderDesc)
}

func (c GeoradiusRoUnitM) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GeoradiusRoUnitM) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type GeoradiusRoUnitMi Incomplete

func (c GeoradiusRoUnitMi) Withcoord() GeoradiusRoWithcoord {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoWithcoord)
}

func (c GeoradiusRoUnitMi) Withdist() GeoradiusRoWithdist {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoWithdist)
}

func (c GeoradiusRoUnitMi) Withhash() GeoradiusRoWithhash {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoWithhash)
}

func (c GeoradiusRoUnitMi) Count(count int64) GeoradiusRoCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoCountCount)
}

func (c GeoradiusRoUnitMi) Asc() GeoradiusRoOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoOrderAsc)
}

func (c GeoradiusRoUnitMi) Desc() GeoradiusRoOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoOrderDesc)
}

func (c GeoradiusRoUnitMi) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GeoradiusRoUnitMi) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type GeoradiusRoWithcoord Incomplete

func (c GeoradiusRoWithcoord) Withdist() GeoradiusRoWithdist {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoWithdist)
}

func (c GeoradiusRoWithcoord) Withhash() GeoradiusRoWithhash {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoWithhash)
}

func (c GeoradiusRoWithcoord) Count(count int64) GeoradiusRoCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoCountCount)
}

func (c GeoradiusRoWithcoord) Asc() GeoradiusRoOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoOrderAsc)
}

func (c GeoradiusRoWithcoord) Desc() GeoradiusRoOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoOrderDesc)
}

func (c GeoradiusRoWithcoord) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GeoradiusRoWithcoord) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type GeoradiusRoWithdist Incomplete

func (c GeoradiusRoWithdist) Withhash() GeoradiusRoWithhash {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoWithhash)
}

func (c GeoradiusRoWithdist) Count(count int64) GeoradiusRoCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoCountCount)
}

func (c GeoradiusRoWithdist) Asc() GeoradiusRoOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoOrderAsc)
}

func (c GeoradiusRoWithdist) Desc() GeoradiusRoOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoOrderDesc)
}

func (c GeoradiusRoWithdist) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GeoradiusRoWithdist) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type GeoradiusRoWithhash Incomplete

func (c GeoradiusRoWithhash) Count(count int64) GeoradiusRoCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoCountCount)
}

func (c GeoradiusRoWithhash) Asc() GeoradiusRoOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoOrderAsc)
}

func (c GeoradiusRoWithhash) Desc() GeoradiusRoOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusRoOrderDesc)
}

func (c GeoradiusRoWithhash) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GeoradiusRoWithhash) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type GeoradiusStoreKey Incomplete

func (c GeoradiusStoreKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type GeoradiusStoredistKey Incomplete

func (c GeoradiusStoredistKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type GeoradiusUnitFt Incomplete

func (c GeoradiusUnitFt) Withcoord() GeoradiusWithcoord {
	_ = "STUB: not implemented"
	return *new(GeoradiusWithcoord)
}

func (c GeoradiusUnitFt) Withdist() GeoradiusWithdist {
	_ = "STUB: not implemented"
	return *new(GeoradiusWithdist)
}

func (c GeoradiusUnitFt) Withhash() GeoradiusWithhash {
	_ = "STUB: not implemented"
	return *new(GeoradiusWithhash)
}

func (c GeoradiusUnitFt) Count(count int64) GeoradiusCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusCountCount)
}

func (c GeoradiusUnitFt) Asc() GeoradiusOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusOrderAsc)
}

func (c GeoradiusUnitFt) Desc() GeoradiusOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusOrderDesc)
}

func (c GeoradiusUnitFt) Store(key string) GeoradiusStoreKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusStoreKey)
}

func (c GeoradiusUnitFt) Storedist(key string) GeoradiusStoredistKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusStoredistKey)
}

func (c GeoradiusUnitFt) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type GeoradiusUnitKm Incomplete

func (c GeoradiusUnitKm) Withcoord() GeoradiusWithcoord {
	_ = "STUB: not implemented"
	return *new(GeoradiusWithcoord)
}

func (c GeoradiusUnitKm) Withdist() GeoradiusWithdist {
	_ = "STUB: not implemented"
	return *new(GeoradiusWithdist)
}

func (c GeoradiusUnitKm) Withhash() GeoradiusWithhash {
	_ = "STUB: not implemented"
	return *new(GeoradiusWithhash)
}

func (c GeoradiusUnitKm) Count(count int64) GeoradiusCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusCountCount)
}

func (c GeoradiusUnitKm) Asc() GeoradiusOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusOrderAsc)
}

func (c GeoradiusUnitKm) Desc() GeoradiusOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusOrderDesc)
}

func (c GeoradiusUnitKm) Store(key string) GeoradiusStoreKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusStoreKey)
}

func (c GeoradiusUnitKm) Storedist(key string) GeoradiusStoredistKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusStoredistKey)
}

func (c GeoradiusUnitKm) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type GeoradiusUnitM Incomplete

func (c GeoradiusUnitM) Withcoord() GeoradiusWithcoord {
	_ = "STUB: not implemented"
	return *new(GeoradiusWithcoord)
}

func (c GeoradiusUnitM) Withdist() GeoradiusWithdist {
	_ = "STUB: not implemented"
	return *new(GeoradiusWithdist)
}

func (c GeoradiusUnitM) Withhash() GeoradiusWithhash {
	_ = "STUB: not implemented"
	return *new(GeoradiusWithhash)
}

func (c GeoradiusUnitM) Count(count int64) GeoradiusCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusCountCount)
}

func (c GeoradiusUnitM) Asc() GeoradiusOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusOrderAsc)
}

func (c GeoradiusUnitM) Desc() GeoradiusOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusOrderDesc)
}

func (c GeoradiusUnitM) Store(key string) GeoradiusStoreKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusStoreKey)
}

func (c GeoradiusUnitM) Storedist(key string) GeoradiusStoredistKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusStoredistKey)
}

func (c GeoradiusUnitM) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type GeoradiusUnitMi Incomplete

func (c GeoradiusUnitMi) Withcoord() GeoradiusWithcoord {
	_ = "STUB: not implemented"
	return *new(GeoradiusWithcoord)
}

func (c GeoradiusUnitMi) Withdist() GeoradiusWithdist {
	_ = "STUB: not implemented"
	return *new(GeoradiusWithdist)
}

func (c GeoradiusUnitMi) Withhash() GeoradiusWithhash {
	_ = "STUB: not implemented"
	return *new(GeoradiusWithhash)
}

func (c GeoradiusUnitMi) Count(count int64) GeoradiusCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusCountCount)
}

func (c GeoradiusUnitMi) Asc() GeoradiusOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusOrderAsc)
}

func (c GeoradiusUnitMi) Desc() GeoradiusOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusOrderDesc)
}

func (c GeoradiusUnitMi) Store(key string) GeoradiusStoreKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusStoreKey)
}

func (c GeoradiusUnitMi) Storedist(key string) GeoradiusStoredistKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusStoredistKey)
}

func (c GeoradiusUnitMi) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type GeoradiusWithcoord Incomplete

func (c GeoradiusWithcoord) Withdist() GeoradiusWithdist {
	_ = "STUB: not implemented"
	return *new(GeoradiusWithdist)
}

func (c GeoradiusWithcoord) Withhash() GeoradiusWithhash {
	_ = "STUB: not implemented"
	return *new(GeoradiusWithhash)
}

func (c GeoradiusWithcoord) Count(count int64) GeoradiusCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusCountCount)
}

func (c GeoradiusWithcoord) Asc() GeoradiusOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusOrderAsc)
}

func (c GeoradiusWithcoord) Desc() GeoradiusOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusOrderDesc)
}

func (c GeoradiusWithcoord) Store(key string) GeoradiusStoreKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusStoreKey)
}

func (c GeoradiusWithcoord) Storedist(key string) GeoradiusStoredistKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusStoredistKey)
}

func (c GeoradiusWithcoord) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type GeoradiusWithdist Incomplete

func (c GeoradiusWithdist) Withhash() GeoradiusWithhash {
	_ = "STUB: not implemented"
	return *new(GeoradiusWithhash)
}

func (c GeoradiusWithdist) Count(count int64) GeoradiusCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusCountCount)
}

func (c GeoradiusWithdist) Asc() GeoradiusOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusOrderAsc)
}

func (c GeoradiusWithdist) Desc() GeoradiusOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusOrderDesc)
}

func (c GeoradiusWithdist) Store(key string) GeoradiusStoreKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusStoreKey)
}

func (c GeoradiusWithdist) Storedist(key string) GeoradiusStoredistKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusStoredistKey)
}

func (c GeoradiusWithdist) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type GeoradiusWithhash Incomplete

func (c GeoradiusWithhash) Count(count int64) GeoradiusCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusCountCount)
}

func (c GeoradiusWithhash) Asc() GeoradiusOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusOrderAsc)
}

func (c GeoradiusWithhash) Desc() GeoradiusOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusOrderDesc)
}

func (c GeoradiusWithhash) Store(key string) GeoradiusStoreKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusStoreKey)
}

func (c GeoradiusWithhash) Storedist(key string) GeoradiusStoredistKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusStoredistKey)
}

func (c GeoradiusWithhash) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Georadiusbymember Incomplete

func (b Builder) Georadiusbymember() (c Georadiusbymember) {
	_ = "STUB: not implemented"
	return *new(Georadiusbymember)
}

func (c Georadiusbymember) Key(key string) GeoradiusbymemberKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberKey)
}

type GeoradiusbymemberCountAny Incomplete

func (c GeoradiusbymemberCountAny) Asc() GeoradiusbymemberOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberOrderAsc)
}

func (c GeoradiusbymemberCountAny) Desc() GeoradiusbymemberOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberOrderDesc)
}

func (c GeoradiusbymemberCountAny) Store(key string) GeoradiusbymemberStoreStoreKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberStoreStoreKey)
}

func (c GeoradiusbymemberCountAny) Storedist(key string) GeoradiusbymemberStoreStoredistKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberStoreStoredistKey)
}

func (c GeoradiusbymemberCountAny) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GeoradiusbymemberCountCount Incomplete

func (c GeoradiusbymemberCountCount) Any() GeoradiusbymemberCountAny {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberCountAny)
}

func (c GeoradiusbymemberCountCount) Asc() GeoradiusbymemberOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberOrderAsc)
}

func (c GeoradiusbymemberCountCount) Desc() GeoradiusbymemberOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberOrderDesc)
}

func (c GeoradiusbymemberCountCount) Store(key string) GeoradiusbymemberStoreStoreKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberStoreStoreKey)
}

func (c GeoradiusbymemberCountCount) Storedist(key string) GeoradiusbymemberStoreStoredistKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberStoreStoredistKey)
}

func (c GeoradiusbymemberCountCount) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GeoradiusbymemberKey Incomplete

func (c GeoradiusbymemberKey) Member(member string) GeoradiusbymemberMember {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberMember)
}

type GeoradiusbymemberMember Incomplete

func (c GeoradiusbymemberMember) Radius(radius float64) GeoradiusbymemberRadius {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRadius)
}

type GeoradiusbymemberOrderAsc Incomplete

func (c GeoradiusbymemberOrderAsc) Store(key string) GeoradiusbymemberStoreStoreKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberStoreStoreKey)
}

func (c GeoradiusbymemberOrderAsc) Storedist(key string) GeoradiusbymemberStoreStoredistKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberStoreStoredistKey)
}

func (c GeoradiusbymemberOrderAsc) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GeoradiusbymemberOrderDesc Incomplete

func (c GeoradiusbymemberOrderDesc) Store(key string) GeoradiusbymemberStoreStoreKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberStoreStoreKey)
}

func (c GeoradiusbymemberOrderDesc) Storedist(key string) GeoradiusbymemberStoreStoredistKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberStoreStoredistKey)
}

func (c GeoradiusbymemberOrderDesc) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GeoradiusbymemberRadius Incomplete

func (c GeoradiusbymemberRadius) M() GeoradiusbymemberUnitM {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberUnitM)
}

func (c GeoradiusbymemberRadius) Km() GeoradiusbymemberUnitKm {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberUnitKm)
}

func (c GeoradiusbymemberRadius) Ft() GeoradiusbymemberUnitFt {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberUnitFt)
}

func (c GeoradiusbymemberRadius) Mi() GeoradiusbymemberUnitMi {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberUnitMi)
}

type GeoradiusbymemberRo Incomplete

func (b Builder) GeoradiusbymemberRo() (c GeoradiusbymemberRo) {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRo)
}

func (c GeoradiusbymemberRo) Key(key string) GeoradiusbymemberRoKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoKey)
}

type GeoradiusbymemberRoCountAny Incomplete

func (c GeoradiusbymemberRoCountAny) Asc() GeoradiusbymemberRoOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoOrderAsc)
}

func (c GeoradiusbymemberRoCountAny) Desc() GeoradiusbymemberRoOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoOrderDesc)
}

func (c GeoradiusbymemberRoCountAny) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c GeoradiusbymemberRoCountAny) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type GeoradiusbymemberRoCountCount Incomplete

func (c GeoradiusbymemberRoCountCount) Any() GeoradiusbymemberRoCountAny {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoCountAny)
}

func (c GeoradiusbymemberRoCountCount) Asc() GeoradiusbymemberRoOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoOrderAsc)
}

func (c GeoradiusbymemberRoCountCount) Desc() GeoradiusbymemberRoOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoOrderDesc)
}

func (c GeoradiusbymemberRoCountCount) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c GeoradiusbymemberRoCountCount) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type GeoradiusbymemberRoKey Incomplete

func (c GeoradiusbymemberRoKey) Member(member string) GeoradiusbymemberRoMember {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoMember)
}

type GeoradiusbymemberRoMember Incomplete

func (c GeoradiusbymemberRoMember) Radius(radius float64) GeoradiusbymemberRoRadius {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoRadius)
}

type GeoradiusbymemberRoOrderAsc Incomplete

func (c GeoradiusbymemberRoOrderAsc) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c GeoradiusbymemberRoOrderAsc) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type GeoradiusbymemberRoOrderDesc Incomplete

func (c GeoradiusbymemberRoOrderDesc) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c GeoradiusbymemberRoOrderDesc) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type GeoradiusbymemberRoRadius Incomplete

func (c GeoradiusbymemberRoRadius) M() GeoradiusbymemberRoUnitM {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoUnitM)
}

func (c GeoradiusbymemberRoRadius) Km() GeoradiusbymemberRoUnitKm {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoUnitKm)
}

func (c GeoradiusbymemberRoRadius) Ft() GeoradiusbymemberRoUnitFt {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoUnitFt)
}

func (c GeoradiusbymemberRoRadius) Mi() GeoradiusbymemberRoUnitMi {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoUnitMi)
}

type GeoradiusbymemberRoUnitFt Incomplete

func (c GeoradiusbymemberRoUnitFt) Withcoord() GeoradiusbymemberRoWithcoord {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoWithcoord)
}

func (c GeoradiusbymemberRoUnitFt) Withdist() GeoradiusbymemberRoWithdist {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoWithdist)
}

func (c GeoradiusbymemberRoUnitFt) Withhash() GeoradiusbymemberRoWithhash {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoWithhash)
}

func (c GeoradiusbymemberRoUnitFt) Count(count int64) GeoradiusbymemberRoCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoCountCount)
}

func (c GeoradiusbymemberRoUnitFt) Asc() GeoradiusbymemberRoOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoOrderAsc)
}

func (c GeoradiusbymemberRoUnitFt) Desc() GeoradiusbymemberRoOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoOrderDesc)
}

func (c GeoradiusbymemberRoUnitFt) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c GeoradiusbymemberRoUnitFt) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type GeoradiusbymemberRoUnitKm Incomplete

func (c GeoradiusbymemberRoUnitKm) Withcoord() GeoradiusbymemberRoWithcoord {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoWithcoord)
}

func (c GeoradiusbymemberRoUnitKm) Withdist() GeoradiusbymemberRoWithdist {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoWithdist)
}

func (c GeoradiusbymemberRoUnitKm) Withhash() GeoradiusbymemberRoWithhash {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoWithhash)
}

func (c GeoradiusbymemberRoUnitKm) Count(count int64) GeoradiusbymemberRoCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoCountCount)
}

func (c GeoradiusbymemberRoUnitKm) Asc() GeoradiusbymemberRoOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoOrderAsc)
}

func (c GeoradiusbymemberRoUnitKm) Desc() GeoradiusbymemberRoOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoOrderDesc)
}

func (c GeoradiusbymemberRoUnitKm) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c GeoradiusbymemberRoUnitKm) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type GeoradiusbymemberRoUnitM Incomplete

func (c GeoradiusbymemberRoUnitM) Withcoord() GeoradiusbymemberRoWithcoord {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoWithcoord)
}

func (c GeoradiusbymemberRoUnitM) Withdist() GeoradiusbymemberRoWithdist {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoWithdist)
}

func (c GeoradiusbymemberRoUnitM) Withhash() GeoradiusbymemberRoWithhash {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoWithhash)
}

func (c GeoradiusbymemberRoUnitM) Count(count int64) GeoradiusbymemberRoCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoCountCount)
}

func (c GeoradiusbymemberRoUnitM) Asc() GeoradiusbymemberRoOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoOrderAsc)
}

func (c GeoradiusbymemberRoUnitM) Desc() GeoradiusbymemberRoOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoOrderDesc)
}

func (c GeoradiusbymemberRoUnitM) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c GeoradiusbymemberRoUnitM) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type GeoradiusbymemberRoUnitMi Incomplete

func (c GeoradiusbymemberRoUnitMi) Withcoord() GeoradiusbymemberRoWithcoord {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoWithcoord)
}

func (c GeoradiusbymemberRoUnitMi) Withdist() GeoradiusbymemberRoWithdist {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoWithdist)
}

func (c GeoradiusbymemberRoUnitMi) Withhash() GeoradiusbymemberRoWithhash {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoWithhash)
}

func (c GeoradiusbymemberRoUnitMi) Count(count int64) GeoradiusbymemberRoCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoCountCount)
}

func (c GeoradiusbymemberRoUnitMi) Asc() GeoradiusbymemberRoOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoOrderAsc)
}

func (c GeoradiusbymemberRoUnitMi) Desc() GeoradiusbymemberRoOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoOrderDesc)
}

func (c GeoradiusbymemberRoUnitMi) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c GeoradiusbymemberRoUnitMi) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type GeoradiusbymemberRoWithcoord Incomplete

func (c GeoradiusbymemberRoWithcoord) Withdist() GeoradiusbymemberRoWithdist {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoWithdist)
}

func (c GeoradiusbymemberRoWithcoord) Withhash() GeoradiusbymemberRoWithhash {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoWithhash)
}

func (c GeoradiusbymemberRoWithcoord) Count(count int64) GeoradiusbymemberRoCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoCountCount)
}

func (c GeoradiusbymemberRoWithcoord) Asc() GeoradiusbymemberRoOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoOrderAsc)
}

func (c GeoradiusbymemberRoWithcoord) Desc() GeoradiusbymemberRoOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoOrderDesc)
}

func (c GeoradiusbymemberRoWithcoord) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c GeoradiusbymemberRoWithcoord) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type GeoradiusbymemberRoWithdist Incomplete

func (c GeoradiusbymemberRoWithdist) Withhash() GeoradiusbymemberRoWithhash {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoWithhash)
}

func (c GeoradiusbymemberRoWithdist) Count(count int64) GeoradiusbymemberRoCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoCountCount)
}

func (c GeoradiusbymemberRoWithdist) Asc() GeoradiusbymemberRoOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoOrderAsc)
}

func (c GeoradiusbymemberRoWithdist) Desc() GeoradiusbymemberRoOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoOrderDesc)
}

func (c GeoradiusbymemberRoWithdist) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c GeoradiusbymemberRoWithdist) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type GeoradiusbymemberRoWithhash Incomplete

func (c GeoradiusbymemberRoWithhash) Count(count int64) GeoradiusbymemberRoCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoCountCount)
}

func (c GeoradiusbymemberRoWithhash) Asc() GeoradiusbymemberRoOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoOrderAsc)
}

func (c GeoradiusbymemberRoWithhash) Desc() GeoradiusbymemberRoOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberRoOrderDesc)
}

func (c GeoradiusbymemberRoWithhash) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c GeoradiusbymemberRoWithhash) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type GeoradiusbymemberStoreStoreKey Incomplete

func (c GeoradiusbymemberStoreStoreKey) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GeoradiusbymemberStoreStoredistKey Incomplete

func (c GeoradiusbymemberStoreStoredistKey) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GeoradiusbymemberUnitFt Incomplete

func (c GeoradiusbymemberUnitFt) Withcoord() GeoradiusbymemberWithcoord {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberWithcoord)
}

func (c GeoradiusbymemberUnitFt) Withdist() GeoradiusbymemberWithdist {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberWithdist)
}

func (c GeoradiusbymemberUnitFt) Withhash() GeoradiusbymemberWithhash {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberWithhash)
}

func (c GeoradiusbymemberUnitFt) Count(count int64) GeoradiusbymemberCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberCountCount)
}

func (c GeoradiusbymemberUnitFt) Asc() GeoradiusbymemberOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberOrderAsc)
}

func (c GeoradiusbymemberUnitFt) Desc() GeoradiusbymemberOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberOrderDesc)
}

func (c GeoradiusbymemberUnitFt) Store(key string) GeoradiusbymemberStoreStoreKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberStoreStoreKey)
}

func (c GeoradiusbymemberUnitFt) Storedist(key string) GeoradiusbymemberStoreStoredistKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberStoreStoredistKey)
}

func (c GeoradiusbymemberUnitFt) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GeoradiusbymemberUnitKm Incomplete

func (c GeoradiusbymemberUnitKm) Withcoord() GeoradiusbymemberWithcoord {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberWithcoord)
}

func (c GeoradiusbymemberUnitKm) Withdist() GeoradiusbymemberWithdist {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberWithdist)
}

func (c GeoradiusbymemberUnitKm) Withhash() GeoradiusbymemberWithhash {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberWithhash)
}

func (c GeoradiusbymemberUnitKm) Count(count int64) GeoradiusbymemberCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberCountCount)
}

func (c GeoradiusbymemberUnitKm) Asc() GeoradiusbymemberOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberOrderAsc)
}

func (c GeoradiusbymemberUnitKm) Desc() GeoradiusbymemberOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberOrderDesc)
}

func (c GeoradiusbymemberUnitKm) Store(key string) GeoradiusbymemberStoreStoreKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberStoreStoreKey)
}

func (c GeoradiusbymemberUnitKm) Storedist(key string) GeoradiusbymemberStoreStoredistKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberStoreStoredistKey)
}

func (c GeoradiusbymemberUnitKm) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GeoradiusbymemberUnitM Incomplete

func (c GeoradiusbymemberUnitM) Withcoord() GeoradiusbymemberWithcoord {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberWithcoord)
}

func (c GeoradiusbymemberUnitM) Withdist() GeoradiusbymemberWithdist {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberWithdist)
}

func (c GeoradiusbymemberUnitM) Withhash() GeoradiusbymemberWithhash {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberWithhash)
}

func (c GeoradiusbymemberUnitM) Count(count int64) GeoradiusbymemberCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberCountCount)
}

func (c GeoradiusbymemberUnitM) Asc() GeoradiusbymemberOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberOrderAsc)
}

func (c GeoradiusbymemberUnitM) Desc() GeoradiusbymemberOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberOrderDesc)
}

func (c GeoradiusbymemberUnitM) Store(key string) GeoradiusbymemberStoreStoreKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberStoreStoreKey)
}

func (c GeoradiusbymemberUnitM) Storedist(key string) GeoradiusbymemberStoreStoredistKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberStoreStoredistKey)
}

func (c GeoradiusbymemberUnitM) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GeoradiusbymemberUnitMi Incomplete

func (c GeoradiusbymemberUnitMi) Withcoord() GeoradiusbymemberWithcoord {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberWithcoord)
}

func (c GeoradiusbymemberUnitMi) Withdist() GeoradiusbymemberWithdist {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberWithdist)
}

func (c GeoradiusbymemberUnitMi) Withhash() GeoradiusbymemberWithhash {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberWithhash)
}

func (c GeoradiusbymemberUnitMi) Count(count int64) GeoradiusbymemberCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberCountCount)
}

func (c GeoradiusbymemberUnitMi) Asc() GeoradiusbymemberOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberOrderAsc)
}

func (c GeoradiusbymemberUnitMi) Desc() GeoradiusbymemberOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberOrderDesc)
}

func (c GeoradiusbymemberUnitMi) Store(key string) GeoradiusbymemberStoreStoreKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberStoreStoreKey)
}

func (c GeoradiusbymemberUnitMi) Storedist(key string) GeoradiusbymemberStoreStoredistKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberStoreStoredistKey)
}

func (c GeoradiusbymemberUnitMi) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GeoradiusbymemberWithcoord Incomplete

func (c GeoradiusbymemberWithcoord) Withdist() GeoradiusbymemberWithdist {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberWithdist)
}

func (c GeoradiusbymemberWithcoord) Withhash() GeoradiusbymemberWithhash {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberWithhash)
}

func (c GeoradiusbymemberWithcoord) Count(count int64) GeoradiusbymemberCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberCountCount)
}

func (c GeoradiusbymemberWithcoord) Asc() GeoradiusbymemberOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberOrderAsc)
}

func (c GeoradiusbymemberWithcoord) Desc() GeoradiusbymemberOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberOrderDesc)
}

func (c GeoradiusbymemberWithcoord) Store(key string) GeoradiusbymemberStoreStoreKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberStoreStoreKey)
}

func (c GeoradiusbymemberWithcoord) Storedist(key string) GeoradiusbymemberStoreStoredistKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberStoreStoredistKey)
}

func (c GeoradiusbymemberWithcoord) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GeoradiusbymemberWithdist Incomplete

func (c GeoradiusbymemberWithdist) Withhash() GeoradiusbymemberWithhash {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberWithhash)
}

func (c GeoradiusbymemberWithdist) Count(count int64) GeoradiusbymemberCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberCountCount)
}

func (c GeoradiusbymemberWithdist) Asc() GeoradiusbymemberOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberOrderAsc)
}

func (c GeoradiusbymemberWithdist) Desc() GeoradiusbymemberOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberOrderDesc)
}

func (c GeoradiusbymemberWithdist) Store(key string) GeoradiusbymemberStoreStoreKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberStoreStoreKey)
}

func (c GeoradiusbymemberWithdist) Storedist(key string) GeoradiusbymemberStoreStoredistKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberStoreStoredistKey)
}

func (c GeoradiusbymemberWithdist) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GeoradiusbymemberWithhash Incomplete

func (c GeoradiusbymemberWithhash) Count(count int64) GeoradiusbymemberCountCount {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberCountCount)
}

func (c GeoradiusbymemberWithhash) Asc() GeoradiusbymemberOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberOrderAsc)
}

func (c GeoradiusbymemberWithhash) Desc() GeoradiusbymemberOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberOrderDesc)
}

func (c GeoradiusbymemberWithhash) Store(key string) GeoradiusbymemberStoreStoreKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberStoreStoreKey)
}

func (c GeoradiusbymemberWithhash) Storedist(key string) GeoradiusbymemberStoreStoredistKey {
	_ = "STUB: not implemented"
	return *new(GeoradiusbymemberStoreStoredistKey)
}

func (c GeoradiusbymemberWithhash) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type Geosearch Incomplete

func (b Builder) Geosearch() (c Geosearch) { _ = "STUB: not implemented"; return *new(Geosearch) }

func (c Geosearch) Key(key string) GeosearchKey {
	_ = "STUB: not implemented"
	return *new(GeosearchKey)
}

type GeosearchCircleBoxBybox Incomplete

func (c GeosearchCircleBoxBybox) Height(height float64) GeosearchCircleBoxHeight {
	_ = "STUB: not implemented"
	return *new(GeosearchCircleBoxHeight)
}

type GeosearchCircleBoxHeight Incomplete

func (c GeosearchCircleBoxHeight) M() GeosearchCircleBoxUnitM {
	_ = "STUB: not implemented"
	return *new(GeosearchCircleBoxUnitM)
}

func (c GeosearchCircleBoxHeight) Km() GeosearchCircleBoxUnitKm {
	_ = "STUB: not implemented"
	return *new(GeosearchCircleBoxUnitKm)
}

func (c GeosearchCircleBoxHeight) Ft() GeosearchCircleBoxUnitFt {
	_ = "STUB: not implemented"
	return *new(GeosearchCircleBoxUnitFt)
}

func (c GeosearchCircleBoxHeight) Mi() GeosearchCircleBoxUnitMi {
	_ = "STUB: not implemented"
	return *new(GeosearchCircleBoxUnitMi)
}

type GeosearchCircleBoxUnitFt Incomplete

func (c GeosearchCircleBoxUnitFt) NumVertices(numVertices int64) GeosearchCirclePolygonNumVertices {
	_ = "STUB: not implemented"
	return *new(GeosearchCirclePolygonNumVertices)
}

func (c GeosearchCircleBoxUnitFt) Asc() GeosearchOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeosearchOrderAsc)
}

func (c GeosearchCircleBoxUnitFt) Desc() GeosearchOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeosearchOrderDesc)
}

func (c GeosearchCircleBoxUnitFt) Count(count int64) GeosearchCountCount {
	_ = "STUB: not implemented"
	return *new(GeosearchCountCount)
}

func (c GeosearchCircleBoxUnitFt) Withcoord() GeosearchWithcoord {
	_ = "STUB: not implemented"
	return *new(GeosearchWithcoord)
}

func (c GeosearchCircleBoxUnitFt) Withdist() GeosearchWithdist {
	_ = "STUB: not implemented"
	return *new(GeosearchWithdist)
}

func (c GeosearchCircleBoxUnitFt) Withhash() GeosearchWithhash {
	_ = "STUB: not implemented"
	return *new(GeosearchWithhash)
}

func (c GeosearchCircleBoxUnitFt) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c GeosearchCircleBoxUnitFt) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type GeosearchCircleBoxUnitKm Incomplete

func (c GeosearchCircleBoxUnitKm) NumVertices(numVertices int64) GeosearchCirclePolygonNumVertices {
	_ = "STUB: not implemented"
	return *new(GeosearchCirclePolygonNumVertices)
}

func (c GeosearchCircleBoxUnitKm) Asc() GeosearchOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeosearchOrderAsc)
}

func (c GeosearchCircleBoxUnitKm) Desc() GeosearchOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeosearchOrderDesc)
}

func (c GeosearchCircleBoxUnitKm) Count(count int64) GeosearchCountCount {
	_ = "STUB: not implemented"
	return *new(GeosearchCountCount)
}

func (c GeosearchCircleBoxUnitKm) Withcoord() GeosearchWithcoord {
	_ = "STUB: not implemented"
	return *new(GeosearchWithcoord)
}

func (c GeosearchCircleBoxUnitKm) Withdist() GeosearchWithdist {
	_ = "STUB: not implemented"
	return *new(GeosearchWithdist)
}

func (c GeosearchCircleBoxUnitKm) Withhash() GeosearchWithhash {
	_ = "STUB: not implemented"
	return *new(GeosearchWithhash)
}

func (c GeosearchCircleBoxUnitKm) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c GeosearchCircleBoxUnitKm) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type GeosearchCircleBoxUnitM Incomplete

func (c GeosearchCircleBoxUnitM) NumVertices(numVertices int64) GeosearchCirclePolygonNumVertices {
	_ = "STUB: not implemented"
	return *new(GeosearchCirclePolygonNumVertices)
}

func (c GeosearchCircleBoxUnitM) Asc() GeosearchOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeosearchOrderAsc)
}

func (c GeosearchCircleBoxUnitM) Desc() GeosearchOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeosearchOrderDesc)
}

func (c GeosearchCircleBoxUnitM) Count(count int64) GeosearchCountCount {
	_ = "STUB: not implemented"
	return *new(GeosearchCountCount)
}

func (c GeosearchCircleBoxUnitM) Withcoord() GeosearchWithcoord {
	_ = "STUB: not implemented"
	return *new(GeosearchWithcoord)
}

func (c GeosearchCircleBoxUnitM) Withdist() GeosearchWithdist {
	_ = "STUB: not implemented"
	return *new(GeosearchWithdist)
}

func (c GeosearchCircleBoxUnitM) Withhash() GeosearchWithhash {
	_ = "STUB: not implemented"
	return *new(GeosearchWithhash)
}

func (c GeosearchCircleBoxUnitM) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c GeosearchCircleBoxUnitM) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type GeosearchCircleBoxUnitMi Incomplete

func (c GeosearchCircleBoxUnitMi) NumVertices(numVertices int64) GeosearchCirclePolygonNumVertices {
	_ = "STUB: not implemented"
	return *new(GeosearchCirclePolygonNumVertices)
}

func (c GeosearchCircleBoxUnitMi) Asc() GeosearchOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeosearchOrderAsc)
}

func (c GeosearchCircleBoxUnitMi) Desc() GeosearchOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeosearchOrderDesc)
}

func (c GeosearchCircleBoxUnitMi) Count(count int64) GeosearchCountCount {
	_ = "STUB: not implemented"
	return *new(GeosearchCountCount)
}

func (c GeosearchCircleBoxUnitMi) Withcoord() GeosearchWithcoord {
	_ = "STUB: not implemented"
	return *new(GeosearchWithcoord)
}

func (c GeosearchCircleBoxUnitMi) Withdist() GeosearchWithdist {
	_ = "STUB: not implemented"
	return *new(GeosearchWithdist)
}

func (c GeosearchCircleBoxUnitMi) Withhash() GeosearchWithhash {
	_ = "STUB: not implemented"
	return *new(GeosearchWithhash)
}

func (c GeosearchCircleBoxUnitMi) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c GeosearchCircleBoxUnitMi) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type GeosearchCircleCircleByradius Incomplete

func (c GeosearchCircleCircleByradius) M() GeosearchCircleCircleUnitM {
	_ = "STUB: not implemented"
	return *new(GeosearchCircleCircleUnitM)
}

func (c GeosearchCircleCircleByradius) Km() GeosearchCircleCircleUnitKm {
	_ = "STUB: not implemented"
	return *new(GeosearchCircleCircleUnitKm)
}

func (c GeosearchCircleCircleByradius) Ft() GeosearchCircleCircleUnitFt {
	_ = "STUB: not implemented"
	return *new(GeosearchCircleCircleUnitFt)
}

func (c GeosearchCircleCircleByradius) Mi() GeosearchCircleCircleUnitMi {
	_ = "STUB: not implemented"
	return *new(GeosearchCircleCircleUnitMi)
}

type GeosearchCircleCircleUnitFt Incomplete

func (c GeosearchCircleCircleUnitFt) Bybox(width float64) GeosearchCircleBoxBybox {
	_ = "STUB: not implemented"
	return *new(GeosearchCircleBoxBybox)
}

func (c GeosearchCircleCircleUnitFt) NumVertices(numVertices int64) GeosearchCirclePolygonNumVertices {
	_ = "STUB: not implemented"
	return *new(GeosearchCirclePolygonNumVertices)
}

func (c GeosearchCircleCircleUnitFt) Asc() GeosearchOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeosearchOrderAsc)
}

func (c GeosearchCircleCircleUnitFt) Desc() GeosearchOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeosearchOrderDesc)
}

func (c GeosearchCircleCircleUnitFt) Count(count int64) GeosearchCountCount {
	_ = "STUB: not implemented"
	return *new(GeosearchCountCount)
}

func (c GeosearchCircleCircleUnitFt) Withcoord() GeosearchWithcoord {
	_ = "STUB: not implemented"
	return *new(GeosearchWithcoord)
}

func (c GeosearchCircleCircleUnitFt) Withdist() GeosearchWithdist {
	_ = "STUB: not implemented"
	return *new(GeosearchWithdist)
}

func (c GeosearchCircleCircleUnitFt) Withhash() GeosearchWithhash {
	_ = "STUB: not implemented"
	return *new(GeosearchWithhash)
}

func (c GeosearchCircleCircleUnitFt) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c GeosearchCircleCircleUnitFt) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type GeosearchCircleCircleUnitKm Incomplete

func (c GeosearchCircleCircleUnitKm) Bybox(width float64) GeosearchCircleBoxBybox {
	_ = "STUB: not implemented"
	return *new(GeosearchCircleBoxBybox)
}

func (c GeosearchCircleCircleUnitKm) NumVertices(numVertices int64) GeosearchCirclePolygonNumVertices {
	_ = "STUB: not implemented"
	return *new(GeosearchCirclePolygonNumVertices)
}

func (c GeosearchCircleCircleUnitKm) Asc() GeosearchOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeosearchOrderAsc)
}

func (c GeosearchCircleCircleUnitKm) Desc() GeosearchOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeosearchOrderDesc)
}

func (c GeosearchCircleCircleUnitKm) Count(count int64) GeosearchCountCount {
	_ = "STUB: not implemented"
	return *new(GeosearchCountCount)
}

func (c GeosearchCircleCircleUnitKm) Withcoord() GeosearchWithcoord {
	_ = "STUB: not implemented"
	return *new(GeosearchWithcoord)
}

func (c GeosearchCircleCircleUnitKm) Withdist() GeosearchWithdist {
	_ = "STUB: not implemented"
	return *new(GeosearchWithdist)
}

func (c GeosearchCircleCircleUnitKm) Withhash() GeosearchWithhash {
	_ = "STUB: not implemented"
	return *new(GeosearchWithhash)
}

func (c GeosearchCircleCircleUnitKm) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c GeosearchCircleCircleUnitKm) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type GeosearchCircleCircleUnitM Incomplete

func (c GeosearchCircleCircleUnitM) Bybox(width float64) GeosearchCircleBoxBybox {
	_ = "STUB: not implemented"
	return *new(GeosearchCircleBoxBybox)
}

func (c GeosearchCircleCircleUnitM) NumVertices(numVertices int64) GeosearchCirclePolygonNumVertices {
	_ = "STUB: not implemented"
	return *new(GeosearchCirclePolygonNumVertices)
}

func (c GeosearchCircleCircleUnitM) Asc() GeosearchOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeosearchOrderAsc)
}

func (c GeosearchCircleCircleUnitM) Desc() GeosearchOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeosearchOrderDesc)
}

func (c GeosearchCircleCircleUnitM) Count(count int64) GeosearchCountCount {
	_ = "STUB: not implemented"
	return *new(GeosearchCountCount)
}

func (c GeosearchCircleCircleUnitM) Withcoord() GeosearchWithcoord {
	_ = "STUB: not implemented"
	return *new(GeosearchWithcoord)
}

func (c GeosearchCircleCircleUnitM) Withdist() GeosearchWithdist {
	_ = "STUB: not implemented"
	return *new(GeosearchWithdist)
}

func (c GeosearchCircleCircleUnitM) Withhash() GeosearchWithhash {
	_ = "STUB: not implemented"
	return *new(GeosearchWithhash)
}

func (c GeosearchCircleCircleUnitM) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c GeosearchCircleCircleUnitM) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type GeosearchCircleCircleUnitMi Incomplete

func (c GeosearchCircleCircleUnitMi) Bybox(width float64) GeosearchCircleBoxBybox {
	_ = "STUB: not implemented"
	return *new(GeosearchCircleBoxBybox)
}

func (c GeosearchCircleCircleUnitMi) NumVertices(numVertices int64) GeosearchCirclePolygonNumVertices {
	_ = "STUB: not implemented"
	return *new(GeosearchCirclePolygonNumVertices)
}

func (c GeosearchCircleCircleUnitMi) Asc() GeosearchOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeosearchOrderAsc)
}

func (c GeosearchCircleCircleUnitMi) Desc() GeosearchOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeosearchOrderDesc)
}

func (c GeosearchCircleCircleUnitMi) Count(count int64) GeosearchCountCount {
	_ = "STUB: not implemented"
	return *new(GeosearchCountCount)
}

func (c GeosearchCircleCircleUnitMi) Withcoord() GeosearchWithcoord {
	_ = "STUB: not implemented"
	return *new(GeosearchWithcoord)
}

func (c GeosearchCircleCircleUnitMi) Withdist() GeosearchWithdist {
	_ = "STUB: not implemented"
	return *new(GeosearchWithdist)
}

func (c GeosearchCircleCircleUnitMi) Withhash() GeosearchWithhash {
	_ = "STUB: not implemented"
	return *new(GeosearchWithhash)
}

func (c GeosearchCircleCircleUnitMi) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c GeosearchCircleCircleUnitMi) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type GeosearchCirclePolygonNumVertices Incomplete

func (c GeosearchCirclePolygonNumVertices) Longitude(longitude float64) GeosearchCirclePolygonVerticesLongitude {
	_ = "STUB: not implemented"
	return *new(GeosearchCirclePolygonVerticesLongitude)
}

type GeosearchCirclePolygonVerticesLatitude Incomplete

func (c GeosearchCirclePolygonVerticesLatitude) Longitude(longitude float64) GeosearchCirclePolygonVerticesLongitude {
	_ = "STUB: not implemented"
	return *new(GeosearchCirclePolygonVerticesLongitude)
}

func (c GeosearchCirclePolygonVerticesLatitude) Asc() GeosearchOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeosearchOrderAsc)
}

func (c GeosearchCirclePolygonVerticesLatitude) Desc() GeosearchOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeosearchOrderDesc)
}

func (c GeosearchCirclePolygonVerticesLatitude) Count(count int64) GeosearchCountCount {
	_ = "STUB: not implemented"
	return *new(GeosearchCountCount)
}

func (c GeosearchCirclePolygonVerticesLatitude) Withcoord() GeosearchWithcoord {
	_ = "STUB: not implemented"
	return *new(GeosearchWithcoord)
}

func (c GeosearchCirclePolygonVerticesLatitude) Withdist() GeosearchWithdist {
	_ = "STUB: not implemented"
	return *new(GeosearchWithdist)
}

func (c GeosearchCirclePolygonVerticesLatitude) Withhash() GeosearchWithhash {
	_ = "STUB: not implemented"
	return *new(GeosearchWithhash)
}

func (c GeosearchCirclePolygonVerticesLatitude) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

func (c GeosearchCirclePolygonVerticesLatitude) Cache() Cacheable {
	_ = "STUB: not implemented"
	return *new(Cacheable)
}

type GeosearchCirclePolygonVerticesLongitude Incomplete

func (c GeosearchCirclePolygonVerticesLongitude) Latitude(latitude float64) GeosearchCirclePolygonVerticesLatitude {
	_ = "STUB: not implemented"
	return *new(GeosearchCirclePolygonVerticesLatitude)
}

type GeosearchCountAny Incomplete

func (c GeosearchCountAny) Withcoord() GeosearchWithcoord {
	_ = "STUB: not implemented"
	return *new(GeosearchWithcoord)
}

func (c GeosearchCountAny) Withdist() GeosearchWithdist {
	_ = "STUB: not implemented"
	return *new(GeosearchWithdist)
}

func (c GeosearchCountAny) Withhash() GeosearchWithhash {
	_ = "STUB: not implemented"
	return *new(GeosearchWithhash)
}

func (c GeosearchCountAny) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GeosearchCountAny) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type GeosearchCountCount Incomplete

func (c GeosearchCountCount) Any() GeosearchCountAny {
	_ = "STUB: not implemented"
	return *new(GeosearchCountAny)
}

func (c GeosearchCountCount) Withcoord() GeosearchWithcoord {
	_ = "STUB: not implemented"
	return *new(GeosearchWithcoord)
}

func (c GeosearchCountCount) Withdist() GeosearchWithdist {
	_ = "STUB: not implemented"
	return *new(GeosearchWithdist)
}

func (c GeosearchCountCount) Withhash() GeosearchWithhash {
	_ = "STUB: not implemented"
	return *new(GeosearchWithhash)
}

func (c GeosearchCountCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GeosearchCountCount) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type GeosearchFrommemberFromlonlat Incomplete

func (c GeosearchFrommemberFromlonlat) Byradius(radius float64) GeosearchCircleCircleByradius {
	_ = "STUB: not implemented"
	return *new(GeosearchCircleCircleByradius)
}

func (c GeosearchFrommemberFromlonlat) Bybox(width float64) GeosearchCircleBoxBybox {
	_ = "STUB: not implemented"
	return *new(GeosearchCircleBoxBybox)
}

func (c GeosearchFrommemberFromlonlat) NumVertices(numVertices int64) GeosearchCirclePolygonNumVertices {
	_ = "STUB: not implemented"
	return *new(GeosearchCirclePolygonNumVertices)
}

type GeosearchFrommemberFrommember Incomplete

func (c GeosearchFrommemberFrommember) Fromlonlat(longitude float64, latitude float64) GeosearchFrommemberFromlonlat {
	_ = "STUB: not implemented"
	return *new(GeosearchFrommemberFromlonlat)
}

func (c GeosearchFrommemberFrommember) Byradius(radius float64) GeosearchCircleCircleByradius {
	_ = "STUB: not implemented"
	return *new(GeosearchCircleCircleByradius)
}

func (c GeosearchFrommemberFrommember) Bybox(width float64) GeosearchCircleBoxBybox {
	_ = "STUB: not implemented"
	return *new(GeosearchCircleBoxBybox)
}

func (c GeosearchFrommemberFrommember) NumVertices(numVertices int64) GeosearchCirclePolygonNumVertices {
	_ = "STUB: not implemented"
	return *new(GeosearchCirclePolygonNumVertices)
}

type GeosearchKey Incomplete

func (c GeosearchKey) Frommember(member string) GeosearchFrommemberFrommember {
	_ = "STUB: not implemented"
	return *new(GeosearchFrommemberFrommember)
}

func (c GeosearchKey) Fromlonlat(longitude float64, latitude float64) GeosearchFrommemberFromlonlat {
	_ = "STUB: not implemented"
	return *new(GeosearchFrommemberFromlonlat)
}

type GeosearchOrderAsc Incomplete

func (c GeosearchOrderAsc) Count(count int64) GeosearchCountCount {
	_ = "STUB: not implemented"
	return *new(GeosearchCountCount)
}

func (c GeosearchOrderAsc) Withcoord() GeosearchWithcoord {
	_ = "STUB: not implemented"
	return *new(GeosearchWithcoord)
}

func (c GeosearchOrderAsc) Withdist() GeosearchWithdist {
	_ = "STUB: not implemented"
	return *new(GeosearchWithdist)
}

func (c GeosearchOrderAsc) Withhash() GeosearchWithhash {
	_ = "STUB: not implemented"
	return *new(GeosearchWithhash)
}

func (c GeosearchOrderAsc) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GeosearchOrderAsc) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type GeosearchOrderDesc Incomplete

func (c GeosearchOrderDesc) Count(count int64) GeosearchCountCount {
	_ = "STUB: not implemented"
	return *new(GeosearchCountCount)
}

func (c GeosearchOrderDesc) Withcoord() GeosearchWithcoord {
	_ = "STUB: not implemented"
	return *new(GeosearchWithcoord)
}

func (c GeosearchOrderDesc) Withdist() GeosearchWithdist {
	_ = "STUB: not implemented"
	return *new(GeosearchWithdist)
}

func (c GeosearchOrderDesc) Withhash() GeosearchWithhash {
	_ = "STUB: not implemented"
	return *new(GeosearchWithhash)
}

func (c GeosearchOrderDesc) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GeosearchOrderDesc) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type GeosearchWithcoord Incomplete

func (c GeosearchWithcoord) Withdist() GeosearchWithdist {
	_ = "STUB: not implemented"
	return *new(GeosearchWithdist)
}

func (c GeosearchWithcoord) Withhash() GeosearchWithhash {
	_ = "STUB: not implemented"
	return *new(GeosearchWithhash)
}

func (c GeosearchWithcoord) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GeosearchWithcoord) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type GeosearchWithdist Incomplete

func (c GeosearchWithdist) Withhash() GeosearchWithhash {
	_ = "STUB: not implemented"
	return *new(GeosearchWithhash)
}

func (c GeosearchWithdist) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GeosearchWithdist) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type GeosearchWithhash Incomplete

func (c GeosearchWithhash) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GeosearchWithhash) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type Geosearchstore Incomplete

func (b Builder) Geosearchstore() (c Geosearchstore) {
	_ = "STUB: not implemented"
	return *new(Geosearchstore)
}

func (c Geosearchstore) Destination(destination string) GeosearchstoreDestination {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreDestination)
}

type GeosearchstoreCircleBoxBybox Incomplete

func (c GeosearchstoreCircleBoxBybox) Height(height float64) GeosearchstoreCircleBoxHeight {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCircleBoxHeight)
}

type GeosearchstoreCircleBoxHeight Incomplete

func (c GeosearchstoreCircleBoxHeight) M() GeosearchstoreCircleBoxUnitM {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCircleBoxUnitM)
}

func (c GeosearchstoreCircleBoxHeight) Km() GeosearchstoreCircleBoxUnitKm {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCircleBoxUnitKm)
}

func (c GeosearchstoreCircleBoxHeight) Ft() GeosearchstoreCircleBoxUnitFt {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCircleBoxUnitFt)
}

func (c GeosearchstoreCircleBoxHeight) Mi() GeosearchstoreCircleBoxUnitMi {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCircleBoxUnitMi)
}

type GeosearchstoreCircleBoxUnitFt Incomplete

func (c GeosearchstoreCircleBoxUnitFt) NumVertices(numVertices int64) GeosearchstoreCirclePolygonNumVertices {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCirclePolygonNumVertices)
}

func (c GeosearchstoreCircleBoxUnitFt) Asc() GeosearchstoreOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreOrderAsc)
}

func (c GeosearchstoreCircleBoxUnitFt) Desc() GeosearchstoreOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreOrderDesc)
}

func (c GeosearchstoreCircleBoxUnitFt) Count(count int64) GeosearchstoreCountCount {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCountCount)
}

func (c GeosearchstoreCircleBoxUnitFt) Storedist() GeosearchstoreStoredist {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreStoredist)
}

func (c GeosearchstoreCircleBoxUnitFt) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GeosearchstoreCircleBoxUnitKm Incomplete

func (c GeosearchstoreCircleBoxUnitKm) NumVertices(numVertices int64) GeosearchstoreCirclePolygonNumVertices {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCirclePolygonNumVertices)
}

func (c GeosearchstoreCircleBoxUnitKm) Asc() GeosearchstoreOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreOrderAsc)
}

func (c GeosearchstoreCircleBoxUnitKm) Desc() GeosearchstoreOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreOrderDesc)
}

func (c GeosearchstoreCircleBoxUnitKm) Count(count int64) GeosearchstoreCountCount {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCountCount)
}

func (c GeosearchstoreCircleBoxUnitKm) Storedist() GeosearchstoreStoredist {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreStoredist)
}

func (c GeosearchstoreCircleBoxUnitKm) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GeosearchstoreCircleBoxUnitM Incomplete

func (c GeosearchstoreCircleBoxUnitM) NumVertices(numVertices int64) GeosearchstoreCirclePolygonNumVertices {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCirclePolygonNumVertices)
}

func (c GeosearchstoreCircleBoxUnitM) Asc() GeosearchstoreOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreOrderAsc)
}

func (c GeosearchstoreCircleBoxUnitM) Desc() GeosearchstoreOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreOrderDesc)
}

func (c GeosearchstoreCircleBoxUnitM) Count(count int64) GeosearchstoreCountCount {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCountCount)
}

func (c GeosearchstoreCircleBoxUnitM) Storedist() GeosearchstoreStoredist {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreStoredist)
}

func (c GeosearchstoreCircleBoxUnitM) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GeosearchstoreCircleBoxUnitMi Incomplete

func (c GeosearchstoreCircleBoxUnitMi) NumVertices(numVertices int64) GeosearchstoreCirclePolygonNumVertices {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCirclePolygonNumVertices)
}

func (c GeosearchstoreCircleBoxUnitMi) Asc() GeosearchstoreOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreOrderAsc)
}

func (c GeosearchstoreCircleBoxUnitMi) Desc() GeosearchstoreOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreOrderDesc)
}

func (c GeosearchstoreCircleBoxUnitMi) Count(count int64) GeosearchstoreCountCount {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCountCount)
}

func (c GeosearchstoreCircleBoxUnitMi) Storedist() GeosearchstoreStoredist {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreStoredist)
}

func (c GeosearchstoreCircleBoxUnitMi) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GeosearchstoreCircleCircleByradius Incomplete

func (c GeosearchstoreCircleCircleByradius) M() GeosearchstoreCircleCircleUnitM {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCircleCircleUnitM)
}

func (c GeosearchstoreCircleCircleByradius) Km() GeosearchstoreCircleCircleUnitKm {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCircleCircleUnitKm)
}

func (c GeosearchstoreCircleCircleByradius) Ft() GeosearchstoreCircleCircleUnitFt {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCircleCircleUnitFt)
}

func (c GeosearchstoreCircleCircleByradius) Mi() GeosearchstoreCircleCircleUnitMi {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCircleCircleUnitMi)
}

type GeosearchstoreCircleCircleUnitFt Incomplete

func (c GeosearchstoreCircleCircleUnitFt) Bybox(width float64) GeosearchstoreCircleBoxBybox {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCircleBoxBybox)
}

func (c GeosearchstoreCircleCircleUnitFt) NumVertices(numVertices int64) GeosearchstoreCirclePolygonNumVertices {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCirclePolygonNumVertices)
}

func (c GeosearchstoreCircleCircleUnitFt) Asc() GeosearchstoreOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreOrderAsc)
}

func (c GeosearchstoreCircleCircleUnitFt) Desc() GeosearchstoreOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreOrderDesc)
}

func (c GeosearchstoreCircleCircleUnitFt) Count(count int64) GeosearchstoreCountCount {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCountCount)
}

func (c GeosearchstoreCircleCircleUnitFt) Storedist() GeosearchstoreStoredist {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreStoredist)
}

func (c GeosearchstoreCircleCircleUnitFt) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GeosearchstoreCircleCircleUnitKm Incomplete

func (c GeosearchstoreCircleCircleUnitKm) Bybox(width float64) GeosearchstoreCircleBoxBybox {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCircleBoxBybox)
}

func (c GeosearchstoreCircleCircleUnitKm) NumVertices(numVertices int64) GeosearchstoreCirclePolygonNumVertices {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCirclePolygonNumVertices)
}

func (c GeosearchstoreCircleCircleUnitKm) Asc() GeosearchstoreOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreOrderAsc)
}

func (c GeosearchstoreCircleCircleUnitKm) Desc() GeosearchstoreOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreOrderDesc)
}

func (c GeosearchstoreCircleCircleUnitKm) Count(count int64) GeosearchstoreCountCount {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCountCount)
}

func (c GeosearchstoreCircleCircleUnitKm) Storedist() GeosearchstoreStoredist {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreStoredist)
}

func (c GeosearchstoreCircleCircleUnitKm) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GeosearchstoreCircleCircleUnitM Incomplete

func (c GeosearchstoreCircleCircleUnitM) Bybox(width float64) GeosearchstoreCircleBoxBybox {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCircleBoxBybox)
}

func (c GeosearchstoreCircleCircleUnitM) NumVertices(numVertices int64) GeosearchstoreCirclePolygonNumVertices {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCirclePolygonNumVertices)
}

func (c GeosearchstoreCircleCircleUnitM) Asc() GeosearchstoreOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreOrderAsc)
}

func (c GeosearchstoreCircleCircleUnitM) Desc() GeosearchstoreOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreOrderDesc)
}

func (c GeosearchstoreCircleCircleUnitM) Count(count int64) GeosearchstoreCountCount {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCountCount)
}

func (c GeosearchstoreCircleCircleUnitM) Storedist() GeosearchstoreStoredist {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreStoredist)
}

func (c GeosearchstoreCircleCircleUnitM) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GeosearchstoreCircleCircleUnitMi Incomplete

func (c GeosearchstoreCircleCircleUnitMi) Bybox(width float64) GeosearchstoreCircleBoxBybox {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCircleBoxBybox)
}

func (c GeosearchstoreCircleCircleUnitMi) NumVertices(numVertices int64) GeosearchstoreCirclePolygonNumVertices {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCirclePolygonNumVertices)
}

func (c GeosearchstoreCircleCircleUnitMi) Asc() GeosearchstoreOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreOrderAsc)
}

func (c GeosearchstoreCircleCircleUnitMi) Desc() GeosearchstoreOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreOrderDesc)
}

func (c GeosearchstoreCircleCircleUnitMi) Count(count int64) GeosearchstoreCountCount {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCountCount)
}

func (c GeosearchstoreCircleCircleUnitMi) Storedist() GeosearchstoreStoredist {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreStoredist)
}

func (c GeosearchstoreCircleCircleUnitMi) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GeosearchstoreCirclePolygonNumVertices Incomplete

func (c GeosearchstoreCirclePolygonNumVertices) Longitude(longitude float64) GeosearchstoreCirclePolygonVerticesLongitude {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCirclePolygonVerticesLongitude)
}

type GeosearchstoreCirclePolygonVerticesLatitude Incomplete

func (c GeosearchstoreCirclePolygonVerticesLatitude) Longitude(longitude float64) GeosearchstoreCirclePolygonVerticesLongitude {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCirclePolygonVerticesLongitude)
}

func (c GeosearchstoreCirclePolygonVerticesLatitude) Asc() GeosearchstoreOrderAsc {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreOrderAsc)
}

func (c GeosearchstoreCirclePolygonVerticesLatitude) Desc() GeosearchstoreOrderDesc {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreOrderDesc)
}

func (c GeosearchstoreCirclePolygonVerticesLatitude) Count(count int64) GeosearchstoreCountCount {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCountCount)
}

func (c GeosearchstoreCirclePolygonVerticesLatitude) Storedist() GeosearchstoreStoredist {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreStoredist)
}

func (c GeosearchstoreCirclePolygonVerticesLatitude) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GeosearchstoreCirclePolygonVerticesLongitude Incomplete

func (c GeosearchstoreCirclePolygonVerticesLongitude) Latitude(latitude float64) GeosearchstoreCirclePolygonVerticesLatitude {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCirclePolygonVerticesLatitude)
}

type GeosearchstoreCountAny Incomplete

func (c GeosearchstoreCountAny) Storedist() GeosearchstoreStoredist {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreStoredist)
}

func (c GeosearchstoreCountAny) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GeosearchstoreCountCount Incomplete

func (c GeosearchstoreCountCount) Any() GeosearchstoreCountAny {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCountAny)
}

func (c GeosearchstoreCountCount) Storedist() GeosearchstoreStoredist {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreStoredist)
}

func (c GeosearchstoreCountCount) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GeosearchstoreDestination Incomplete

func (c GeosearchstoreDestination) Source(source string) GeosearchstoreSource {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreSource)
}

type GeosearchstoreFrommemberFromlonlat Incomplete

func (c GeosearchstoreFrommemberFromlonlat) Byradius(radius float64) GeosearchstoreCircleCircleByradius {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCircleCircleByradius)
}

func (c GeosearchstoreFrommemberFromlonlat) Bybox(width float64) GeosearchstoreCircleBoxBybox {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCircleBoxBybox)
}

func (c GeosearchstoreFrommemberFromlonlat) NumVertices(numVertices int64) GeosearchstoreCirclePolygonNumVertices {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCirclePolygonNumVertices)
}

type GeosearchstoreFrommemberFrommember Incomplete

func (c GeosearchstoreFrommemberFrommember) Fromlonlat(longitude float64, latitude float64) GeosearchstoreFrommemberFromlonlat {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreFrommemberFromlonlat)
}

func (c GeosearchstoreFrommemberFrommember) Byradius(radius float64) GeosearchstoreCircleCircleByradius {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCircleCircleByradius)
}

func (c GeosearchstoreFrommemberFrommember) Bybox(width float64) GeosearchstoreCircleBoxBybox {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCircleBoxBybox)
}

func (c GeosearchstoreFrommemberFrommember) NumVertices(numVertices int64) GeosearchstoreCirclePolygonNumVertices {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCirclePolygonNumVertices)
}

type GeosearchstoreOrderAsc Incomplete

func (c GeosearchstoreOrderAsc) Count(count int64) GeosearchstoreCountCount {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCountCount)
}

func (c GeosearchstoreOrderAsc) Storedist() GeosearchstoreStoredist {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreStoredist)
}

func (c GeosearchstoreOrderAsc) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GeosearchstoreOrderDesc Incomplete

func (c GeosearchstoreOrderDesc) Count(count int64) GeosearchstoreCountCount {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreCountCount)
}

func (c GeosearchstoreOrderDesc) Storedist() GeosearchstoreStoredist {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreStoredist)
}

func (c GeosearchstoreOrderDesc) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GeosearchstoreSource Incomplete

func (c GeosearchstoreSource) Frommember(member string) GeosearchstoreFrommemberFrommember {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreFrommemberFrommember)
}

func (c GeosearchstoreSource) Fromlonlat(longitude float64, latitude float64) GeosearchstoreFrommemberFromlonlat {
	_ = "STUB: not implemented"
	return *new(GeosearchstoreFrommemberFromlonlat)
}

type GeosearchstoreStoredist Incomplete

func (c GeosearchstoreStoredist) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

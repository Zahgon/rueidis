// Code generated DO NOT EDIT

package cmds

type Vadd Incomplete

func (b Builder) Vadd() (c Vadd) { _ = "STUB: not implemented"; return *new(Vadd) }

func (c Vadd) Key(key string) VaddKey { _ = "STUB: not implemented"; return *new(VaddKey) }

type VaddCas Incomplete

func (c VaddCas) Noquant() VaddQuantizationNoquant {
	_ = "STUB: not implemented"
	return *new(VaddQuantizationNoquant)
}

func (c VaddCas) Q8() VaddQuantizationQ8 {
	_ = "STUB: not implemented"
	return *new(VaddQuantizationQ8)
}

func (c VaddCas) Bin() VaddQuantizationBin {
	_ = "STUB: not implemented"
	return *new(VaddQuantizationBin)
}

func (c VaddCas) Ef(buildExplorationFactor int64) VaddEf {
	_ = "STUB: not implemented"
	return *new(VaddEf)
}

func (c VaddCas) Setattr(attributes string) VaddSetattr {
	_ = "STUB: not implemented"
	return *new(VaddSetattr)
}

func (c VaddCas) M(numlinks int64) VaddM { _ = "STUB: not implemented"; return *new(VaddM) }

func (c VaddCas) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type VaddEf Incomplete

func (c VaddEf) Setattr(attributes string) VaddSetattr {
	_ = "STUB: not implemented"
	return *new(VaddSetattr)
}

func (c VaddEf) M(numlinks int64) VaddM { _ = "STUB: not implemented"; return *new(VaddM) }

func (c VaddEf) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type VaddElement Incomplete

func (c VaddElement) Cas() VaddCas { _ = "STUB: not implemented"; return *new(VaddCas) }

func (c VaddElement) Noquant() VaddQuantizationNoquant {
	_ = "STUB: not implemented"
	return *new(VaddQuantizationNoquant)
}

func (c VaddElement) Q8() VaddQuantizationQ8 {
	_ = "STUB: not implemented"
	return *new(VaddQuantizationQ8)
}

func (c VaddElement) Bin() VaddQuantizationBin {
	_ = "STUB: not implemented"
	return *new(VaddQuantizationBin)
}

func (c VaddElement) Ef(buildExplorationFactor int64) VaddEf {
	_ = "STUB: not implemented"
	return *new(VaddEf)
}

func (c VaddElement) Setattr(attributes string) VaddSetattr {
	_ = "STUB: not implemented"
	return *new(VaddSetattr)
}

func (c VaddElement) M(numlinks int64) VaddM { _ = "STUB: not implemented"; return *new(VaddM) }

func (c VaddElement) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type VaddKey Incomplete

func (c VaddKey) Reduce(dim int64) VaddReduce { _ = "STUB: not implemented"; return *new(VaddReduce) }

func (c VaddKey) Fp32() VaddNumFp32Fp32 { _ = "STUB: not implemented"; return *new(VaddNumFp32Fp32) }

func (c VaddKey) Values(num int64) VaddNumValuesValues {
	_ = "STUB: not implemented"
	return *new(VaddNumValuesValues)
}

type VaddM Incomplete

func (c VaddM) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type VaddNumFp32Fp32 Incomplete

func (c VaddNumFp32Fp32) Vector(vector string) VaddNumFp32Vector {
	_ = "STUB: not implemented"
	return *new(VaddNumFp32Vector)
}

type VaddNumFp32Vector Incomplete

func (c VaddNumFp32Vector) Element(element string) VaddElement {
	_ = "STUB: not implemented"
	return *new(VaddElement)
}

type VaddNumValuesValues Incomplete

func (c VaddNumValuesValues) Vector(vector ...float32) VaddNumValuesVector {
	_ = "STUB: not implemented"
	return *new(VaddNumValuesVector)
}

type VaddNumValuesVector Incomplete

func (c VaddNumValuesVector) Vector(vector ...float32) VaddNumValuesVector {
	_ = "STUB: not implemented"
	return *new(VaddNumValuesVector)
}

func (c VaddNumValuesVector) Element(element string) VaddElement {
	_ = "STUB: not implemented"
	return *new(VaddElement)
}

type VaddQuantizationBin Incomplete

func (c VaddQuantizationBin) Ef(buildExplorationFactor int64) VaddEf {
	_ = "STUB: not implemented"
	return *new(VaddEf)
}

func (c VaddQuantizationBin) Setattr(attributes string) VaddSetattr {
	_ = "STUB: not implemented"
	return *new(VaddSetattr)
}

func (c VaddQuantizationBin) M(numlinks int64) VaddM { _ = "STUB: not implemented"; return *new(VaddM) }

func (c VaddQuantizationBin) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type VaddQuantizationNoquant Incomplete

func (c VaddQuantizationNoquant) Ef(buildExplorationFactor int64) VaddEf {
	_ = "STUB: not implemented"
	return *new(VaddEf)
}

func (c VaddQuantizationNoquant) Setattr(attributes string) VaddSetattr {
	_ = "STUB: not implemented"
	return *new(VaddSetattr)
}

func (c VaddQuantizationNoquant) M(numlinks int64) VaddM {
	_ = "STUB: not implemented"
	return *new(VaddM)
}

func (c VaddQuantizationNoquant) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type VaddQuantizationQ8 Incomplete

func (c VaddQuantizationQ8) Ef(buildExplorationFactor int64) VaddEf {
	_ = "STUB: not implemented"
	return *new(VaddEf)
}

func (c VaddQuantizationQ8) Setattr(attributes string) VaddSetattr {
	_ = "STUB: not implemented"
	return *new(VaddSetattr)
}

func (c VaddQuantizationQ8) M(numlinks int64) VaddM { _ = "STUB: not implemented"; return *new(VaddM) }

func (c VaddQuantizationQ8) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type VaddReduce Incomplete

func (c VaddReduce) Fp32() VaddNumFp32Fp32 { _ = "STUB: not implemented"; return *new(VaddNumFp32Fp32) }

func (c VaddReduce) Values(num int64) VaddNumValuesValues {
	_ = "STUB: not implemented"
	return *new(VaddNumValuesValues)
}

type VaddSetattr Incomplete

func (c VaddSetattr) M(numlinks int64) VaddM { _ = "STUB: not implemented"; return *new(VaddM) }

func (c VaddSetattr) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Vcard Incomplete

func (b Builder) Vcard() (c Vcard) { _ = "STUB: not implemented"; return *new(Vcard) }

func (c Vcard) Key(key string) VcardKey { _ = "STUB: not implemented"; return *new(VcardKey) }

type VcardKey Incomplete

func (c VcardKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Vdim Incomplete

func (b Builder) Vdim() (c Vdim) { _ = "STUB: not implemented"; return *new(Vdim) }

func (c Vdim) Key(key string) VdimKey { _ = "STUB: not implemented"; return *new(VdimKey) }

type VdimKey Incomplete

func (c VdimKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Vemb Incomplete

func (b Builder) Vemb() (c Vemb) { _ = "STUB: not implemented"; return *new(Vemb) }

func (c Vemb) Key(key string) VembKey { _ = "STUB: not implemented"; return *new(VembKey) }

type VembElement Incomplete

func (c VembElement) Raw() VembRaw { _ = "STUB: not implemented"; return *new(VembRaw) }

func (c VembElement) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type VembKey Incomplete

func (c VembKey) Element(element string) VembElement {
	_ = "STUB: not implemented"
	return *new(VembElement)
}

type VembRaw Incomplete

func (c VembRaw) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Vgetattr Incomplete

func (b Builder) Vgetattr() (c Vgetattr) { _ = "STUB: not implemented"; return *new(Vgetattr) }

func (c Vgetattr) Key(key string) VgetattrKey { _ = "STUB: not implemented"; return *new(VgetattrKey) }

type VgetattrElement Incomplete

func (c VgetattrElement) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type VgetattrKey Incomplete

func (c VgetattrKey) Element(element string) VgetattrElement {
	_ = "STUB: not implemented"
	return *new(VgetattrElement)
}

type Vinfo Incomplete

func (b Builder) Vinfo() (c Vinfo) { _ = "STUB: not implemented"; return *new(Vinfo) }

func (c Vinfo) Key(key string) VinfoKey { _ = "STUB: not implemented"; return *new(VinfoKey) }

type VinfoKey Incomplete

func (c VinfoKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Vlinks Incomplete

func (b Builder) Vlinks() (c Vlinks) { _ = "STUB: not implemented"; return *new(Vlinks) }

func (c Vlinks) Key(key string) VlinksKey { _ = "STUB: not implemented"; return *new(VlinksKey) }

type VlinksElement Incomplete

func (c VlinksElement) Withscores() VlinksWithscores {
	_ = "STUB: not implemented"
	return *new(VlinksWithscores)
}

func (c VlinksElement) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type VlinksKey Incomplete

func (c VlinksKey) Element(element string) VlinksElement {
	_ = "STUB: not implemented"
	return *new(VlinksElement)
}

type VlinksWithscores Incomplete

func (c VlinksWithscores) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Vrandmember Incomplete

func (b Builder) Vrandmember() (c Vrandmember) { _ = "STUB: not implemented"; return *new(Vrandmember) }

func (c Vrandmember) Key(key string) VrandmemberKey {
	_ = "STUB: not implemented"
	return *new(VrandmemberKey)
}

type VrandmemberCount Incomplete

func (c VrandmemberCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type VrandmemberKey Incomplete

func (c VrandmemberKey) Count(count int64) VrandmemberCount {
	_ = "STUB: not implemented"
	return *new(VrandmemberCount)
}

func (c VrandmemberKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Vrem Incomplete

func (b Builder) Vrem() (c Vrem) { _ = "STUB: not implemented"; return *new(Vrem) }

func (c Vrem) Key(key string) VremKey { _ = "STUB: not implemented"; return *new(VremKey) }

type VremElement Incomplete

func (c VremElement) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type VremKey Incomplete

func (c VremKey) Element(element string) VremElement {
	_ = "STUB: not implemented"
	return *new(VremElement)
}

type Vsetattr Incomplete

func (b Builder) Vsetattr() (c Vsetattr) { _ = "STUB: not implemented"; return *new(Vsetattr) }

func (c Vsetattr) Key(key string) VsetattrKey { _ = "STUB: not implemented"; return *new(VsetattrKey) }

type VsetattrElement Incomplete

func (c VsetattrElement) Json(json string) VsetattrJson {
	_ = "STUB: not implemented"
	return *new(VsetattrJson)
}

type VsetattrJson Incomplete

func (c VsetattrJson) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type VsetattrKey Incomplete

func (c VsetattrKey) Element(element string) VsetattrElement {
	_ = "STUB: not implemented"
	return *new(VsetattrElement)
}

type Vsim Incomplete

func (b Builder) Vsim() (c Vsim) { _ = "STUB: not implemented"; return *new(Vsim) }

func (c Vsim) Key(key string) VsimKey { _ = "STUB: not implemented"; return *new(VsimKey) }

type VsimCount Incomplete

func (c VsimCount) Epsilon(delta float32) VsimEpsilon {
	_ = "STUB: not implemented"
	return *new(VsimEpsilon)
}

func (c VsimCount) Ef(searchExplorationFactor int64) VsimEf {
	_ = "STUB: not implemented"
	return *new(VsimEf)
}

func (c VsimCount) Filter(expression string) VsimFilter {
	_ = "STUB: not implemented"
	return *new(VsimFilter)
}

func (c VsimCount) FilterEf(maxFilteringEffort int64) VsimFilterEf {
	_ = "STUB: not implemented"
	return *new(VsimFilterEf)
}

func (c VsimCount) Truth() VsimTruth { _ = "STUB: not implemented"; return *new(VsimTruth) }

func (c VsimCount) Nothread() VsimNothread { _ = "STUB: not implemented"; return *new(VsimNothread) }

func (c VsimCount) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type VsimEf Incomplete

func (c VsimEf) Filter(expression string) VsimFilter {
	_ = "STUB: not implemented"
	return *new(VsimFilter)
}

func (c VsimEf) FilterEf(maxFilteringEffort int64) VsimFilterEf {
	_ = "STUB: not implemented"
	return *new(VsimFilterEf)
}

func (c VsimEf) Truth() VsimTruth { _ = "STUB: not implemented"; return *new(VsimTruth) }

func (c VsimEf) Nothread() VsimNothread { _ = "STUB: not implemented"; return *new(VsimNothread) }

func (c VsimEf) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type VsimEpsilon Incomplete

func (c VsimEpsilon) Ef(searchExplorationFactor int64) VsimEf {
	_ = "STUB: not implemented"
	return *new(VsimEf)
}

func (c VsimEpsilon) Filter(expression string) VsimFilter {
	_ = "STUB: not implemented"
	return *new(VsimFilter)
}

func (c VsimEpsilon) FilterEf(maxFilteringEffort int64) VsimFilterEf {
	_ = "STUB: not implemented"
	return *new(VsimFilterEf)
}

func (c VsimEpsilon) Truth() VsimTruth { _ = "STUB: not implemented"; return *new(VsimTruth) }

func (c VsimEpsilon) Nothread() VsimNothread { _ = "STUB: not implemented"; return *new(VsimNothread) }

func (c VsimEpsilon) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type VsimFilter Incomplete

func (c VsimFilter) FilterEf(maxFilteringEffort int64) VsimFilterEf {
	_ = "STUB: not implemented"
	return *new(VsimFilterEf)
}

func (c VsimFilter) Truth() VsimTruth { _ = "STUB: not implemented"; return *new(VsimTruth) }

func (c VsimFilter) Nothread() VsimNothread { _ = "STUB: not implemented"; return *new(VsimNothread) }

func (c VsimFilter) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type VsimFilterEf Incomplete

func (c VsimFilterEf) Truth() VsimTruth { _ = "STUB: not implemented"; return *new(VsimTruth) }

func (c VsimFilterEf) Nothread() VsimNothread { _ = "STUB: not implemented"; return *new(VsimNothread) }

func (c VsimFilterEf) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type VsimKey Incomplete

func (c VsimKey) Ele() VsimQueryTypeEleEle {
	_ = "STUB: not implemented"
	return *new(VsimQueryTypeEleEle)
}

func (c VsimKey) Fp32() VsimQueryTypeFp32Fp32 {
	_ = "STUB: not implemented"
	return *new(VsimQueryTypeFp32Fp32)
}

func (c VsimKey) Values(num int64) VsimQueryTypeValuesValues {
	_ = "STUB: not implemented"
	return *new(VsimQueryTypeValuesValues)
}

type VsimNothread Incomplete

func (c VsimNothread) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type VsimQueryTypeEleEle Incomplete

func (c VsimQueryTypeEleEle) Element(element string) VsimQueryTypeEleElement {
	_ = "STUB: not implemented"
	return *new(VsimQueryTypeEleElement)
}

type VsimQueryTypeEleElement Incomplete

func (c VsimQueryTypeEleElement) Withscores() VsimWithscores {
	_ = "STUB: not implemented"
	return *new(VsimWithscores)
}

func (c VsimQueryTypeEleElement) Withattribs() VsimWithattribs {
	_ = "STUB: not implemented"
	return *new(VsimWithattribs)
}

func (c VsimQueryTypeEleElement) Count(num int64) VsimCount {
	_ = "STUB: not implemented"
	return *new(VsimCount)
}

func (c VsimQueryTypeEleElement) Epsilon(delta float32) VsimEpsilon {
	_ = "STUB: not implemented"
	return *new(VsimEpsilon)
}

func (c VsimQueryTypeEleElement) Ef(searchExplorationFactor int64) VsimEf {
	_ = "STUB: not implemented"
	return *new(VsimEf)
}

func (c VsimQueryTypeEleElement) Filter(expression string) VsimFilter {
	_ = "STUB: not implemented"
	return *new(VsimFilter)
}

func (c VsimQueryTypeEleElement) FilterEf(maxFilteringEffort int64) VsimFilterEf {
	_ = "STUB: not implemented"
	return *new(VsimFilterEf)
}

func (c VsimQueryTypeEleElement) Truth() VsimTruth {
	_ = "STUB: not implemented"
	return *new(VsimTruth)
}

func (c VsimQueryTypeEleElement) Nothread() VsimNothread {
	_ = "STUB: not implemented"
	return *new(VsimNothread)
}

func (c VsimQueryTypeEleElement) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type VsimQueryTypeFp32Fp32 Incomplete

func (c VsimQueryTypeFp32Fp32) Vector(vector string) VsimQueryTypeFp32Vector {
	_ = "STUB: not implemented"
	return *new(VsimQueryTypeFp32Vector)
}

type VsimQueryTypeFp32Vector Incomplete

func (c VsimQueryTypeFp32Vector) Withscores() VsimWithscores {
	_ = "STUB: not implemented"
	return *new(VsimWithscores)
}

func (c VsimQueryTypeFp32Vector) Withattribs() VsimWithattribs {
	_ = "STUB: not implemented"
	return *new(VsimWithattribs)
}

func (c VsimQueryTypeFp32Vector) Count(num int64) VsimCount {
	_ = "STUB: not implemented"
	return *new(VsimCount)
}

func (c VsimQueryTypeFp32Vector) Epsilon(delta float32) VsimEpsilon {
	_ = "STUB: not implemented"
	return *new(VsimEpsilon)
}

func (c VsimQueryTypeFp32Vector) Ef(searchExplorationFactor int64) VsimEf {
	_ = "STUB: not implemented"
	return *new(VsimEf)
}

func (c VsimQueryTypeFp32Vector) Filter(expression string) VsimFilter {
	_ = "STUB: not implemented"
	return *new(VsimFilter)
}

func (c VsimQueryTypeFp32Vector) FilterEf(maxFilteringEffort int64) VsimFilterEf {
	_ = "STUB: not implemented"
	return *new(VsimFilterEf)
}

func (c VsimQueryTypeFp32Vector) Truth() VsimTruth {
	_ = "STUB: not implemented"
	return *new(VsimTruth)
}

func (c VsimQueryTypeFp32Vector) Nothread() VsimNothread {
	_ = "STUB: not implemented"
	return *new(VsimNothread)
}

func (c VsimQueryTypeFp32Vector) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type VsimQueryTypeValuesValues Incomplete

func (c VsimQueryTypeValuesValues) Vector(vector ...float32) VsimQueryTypeValuesVector {
	_ = "STUB: not implemented"
	return *new(VsimQueryTypeValuesVector)
}

type VsimQueryTypeValuesVector Incomplete

func (c VsimQueryTypeValuesVector) Vector(vector ...float32) VsimQueryTypeValuesVector {
	_ = "STUB: not implemented"
	return *new(VsimQueryTypeValuesVector)
}

func (c VsimQueryTypeValuesVector) Withscores() VsimWithscores {
	_ = "STUB: not implemented"
	return *new(VsimWithscores)
}

func (c VsimQueryTypeValuesVector) Withattribs() VsimWithattribs {
	_ = "STUB: not implemented"
	return *new(VsimWithattribs)
}

func (c VsimQueryTypeValuesVector) Count(num int64) VsimCount {
	_ = "STUB: not implemented"
	return *new(VsimCount)
}

func (c VsimQueryTypeValuesVector) Epsilon(delta float32) VsimEpsilon {
	_ = "STUB: not implemented"
	return *new(VsimEpsilon)
}

func (c VsimQueryTypeValuesVector) Ef(searchExplorationFactor int64) VsimEf {
	_ = "STUB: not implemented"
	return *new(VsimEf)
}

func (c VsimQueryTypeValuesVector) Filter(expression string) VsimFilter {
	_ = "STUB: not implemented"
	return *new(VsimFilter)
}

func (c VsimQueryTypeValuesVector) FilterEf(maxFilteringEffort int64) VsimFilterEf {
	_ = "STUB: not implemented"
	return *new(VsimFilterEf)
}

func (c VsimQueryTypeValuesVector) Truth() VsimTruth {
	_ = "STUB: not implemented"
	return *new(VsimTruth)
}

func (c VsimQueryTypeValuesVector) Nothread() VsimNothread {
	_ = "STUB: not implemented"
	return *new(VsimNothread)
}

func (c VsimQueryTypeValuesVector) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type VsimTruth Incomplete

func (c VsimTruth) Nothread() VsimNothread { _ = "STUB: not implemented"; return *new(VsimNothread) }

func (c VsimTruth) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type VsimWithattribs Incomplete

func (c VsimWithattribs) Count(num int64) VsimCount {
	_ = "STUB: not implemented"
	return *new(VsimCount)
}

func (c VsimWithattribs) Epsilon(delta float32) VsimEpsilon {
	_ = "STUB: not implemented"
	return *new(VsimEpsilon)
}

func (c VsimWithattribs) Ef(searchExplorationFactor int64) VsimEf {
	_ = "STUB: not implemented"
	return *new(VsimEf)
}

func (c VsimWithattribs) Filter(expression string) VsimFilter {
	_ = "STUB: not implemented"
	return *new(VsimFilter)
}

func (c VsimWithattribs) FilterEf(maxFilteringEffort int64) VsimFilterEf {
	_ = "STUB: not implemented"
	return *new(VsimFilterEf)
}

func (c VsimWithattribs) Truth() VsimTruth { _ = "STUB: not implemented"; return *new(VsimTruth) }

func (c VsimWithattribs) Nothread() VsimNothread {
	_ = "STUB: not implemented"
	return *new(VsimNothread)
}

func (c VsimWithattribs) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type VsimWithscores Incomplete

func (c VsimWithscores) Withattribs() VsimWithattribs {
	_ = "STUB: not implemented"
	return *new(VsimWithattribs)
}

func (c VsimWithscores) Count(num int64) VsimCount {
	_ = "STUB: not implemented"
	return *new(VsimCount)
}

func (c VsimWithscores) Epsilon(delta float32) VsimEpsilon {
	_ = "STUB: not implemented"
	return *new(VsimEpsilon)
}

func (c VsimWithscores) Ef(searchExplorationFactor int64) VsimEf {
	_ = "STUB: not implemented"
	return *new(VsimEf)
}

func (c VsimWithscores) Filter(expression string) VsimFilter {
	_ = "STUB: not implemented"
	return *new(VsimFilter)
}

func (c VsimWithscores) FilterEf(maxFilteringEffort int64) VsimFilterEf {
	_ = "STUB: not implemented"
	return *new(VsimFilterEf)
}

func (c VsimWithscores) Truth() VsimTruth { _ = "STUB: not implemented"; return *new(VsimTruth) }

func (c VsimWithscores) Nothread() VsimNothread {
	_ = "STUB: not implemented"
	return *new(VsimNothread)
}

func (c VsimWithscores) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

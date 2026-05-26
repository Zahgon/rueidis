// Code generated DO NOT EDIT

package cmds

type GraphConfigGet Incomplete

func (b Builder) GraphConfigGet() (c GraphConfigGet) {
	_ = "STUB: not implemented"
	return *new(GraphConfigGet)
}

func (c GraphConfigGet) Name(name string) GraphConfigGetName {
	_ = "STUB: not implemented"
	return *new(GraphConfigGetName)
}

type GraphConfigGetName Incomplete

func (c GraphConfigGetName) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type GraphConfigSet Incomplete

func (b Builder) GraphConfigSet() (c GraphConfigSet) {
	_ = "STUB: not implemented"
	return *new(GraphConfigSet)
}

func (c GraphConfigSet) Name(name string) GraphConfigSetName {
	_ = "STUB: not implemented"
	return *new(GraphConfigSetName)
}

type GraphConfigSetName Incomplete

func (c GraphConfigSetName) Value(value string) GraphConfigSetValue {
	_ = "STUB: not implemented"
	return *new(GraphConfigSetValue)
}

type GraphConfigSetValue Incomplete

func (c GraphConfigSetValue) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type GraphConstraintCreate Incomplete

func (b Builder) GraphConstraintCreate() (c GraphConstraintCreate) {
	_ = "STUB: not implemented"
	return *new(GraphConstraintCreate)
}

func (c GraphConstraintCreate) Key(key string) GraphConstraintCreateKey {
	_ = "STUB: not implemented"
	return *new(GraphConstraintCreateKey)
}

type GraphConstraintCreateEntityNode Incomplete

func (c GraphConstraintCreateEntityNode) Properties(properties int64) GraphConstraintCreateProperties {
	_ = "STUB: not implemented"
	return *new(GraphConstraintCreateProperties)
}

type GraphConstraintCreateEntityRelationship Incomplete

func (c GraphConstraintCreateEntityRelationship) Properties(properties int64) GraphConstraintCreateProperties {
	_ = "STUB: not implemented"
	return *new(GraphConstraintCreateProperties)
}

type GraphConstraintCreateKey Incomplete

func (c GraphConstraintCreateKey) Mandatory() GraphConstraintCreateModeMandatory {
	_ = "STUB: not implemented"
	return *new(GraphConstraintCreateModeMandatory)
}

func (c GraphConstraintCreateKey) Unique() GraphConstraintCreateModeUnique {
	_ = "STUB: not implemented"
	return *new(GraphConstraintCreateModeUnique)
}

type GraphConstraintCreateModeMandatory Incomplete

func (c GraphConstraintCreateModeMandatory) Node(node string) GraphConstraintCreateEntityNode {
	_ = "STUB: not implemented"
	return *new(GraphConstraintCreateEntityNode)
}

func (c GraphConstraintCreateModeMandatory) Relationship(relationship string) GraphConstraintCreateEntityRelationship {
	_ = "STUB: not implemented"
	return *new(GraphConstraintCreateEntityRelationship)
}

type GraphConstraintCreateModeUnique Incomplete

func (c GraphConstraintCreateModeUnique) Node(node string) GraphConstraintCreateEntityNode {
	_ = "STUB: not implemented"
	return *new(GraphConstraintCreateEntityNode)
}

func (c GraphConstraintCreateModeUnique) Relationship(relationship string) GraphConstraintCreateEntityRelationship {
	_ = "STUB: not implemented"
	return *new(GraphConstraintCreateEntityRelationship)
}

type GraphConstraintCreateProp Incomplete

func (c GraphConstraintCreateProp) Prop(prop ...string) GraphConstraintCreateProp {
	_ = "STUB: not implemented"
	return *new(GraphConstraintCreateProp)
}

func (c GraphConstraintCreateProp) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GraphConstraintCreateProperties Incomplete

func (c GraphConstraintCreateProperties) Prop(prop ...string) GraphConstraintCreateProp {
	_ = "STUB: not implemented"
	return *new(GraphConstraintCreateProp)
}

type GraphConstraintDrop Incomplete

func (b Builder) GraphConstraintDrop() (c GraphConstraintDrop) {
	_ = "STUB: not implemented"
	return *new(GraphConstraintDrop)
}

func (c GraphConstraintDrop) Key(key string) GraphConstraintDropKey {
	_ = "STUB: not implemented"
	return *new(GraphConstraintDropKey)
}

type GraphConstraintDropEntityNode Incomplete

func (c GraphConstraintDropEntityNode) Properties(properties int64) GraphConstraintDropProperties {
	_ = "STUB: not implemented"
	return *new(GraphConstraintDropProperties)
}

type GraphConstraintDropEntityRelationship Incomplete

func (c GraphConstraintDropEntityRelationship) Properties(properties int64) GraphConstraintDropProperties {
	_ = "STUB: not implemented"
	return *new(GraphConstraintDropProperties)
}

type GraphConstraintDropKey Incomplete

func (c GraphConstraintDropKey) Mandatory() GraphConstraintDropModeMandatory {
	_ = "STUB: not implemented"
	return *new(GraphConstraintDropModeMandatory)
}

func (c GraphConstraintDropKey) Unique() GraphConstraintDropModeUnique {
	_ = "STUB: not implemented"
	return *new(GraphConstraintDropModeUnique)
}

type GraphConstraintDropModeMandatory Incomplete

func (c GraphConstraintDropModeMandatory) Node(node string) GraphConstraintDropEntityNode {
	_ = "STUB: not implemented"
	return *new(GraphConstraintDropEntityNode)
}

func (c GraphConstraintDropModeMandatory) Relationship(relationship string) GraphConstraintDropEntityRelationship {
	_ = "STUB: not implemented"
	return *new(GraphConstraintDropEntityRelationship)
}

type GraphConstraintDropModeUnique Incomplete

func (c GraphConstraintDropModeUnique) Node(node string) GraphConstraintDropEntityNode {
	_ = "STUB: not implemented"
	return *new(GraphConstraintDropEntityNode)
}

func (c GraphConstraintDropModeUnique) Relationship(relationship string) GraphConstraintDropEntityRelationship {
	_ = "STUB: not implemented"
	return *new(GraphConstraintDropEntityRelationship)
}

type GraphConstraintDropProp Incomplete

func (c GraphConstraintDropProp) Prop(prop ...string) GraphConstraintDropProp {
	_ = "STUB: not implemented"
	return *new(GraphConstraintDropProp)
}

func (c GraphConstraintDropProp) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type GraphConstraintDropProperties Incomplete

func (c GraphConstraintDropProperties) Prop(prop ...string) GraphConstraintDropProp {
	_ = "STUB: not implemented"
	return *new(GraphConstraintDropProp)
}

type GraphDelete Incomplete

func (b Builder) GraphDelete() (c GraphDelete) { _ = "STUB: not implemented"; return *new(GraphDelete) }

func (c GraphDelete) Graph(graph string) GraphDeleteGraph {
	_ = "STUB: not implemented"
	return *new(GraphDeleteGraph)
}

type GraphDeleteGraph Incomplete

func (c GraphDeleteGraph) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type GraphExplain Incomplete

func (b Builder) GraphExplain() (c GraphExplain) {
	_ = "STUB: not implemented"
	return *new(GraphExplain)
}

func (c GraphExplain) Graph(graph string) GraphExplainGraph {
	_ = "STUB: not implemented"
	return *new(GraphExplainGraph)
}

type GraphExplainGraph Incomplete

func (c GraphExplainGraph) Query(query string) GraphExplainQuery {
	_ = "STUB: not implemented"
	return *new(GraphExplainQuery)
}

type GraphExplainQuery Incomplete

func (c GraphExplainQuery) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type GraphList Incomplete

func (b Builder) GraphList() (c GraphList) { _ = "STUB: not implemented"; return *new(GraphList) }

func (c GraphList) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type GraphProfile Incomplete

func (b Builder) GraphProfile() (c GraphProfile) {
	_ = "STUB: not implemented"
	return *new(GraphProfile)
}

func (c GraphProfile) Graph(graph string) GraphProfileGraph {
	_ = "STUB: not implemented"
	return *new(GraphProfileGraph)
}

type GraphProfileGraph Incomplete

func (c GraphProfileGraph) Query(query string) GraphProfileQuery {
	_ = "STUB: not implemented"
	return *new(GraphProfileQuery)
}

type GraphProfileQuery Incomplete

func (c GraphProfileQuery) Timeout(timeout int64) GraphProfileTimeout {
	_ = "STUB: not implemented"
	return *new(GraphProfileTimeout)
}

func (c GraphProfileQuery) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type GraphProfileTimeout Incomplete

func (c GraphProfileTimeout) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type GraphQuery Incomplete

func (b Builder) GraphQuery() (c GraphQuery) { _ = "STUB: not implemented"; return *new(GraphQuery) }

func (c GraphQuery) Graph(graph string) GraphQueryGraph {
	_ = "STUB: not implemented"
	return *new(GraphQueryGraph)
}

type GraphQueryGraph Incomplete

func (c GraphQueryGraph) Query(query string) GraphQueryQuery {
	_ = "STUB: not implemented"
	return *new(GraphQueryQuery)
}

type GraphQueryQuery Incomplete

func (c GraphQueryQuery) Timeout(timeout int64) GraphQueryTimeout {
	_ = "STUB: not implemented"
	return *new(GraphQueryTimeout)
}

func (c GraphQueryQuery) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type GraphQueryTimeout Incomplete

func (c GraphQueryTimeout) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type GraphRoQuery Incomplete

func (b Builder) GraphRoQuery() (c GraphRoQuery) {
	_ = "STUB: not implemented"
	return *new(GraphRoQuery)
}

func (c GraphRoQuery) Graph(graph string) GraphRoQueryGraph {
	_ = "STUB: not implemented"
	return *new(GraphRoQueryGraph)
}

type GraphRoQueryGraph Incomplete

func (c GraphRoQueryGraph) Query(query string) GraphRoQueryQuery {
	_ = "STUB: not implemented"
	return *new(GraphRoQueryQuery)
}

type GraphRoQueryQuery Incomplete

func (c GraphRoQueryQuery) Timeout(timeout int64) GraphRoQueryTimeout {
	_ = "STUB: not implemented"
	return *new(GraphRoQueryTimeout)
}

func (c GraphRoQueryQuery) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GraphRoQueryQuery) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type GraphRoQueryTimeout Incomplete

func (c GraphRoQueryTimeout) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

func (c GraphRoQueryTimeout) Cache() Cacheable { _ = "STUB: not implemented"; return *new(Cacheable) }

type GraphSlowlog Incomplete

func (b Builder) GraphSlowlog() (c GraphSlowlog) {
	_ = "STUB: not implemented"
	return *new(GraphSlowlog)
}

func (c GraphSlowlog) Graph(graph string) GraphSlowlogGraph {
	_ = "STUB: not implemented"
	return *new(GraphSlowlogGraph)
}

type GraphSlowlogGraph Incomplete

func (c GraphSlowlogGraph) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

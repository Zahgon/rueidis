// Code generated DO NOT EDIT

package cmds

type Psubscribe Incomplete

func (b Builder) Psubscribe() (c Psubscribe) { _ = "STUB: not implemented"; return *new(Psubscribe) }

func (c Psubscribe) Pattern(pattern ...string) PsubscribePattern {
	_ = "STUB: not implemented"
	return *new(PsubscribePattern)
}

type PsubscribePattern Incomplete

func (c PsubscribePattern) Pattern(pattern ...string) PsubscribePattern {
	_ = "STUB: not implemented"
	return *new(PsubscribePattern)
}

func (c PsubscribePattern) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Publish Incomplete

func (b Builder) Publish() (c Publish) { _ = "STUB: not implemented"; return *new(Publish) }

func (c Publish) Channel(channel string) PublishChannel {
	_ = "STUB: not implemented"
	return *new(PublishChannel)
}

type PublishChannel Incomplete

func (c PublishChannel) Message(message string) PublishMessage {
	_ = "STUB: not implemented"
	return *new(PublishMessage)
}

type PublishMessage Incomplete

func (c PublishMessage) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type PubsubChannels Incomplete

func (b Builder) PubsubChannels() (c PubsubChannels) {
	_ = "STUB: not implemented"
	return *new(PubsubChannels)
}

func (c PubsubChannels) Pattern(pattern string) PubsubChannelsPattern {
	_ = "STUB: not implemented"
	return *new(PubsubChannelsPattern)
}

func (c PubsubChannels) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type PubsubChannelsPattern Incomplete

func (c PubsubChannelsPattern) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type PubsubHelp Incomplete

func (b Builder) PubsubHelp() (c PubsubHelp) { _ = "STUB: not implemented"; return *new(PubsubHelp) }

func (c PubsubHelp) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type PubsubNumpat Incomplete

func (b Builder) PubsubNumpat() (c PubsubNumpat) {
	_ = "STUB: not implemented"
	return *new(PubsubNumpat)
}

func (c PubsubNumpat) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type PubsubNumsub Incomplete

func (b Builder) PubsubNumsub() (c PubsubNumsub) {
	_ = "STUB: not implemented"
	return *new(PubsubNumsub)
}

func (c PubsubNumsub) Channel(channel ...string) PubsubNumsubChannel {
	_ = "STUB: not implemented"
	return *new(PubsubNumsubChannel)
}

func (c PubsubNumsub) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type PubsubNumsubChannel Incomplete

func (c PubsubNumsubChannel) Channel(channel ...string) PubsubNumsubChannel {
	_ = "STUB: not implemented"
	return *new(PubsubNumsubChannel)
}

func (c PubsubNumsubChannel) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type PubsubShardchannels Incomplete

func (b Builder) PubsubShardchannels() (c PubsubShardchannels) {
	_ = "STUB: not implemented"
	return *new(PubsubShardchannels)
}

func (c PubsubShardchannels) Pattern(pattern string) PubsubShardchannelsPattern {
	_ = "STUB: not implemented"
	return *new(PubsubShardchannelsPattern)
}

func (c PubsubShardchannels) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type PubsubShardchannelsPattern Incomplete

func (c PubsubShardchannelsPattern) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type PubsubShardnumsub Incomplete

func (b Builder) PubsubShardnumsub() (c PubsubShardnumsub) {
	_ = "STUB: not implemented"
	return *new(PubsubShardnumsub)
}

func (c PubsubShardnumsub) Channel(channel ...string) PubsubShardnumsubChannel {
	_ = "STUB: not implemented"
	return *new(PubsubShardnumsubChannel)
}

func (c PubsubShardnumsub) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type PubsubShardnumsubChannel Incomplete

func (c PubsubShardnumsubChannel) Channel(channel ...string) PubsubShardnumsubChannel {
	_ = "STUB: not implemented"
	return *new(PubsubShardnumsubChannel)
}

func (c PubsubShardnumsubChannel) Build() Completed {
	_ = "STUB: not implemented"
	return *new(Completed)
}

type Punsubscribe Incomplete

func (b Builder) Punsubscribe() (c Punsubscribe) {
	_ = "STUB: not implemented"
	return *new(Punsubscribe)
}

func (c Punsubscribe) Pattern(pattern ...string) PunsubscribePattern {
	_ = "STUB: not implemented"
	return *new(PunsubscribePattern)
}

func (c Punsubscribe) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type PunsubscribePattern Incomplete

func (c PunsubscribePattern) Pattern(pattern ...string) PunsubscribePattern {
	_ = "STUB: not implemented"
	return *new(PunsubscribePattern)
}

func (c PunsubscribePattern) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Spublish Incomplete

func (b Builder) Spublish() (c Spublish) { _ = "STUB: not implemented"; return *new(Spublish) }

func (c Spublish) Channel(channel string) SpublishChannel {
	_ = "STUB: not implemented"
	return *new(SpublishChannel)
}

type SpublishChannel Incomplete

func (c SpublishChannel) Message(message string) SpublishMessage {
	_ = "STUB: not implemented"
	return *new(SpublishMessage)
}

type SpublishMessage Incomplete

func (c SpublishMessage) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Ssubscribe Incomplete

func (b Builder) Ssubscribe() (c Ssubscribe) { _ = "STUB: not implemented"; return *new(Ssubscribe) }

func (c Ssubscribe) Channel(channel ...string) SsubscribeChannel {
	_ = "STUB: not implemented"
	return *new(SsubscribeChannel)
}

type SsubscribeChannel Incomplete

func (c SsubscribeChannel) Channel(channel ...string) SsubscribeChannel {
	_ = "STUB: not implemented"
	return *new(SsubscribeChannel)
}

func (c SsubscribeChannel) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Subscribe Incomplete

func (b Builder) Subscribe() (c Subscribe) { _ = "STUB: not implemented"; return *new(Subscribe) }

func (c Subscribe) Channel(channel ...string) SubscribeChannel {
	_ = "STUB: not implemented"
	return *new(SubscribeChannel)
}

type SubscribeChannel Incomplete

func (c SubscribeChannel) Channel(channel ...string) SubscribeChannel {
	_ = "STUB: not implemented"
	return *new(SubscribeChannel)
}

func (c SubscribeChannel) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Sunsubscribe Incomplete

func (b Builder) Sunsubscribe() (c Sunsubscribe) {
	_ = "STUB: not implemented"
	return *new(Sunsubscribe)
}

func (c Sunsubscribe) Channel(channel ...string) SunsubscribeChannel {
	_ = "STUB: not implemented"
	return *new(SunsubscribeChannel)
}

func (c Sunsubscribe) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type SunsubscribeChannel Incomplete

func (c SunsubscribeChannel) Channel(channel ...string) SunsubscribeChannel {
	_ = "STUB: not implemented"
	return *new(SunsubscribeChannel)
}

func (c SunsubscribeChannel) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Unsubscribe Incomplete

func (b Builder) Unsubscribe() (c Unsubscribe) { _ = "STUB: not implemented"; return *new(Unsubscribe) }

func (c Unsubscribe) Channel(channel ...string) UnsubscribeChannel {
	_ = "STUB: not implemented"
	return *new(UnsubscribeChannel)
}

func (c Unsubscribe) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type UnsubscribeChannel Incomplete

func (c UnsubscribeChannel) Channel(channel ...string) UnsubscribeChannel {
	_ = "STUB: not implemented"
	return *new(UnsubscribeChannel)
}

func (c UnsubscribeChannel) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

// Code generated DO NOT EDIT

package cmds

type Discard Incomplete

func (b Builder) Discard() (c Discard) { _ = "STUB: not implemented"; return *new(Discard) }

func (c Discard) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Exec Incomplete

func (b Builder) Exec() (c Exec) { _ = "STUB: not implemented"; return *new(Exec) }

func (c Exec) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Multi Incomplete

func (b Builder) Multi() (c Multi) { _ = "STUB: not implemented"; return *new(Multi) }

func (c Multi) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Unwatch Incomplete

func (b Builder) Unwatch() (c Unwatch) { _ = "STUB: not implemented"; return *new(Unwatch) }

func (c Unwatch) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

type Watch Incomplete

func (b Builder) Watch() (c Watch) { _ = "STUB: not implemented"; return *new(Watch) }

func (c Watch) Key(key ...string) WatchKey { _ = "STUB: not implemented"; return *new(WatchKey) }

type WatchKey Incomplete

func (c WatchKey) Key(key ...string) WatchKey { _ = "STUB: not implemented"; return *new(WatchKey) }

func (c WatchKey) Build() Completed { _ = "STUB: not implemented"; return *new(Completed) }

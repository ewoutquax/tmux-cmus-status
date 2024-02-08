package cmd

type Executor struct {
	CommandExecutor
}

type CommandExecutor interface {
	ExecInternal(name string, arg ...string)
}

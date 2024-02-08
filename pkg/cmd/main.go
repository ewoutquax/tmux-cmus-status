package cmd

import "os/exec"

type ExecutorOptsFunc func(*Executor)

type CommandExecutor interface {
	ExecInternal(name string, arg ...string)
}

type CommandExecutorDefault struct{}

func (e CommandExecutorDefault) ExecInternal(name string, arg ...string) {
	run := exec.Command(name, arg...)
	run.Run()
}

type Executor struct {
	CommandExecutor
}

func (e *Executor) Exec(params []string) {
	e.CommandExecutor.ExecInternal(params[0], params[1:]...)
}

func Building(optFuncs ...ExecutorOptsFunc) *Executor {
	p := defaultExecutor()
	for _, fn := range optFuncs {
		fn(p)
	}

	return p
}

func WithCommandExecutor(e CommandExecutor) ExecutorOptsFunc {
	return func(p *Executor) {
		p.CommandExecutor = e
	}
}

func defaultExecutor() *Executor {
	return &Executor{
		CommandExecutorDefault{},
	}
}

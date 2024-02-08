package cmd

import "os/exec"

func (e *Executor) Exec(params []string) {
	e.CommandExecutor.ExecInternal(params[0], params[1:]...)
}

type CommandExecutorDefault struct{}

func (e CommandExecutorDefault) ExecInternal(name string, arg ...string) {
	run := exec.Command(name, arg...)
	run.Run()
}

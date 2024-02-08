package cmd_test

import (
	"testing"

	"github.com/ewoutquax/tmux-cmus-status/pkg/cmd"
	"github.com/stretchr/testify/assert"
)

func TestCommandExec(t *testing.T) {
	params := []string{
		"tmux",
		"rename-window",
		"-t",
		"0",
		"playing: 'Carillion'|cmus",
	}

	executor := cmd.BuildExecutor(
		cmd.WithCommandExecutor(CommandExecutorMock{assert.New(t)}),
	)

	executor.Exec(params)
}

type CommandExecutorMock struct {
	assert *assert.Assertions
}

func (e CommandExecutorMock) ExecInternal(name string, arg ...string) {
	expectedArgs := []string{
		"rename-window",
		"-t",
		"0",
		"playing: 'Carillion'|cmus",
	}

	e.assert.Equal("tmux", name)
	e.assert.Equal(expectedArgs, arg)
}

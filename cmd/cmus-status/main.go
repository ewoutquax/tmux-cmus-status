package main

import (
	"os"
	"sync"

	"github.com/ewoutquax/tmux-cmus-status/internal/player"
	"github.com/ewoutquax/tmux-cmus-status/internal/status"
	"github.com/ewoutquax/tmux-cmus-status/pkg/cmd"
	"github.com/ewoutquax/tmux-cmus-status/pkg/envvars"
)

var once sync.Once

func main() {
	once.Do(func() {
		envvars.LoadEnvVars()
	})

	executor := cmd.BuildExecutor()
	params := player.ExtractCmusData(os.Args)
	executor.Exec(status.ForTmux(params))
}

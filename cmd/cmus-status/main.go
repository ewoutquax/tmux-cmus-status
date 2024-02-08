package main

import (
	"os"
	"sync"

	"github.com/ewoutquax/tmux-cmus-status/internal/cmus"
	"github.com/ewoutquax/tmux-cmus-status/internal/status"
	"github.com/ewoutquax/tmux-cmus-status/pkg/cmd"
	"github.com/ewoutquax/tmux-cmus-status/pkg/envvars"
)

var once sync.Once

func main() {
	once.Do(func() {
		envvars.LoadEnvVars()
	})

	executor := cmd.Building()
	params := cmus.ExtractData(os.Args)
	executor.Exec(status.ForTmux(params))
}

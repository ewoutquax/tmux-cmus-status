package main

import (
	"fmt"
	"os"

	"github.com/ewoutquax/tmux-cmus-status/internal/player"
	"github.com/ewoutquax/tmux-cmus-status/internal/status"
	"github.com/ewoutquax/tmux-cmus-status/pkg/cmd"
	"github.com/ewoutquax/tmux-cmus-status/pkg/config"
)

func main() {
	config.InitConfig()

	debugPrintArgs()

	executor := cmd.BuildExecutor()
	params := player.ExtractCmusData(os.Args)
	executor.Exec(status.ForTmux(params))
}

func debugPrintArgs() {
	if config.Config.Debug {
		os.Mkdir("/tmp/cmus-status", 0777)
		ptrFile, _ := os.OpenFile("/tmp/cmus-status/log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)
		defer ptrFile.Close()

		ptrFile.Write([]byte("\nCMUS arguments\n"))
		ptrFile.Write([]byte("--------------\n"))

		for _, value := range os.Args {
			ptrFile.Write([]byte(fmt.Sprintf("v: %v\n", value)))
		}
	}
}

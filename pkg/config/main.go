package config

import (
	"os"
	"strconv"
	"sync"

	"github.com/ewoutquax/tmux-cmus-status/pkg/envvars"
)

type AppConfig struct {
	Debug         bool
	TmuxCmusIndex int
}

var once sync.Once
var Config AppConfig

func InitConfig() {
	once.Do(func() {
		envvars.LoadEnvVars()
	})

	Config = AppConfig{
		Debug:         os.Getenv("DEBUG") == "true",
		TmuxCmusIndex: convAtoI(os.Getenv("TMUX_CMUS_INDEX")),
	}
}

func convAtoI(input string) int {
	nmbr, err := strconv.Atoi(input)
	if err != nil {
		return 0
	}

	return nmbr
}

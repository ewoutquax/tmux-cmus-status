package envvars

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ewoutquax/tmux-cmus-status/pkg/rootdir"
	"github.com/joho/godotenv"
)

const defaultEnv string = "DEV"

func LoadEnvVars() {
	env := os.Getenv("GOENV")
	if env == "" {
		env = defaultEnv
	}

	baseDir := rootdir.Get()

	envFileDefault, _ := filepath.Abs(baseDir + "/.env")
	envFileCustom, _ := filepath.Abs(baseDir + "/.env." + strings.ToLower(env))

	fmt.Printf("envFileDefault: %v\n", envFileDefault)
	fmt.Printf("envFileCustom: %v\n", envFileCustom)

	godotenv.Load(envFileDefault, envFileCustom)
}

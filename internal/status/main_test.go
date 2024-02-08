package status_test

import (
	"os"
	"testing"

	"github.com/ewoutquax/tmux-cmus-status/internal/cmus"
	. "github.com/ewoutquax/tmux-cmus-status/internal/status"
	"github.com/ewoutquax/tmux-cmus-status/pkg/envvars"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Setenv("GOENV", "TEST")
	envvars.LoadEnvVars()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestForTmux(t *testing.T) {
	expectedResult := []string{
		"tmux",
		"rename-window",
		"-t",
		"0",
		"paused: 'Scooby Snacks'|cmus",
	}

	data := cmus.ExtractedData{
		Status: "paused",
		Title:  "Scooby Snacks",
	}

	assert.Equal(t, expectedResult, ForTmux(data))
}

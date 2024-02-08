package status

import (
	"os"

	"github.com/ewoutquax/tmux-cmus-status/internal/player"
)

func ForTmux(data player.ExtractedData) (params []string) {
	return []string{
		"tmux",
		"rename-window",
		"-t",
		os.Getenv("TMUX_CMUS_INDEX"),
		data.Status + ": '" + data.Title + "'|cmus",
	}
}

package status

import (
	"strconv"

	"github.com/ewoutquax/tmux-cmus-status/internal/player"
	"github.com/ewoutquax/tmux-cmus-status/pkg/config"
)

func ForTmux(data player.ExtractedData) (params []string) {
	return []string{
		"tmux",
		"rename-window",
		"-t",
		strconv.Itoa(config.Config.TmuxCmusIndex),
		data.Status + ": '" + data.Title + "'|cmus",
	}
}

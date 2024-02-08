package cmus_test

import (
	"testing"

	. "github.com/ewoutquax/tmux-cmus-status/internal/cmus"
	"github.com/stretchr/testify/assert"
)

func TestExtractData(t *testing.T) {
	expectedResult := ExtractedData{
		Status: "paused",
		Title:  "Scooby Snacks",
	}

	assert.Equal(t, expectedResult, ExtractData(testInput()))
}

func testInput() []string {
	return []string{
		"/home/ewout/sites/go/tmux-cmus-status/main",
		"status",
		"paused",
		"file",
		"/home/ewout/Music/Fun Lovin Criminals - Scooby Snacks (Official Music Video).mp3",
		"artist",
		"Fun Lovin' Criminals",
		"album",
		"Come Find Yourself",
		"title",
		"Scooby Snacks",
		"duration",
		"205",
	}
}

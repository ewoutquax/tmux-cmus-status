package player_test

import (
	"testing"

	. "github.com/ewoutquax/tmux-cmus-status/internal/player"
	"github.com/stretchr/testify/assert"
)

func TestExtractDataLocal(t *testing.T) {
	expectedResult := ExtractedData{
		IsStream: false,
		Status:   "paused",
		Title:    "Scooby Snacks",
	}

	assert.Equal(t, expectedResult, ExtractCmusData(testInput()))
}

func TestExtractDataStream(t *testing.T) {
	expectedResult := ExtractedData{
		IsStream: true,
		Status:   "playing",
		Title:    "Adele - Hometown Glory",
	}

	assert.Equal(t, expectedResult, ExtractCmusData(testInputStream()))
}

func TestFormatTextPlaying(t *testing.T) {
	assert.Equal(t, WindowText("playing: 'Scooby Snacks'|cmus"), FormatText(testDataLocal("playing")))
}

func TestFormatTextPaused(t *testing.T) {
	assert.Equal(t, WindowText("paused: 'Scooby Snacks'|cmus"), FormatText(testDataLocal("paused")))
}

func TestFormatTextExiting(t *testing.T) {
	assert.Equal(t, WindowText("cmus"), FormatText(testDataLocal("exiting")))
}

func TestFormatTextStreamPlaying(t *testing.T) {
	assert.Equal(t, WindowText("streaming: 'Adele - Hometown Glory'|cmus"), FormatText(testDataStream()))
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

func testInputStream() []string {
	return []string{
		"/home/ewout/projects/tmux-cmus-status/bin/cmus-status",
		"status",
		"playing",
		"url",
		"http://playerservices.streamtheworld.com/api/livestream-redirect/RADIO10.mp3",
		"title",
		"Adele - Hometown Glory",
	}
}

func testDataLocal(status string) ExtractedData {
	return ExtractedData{
		IsStream: false,
		Status:   status,
		Title:    "Scooby Snacks",
	}
}

func testDataStream() ExtractedData {
	return ExtractedData{
		IsStream: true,
		Status:   "playing",
		Title:    "Adele - Hometown Glory",
	}
}

package player

type WindowText string

type ExtractedData struct {
	IsStream bool   // Is this data from a stream, e.g. 'Radio10'
	Status   string // The status of the CMUS-player: 'paused', exiting or playing
	Title    string // The title of the playing song.
}

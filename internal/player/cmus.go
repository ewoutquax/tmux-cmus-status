package player

func ExtractCmusData(args []string) ExtractedData {
	var data = ExtractedData{
		IsStream: false,
		Status:   "",
		Title:    "",
	}

	for idx := 1; idx < len(args)-1; idx += 2 {
		switch args[idx] {
		case "url":
			data.IsStream = true
		case "status":
			data.Status = args[idx+1]
		case "title":
			data.Title = args[idx+1]
		}
	}

	return data
}

func FormatText(data ExtractedData) WindowText {
	var localStatus string = data.Status

	switch data.Status {
	case "exiting":
		return WindowText("cmus")
	case "playing":
		if data.IsStream {
			localStatus = "streaming"
		}
	}
	return WindowText(localStatus + ": '" + data.Title + "'|cmus")
}

package player

func ExtractCmusData(args []string) (data ExtractedData) {
	for idx := 1; idx < len(args)-1; idx += 2 {
		switch args[idx] {
		case "status":
			data.Status = args[idx+1]
		case "title":
			data.Title = args[idx+1]
		}
	}

	return
}

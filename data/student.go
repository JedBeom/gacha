package data

func studentIDVaildation(id string) (ok bool) {

	// StudentID must be 5 letters
	if len(id) != 5 {
		return
	}

	// Grade Check (1-3)
	switch id[0:2] {
	case "10", "20", "30":
		break
	default:
		return
	}

	// Class Check (1-9)
	switch string(id[2]) {
	case "1", "2", "3", "4", "5", "6", "7", "8", "9":
		break
	default:
		return
	}

	// Class Check for Grade 1 (1-7)
	switch id[0:3] {
	case "108", "109":
		return
	default:
		break
	}

	// Student Number tens Check (0-3)
	switch string(id[3]) {
	case "0", "1", "2", "3":
		break
	default:
		return
	}

	// Student Number units Check (0-9)
	switch string(id[4]) {
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		ok = true
		return
	default:
		return
	}

}

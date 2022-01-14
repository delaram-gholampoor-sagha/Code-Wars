package kata

func BoolToWord(word bool) string {
	result := ""
	if word {
		result = "Yes"
	} else {
		result = "No"
	}
	return result
}

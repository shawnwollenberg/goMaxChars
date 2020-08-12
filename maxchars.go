package maxchars

// --- Directions
// Given a string, return the character that is most
// commonly used in the string.
// --- Examples
// maxChar("abcccccccd") === "c"
// maxChar("apple 1231111") === "1"

func maxChars(str string) string {
	var m = make(map[string]int)
	var maxVal = 0
	var maxChar = ""
	for _, v := range str {
		if m[string(v)] > 0 {
			m[string(v)] = m[string(v)] + 1
		} else {
			m[string(v)] = 1

		}
	}
	for key, val := range m {
		if val > maxVal {
			maxVal = val
			maxChar = key
		}
	}

	return maxChar
}



func isVowel(b rune) bool {
	return b == 65 || b == 69 || b == 73 || b == 79 || b == 85 ||
		b == 97 || b == 101 || b == 105 || b == 111 || b == 117
}

func isEnglishLetter(b rune) bool {
	if b >= 65 && b <= 90 || b >= 97 && b <= 122 {
		return true
	}
	return false
}

func isNumber(b rune) bool {
	return b >= '0' && b <= '9'
}

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}

	vowels := 0
	consonant := 0

	for _, char := range word {
		if isVowel(char) {
			vowels = vowels + 1
		} else if isEnglishLetter(char) {
			consonant = consonant + 1
		} else if !isNumber(char) {
			return false
		}
	}

	return vowels >= 1 && consonant >= 1
}
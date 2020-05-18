func reverseOnlyLetters(S string) string {
	var left,right int = 0,len(S) - 1
	for left < right {
		if isLetter(S[left]) && isLetter(S[right]) {
			S = S[:left] + string(S[right]) + S[left + 1:right] + string(S[left]) + S[right + 1:]
			left++
			right--
		}
		if !isLetter(S[left]) {
			left++
		}
		if !isLetter(S[right]) {
			right--
		}
	}

	return S
}
func isLetter(a byte) bool {
	if (a >= 65 && a <= 90) || (a >= 97 && a <= 122) {
		return true
	}
	return false
}

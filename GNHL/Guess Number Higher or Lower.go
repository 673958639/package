func guessNumber(n int) int {
	a , b := 1, n
	for a < b {
		mid := a + (b - a) / 2
		if guess(mid) == 0{
			return mid
		}else if guess(mid) == 1{
			a = mid + 1
		}else{
			b = mid - 1
		}
	}
	return a
}
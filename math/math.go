package math

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isFoo(x int) bool {
	if x == 42 {
		return true
	}
	return false
}

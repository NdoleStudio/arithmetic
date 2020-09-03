package arithmethic

// DivideIntsReturnFloat divides 2 integers and returns a float64 as output
func DivideIntsReturnFloat(numerator int, denominator int) float64 {
	return float64(numerator) / float64(denominator)
}

// MinInt returns the minimum of 2 integers
func MinInt(i, j int) int {
	if i < j {
		return i
	}

	return j
}

// MinInts returns the minimum of the input integers.
func MinInts(min int, others ...int) int {
	for i := 0; i < len(others); i++ {
		if others[i] < min {
			min = others[i]
		}
	}

	return min
}

// MaxInt returns the maximum of 2 integers
func MaxInt(i, j int) int {
	if i > j {
		return i
	}

	return j
}

// MaxInts returns the maximum of the input integers.
func MaxInts(max int, others ...int) int {
	for i := 0; i < len(others); i++ {
		if others[i] > max {
			max = others[i]
		}
	}

	return max
}

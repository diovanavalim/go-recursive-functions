package fibonacci

func GetFibonacciNumber(position uint) uint {
	if position <= 1 {
		return position
	}

	return GetFibonacciNumber(position-2) + GetFibonacciNumber(position-1)
}

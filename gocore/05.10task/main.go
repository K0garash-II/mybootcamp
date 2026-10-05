package main

type Number interface{
	int | float64 | int64
}

func findMax[T Number](values []T) T {
	if len(values) == 0 {
		return 0
	}
	var max T = values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > max{
			max = values[i]
		}
	}
	return max
}

func main() {
	findMax([]int{1, 2, 3, 4})
}
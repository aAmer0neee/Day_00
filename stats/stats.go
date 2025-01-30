package stats

import (
	"math"
	"strconv"
)

func Mode(input []int) float64 {
	result := 100001.0
	countMap := make(map[int]int)

	maxVal := 0

	for _, value := range input {
		countMap[value]++
		if countMap[value] > maxVal {
			result = float64(value)
			maxVal = countMap[value]
		} else if countMap[value] == maxVal {
			if result > float64(value) {
				result = float64(value)
			}
		}

	}
	return result
}

func StandardDeviation(input []int) float64 {
	mean := Mean(input)
	res := 0.0

	for _, value := range input {

		res += math.Pow(float64(value)-mean, 2)
	}
	return math.Sqrt(res / float64(len(input)))
}

func Mean(input []int) float64 {
	var result float64
	for _, value := range input {
		result += float64(value)
	}
	return result / float64(len(input))
}

func Median(input []int) float64 {
	var result float64
	if n := len(input); n%2 == 1 {
		result = float64(input[n/2])
	} else {
		result = float64(input[n/2-1]+input[n/2]) / 2.0
	}
	return result
}

func SingleNumberSequence(num []int) []int {
	digits := []int{}
	str := strconv.Itoa(num[0])

	for _, char := range str {
		digit := int(char - '0')
		digits = append(digits, digit)
	}
	return digits
}
package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum) //quantos "slice de int" foram passados
	sums := make([]int, lengthOfNumbers) //cria um slice do tamalho de lengthOfNumbers

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}

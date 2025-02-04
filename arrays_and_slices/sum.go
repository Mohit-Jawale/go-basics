package main

func Sum(numbers []int) int {

	sum := 0
	// for i := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }
	// return sum

	for _, number := range numbers {
		sum += number
	}
	return sum
}

// func sumAll(numbersToSum ...[]int) []int {

// 	// lengthofNumbers := len(numbersToSum)
// 	// sums := make([]int, lengthofNumbers)
// 	var sums []int

// 	for _, number := range numbersToSum {
// 		sums = append(sums, Sum(number))
// 	}

// 	return sums

// }
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

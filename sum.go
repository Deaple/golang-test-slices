package main

func Sum(numbers[] int) (sum int){
	for _,number := range numbers {
		sum += number
	}
	return
}

//variadic function
//receives a slice of int slices
//sum the elements of each inner slices
//and return a new slice of all sums
func SumAll(numbersToSum ...[]int) []int {
	length := len(numbersToSum)
	//create a new empty slice to record the sum
	//with the length of total slices passed
	sums := make([]int, length)

	//sum each slice and record the total in slice sum
	for i, numbers := range numbersToSum{
		sums[i] = Sum(numbers)
	}
	
	return sums
}
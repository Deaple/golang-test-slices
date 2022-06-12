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
	//create a new empty slice to record the sum
	var sums []int

	//sum each slice and record the total in slice sum
	for _, numbers := range numbersToSum{
		sums = append(sums,Sum(numbers))
	}
	
	return sums
}
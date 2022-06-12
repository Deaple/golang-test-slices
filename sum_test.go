package main

import "testing"

func TestSum(t *testing.T){
	numbers := [5]int {1,2,3,4,10}

	got := Sum(numbers)

	want := 20

	if got != want{
		t.Errorf("got %d, want %d, given %v",got,want,numbers)
	}
}
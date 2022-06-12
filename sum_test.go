package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T){
	t.Run("collection of 5 numbers", func (t *testing.T){
		numbers := []int {1,2,3,4,5}
	
		got := Sum(numbers)
	
		want := 15
	
		if got != want{
			t.Errorf("got %d, want %d, given %v",got,want,numbers)
		}
		
	})

}

func TestSumAll(t *testing.T) {
	slice1 := []int {1,2}
	slice2 := []int {0,9,1}
	got  := SumAll(slice1, slice2)
	want := []int {3,10}

	if !reflect.DeepEqual(got,want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestSumAllTails(t *testing.T){
	got := SumAllTails([]int{1,2},[]int{5,9},[]int{4,5})
	want := []int{2,9,5}

	if !reflect.DeepEqual(got,want){
		t.Errorf("got %v, wanted %v",got,want)
	}

}
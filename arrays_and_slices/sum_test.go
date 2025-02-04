package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers array", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given,%v", got, want, numbers)

		}

	})

	t.Run("tail of arrays", func(t *testing.T) {

		got := SumAllTails([]int{1, 2}, []int{1, 3, 5})

		want := []int{2, 8}

		assertCorrectMessage(t, got, want)

	})

}

// func TestSumall(t *testing.T) {

// 	got := sumAll([]int{1, 2}, []int{0, 8})

// 	want := []int{3, 8}

// 	assertCorrectMessage(t, got, want)

// }

func assertCorrectMessage(t testing.TB, got, want []int) {

	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v,", got, want)

	}

}

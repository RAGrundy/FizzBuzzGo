package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	tables := []struct {
		x int
		y string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{6, "Fizz"},
		{7, "7"},
		{8, "8"},
		{9, "Fizz"},
		{10, "Buzz"},
		{11, "11"},
		{12, "Fizz"},
		{13, "13"},
		{14, "14"},
		{15, "FizzBuzz"},
		{30, "FizzBuzz"},
		{45, "FizzBuzz"},
	}

	for _, table := range tables {
		returnValue := FizzBuzz(table.x)

		if returnValue != table.y {
			t.Errorf("Didn't Fizz, got: %s, wanted %s.", returnValue, table.y)
		}
	}
}

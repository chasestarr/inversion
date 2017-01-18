package inversion

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestReverse4(t *testing.T) {
	input := []int{4, 3, 2, 1}
	correct := 6
	if output := Count(input); output != correct {
		t.Fatal("count incorrect, expected:", correct, "received:", output)
	}
}

func TestSimple(t *testing.T) {
	input := []int{1, 3, 5, 2, 4, 6}
	correct := 3
	if output := Count(input); output != correct {
		t.Fatal("count incorrect, expected:", correct, "received:", output)
	}
}

func TestReverse(t *testing.T) {
	input := []int{6, 5, 4, 3, 2, 1}
	correct := 15
	if output := Count(input); output != correct {
		t.Fatal("count incorrect, expected:", correct, "received:", output)
	}
}

func readInput(name string) []int {
	input, _ := ioutil.ReadFile(name)
	s := strings.Split(string(input), "\n")
	output := make([]int, len(s))
	for i, item := range s {
		intItem, _ := strconv.Atoi(item)
		output[i] = intItem
	}
	return output
}

func TestLargeSet(t *testing.T) {
	input := readInput("test_input.txt")
	correct := 2407905288
	if output := Count(input); output != correct {
		t.Fatal("count incorrect, expected:", correct, "received:", output)
	}
}

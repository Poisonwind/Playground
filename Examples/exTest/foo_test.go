package main

import "testing"

type testPair struct {
	testValue []int
	testResult int
}

var testsMin = []testPair{
	{[]int{3,4,1,5}, 1},
	{[]int{10,13,15,21}, 10},
}

var testsMax = []testPair{
	{[]int{3,4,1,5}, 5},
	{[]int{10,13,21,15}, 21},
}

func TestMin(t *testing.T) {

	for _, pair := range(testsMin) {
		result := Min(pair.testValue)
		if result != pair.testResult {
			t.Error("For", pair.testValue, "expected: ", pair.testResult, "got :", result)
		}
	}

}

func TestMax(t *testing.T) {

	for _, pair := range(testsMax) {
		result := Max(pair.testValue)
		if result != pair.testResult {
			t.Error("For", pair.testValue, "expected: ", pair.testResult, "got :", result)
		}
	}

}
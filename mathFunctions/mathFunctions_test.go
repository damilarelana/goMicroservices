package mathFunctions

import (
	"reflect"
	"testing"
)

// Define testValues struct i.e. a fieldset of testValues
type testValues struct {
	testSlice    []float64
	averageValue float64
	maxValue     float64
	minValue     float64
	sortedSlice  []float64
}

// Implement the testValues struct with actual testPairs
var testValuesFieldSet = []testValues{
	{[]float64{2, 1}, 1.5, 2, 1, []float64{1, 2}},
	{[]float64{1, 1, 1, 1, 1}, 1, 1, 1, []float64{1, 1, 1, 1, 1}},
	{[]float64{-1, 1}, 0, 1, -1, []float64{-1, 1}},
	{[]float64{1, 2, 4, 3, 5}, 3, 5, 1, []float64{1, 2, 3, 4, 5}},
	{[]float64{}, 0, 0, 0, []float64{}},
}

// TestAverage function tets the Average function in mathPackage
func TestAverage(t *testing.T) {
	for _, testValuesExtract := range testValuesFieldSet {
		returnedValue := Average(testValuesExtract.testSlice)
		if returnedValue != testValuesExtract.averageValue {
			t.Error(
				"For the test slice", testValuesExtract.testSlice,
				"The expected average was", testValuesExtract.averageValue,
				"But we instead got", returnedValue,
			)
		}
	}
}

// TestMin function tests the Min function in mathPackage
func TestMin(t *testing.T) {
	for _, testValuesExtract := range testValuesFieldSet {
		returnedValue := Min(testValuesExtract.testSlice)
		if returnedValue != testValuesExtract.minValue {
			t.Error(
				"For the test slice", testValuesExtract.testSlice,
				"The expected minimum value was", testValuesExtract.minValue,
				"But we instead got", returnedValue,
			)
		}
	}
}

// TestMax function tests the Max function in mathPackage
func TestMax(t *testing.T) {
	for _, testValuesExtract := range testValuesFieldSet {
		returnedValue := Max(testValuesExtract.testSlice)
		if returnedValue != testValuesExtract.maxValue {
			t.Error(
				"For the test slice", testValuesExtract.testSlice,
				"The expected maximum value was", testValuesExtract.maxValue,
				"But we instead got", returnedValue,
			)
		}
	}
}

// TestBubblesort function tests the Bubblesort function in mathPackage
func TestBubblesort(t *testing.T) {
	for _, testValuesExtract := range testValuesFieldSet {
		returnedValue := Bubblesort(testValuesExtract.testSlice)
		// if returnedValue != testValuesExtract.sortedSlice {
		if !reflect.DeepEqual(returnedValue, testValuesExtract.sortedSlice) {
			t.Error(
				"For the test slice", testValuesExtract.testSlice,
				"The expected sorted slice was", testValuesExtract.sortedSlice,
				"But we instead got", returnedValue,
			)
		}
	}
}

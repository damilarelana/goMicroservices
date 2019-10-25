package mathFunctions

// Add function adds up two integers and returns an integer
func Add(x int64, y int64) int64 {
	return x + y
}

// Average function calculates the average of the total sum of the array values
func Average(dynamicArray []float64) float64 { // argument is a slice here
	if len(dynamicArray) != 0 {
		var arrayTotal float64
		for _, value := range dynamicArray {
			arrayTotal += value
		}
		return arrayTotal / float64(len(dynamicArray))
	}
	return 0
}

// Bubblesort function sorts the array values using Bubblesort
// upon completion of the sort , it returns the sorted array
// also note that variadic parameters are being used as arguments
func Bubblesort(dynamicArray []float64) []float64 {
	if len(dynamicArray) != 0 {
		outerCount := 0
	OuterForLoop:
		for outerCount < len(dynamicArray) {
			/* initial flag handles: sorted input, sorting completion, and bubbling loop */
			swapflag := false
			innerCount := 0
			for innerCount < (len(dynamicArray) - 1) {
				if dynamicArray[innerCount] > dynamicArray[innerCount+1] {
					temp := dynamicArray[innerCount+1]
					dynamicArray[innerCount+1] = dynamicArray[innerCount]
					dynamicArray[innerCount] = temp
					swapflag = true
				}
				innerCount++
			}
			/* exiting from loop when already sorted input and sorting completion */
			if !swapflag {
				break OuterForLoop
			}
			outerCount++
		}
		return dynamicArray
	}
	return []float64{}
}

// Max function calls the "bubblesort" function to sort the array values.
// upon completion of the sort smallest value would be at index 0
// it then returns the smallest value at  dynamicArray[len(dynamicArray)-1
// also note that variadic parameters are being used as arguments
func Max(dynamicArray []float64) float64 {
	if len(dynamicArray) != 0 {
		var sortedArray = Bubblesort(dynamicArray)
		return sortedArray[len(sortedArray)-1]
	}
	return 0
}

// Min function sorts the array values using maxBubblesort
// upon completion of the sort smallest value would be at index 0
// it then returns the smallest value at dynamicArray[0]
// also note that variadic parameters are being used as arguments
func Min(dynamicArray []float64) float64 {
	if len(dynamicArray) != 0 {
		var sortedArray = Bubblesort(dynamicArray)
		return sortedArray[0]
	}
	return 0
}

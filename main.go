package main

import "fmt"

func BuildSet(arr []int) []int {
	finalArr := []int{}
	uniqueNums := map[int]bool{}
	for _, v := range arr {
		if ok, _ := uniqueNums[v]; !ok {
			uniqueNums[v] = true
			finalArr = append(finalArr, v)
		}
	}
	return finalArr
}

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	smallerThan := []int{}
	biggerThan := []int{}

	pivot := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] <= pivot {
			smallerThan = append(smallerThan, arr[i])
		} else {
			biggerThan = append(biggerThan, arr[i])
		}
	}
	smallerThan = QuickSort(smallerThan)
	biggerThan = QuickSort(biggerThan)

	smallerThan = append(smallerThan, pivot)
	smallerThan = append(smallerThan, biggerThan...)
	return smallerThan
}

func BinarySearch(arr []int, elementToFind int) bool {

	if len(arr) <= 1 {
		return false
	}

	SortedArray := QuickSort(arr)

	pivot := SortedArray[int(len(SortedArray)/2)]

	foundIt := false

	if pivot == elementToFind {
		foundIt = true
	} else {
		if pivot < elementToFind {
			foundIt = BinarySearch(SortedArray[len(SortedArray)/2:], elementToFind)
		} else {
			foundIt = BinarySearch(SortedArray[:len(SortedArray)/2], elementToFind)
		}
	}
	return foundIt
}

func reverseArray(arr []int) []int {
	newArr := make([]int, len(arr))
	for i := 1; i < int(len(arr)/2)+1; i++ {
		tmp := arr[i-1]
		newArr[i-1] = arr[len(arr)-i]
		newArr[len(arr)-i] = tmp
	}
	return newArr
}

func BuildMatrix(arr [9]int) [][]int {
	row := 0
	column := 0
	newMatrix := make([][]int, 3)
	for i := 0; i < len(newMatrix); i++ {
		newMatrix[i] = make([]int, 3)
	}

	for i := 0; i < len(arr); i++ {
		column = i % 3
		if column == 0 && i > 0 {
			row++
		}
		newMatrix[row][column] = arr[i]
	}

	return newMatrix
}

//Clockwise
func rotateMatrix(oldMatrix [][]int) [][]int {

	newMatrix := make([][]int, 3)
	for i := 0; i < len(newMatrix); i++ {
		newMatrix[i] = make([]int, 3)
	}

	for row := 0; row < len(oldMatrix); row++ {
		for column := 0; column < len(oldMatrix[row]); column++ {
			newMatrix[len(oldMatrix)-1-row][column] = oldMatrix[column][row]
		}
	}

	return newMatrix
}

//Rotate entirely outern side of the matrix in clockwise order
func rotateOuterMatrix(n int) {

	initialMatrix := make([][]int, n)
	currentValue := 1
	for i := 0; i < n; i++ {
		initialMatrix[i] = make([]int, n)
		for j := 0; j < len(initialMatrix[i]); j++ {
			initialMatrix[i][j] = currentValue
			currentValue++
		}
	}
	// Initial Status print
	printMatrix(initialMatrix)

}

func printMatrix(initialMatrix [][]int) {
	for i := 0; i < len(initialMatrix); i++ {
		for j := 0; j < len(initialMatrix); j++ {
			if j == 0 || j == len(initialMatrix)-1 || i == 0 || i == len(initialMatrix)-1 {
				fmt.Printf("\x1b[0m")
			} else {
				fmt.Printf("\x1b[31;1m")
			}
			fmt.Printf("%  v ", initialMatrix[i][j])
			if initialMatrix[i][j] <= 9 {
				fmt.Print(" ")
			}

		}
		fmt.Println("\n")
	}
}

func main() {
	fmt.Println("\n ---------- Start -----------")
	//Arr
	arr := []int{9, 3, 10, 3, 9, 7, 9}
	fmt.Println("Set: ", BuildSet(arr))
	sortedArray := QuickSort(arr)
	fmt.Println("Sorted array: ", sortedArray)
	elementToFind := 7
	fmt.Printf("Binary search of %d: %t \n", elementToFind, BinarySearch(sortedArray, elementToFind))
	reversedArray := reverseArray(sortedArray)
	fmt.Printf("Reversed array %v: %v \n", sortedArray, reversedArray)

	//Matrix
	numsInMatrix := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	generatedMatrix := BuildMatrix(numsInMatrix)
	for i := 0; i < len(generatedMatrix); i++ {
		fmt.Printf("%v \n", generatedMatrix[i])
	}

	fmt.Println("")

	rotatedMatrix := rotateMatrix(generatedMatrix)
	for i := 0; i < len(rotatedMatrix); i++ {
		fmt.Printf("%v \n", rotatedMatrix[i])
	}

	fmt.Println("")
	rotateOuterMatrix(5)

	fmt.Println("\n ---------- Finish -----------")
}

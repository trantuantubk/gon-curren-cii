package main
import (
	"fmt"
	//"sync"
)

// Find the leftmost index of "target" in arr
// if there is no target within arr, 
// return the index of the most recent element with target
func findLowerBound(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + (right - left)/2
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

//
// asynchronous implementation 
// using Goroutines
func asynFindLowerBound(arr []int, target int, leftBorder chan int) {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + (right - left)/2
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	leftBorder <- left
}

func main(){
	fmt.Printf("== Given a sorted array, find the range of elements with value = target\n")
	arr := []int {1,2,3,4,4,4,4,5,6,7,7,8,9}
	target := 4
	// display 
	
	//
	for _, num := range arr {
		fmt.Print(num, " ")
	}
	fmt.Print("\n")
	chLeftBound := make(chan int)
	chRightBound := make(chan int)
	go asynFindLowerBound(arr, target, chLeftBound)
	go asynFindLowerBound(arr, target+1, chRightBound)
	leftBound := <-chLeftBound
	rightBound := <-chRightBound - 1
	//
	if leftBound > rightBound {
		fmt.Printf("[%d,%d]\n",-1,-1)
	} else {
		fmt.Printf("[%d,%d]\n",leftBound,rightBound)
	}
}
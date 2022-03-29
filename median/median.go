package main

import "sort"

func main() {
	arr := [] int32 {1,2}
	println(findMedian(arr))
}

func findMedian(arr []int32) int32 {
    // Write your code here
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	idx := len(arr)/2 + 1
	return arr[idx]
}
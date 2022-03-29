package main

import "fmt"

func main() {
	arr := [] int32 {256741038, 623958417 ,467905213, 714532089, 938071625}
	miniMaxSum(arr)
} 

func miniMaxSum(arr []int32) {
    // Write your code here
	sum := int64(0)
	min, max := arr[0],arr[0]
	for _, v := range arr {
		sum+=int64(v)
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	fmt.Printf("%v %v\n", (sum - int64(max)), (sum - int64(min)))


}
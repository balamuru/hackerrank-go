package main

import "fmt"

func main() {
	arr := [] int32 {0,1,2,3,-4,-5,6}
	plusMinus(arr)
}

func plusMinus(arr []int32) {
    // Write your code here
	sumP, sumZ, sumN := 0.0,0.0,0.0

	for _, v := range arr {
		if (v >0 ) {
			sumP++
		} else if (v < 0 ) {
			sumN++
		} else {
			sumZ++
		}
	}
	fmt.Printf("%.6f\n", float64(sumP)/float64(len(arr)))
	fmt.Printf("%.6f\n", float64(sumN)/float64(len(arr)))
	fmt.Printf("%.6f\n", float64(sumZ)/float64(len(arr)))
}
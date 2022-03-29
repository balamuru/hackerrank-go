package main

	import "math"

func main() {

	
	arr := [][] int32 {
		{1,2,3},
		{4,5,6},
		{9,8,9},

	}
	println (diagonalDifference(arr))
}
 
func diagonalDifference(arr [][]int32) int32 {
    // Write your code here
	diag1, diag2 := int32(0), int32(0)

	for i,_ := range arr{
		diag1 += arr[i][i]
		diag2 += arr[i][len(arr)-i-1]	
	}

	return int32(math.Abs(float64(diag1-diag2)))
}

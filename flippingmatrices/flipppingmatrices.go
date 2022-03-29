package main

func main() {

	matrix := [][]int32{
		{112, 42, 83, 119},
		{56, 125, 56, 49},
		{15, 78, 101, 43},
		{62, 98, 114, 108},
	}
	println(flippingMatrix(matrix))
}

func flippingMatrix(matrix [][]int32) int32 {
	// Write your code here
	total := int32(0)

	n := len(matrix) / 2
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			//there are 3 elements that could possibly be swapped with an element matrix[i,j]
			//compare with each of them to find the max value and add to the total
			
			//compare with rightmost element of row
			max := maximum(matrix[i][j], matrix[i][2*n-1-j])
			//compare with bottommost element of column
			max = maximum(max, matrix[2*n-1-i][j])
			//compare with furthest diagonal element 
			max = maximum(max, matrix[2*n-1-i][2*n-1-j])
			total += max
		}
	}
	return total
}

func maximum(a int32, b int32) int32 {
	if a > b {
		return a
	} else {
		return b
	}
}

package main

func main() {
	arr := []int32{5,2,1,3,1,8,2,2}
	res := countingSort(arr)
	for _,v := range res{
		print(v)
		print(" ")
		
	}
	println()
}


func countingSort(arr []int32) []int32 {
    // Write your code here
	var res [] int32  = make([]int32, 100)
	for _,v := range arr{
		res[v]++
	}

	return res

}

package main

func main() {
	arr := []int32{1,2,1,2,4}
	println(lonelyinteger(arr))
	
}

func lonelyinteger(a []int32) int32 {
    // Write your code here
	var answer int32 = 0
	for _,v := range a{
		answer ^= v
	}

	return answer
}
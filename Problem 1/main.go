package main

import "fmt"

var mySlice = []int{5, 6, 7, 4}
var sum = 9

func main() {
	m := make(map[int]int)

	for i, num := range mySlice {
		complement := sum - num
		if j, found := m[complement]; found {
			fmt.Printf("The sum of %v and %v is %v.\n", mySlice[j], num, sum)
		}
		m[num] = i
	}
}

/*
The time complexity of this algorithm is O(n),
where n is the number of elements in the array "mySlice".
This is because it has only one for loop that iterates through
the array only once, and each lookup in the hash map takes O(1)
time on average. Therefore, the overall time complexity is,O(n) + O(1) = O(n).
*/

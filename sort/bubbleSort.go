package main

import "fmt"

func main() {
	arr := []int{4, 9, 6, 1, 3, 7, 5, 0}
	bubbleSort(arr)
	fmt.Println(arr)
}

//冒泡排序O(n²)
//https://www.cnblogs.com/bigdata-stone/p/10464243.html  详细理解
func bubbleSort(arr []int) {

	n := len(arr)

	//最外层代表循环次数，8个数就循环7次,没有索引指代
	for i := 0; i < n-1; i++ {

		//一共进行n-i-1次比较
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

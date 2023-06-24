package main

import "fmt"

func main() {
	arr := []int{4, 9, 6, 1, 3, 7, 5, 0}
	insertionSort(arr)
	fmt.Println(arr)
}

//插入排序O(n)-O(n²)
func insertionSort(arr []int) {
	n := len(arr)

	//确定循环次数
	for i := 1; i < n; i++ {

		//记录前一个索引
		preIndex := i - 1
		//确定比较次数，有可能不用比较
		for preIndex >= 0 && arr[preIndex] > arr[preIndex+1] {

			arr[preIndex], arr[preIndex+1] = arr[preIndex+1], arr[preIndex]
			preIndex -= 1
		}

	}
}

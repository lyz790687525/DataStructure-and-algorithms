package main

import "fmt"

func main() {
	arr := []int{4, 9, 6, 1, 3, 7, 5, 0}
	selectionSort(arr)
	fmt.Println(arr)
}

//选择排序O(n²)
//第一轮，将数组里的第一个元素和余下元素对比，小的放第一，再将余下的与第一个作对比，小的放第一位。第二轮从第二个元素开始依次类推到最后一位
//直接改变底层数组，不需要额外返回
func selectionSort(arr []int) {

	//都有用到就可以简化
	n := len(arr)
	//这里可以简化将n改为n-1,因为最后一个数不需要再进行比较了
	//这里指找到第一个数开始，有索引指代（冒泡排序最外层循环指的是循环次数递减）
	for i := 0; i < n-1; i++ {

		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

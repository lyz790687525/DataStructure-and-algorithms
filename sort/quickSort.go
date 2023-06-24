package main

import "fmt"

func main() {
	slice := []int{39, 10, 33, 25, 17, 25, 27, 29, 44, 41, 21, 20, 32}
	quickSort(slice)
	fmt.Println(slice)
}

//快排绝大部分情况平均下来O(nlogn)，最差O(n²)几率很低
//通常为最快的排序
func quickSort(slice []int) {

}

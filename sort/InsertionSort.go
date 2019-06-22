package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := []int{ 5, 2, 4, 6, 1, 3}
	insertSort(&arr)

	fmt.Println(arr)
}


func insertSort(arr *[]int) {

	var tempArr = *arr
	var i int
	var j int

	fmt.Println(tempArr)

	var order int = 0
	for i = 1; i < len(tempArr); i++ {
		for j = 0; j < i; j++ {
			if tempArr[j] > tempArr[i] {
				swap(&tempArr[i],&tempArr[j])
			}
			order ++
		}

		fmt.Println(tempArr)
	}

	fmt.Println("order: " + strconv.Itoa(order))

}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*数组练习*/
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Print(arr)

	var arr2 = [3]int{4, 5, 6}
	fmt.Print(arr2)
	var arr3 [3]int = [3]int{7, 8, 9}
	fmt.Print(arr3)

	var arr4 = [...]int{11, 12, 13}
	fmt.Print(arr4)
	var arr5 = [...]int{1: 100, 0: 90, 2: 13}
	fmt.Print(arr5)

	var arr6 [26]byte
	for i := 0; i < 26; i++ {
		arr6[i] = 'A' + byte(i)
	}
	for i := 0; i < 26; i++ {
		fmt.Print(arr6[i])
		fmt.Printf("%c ", arr6[i])
	}
	fmt.Print(arr6)

	var arr7 = [...]int{1, 9, -1, 100, 23}

	max := 0
	suoyin := 0
	total := 0
	for index, value := range arr7 {
		if value > max {
			max = value
			suoyin = index
		}
		total += value
	}
	fmt.Print("\n最大值： ", max)
	fmt.Print("\n索引：", suoyin)
	fmt.Print("\n求和：", total)

	fmt.Print("\n平均数：", total/len(arr7))

	var arr8 [5]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr8); i++ {
		arr8[i] = rand.Intn(100)
	}
	fmt.Print(arr8)
	for i := len(arr8); i > 0; i-- {
		fmt.Print("\n", arr8[i-1])
	}

}

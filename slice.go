package main

import "fmt"

func main() {
	var arr = [5]int{1, 22, 33, 66, 99}
	//表示slice引用arr数组，从下标为1开始，到下标为3结束，但是不包含下标3
	//slice引用数组创建，数组对开发人员是可见的，存在暴露的内存地址的
	var slice = arr[1:3]
	fmt.Print(slice)

	var slice2 = make([]int, 3, 10)//直接创建的slice，同样是引用数组，但是数组地址对开发是不可见的，没有暴露内存地址

	fmt.Print(slice2)

	var slice3 = []int{1, 2, 3}
	fmt.Print(slice3)

	//切片的遍历

	var slice4 = []int{5,22,5,73}

	for i:=0;i<len(slice4);i++{
		fmt.Println("slice4[%v]=%v ",i,slice4[i])
	}

	for k,v :=range slice4{
		fmt.Println("slice4[%v]=%v ",k,v)
	}

	//append给slice追加元素
	slice4 = append(slice4,slice4...)
	fmt.Println("slice4 ",slice4)

	//切片的拷贝
	var slice5 = make([]int ,10)
	copy(slice5,slice4)
	fmt.Println("slice5 ",slice5)
}

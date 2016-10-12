package main

import "fmt"

func main() {

	// 数组定义后，如果没有初始值，则以数组类型的空值填充，例如int类型空值为0
	var arr1 [10]int
	arr2 := arr1
	fmt.Printf("No.1 Output:  Array 1 is: %v, Array 2 is: %v\n", arr1, arr2)
	fmt.Printf("No.2 Output:  Array 1 length is: %d, Array 2 length is %d\n", len(arr1), len(arr2))

	// array 是固定长度的，不可以动态改变长度
	arr1[5] = 10
	arr2[2] = 8
	// arr1[10] = 9   out of bounds for array
	fmt.Printf("No.3 Output:  Array 1 is: %v, Array 2 is: %v\n", arr1, arr2)

	// array 是值类型而非引用类型
	arr3 := arr2
	arr3[5] = 18
	fmt.Printf("No.4 Output:  Array 1 is: %v, Array 2 is: %v, Array 3 is: %v\n", arr1, arr2, arr3)

	// 二维数组定义，设置元素值，数组元素访问
	arr4 := [5][5]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	fmt.Printf("No.5 Output:  Array 4 is: %v\n", arr4)
	arr4[4][2] = 20
	fmt.Printf("No.6 Output:  Array 4 is: %v\n", arr4)
	fmt.Printf("No.7 Output:  Array 4 the element of 1 row and 2 column is: %d\n", arr4[0][1])

	// 可以通过...代替声明数组长度，此时数组将根据元素个数自动计算长度
	arr5 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("No.8 Output:  Array 5 is: %v, length is %d\n", arr5, len(arr5))
}

package main

import "fmt"

func main() {

	// slice定义
	var slice1 []int
	slice2 := slice1
	fmt.Printf("No.1 Output:   Slice 1 is %v, Slice 2 is %v\n", slice1, slice2)
	fmt.Printf("No.2 Output:   Slice 1 length is %d, Slice 1 capacity is %d\n", len(slice1), cap(slice1))

	// 可以通过append方法在slice中追加单个或者多个元素
	slice1 = append(slice1, 1, 2, 3)
	slice2 = append(slice2, 4, 5)
	slice3 := append(slice1, 6, 7, 8, 9)
	fmt.Printf("No.3 Output:   Slice 1 is %v, Slice 2 is %v, Slice 3 is %v\n", slice1, slice2, slice3)

	// append后的结果为一个新的slice对象
	slice1[2] = 100
	slice2[1] = 200
	slice3[4] = 300
	fmt.Printf("No.4 Output:   Slice 1 is %v, Slice 2 is %v, Slice 3 is %v\n", slice1, slice2, slice3)

	// 可以通过make关键字声明一个新的slice，也可以通过索引起止位置声明一个新的slice
	slice4 := make([]int, 3)
	slice4[0] = 1
	slice4[1] = 2
	slice4[2] = 3
	slice5 := slice4[1:]
	fmt.Printf("No.5 Output:   Slice 4 is %v, Slice 5 is %v\n", slice4, slice5)

	// slice为引用类型，源相同的slice均指向同一个内存地址
	slice4[1] = 100
	fmt.Printf("No.6 Output:   Slice 4 is %v, Slice 5 is %v\n", slice4, slice5)

	// len(slice)可以获取slice的长度，cap(slice)可以获取slice的容量
	// 当追加元素时长度超过容量时，容量扩充为`添加元素的个数或者当前元素的个数（去较大者）`的两倍
	fmt.Printf("No.7 Output:   Slice 4 is %v, length is %d, capacity is %d\n", slice4, len(slice4), cap(slice4))
	slice4 = append(slice4, 4, 5)
	fmt.Printf("No.8 Output:   Slice 4 is %v, length is %d, capacity is %d\n", slice4, len(slice4), cap(slice4))
	slice4 = append(slice4, 6, 7, 8, 9, 10, 11, 12, 13, 14)
	fmt.Printf("No.9 Output:   Slice 4 is %v, length is %d, capacity is %d\n", slice4, len(slice4), cap(slice4))

	// copy(dest, src)可以拷贝一个新的slice，两个slice不指向同一个内存
	slice6 := make([]int, len(slice4))
	copy(slice6, slice4)
	fmt.Printf("No.10 Output:  Slice 6 is %v\n", slice6)
	slice4[8] = 500
	fmt.Printf("No.11 Output:  Slice 6 is %v\n", slice6)
}

# Golang 学习笔记

## 语法笔记

### Hello world

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}
```

##### 知识点：

- `main`包下的`main`函数可以运行
- 类库导出的方法首字母为大写，例如`Println`
- `fmt`为标准输入输出库，包含各种输入输出方法，和`C`, `Java`中的方法名类似
- 字符串使用双引号`""`或者反引号包裹，不可以使用单引号`''`
- Go语言中大部分条件下可以省略语句末尾分号
- Go语言中对于代码风格是语法级的强制要求


------

### 变量

```go
var str1 string

var str2, str3 string

var str4 = "variable 4"

var str5, str6 string = "variable str 5", "variable str 6"

str7 := "variable str 7"

str8, str9 := "variable str 8", "variable str 9"
```

##### 知识点：

- 定义变量的关键字为`var`
- 定义变量可以在变量名后面声明变量类型
- 如果没有声明变量类型，编译器会根据变量值分析出变量类型，**Go是强类型语言**
- 可以同时定义多个变量并且给每个变量赋值，例如`str5`和`str6`
- ***可以使用`:=`的简单形式在定义变量的时候省略`var`***


------

### 常量

```go
const cst1 string = "CONSTANT"
```

##### 知识点：

- 常量的定义和变量除去关键字以外基本一致。
- 常量定义没有简略方式
- 常量值在定义之后不可以改变


------

### 分组

```go
import (
	"fmt"
	"os"
)
const (
	groupConstantA = "constant a"
	groupConstantB = "constant b"
)
var (
	groupVariableA = "variable a"
	groupVariableB = "variable b"
)
```

##### 知识点：

- 分组可以用于`import`, `const`和`var`三种定义和引用
- 其他和单独定义一样

------

### iota 累加常量

```go
const a = iota         // 0
const b = iota         // 0
const (
	c = iota           // 0
	d = iota           // 1
	e, f = iota, iota  // 2, 2
)
const (
	g = iota           // 0
	h = iota           // 1
)
const (
	Mon = iota + 2     // 2
	Tue                // 3
	Wed                // 4
	Thu = 100          // 100
	Fri                // 100
	Sat = iota + 3     // 8
	Sun                // 9
)
```

##### 知识点：

- `iota`为累加变量，只可以在常量中使用
- `iota`从`0`开始，并且只在分组定义常量中累加
- 当一个分组定义结束时，`iota`再次从`0`开始
- 当两个变量同时在一行定义时，`iota`的值不进行累加
- 当之定义第一`iota`时，下面定义的常量和第一个的表达式相同，只是iota做累加，并且在`iota`中断时，`iota`本身还会自动累加，只是不再显示，如果需要显示，需要手动恢复


------

### 数组 Array

```go

// 定义数组
var arr1, arr2 = [5]string{}, [5]string{"a", "b", "c"}

// 简略方式定义数组
arr3 := [5]string{"a", "b", "c", "e"}

arr4 := [...]int{1, 2, 3, 4}

// 二维数组
arr5 := [4][4]string{{"a", "b", "c"}, {"e", "f", "g"}, {"h", "i", "j", "k"}, {"x", "y"}}

len(arr3)  // 4
```

##### 知识点：

- ***数组是定长的，长度在定义数组是需要声明***
- ***可以通过`...`来自动计算数组长度***
- 可以通过`len(arr)`来获取数组长度
- 定义数组和定义变量类似，可以使用简略方式定义
- 数组赋值采用索引赋值，即`arr[0] = "xxx"`
- 可以采用`start:end`的形式访问连续下标的元素，包含`start`，不包含`end`
- `arr[:5]`代表`arr[0:5]`, `arr[1:]`代表`arr[1:length]`
- 二维数组采用二维索引值进行访问和赋值，即`arr[0][0]`


------

### 切片 Slice

```go
// slice定义
var slice1 []int
slice2 := make([]int, 10)

// 可以通过append方法在slice中追加单个或者多个元素
slice1 = append(slice1, 1, 2, 3)
slice2 = append(slice2, 4, 5)
slice3 := append(slice1, 6, 7, 8, 9)
/**
 * results: 
 * slice1: [1, 2, 3]
 * slice2: [4, 5]
 * slice3: [1, 2, 3, 6, 7, 8, 9]
 */

// append后的结果为一个新的slice对象
slice1[2] = 100
slice2[1] = 200
slice3[4] = 300
/** 
 * results:
 * slice1: [1, 2, 100]
 * slice2: [4, 200]
 * slice3: [1, 2, 3, 6, 300, 8, 9]
 */

// 可以通过make关键字声明一个新的slice，也可以通过索引起止位置声明一个新的slice
slice4 := make([]int, 3)
slice4[0] = 1
slice4[1] = 2
slice4[2] = 3
slice5 := slice4[1:]
/**
 * results:
 * slice4: [1, 2, 3]
 * slice5: [2, 3]
 */


// slice为引用类型，源相同的slice均指向同一个内存地址
slice4[1] = 100
/**
 * results:
 * slice4: [1, 100, 3]
 * slice5: [100, 3]
 */

// len(slice)可以获取slice的长度，cap(slice)可以获取slice的容量
// 当追加元素时长度超过容量时，容量扩充为`添加元素的个数或者当前元素的个数（去较大者）`的两倍

/* slice4 length is 3, capacity is 3 */

slice4 = append(slice4, 4, 5)

/* slice4 length is 5, capacity is 6 */

slice4 = append(slice4, 6, 7, 8, 9, 10, 11, 12, 13, 14)

/* slice4 length is 14, capacity is 14 */

// copy(dest, src)可以拷贝一个新的slice，两个slice不指向同一个内存
slice6 := []int{1, 2, 3}
slice7 := make([]int, len(slice6))
copy(slice7, slice6)
/**
 * results:
 * slice6: [1, 2, 3]
 * slice7: [1, 2, 3]
 */

slice6[1] = 500
/**
 * results:
 * slice6: [1, 500, 3]
 * slice7: [1, 2, 3]
 */
```


------

### Map

```go
map1 := make(map[string]int)
map2 := map[string]int {"a": 1, "b": 2, "c": 3}

map1["first"] = 1
map1["second"] = 2

e1, exist1 := map2["a"]  // 1, true
e2, exist2 := map2["d"]  // 0, false

delete(map2, "b")
len(map2) // the number of keys

```

##### 知识点：

- 声明`map`和声明数组类似，可以通过`make`定义，结构为`map[keyType]valueType`，也可以定义并且初始化值
- 设置`map`中的值可以通过`map[key] = value`的形式
- 访问`map`中元素的值可以通过`key`访问，返回结果又两个值，第一个是`key对应的value`，第二个是`该key是否存在，存在返回true，不存在返回false`
- 可以通过`len(map)`方法查询`map`中`key`的数量
- 可以通过`delete(map, key)`方法来删除`map`中的`key`

### 流程控制

### 条件判断 if...else...

```go
if x := 10; x > 5 {
	do something
} else if x == 5 {
	do something
} else {
	do something
}

```

##### 知识点

- 逻辑和其他数语言相同，通过`if`判断条件是否为`true`，不为`true`则进入`else`语句
- Go语言的语句不要小括号`()`包裹
- **在语句中可以定义变量，该变量只在当前作用域生效**

### 循环 for

```go

```
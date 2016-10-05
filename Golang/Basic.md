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
```

##### 知识点：

- `iota`为累加变量，只可以在常量中使用
- `iota`从`0`开始，并且只在分组定义常量中累加
- 当一个分组定义结束时，`iota`再次从`0`开始
- 当两个变量同时在一行定义时，`iota`的值不进行累加


------

### 数组 Array

```go

// 定义数组
var arr1, arr2 = [5]string{}, [5]string{"a", "b", "c"}

// 简略方式定义数组
arr3 := [5]string{"a", "b", "c", "e"}

arr4 := arr3[:3]

arr4[2] = "x"

fmt.Printf("%s\n%s", arr3, arr4) 
// ["a", "b", "x", "e", ""]
// ["a", "b", "x"]

// 二维数组
arr5 := [4][4]string{{"a", "b", "c"}, {"e", "f", "g"}, {"h", "i", "j", "k"}, {"x", "y"}}

len(arr3)  // 4
len(arr4)  // 3
```

##### 知识点：

- ***数组是定长的，长度在定义数组是需要声明***
- 可以通过`len(arr)`来获取数组长度
- 定义数组和定义变量类似，可以使用简略方式定义
- 数组赋值采用索引赋值，即`arr[0] = "xxx"`
- 数组传值为引用传值，例如`arr3`和`arr4`(这种实质上是定义切片，参照下一节)
- 可以采用`start:end`的形式访问连续下标的元素，包含`start`，不包含`end`
- `arr[:5]`代表`arr[0:5]`, `arr[1:]`代表`arr[1:length]`
- 二维数组采用二维索引值进行访问和赋值，即`arr[0][0]`


------

### 切片 Slice

```go

```
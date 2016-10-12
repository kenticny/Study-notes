package main

import (
	"fmt"
)

func main() {

	// iota只可以在常量中使用
	// iota的值是从0开始累加的，同一行使用逗号分隔的iota值相同，并且同一行内不会造成累加
	const (
		a    = iota
		b    = iota
		c, d = iota, iota
		x    = iota
	)
	fmt.Printf("No.1 Output:  Const a is %d, Const b is %d, Const c is %d, Const d is %d, Const x is %d\n",
		a, b, c, d, x)

	// 使用单独常量赋值iota不会造成累加，每一次定义常量都会重新累计
	const e = iota
	const f = iota
	fmt.Printf("No.2 Output:  Const e is %d, Const f is %d\n", e, f)

	// 常量组中省略赋值的部分，将和上一个定义的表达式相同，iota正常累加
	const (
		Mon = iota
		Tue
		Wed
		Thu
		Fri
		Sat
		Sun
	)
	fmt.Printf("No.3 Output:  Mon: %d, Tue: %d, Wed: %d, Thu: %d, Fri: %d, Sat: %d, Sun: %d\n",
		Mon, Tue, Wed, Thu, Fri, Sat, Sun)

	// 当iota不为显式时，累加正常进行
	// 当需要iota为显式时，需要手动恢复
	const (
		g = iota * 3
		h
		i
		j
		k = "Hello"
		l
		m = iota + 5
		n
	)
	fmt.Printf("No.4 Output:  g: %v, h: %v, i: %v, j: %v, k: %v, l: %v, m: %v, n: %v\n",
		g, h, i, j, k, l, m, n)

}

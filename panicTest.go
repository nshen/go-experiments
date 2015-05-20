package main

import (
	"fmt"
)

func panicTest() {
	defer func() {
		if x := recover(); x != nil {
			switch x.(type) {
			case int: //捕捉自己抛出的错误
				fmt.Println("Function panicked with a very unhelpful error", x)
			default:
				panic(x) //runtime 抛出的错误,重新抛出不处理
			}
		}
	}() //立即执行

	badFunction()
}

func badFunction() {
	fmt.Println("选择panic类型(0=no panic, 1=int, 2=runtime panic)")
	var choice int
	fmt.Scanf("%d", &choice) //%d 10进制整数
	switch choice {
	case 1:
		panic(0)
	case 2:
		var invalid func()
		invalid()
	}
}

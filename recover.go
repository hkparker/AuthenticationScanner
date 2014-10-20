package main

import (
	"fmt"
)

func recovery_test(i bool) (exit_code int) {
	defer func() {
        if r := recover(); r != nil {
            fmt.Println("Error in func")
            exit_code = 1
        }
    }()
    if i {
		panic("panic")
	}
    exit_code = 0
    return
}

func main() {
	a := recovery_test(false)
	fmt.Println("func returned", a)
}



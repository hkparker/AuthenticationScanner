package main

import (
	"fmt"
)

func recovery_test(i bool) (exit_code int) {
	defer func() {
        if r := recover(); r != nil {
            fmt.Println("Error in func,", r)
            exit_code = 1
        }
    }()
    if i {
		panic("panic on line 15 of recover.go")
	}
    exit_code = 0
    return
}

func main() {
	a := recovery_test(true)
	fmt.Println("func returned", a)
}



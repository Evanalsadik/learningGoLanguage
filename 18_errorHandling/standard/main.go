package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := devide(67, 0)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	fmt.Println("result :", result)
}

func devide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot devided by zero")
	}
	return a / b, nil
}

package main

import (
	"fmt"
	"io"
	"os"
)

var block = "package"

func main() {
	block := "function"
	{
		block := "inner"
		fmt.Printf("The block is %s.\n", block)
	}
	fmt.Printf("The block is %s.\n", block)
	declare()
}

// 可重名变量

func declare() {
	var err error
	n, err := io.WriteString(os.Stdout, "Hello, everyone!\n")
	fmt.Println(n)
	fmt.Println(err)
}

// 变量重声明

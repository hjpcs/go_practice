package main

import (
	"flag"
	"fmt"
	"os"
)

var name string
var number int
var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init() {
	//flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	//flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	cmdLine.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "Usage of %s:\n", "questions")
		cmdLine.PrintDefaults()
	}
	//flag.StringVar(&name, "name", "everyone", "The greeting object.")
	cmdLine.StringVar(&name, "name", "everyone", "The greeting object.")
	cmdLine.IntVar(&number, "number", 0, "The greeting int.")
}

func main() {
	// 自定义参数使用说明
	//flag.Usage = func() {
	//	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	//	flag.PrintDefaults()
	//}
	//flag.Parse()
	_ = cmdLine.Parse(os.Args[1:])
	fmt.Printf("Hello, %s!\n", name)
	fmt.Printf("Hello, %d!\n", number)
}

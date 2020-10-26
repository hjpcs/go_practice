package main

import (
	//. "package_import/p1/example"
	example1 "package_import/p1/example"
	example2 "package_import/p2/example"
)

// 如果example中声明的包名一样，则会有冲突，如果不一样，则不会冲突
// 如果有冲突的话，可以在导入时给包起别名避免冲突
// 如果使用 . 操作引入包，调用函数时直接使用函数名即可

func main() {
	var name = "package"
	example1.Hello(name)
	example2.Hello(name)
	//Hello(name)
}

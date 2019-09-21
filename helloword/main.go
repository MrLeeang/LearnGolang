package main

import "fmt"

// 全局变量
var dir = "/home/lihw"
// 常量
const dirname = "/root"
// 全局变量和常量都可以批量声明
var (
    dir1 = "dir1"
    dir2 = "dir2"
)

const (
    dir3 = "dir3"
    dir4 = "dir4"
)

// iota是go语言的常量计数器，只能在常量的表达式中使用'
const (
    _  = iota
    x = 10 * iota
    y = 10 * iota
    z = 10 * iota
    v = 10 * iota
    w = 10 * iota
)


func test()(string, int){
    // 声明返回值类型
    return "12312", 123
}


func main(){
    fmt.Println("hello world")
    // 标准类型
    var name string
    name = "sDS"
    fmt.Println(name)
    // 批量声明
    var (
        name1 string
        name2 string
    )
    name1 = "xiaoming"
    name2 = "xiaohong"
    fmt.Println(name1, name2)
    // 类型推导
    var name3 = "lihongwei"
    fmt.Println(name3)
    // 短变量声明
    a := "aaaa"
    fmt.Println(a)
    // 声明多个
    var name4, age = "Q1mi", 20
    name5, age1 := "Q1mi", 20
    fmt.Println(name4, age)
    fmt.Println(name5, age1)
    // 匿名变量
    ba, _ := test()
    fmt.Println(ba)
    fmt.Println(x, y, z, v, w)
}

package main

import (
	"fmt"
	"strings"
)


// byte和rune类型
// 遍历字符串
func traversalString() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
	// 反向遍历字符串
	s1 := []byte(s) 
	for i := len(s1)-1; i >= 0; i-- { //byte
		fmt.Printf("%v(%c) ", s1[i], s1[i])
	}
	fmt.Println()
}

// 修改字符串
func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}


func main()  {
	// 单行和多行
	a := "adfafa"
	b := `
		fda
		fda
		fdas
		fdas
		fads
		d
	`
	fmt.Println(a)
	fmt.Println(b)

	// len 求长度，字节类型
	c := "hello"
	d := "hello沙河" // utf8 编码中文占三个字节
	fmt.Println(len(c))
	fmt.Println(len(d))
	// 拼接字符串
	e := "e"
	f := "f"
	fmt.Println(e+f)
	g := fmt.Sprintf(e+f)
	fmt.Println(g)

	// 字符串分割
	name := "li-hong-wei"
	h := strings.Split(name, "-")
	fmt.Println(h)
	// 判断是否包含
	fmt.Println(strings.Contains(name, "li"))

	// 前缀/后缀判断
	fmt.Println(strings.HasPrefix(name, "li"))
	fmt.Println(strings.HasSuffix(name, "li"))

	// 子串出现的位置
	fmt.Println(strings.Index(name, "i"))
	fmt.Println(strings.LastIndex(name, "i"))

	//join操作
	fmt.Println(strings.Join(h, ";"))
	traversalString()
}

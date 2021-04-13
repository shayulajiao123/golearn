package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)


func main(){
	b:=os.Args[1]
	c:=os.Args[2]

	fmt.Println(b)
	fmt.Println(c)
	os.Exit(222)

	//1.1 hello world go build  learn1.go生成了一个可以执行的二进制文件（windows下是exe文件）
	//笔记:go语言的代码通过包的形式组织，一个包位由位于单个目录下的一个或多个.go源代码文件组成。目录定义包的作用，每个源文件（包）都以
	//一条package声明语句开始。
	//main包比较特殊，它定义了一个独立可执行的程序，而不是库。在main包中的main函数，是整个程序的入口。
	//导入了不需要的包和缺少了必要的包，程序都无法编译。import用来导入包
	//fmt.Print("\n","hello world!")
	//1.2 命令行参数
	//
	//var s,sep string
	//for i := 1;i<len(os.Args);i++{
	//	s += sep + os.Args[i]
	//	sep = " "
	//}
	//fmt.Println("\n",s)
	////fmt.Printf("1----\n",time.Now().Unix())
	//os.Exit(123)
	//
	////另一种写法
	//s,sep = ""," "
	//for _,arg :=range os.Args[1:]{
	//	s += sep + arg
	//	sep = " "
	//}
	//fmt.Println("\n",s)
	//fmt.Printf("2----\n",time.Now().Unix())
	//
	////简单写法
	//fmt.Println(strings.Join(os.Args[1:], " "))
	//练习 修改上边程序，使其能够打印os.Args[0]，即被执行命令本身的名字。
	//fmt.Printf("3----\n",time.Now().Unix())
	//fmt.Println(strings.Join(os.Args, " "))
	//fmt.Printf("4-----\n",time.Now().Unix())
	//练习 修改上边程序，其打印每个参数的索引和值，每个一行。
	//练习 做实验测量潜在低效的版本(1和2写法)和使用了strings.Join(简单写法)的版本的运行时间差异。

	//os包接收命令行参数，_ 下划线用来标识无用变量
	//1.3 查找重复的行 使用map。 map存储量key/value的集合，make用来创建空map
	// fmt 的打印
	// %d          十进制整数
	//%x, %o, %b  十六进制，八进制，二进制整数。
	//%f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
	//%t          布尔：true或false
	//%c          字符（rune） (Unicode码点)
	//%s          字符串
	//%q          带双引号的字符串"abc"或带单引号的字符'c'
	//%v          变量的自然形式（natural format）
	//%T          变量的类型
	//%%          字面上的百分号标志（无操作数）
	//
	//fmt.Printf("--1.3例子1----\n")
	//counts := make(map[string]int)
	//input :=bufio.NewScanner(os.Stdin)
	//for input.Scan(){
	//	counts[input.Text()]++
	//}
    //for line,n := range counts{
    //	if n>1{
    //		fmt.Printf("%d\t%n",n,line)
	//	}
	//}
	//1.4 gif动画 见gif.go文件
	//1.5 获取url
	fmt.Printf("--1.5例子1----\n")
	for _,url := range os.Args[1:]{
		fmt.Printf("--1.5例子输出----\n",url)
		resp,err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr,"fetch: %v\n",err)
		}
		b,err := ioutil.ReadAll(resp.Body)
    	resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr,"fetch:reading %s : %v\n",url,err)
		}
		fmt.Printf("%s",b)
	}
	//编译 go build learn1.go
	//执行编译文件输出结果到文件 learn1.exe http://gopl.io > aa.html

	//1.6 并发获取多个url


}

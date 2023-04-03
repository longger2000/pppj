package main

import (
	"fmt" //输出包
	"math/rand"
	"strconv" //基本数据类型转字符串，字符串转基本数据类型。//注意转成有效数据类型
	"strings" //字符串的转换 ，查询 ...等。
	"time"    //时间包
	"unsafe"  //查看字节
)

// for range遍历数组
func ccc() {
	var heroes [3]int = [3]int{1, 2, 3}
	for i, v := range heroes { //i = 下标 v= 下标对应的值
		fmt.Printf("i=%v v=%v", i, v)
	}
	return
}

//unix nuixnano 时间戳

//捕获异常 defer + recover
//defer func (){
//	err := recover()
//	if err != nil{  捕获到错误
//		fmt.Printf("err=",err)
//		fmt.Println("发送给邮件管理@")
//	}
//}

func cal(n int) int {
	var num int
	if n == 10 {
		//fmt.Printf("n==10", n)
		num = 1
	} else {
		//println("else->", n)
		num = (cal(n+1) + 1) * 2
	}

	return num
}

func call(n1 int) {
	n1 = n1 + 1
	fmt.Println("n1-call=", n1)
}

func main() {

	fmt.Printf("第一天是多少", cal(1), 000000)

	n2 := 10
	call(n2)
	fmt.Println("n1-main=", n2)

	var (
		j uint8 = 255
		i int   = 10
	)

	//字符串拼接
	var v string = "炸麻花" + "你好"
	v += "111"
	var g = 1000
	var str string
	fmt.Println("hello,world")
	fmt.Println("i=", i)
	fmt.Println("J=", j)
	fmt.Println(v)

	//查看使用了多少字节
	fmt.Println(unsafe.Sizeof(j))
	fmt.Printf("使用的数据类型%T", g)
	fmt.Println("g")

	//数据类型转字符串，定义一个空的字符串,使用的strconv函数
	str = strconv.FormatInt(int64(i), 10)
	fmt.Printf("str type%T str=%q\n", str, str)

	//
	str = fmt.Sprintln("%d", i)
	fmt.Println("str type %T str=%v", str, str)

	//使用strconv中的函数的Itoa
	str = strconv.Itoa(i)
	fmt.Printf("str type%T str=%q\n", str, str)

	//string转基本数据类型 strconv.parse
	var str1 string = "ture"
	var b bool
	b, _ = strconv.ParseBool(str1)
	fmt.Printf("b type%T b=%v\n", b, b)

	//字符串转整数strconv.Atoi("")
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换成功", n)
	}

	//整数转字符串strconv.Itoa()
	str33 := strconv.Itoa(123)
	fmt.Println("str33=", str33)

	//字符串转byte   []byte("")
	var bytes = []byte("123")
	fmt.Println("bytes=", bytes)

	//查找字符串 是否在 指定字符串内 strings.Contains("","")
	c1 := strings.Contains("seafood", "foo")
	fmt.Println("", c1)

	//统计一个字符串里面有几个字符串 strings.Count("","")
	c2 := strings.Count("seafood", "foo")
	fmt.Println("", c2)

	/*指针，
	1.ptr是一个指针模型
	2.ptr的类型是*int
	3.ptr本身的值是&i
	*/
	//基本数据类型布局
	var c int = 10
	var ptr *int = &c
	fmt.Printf("prt=%v", ptr)
	fmt.Printf("prt指向的值=%v\n", *ptr)

	//从键盘输入
	var name string
	var age int32
	var sal float64
	fmt.Println("请输入名字")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水")
	fmt.Scanln(&sal)
	fmt.Printf("名字 %v \n 年龄 %v  \n 薪水 %v \n", name, age, sal)

	//方法2
	//fmt.Scanf

	//switch case default
	var char byte
	fmt.Println("请输入一个字符")
	fmt.Scanln(&char)
	switch char {
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
	case 'c':
		fmt.Println("C")
	case 'd':
		fmt.Println("D")
	default:
		fmt.Println("输入错误")

		//for遍歷字符串
		var str string = "hello,world"
		str2 := []rune(str)
		for i := 0; i < len(str2); i++ {
			fmt.Printf("%c \n", str2[i]) //使用的下标
		}
		fmt.Println()
	}

	//s使用for rang來遍歷
	str = "abc~ok"
	for index, val := range str {
		fmt.Printf("index= %d val=%c \n", index, val)

		//生成随机数
		rand.Seed(time.Now().UnixNano()) //返回一个从1970年到现在的秒数
		n := rand.Intn(100) + 1          //生成一个1-100的随机数
		fmt.Println("n=", n)

		//break 后面加标签  可以选着跳出的循环
		// lable1
		//lable2
		// break lable1；跳出1的循环

		//continue 结束一次循环

		//字符串遍历 len []rune()
		str11 := "hello"
		fmt.Println("str11 len=", len(str11))
		str22 := "hello张"
		l := []rune(str22)
		for i := 0; i < len(l); i++ {
			fmt.Printf("字符串=%c\n", l[i])

			//切片slice  make
			var intarr [5]int = [...]int{1, 2, 3, 4, 5}
			slice := intarr[1:3]
			fmt.Println("slice= ", slice) //输出下标1到3的
			//var slice []int = make([]int, 2, 3z)//通过make创造切片
			//slice[1] = 10
			//slice[3] = 30
			//fmt.Println("slice= "slice)
			slice = append(slice, 6) //appnd 动态增加
			var slice2 = make([]int, 10)
			copy(slice2, slice) //把slice拷贝个slice2
		}
	}
}

package main

import "fmt"

//强调:go语言程序的入口是main函数，go语言规定，main函数所在的软件包包名必须是main
func main() {
	/**
	   单变量的定义和声明:
	    var age int
		age = 18
		fmt.Println(age)
	 */
	var a, b, c int
	// 3、4、5
	a = 3
	b = 4
	c = 5
	fmt.Println(a,b,c)

	//对声明的变量值进行修改
    a = 6
    b = 8
    c = 10
    fmt.Println(a,b,c)

    //a = "于洪伟" // 表示 语法错误
    // a = 于洪伟 双引号引起来的表示的是字符串

    /**
     格式二：var 变量1，变量2，变量3 int = 3,4,5
             var   st1, str2, str3  string = "yu", "hong", "wei"
     */
    var st1, st2, st3 string = "yu", "hong", "wei"
    fmt.Println(st1,st2,st3)

    /**
      格式三：var 变量1，变量2，变量3  = 数值1，数值2，数值3
      英文单词：address：地址
	*/
    var name, age, address = "于洪伟", 18, "山东省泰安市泰山风景区" // 不同的数据类型的多变量声明方式
    fmt.Println(name,age,address)
	fmt.Printf("name的数据类型是:%T\n",name)
    // name = 20 // go语言是强类型，因为name已经在初始化的时候被设置成了字符串，所以其类型是字符串类型
    //格式化打印 %s是代表获取字符串的内容值
    fmt.Printf("name的内容是:%s,  其数据类型是:%T\n",name,name)
    fmt.Printf("age的数值是:%d,其数据类型是：%T\n",age,age)

    /**
       格式四：  变量1，变量2, 变量3 := 数值1，数值2，数值3
     */
    // 两个数之和  num1, num2
    num1, num2 := 1, 3
    fmt.Println(num1 + num2)
    //注意1： 只有已经被声明的变量才能被使用
    //fmt.Println(num3) // 没有声明和定义的变量不能够被使用

    num1, num2  = 5,8// 多个变量的赋值
    fmt.Println(num1,num2)

    num1, num3 := 10, 20// 在计算机当中有几个num1 答案是：1个
    fmt.Println(num1,num3)

    //num1 := 50 //为什么报错? 原因 := 表示的是新定义一个变量, num1已经在前面定义过了，在计算机中变量的名字保持唯一
    //fmt.Println(num1)
}

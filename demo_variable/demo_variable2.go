package main

import "fmt"

func main() {
	/**
	   变量的默认值
	 */
	var age0 int //默认值是10
	fmt.Println(age0)

	var name0 string //  默认值是空字符串 ""
	fmt.Println("name的值是：",name0)

	var grade float32 // grade 成绩
     fmt.Println(grade)//float 数据类型默认值0

     var st1,st2  = "yu","hongwei"
     fmt.Println(st1,st2)
     
     //变量的集合
     var (
     	name = "yuhongwei"
     	sex = "男"
     	address = "山东省泰安市泰山风景区。。。。"
     	age = 18
	 )
     fmt.Println(name,sex,address,age)

     // 舍弃变量
     num1, _ := 1,2 // _符号表示的将该位置上的变量进行舍弃
     //fmt.Println(num1,num2)
     fmt.Println(num1)
}

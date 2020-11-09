package main

import "fmt"

func main() {
	// 数学知识: π 是一个固定数值：3.1415926 - 3.1415927 , 往往设置 π = 3.14。
	//格式一：const 常量名字 数据类型 = 内容值
	const PAI float32 = 3.14 // 定义一个常量，名字为PAI、类型是float32，具体数值为3.14
	// PAI = 3.1415926 在定义的时候const表明 该变量是一个常量，常量是不可修改的
    const BAIDU = "http://www.baidu.com" //定义个常量，名字为BAIDU，具体的值为 ...
    var address = "山东省泰安市泰山风景区.."

	//规范: 常量值在定义的时候，其名字往往会使用全部大写来进行和变量区分
	fmt.Println(BAIDU)//全部大写的通常是常量
	fmt.Println(PAI)
    fmt.Println(address)//小写的通常就是变量
    //驼峰命名法：变量、函数、结构体

    //常量的集合
    const (
    	num = 100
    	name = "yuhongwei"
    	sex = "男"
	)

    // 常量的举例
    const (
		MONDAY = "星期一"
		Tuesday = "星期二"
		Wednesday = "星期三"
		Thursday = "星期四"
	)

    // 四季：穿下秋冬、一年十二月、性别

    // 常量组
    const (
    	x int = 10 // x常量，int类型，数值10
		y
		z
    	//y int = 10 // y常量，int类型，数值10
    	//z int = 10 // z常量, int类型, 数值10
    	name1 string = "yuhongwei"
    	sex1 string = "男"
	)
    fmt.Println(x,y,z)

}

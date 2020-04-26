package main


import "fmt"

func main() {
	//获取输入的写法
	var (
		name string
		age int
		married bool
	)

	//fmt.Scan(&name, &age, &married)
	//fmt.Printf("扫描结果是 name:%s, age:%d, married:%t", name, age, married)


	//fmt.Scanln(&name, &age, &married)  遇到换行才结束
	fmt.Scan(&name, &age, &married)

	//scanf 执行了输入的内容的格式
	//fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)

	//需要扎实的基础
	//命令行获取参数
	//if(len(os.Args) > 0){
	//	for index,arg := range os.Args{
	//		//fmt.Println(index, arg )
	//		fmt.Printf("args[%d]=%v\n", index, arg )
	//	}
	//}

	//flag包的使用 有以下两种常用的定义 flag参数的方法

	//name := flag.String("name", "张三", "姓名")
	//fmt.Println(&name)
	//fmt.Println(flag.Args())

	//var name string
	//flag.StringVar(&name, "name", "liuwei", "name")
	//
	//var age int
	//flag.IntVar(&age, "age", 18, "年龄")
	//
	//var married bool
	//flag.BoolVar(&married, "married", false, "婚否")
	//
	//var delay time.Duration
	//flag.DurationVar(&delay, "delay", 1, "时间间隔")
	////======================================================
	////解析命令行参数
	//flag.Parse()
	//fmt.Println(name, age, married, delay)
	//
	////返回命令行参数后的其他参数
	//fmt.Println(flag.Args())
	////返回命令行参数后的其他参数个数
	//fmt.Println(flag.NArg())
	////返回使用的命令行参数个数
	//fmt.Println(flag.NFlag())
}


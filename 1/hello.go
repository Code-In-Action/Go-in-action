//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println("hello world")
//}

//指针
//package main
//
//import "fmt"
//
//func main() {
//	i, j := 42, 2701
//
//	p := &i         // point to i
//	fmt.Println(*p) // read i through the pointer
//	*p = 21         // set i through the pointer
//	fmt.Println(i)  // see the new value of i
//
//	p = &j         // point to j
//	*p = *p / 37   // divide j through the pointer
//	fmt.Println(j) // see the new value of j
//}

//结构体
//package main
//
//import "fmt"
//
//type Vertex struct {
//	X int
//	Y int
//}
//
//func main() {
//	fmt.Println(Vertex{1, 2})
//}

//结构体字段
//package main
//
//import "fmt"
//
//type Vertex struct {
//	X int
//	Y int
//}
//
//func main() {
//	v := Vertex{1, 2}
//	//v.X = 4
//	fmt.Println(v.X)
//}

// 结构指针
// 结构体字段可以通过结构体指针来访问。

//通过指针间接的访问是透明的。
//package main
//
//import "fmt"
//
//type Vertex struct {
//	X int
//	Y int
//}
//
//func main() {
//	v := Vertex{1, 2}
//	p := &v
//	p.X = 1e9
//	fmt.Println(v)
//}

// 数组
//package main
//
//import "fmt"
//
//func main() {
//	var a [3]string
//	a[0] = "Hello"
//	a[1] = "World"
//	fmt.Println(a[0], a[1])
//	fmt.Println(a)
//}

// 切片
//package main
//import "fmt"
//
//func main() {
//	p := []int{2, 3, 5, 7, 11, 13}
//	fmt.Println("p ==", p)
//
//	for i := 0; i < len(p); i++ {
//		fmt.Printf("p[%d] == %d\n", i, p[i])
//	}
//}
//package main
//
//import "fmt"
//
//func main() {
//	p := []int{2, 3, 5, 7, 11, 13}
//	fmt.Println("p ==", p)
//	fmt.Println("p[1:4] ==", p[1:4])
//
//	// 省略下标代表从 0 开始
//	fmt.Println("p[:3] ==", p[:3])
//
//	// 省略上标代表到 len(s) 结束
//	fmt.Println("p[4:] ==", p[4:])
//}


// 构造 slice
//package main
//
//import "fmt"
//
//func main() {
//	a := make([]int, 5)
//	printSlice("a", a)
//	b := make([]int, 0, 5)
//	printSlice("b", b)
//	c := b[:2]
//	printSlice("c", c)
//	d := c[2:5]
//	printSlice("d", d)
//}
//
//func printSlice(s string, x []int) {
//	fmt.Printf("%s len=%d cap=%d %v\n",
//		s, len(x), cap(x), x)
//}

//package main
//
//import "fmt"
//
//func main() {
//	var z []int
//	fmt.Println(z, len(z), cap(z))
//	if z == nil {
//		fmt.Println("nil!")
//	}
//}

//package main
//
//import "fmt"
//
//var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
//
//func main() {
//	for i, v := range pow {
//		fmt.Printf("2**%d = %d\n", i, v)
//	}
//}


//package main
//
//import "fmt"
//
//func main() {
//	pow := make([]int, 10)
//	for i := range pow {
//		pow[i] = 1 << uint(i)
//	}
//	fmt.Printf("%d\n", pow)
//	for _, value := range pow {
//		fmt.Printf("%d\n", value)
//	}
//}

// map 的文法
//package main
//
//import "fmt"
//
//type Vertex struct {
//	Lat, Long float64
//}
//
//var m = map[string]Vertex{
//	"Bell Labs": Vertex{
//		40.68433, -74.39967,
//	},
//	"Google": Vertex{
//		37.42202, -122.08408,
//	},
//}
//
//func main() {
//	fmt.Println(m)
//}
//
//package main
//
//import "fmt"
//
//func main() {
//	m := make(map[string]int)
//
//	m["Answer"] = 42
//	fmt.Println("The value:", m["Answer"])
//
//	m["Answer"] = 48
//	fmt.Println("The value:", m["Answer"])
//
//	delete(m, "Answer")
//	fmt.Println("The value:", m["Answer"])
//
//	v, ok := m["Answer"]
//	fmt.Println("The value:", v, "Present?", ok)
//}

//函数
//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//func main() {
//	hypot := func(x, y float64) float64 {
//		return math.Sqrt(x*x + y*y)
//	}
//
//	fmt.Println(hypot(3, 4))
//}

//函数的闭包
//Go 函数可以是闭包的。闭包是一个函数值，它来自函数体的外部的变量引用。 函数可以对这个引用值进行访问和赋值；换句话说这个函数被“绑定”在这个变量上。
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

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
package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

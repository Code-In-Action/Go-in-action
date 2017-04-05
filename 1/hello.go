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
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
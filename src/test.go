package main

import (
	"fmt"
	//"parser"
	//"errors"
)

type test struct {
	a string
	b string
}

func main()  {
	//parser.Parse()
	//fmt.Println("111")
	//test := make([]int, 0)
	//test = append(test, 2)

	//error := errors.New("1123321")
	//errorMap := make(map[int]string, 0)
	//fmt.Println(error)
	v := "1234"
	fmt.Println(v[0] == '1')
	ss := test {
		a:"1",
		b:"2",
	}
	fmt.Println(ss.a)
}

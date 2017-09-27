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

	a := "\"1\\\"2\""
	b := []byte(a)
	fmt.Println(len(b))
	for i := 0; i < len(b) ;i++  {
		fmt.Printf("%c", b[i])
	}
	//a := "12342"
	//b := a[2:]
	//fmt.Println(a[2:])
	//for i,_ := range b{
	//	fmt.Println(i + 2)
	//}
	//var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
	//
	//for {
	//	if piece, ok := getNextString(); ok {
	//		count++
	//		/*
	//		   func (b *Buffer) WriteString(s string) (n int, err error)
	//		   Write将s的内容写入缓冲中，如必要会增加缓冲容量。返回值n为len(p)，err总是nil。如果缓冲变得太大，Write会采用错误值ErrTooLarge引发panic。
	//		*/
	//		buffer.WriteByte(piece)
	//	} else {
	//		break
	//	}
	//}
	//
	//fmt.Println("拼接后的结果为-->", buffer.String())
}
var count int = 0
var teststr[] byte = []byte{'1','2','3'}
func getNextString()(byte, bool)  {
	if count > 2{
		return ',',false
	} else {
		return teststr[count], true
	}

}

package main

import (
	"fmt"
	//"parser"
	//"errors"
	"bytes"
)

type test struct {
	a string
	b string
}

func main()  {
	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲

	for {
		if piece, ok := getNextString(); ok {
			count++
			/*
			   func (b *Buffer) WriteString(s string) (n int, err error)
			   Write将s的内容写入缓冲中，如必要会增加缓冲容量。返回值n为len(p)，err总是nil。如果缓冲变得太大，Write会采用错误值ErrTooLarge引发panic。
			*/
			buffer.WriteByte(piece)
		} else {
			break
		}
	}

	fmt.Println("拼接后的结果为-->", buffer.String())
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

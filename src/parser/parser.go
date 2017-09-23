package parser

import (
	"errors"
	"bytes"
)

/**
	null -> null
	boolean -> true or false
	number -> float or number
	string -> [...]
	object -> {...}

	Production type
	object -> {
	@Author : royruan
 */

var (
	UnKnownValueTypeError		= errors.New("unknown value type")
)

type root struct {
	name string
	child[] node
}

type node struct {
	leaf bool
	name string
	root bool
	dataType string
	child[] node
	key string
}

var pos = 0
func Parse(input string)  {
	root := root {
		name:"root",
	}
	var buffer bytes.Buffer
	var key string
	for input[pos] != '}' {
		pos++
		if input[pos] == '"' {
			buffer.WriteByte(input[pos])
			for input[pos] != '"' {
			}
		}
	}
}

func object(input string) (int){
	for input[pos] != '}' {
		if input[pos] == '{' {
			pos++

		}
	}
	return 0
}

func typeNull()  {

}


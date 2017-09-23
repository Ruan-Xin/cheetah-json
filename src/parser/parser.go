package parser

import (
	"errors"
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
	child node
}

type node struct {
	leaf bool
	name string
	root bool
	dataType string
	child[] node
}

var pos = 0
func Parse(input string)  {
	root := node{
		leaf: false,
		name:"root",
		dataType:"object",
		kv:nil,
	}



	for input[pos] != '}' {
		if input[pos] == '{' {
			pos++
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


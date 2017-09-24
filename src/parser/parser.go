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
	OutOfBoundsError			= errors.New("out of bounds errors")
	WrongSymbol					= errors.New("wrong symbol which can't resolve")
)

const (
	NotExist
	String
	Number
)

func ()  {
	
}

type node struct {
	leaf bool
	name string
	root bool
	dataType string
	key string
	value_obj node
	value_arr[] string
	value string
}

var pos = 0
var length int
func Parse(input string) error {
	length = len(input)

	root := node {
		name:"root",
		leaf:false,
		root:true,
	}
	var buffer bytes.Buffer
	var key string
	for input[pos] != '}' {
		pos++
		if input[pos] == '"' {
			buffer.WriteByte(input[pos])
			for input[pos] != '"' {
				buffer.WriteByte(input[pos])
				pos++
			}
		}
		buffer.WriteByte(input[pos])
		key = buffer.String()
		pos++

		for input[pos] != ' ' {
			pos++
		}
		if input[pos] == ':' {
			pos++
		} else {
			return WrongSymbol
		}

		for input[pos] != ' ' {
			pos++
		}

		if input[pos] == 'n' {
			nilResolve(key)
		} else if input[pos] == '"' {
			stringResolve(key)
		} else if input[pos] == '{' {
			objectResolve(key)
		} else if input[pos] == '[' {
			arrayResolve(key)
		} else if input[pos] >= '0' &&
			input[pos] <= '9' && input[pos] == '-'{
			numResolve(key)
		} else if input[pos] == 't' && input[pos] == 'f' {
			boolResolve(key)
		} else {
			return UnKnownValueTypeError
		}
	}
	return nil
}

func nilResolve(key string)  {
	
}

func arrayResolve(key string)  {
	
}

func boolResolve(key string)  {
	
}

func numResolve(key string)  {
	
}
func stringResolve(key string)  {

}
func objectResolve(input string) (int){
	for input[pos] != '}' {
		if input[pos] == '{' {
			pos++

		}
	}
	return 0
}

func next() error {
	pos++
	if pos >= length {
		return OutOfBoundsError
	}
	return nil
}

func typeNull()  {

}


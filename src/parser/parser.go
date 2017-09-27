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
	MalformedStringError       = errors.New("Value is string, but can't find closing '\"' symbol")
	KeyPathNotFoundError       = errors.New("Key path not found")
)

type node struct {
	leaf bool
	name string
	root bool
	dataType ValueType
	key string
	value_obj node
	value_arr[] string
	value string
}

var pos = 0
var length int

func findKeyStart(data []byte, key string) (int ,error) {
	for i, _ := range data{
		switch data[i] {
		case '"':
			i++
			keyBegin := i

			strEnd, _ := stringEnd(data[i:])
			if strEnd == -1 {
				break
			}
			i += strEnd

			valueOffset := nextToken(data[i:])
			if valueOffset == -1 {
				break
			}

			i += valueOffset

			if data[i] == ':' {
				return keyBegin - 1,nil
			}
		}
		i++
	}
	return -1, KeyPathNotFoundError
}
func Parse(data[] byte) error {
	length = len(data)

	root := node {
		name:"root",
		leaf:false,
		root:true,
	}
	var buffer bytes.Buffer
	var key string
	if pos != '{' {
		return WrongSymbol
	}
	pos = nextToken(data[pos:])
	for data[pos] != '}' {
		pos = nextToken(data)
		if data[pos] == '"' {
			buffer.WriteByte(data[pos])
			for data[pos] != '"' {
				buffer.WriteByte(data[pos])
				pos = nextToken(data[pos:])
			}
		}
		buffer.WriteByte(data[pos])
		key = buffer.String()

		pos = nextToken(data[pos:])
		if data[pos] == ':' {
			pos++
		} else {
			return WrongSymbol
		}

		for data[pos] != ' ' {
			pos++
		}

		if data[pos] == 'n' {
			nilResolve(key)
		} else if data[pos] == '"' {
			stringResolve(key)
		} else if data[pos] == '{' {
			objectResolve(key)
		} else if data[pos] == '[' {
			arrayResolve(key)
		} else if data[pos] >= '0' &&
			data[pos] <= '9' && data[pos] == '-'{
			numResolve(key)
		} else if data[pos] == 't' && data[pos] == 'f' {
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
func stringResolve(key string, parent node)  {

}

func keyResolve(data string)  {

}
func objectResolve(data string) (int){
	for data[pos] != '}' {
		if data[pos] == '{' {
			pos++

		}
	}
	return 0
}

func nextToken(data[] byte) int {
	for i,c := range data{
		switch c {
		case ' ', '\n', '\r', '\t':
			continue
		default:
			return i
		}
	}

	return -1
}

func stringEnd(data[] byte) (int, bool) {
	escaped := false
	for i, c := range data{
		if c == '"' {
			if !escaped {
				return i + 1, false
			} else {
				j := i - 1
				for  {
					//string like this \\\\"
					if j < 0 || data[j] != '\\' {
						return i + 1, true
					}
					j--
					if j < 0 || data[j] != '\\' {
						break
					}
					j--
				}
			}
		} else if c == '\\' {
			escaped = true
		}
	}
	return -1, escaped
}

func blockEnd(data []byte, openSymbol byte, closeSymbol byte) int {
	match := 0

	for i, c := range data{
		switch c {
		case '"':
			se, _ := stringEnd(data[i + 1:])
			if se == -1 {
				return -1
			}
			i += se
		case openSymbol:
			match++
		case closeSymbol:
			match--

			if match == 0 {
				return i + 1
			}
		}
		i++
	}

	return -1
}

func tokenEnd(data[] byte) int {
	for i, c := range data {
		switch c {
		case ' ', '\n', '\r', '\t', ',', '}',']':
			return i
		}
	}
	return len(data)
}

func getType(data[] byte, offset int) ([]byte, ValueType, int, error) {
	var dataType ValueType
	endOffset := offset

	if data[offset] == '"' {
		dataType = String
		if idx,_ := stringEnd(data[offset + 1:]); idx != -1 {
			endOffset += idx + 1
		} else {
			return nil, dataType, offset, MalformedStringError
		}
	} else if data[offset] == '[' {
		dataType = Array
		endOffset = blockEnd(data[offset:], '[', ']')

		if endOffset == -1 {
			return nil, dataType, offset, MalformedStringError
		}

		endOffset += offset
	} else if data[offset] == '{' {
		dataType = Object

		endOffset = blockEnd(data[offset:], '{', '}')

		if endOffset == -1 {
			return nil, dataType, offset, MalformedStringError
		}

		endOffset += offset
	} else {
		end := tokenEnd(data[endOffset:])

		if end == -1 {
			return nil, dataType, offset, MalformedStringError
		}

		value := data[offset : endOffset + end]

		switch data[offset] {
		case 't', 'f'://true or false
			if bytes.Equal(value, trueLiteral) || bytes.Equal(value, falseLiteral){
				dataType = Boolean
			} else {
				return nil, Unknown, offset, UnKnownValueTypeError
			}
		case 'u', 'n':// undefine or null
			if bytes.Equal(value, nullLiteral) {
				dataType = Null
			} else {
				return nil, Unknown, offset, UnKnownValueTypeError
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-':
			dataType = Number
		default:
			return nil, Unknown, offset, UnKnownValueTypeError
		}
		endOffset += end
	}
	return data[offset : endOffset], dataType, endOffset, nil
}

type ValueType int 

const (
	NotExist = ValueType(iota)
	String
	Number
	Object
	Array
	Boolean
	Null
	Unknown
)

var (
	trueLiteral = []byte("true")
	falseLiteral = []byte("false")
	nullLiteral = []byte("null")
)

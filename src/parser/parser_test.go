package parser

import (
	"testing"
	"fmt"
)

func TestParse(t *testing.T) {
	root := node{
		leaf: false,
		name:"root",
		dataType:"object",
		kv:nil,
	}

	fmt.Println(root.child)
}

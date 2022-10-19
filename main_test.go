package main

import (
	"fmt"
	"testing"
)

func TestTemp(t *testing.T) {
	a := 1
	switch a {
	case 1:
		fmt.Println(1)
		fmt.Println(3)
	case 2:
		fmt.Println(2)
	}
}

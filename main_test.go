package main

import (
	"fmt"
	"testing"
)

type A struct {
	AbCdEf int
}

func TestTemp(t *testing.T) {
	a := []A{
		{
			AbCdEf: 1,
		},
		{
			AbCdEf: 2,
		},
	}
	T(a)
	fmt.Print(a)
}

func T(a []A) {
	a[0].AbCdEf = 9
	a[1].AbCdEf = 10
}

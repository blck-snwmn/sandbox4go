package main

import (
	"fmt"

	"github.com/blck-snwmn/sandbox4go/20200911_packages/testdata/foo"
)

type T string

func (t T) Do(xx string) {
	fmt.Println("dodododo", xx)
}
func do() {
	do2()
	T("a").Do("r")
	foo.DoFoo()
}

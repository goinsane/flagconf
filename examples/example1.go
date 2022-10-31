//go:build ignore
// +build ignore

package main

import (
	"flag"
	"fmt"

	"github.com/goinsane/flagconf"
)

func main() {
	var boolVar bool
	var intVar int
	var strVar string
	s := `
bool-var true
int-var 50
str-var efgh
`
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	fs.BoolVar(&boolVar, "bool-var", false, "bool variable")
	fs.IntVar(&intVar, "int-var", 10, "int variable")
	fs.StringVar(&strVar, "str-var", "abcd", "string variable")
	if e := flagconf.ParseString(fs, s); e != nil {
		panic(e)
	}

	fmt.Println("bool-var", boolVar)
	fmt.Println("int-var", intVar)
	fmt.Println("string-var", strVar)
}

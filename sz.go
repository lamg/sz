package main

import (
	"fmt"
	"os"
)

var usage = "Usage: %s file\n"

func main(){
	if len(os.Args) == 2 {
		fi, e := os.Lstat(os.Args[1])
		if e == nil {
			println(fi.Size())
		} else {
			fmt.Fprintf(os.Stderr, "%s\n", e.Error())
		}
	} else {
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
	}
}

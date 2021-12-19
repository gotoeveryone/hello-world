package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	fmt.Println("Please enter the singular word:")

	var w string
	fmt.Scan(&w)

	m, e := regexp.MatchString("[s|sh|ch|o|x]$", w)
	if e != nil {
		os.Exit(1)
	}
	if m {
		fmt.Println(w + "es")
		os.Exit(0)
	}

	r, e := regexp.Compile("(.*)(fe?)$")
	if e != nil {
		os.Exit(1)
	}
	f := r.FindStringSubmatch(w)
	if f != nil {
		fmt.Println(f[1] + "ves")
		os.Exit(0)
	}

	r2, e := regexp.Compile("(.*[^aiueo])y$")
	if e != nil {
		os.Exit(1)
	}
	f2 := r2.FindStringSubmatch(w)
	if f2 != nil {
		fmt.Println(f2[1] + "ies")
		os.Exit(0)
	}

	fmt.Println(w + "s")
}

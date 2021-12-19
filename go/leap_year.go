package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func isLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	}
	if year%100 == 0 {
		return year%400 == 0
	}
	return true
}

func main() {
	fmt.Println("Please enter the year")

	var w string
	fmt.Scan(&w)

	r, e := regexp.Compile("(\\d+)")
	if e != nil {
		os.Exit(1)
	}
	f := r.FindStringSubmatch(w)
	if len(f) < 1 {
		fmt.Println(fmt.Sprintf("'%s' is not integer", w))
		os.Exit(1)
	}

	v, e := strconv.Atoi(w)
	if e != nil {
		os.Exit(1)
	}

	if isLeapYear(v) {
		fmt.Println(fmt.Sprintf("'%d' is leap year", v))
	} else {
		fmt.Println(fmt.Sprintf("'%d' is not leap year", v))
	}
}

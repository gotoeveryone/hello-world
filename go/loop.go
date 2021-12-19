package main

import "fmt"

func while() {
	count := 0
	for count < 10 {
		fmt.Println(count)
		count++
	}
}

func foreach() {
	arr := []int{1, 2, 3}
	for _, v := range arr {
		fmt.Println(v)
	}
}

func main() {
	while()
	foreach()
}

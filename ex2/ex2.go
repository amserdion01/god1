package main

import "fmt"

type intern struct {
	name   string
	gender string
	age    int
}

func main() {
	var n string
	var g string
	var a int
	fmt.Scanln(&n)
	fmt.Scanln(&g)
	fmt.Scanln(&a)
	var newIntern = intern{
		name:   n,
		gender: g,
		age:    a,
	}
	fmt.Println(newIntern)
}

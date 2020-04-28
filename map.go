package main

import "fmt"

func printMap(c map[int]string) {
	for hash, value := range c {
		fmt.Println(hash, ":", value)
	}
}

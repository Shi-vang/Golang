package main

import "fmt"

type bot interface {
	getGreeting() string
}

type english struct{}
type spanish struct{}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (en english) getGreeting() string {
	return "Hi English"
}

func (sp spanish) getGreeting() string {
	return "Hola"
}

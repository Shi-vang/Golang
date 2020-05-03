package main

import (
	"fmt"
	"net/http"
)

func checkLinks(link string) {
	_, err := http.Get(link)
	if err == nil {
		fmt.Println(link, " is up")
		return
	}
	fmt.Println(link, " is Down")
}

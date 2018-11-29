package main

import (
	"os"
	"strings"

	"services/translate"
)

func main() {
	var word string

	argsCount := len(os.Args)-1
	argsData  := os.Args[1:]

	if argsCount < 1 {
		return
	}

	ch := make(chan bool, 1)
	translate.SetpIqaD(ch)

	for i, v := range argsData {
		word = word + " " + v
	}
	word = strings.Trim(word, " ")

	<- ch
    close(ch)

	hex := translate.Klingon(word)
}

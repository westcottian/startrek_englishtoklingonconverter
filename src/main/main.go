package main

import (
    "os"
    "strings"
    "fmt"

    "services/character"
    "services/translate"
)

func main() {
    var hex, specie, word string

    argsCount := len(os.Args) - 1
    argsData := os.Args[1: ]

    if argsCount < 1 {
        fmt.Println("Invalid: Nothing to transalate.")
        return
    }

    for _, v := range argsData {
        word = word + " " + v
    }
    word = strings.Trim(word, " ")

    translate.SetpIqaD()

    ch := make(chan bool, 2)
    go translate.Klingon(ch, word, &hex)
    go character.GetSpecie(ch, word, &specie) 
	<-ch
	<-ch
    close(ch)
    fmt.Println(hex)
    fmt.Println(specie)
}
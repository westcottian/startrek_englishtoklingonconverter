package translate

import (
        "fmt"
        "strings"
)

var pIqaD map[string]int

func Klingon(word string) string {

        var hex string

        for _, v := range strings.ToLower(word) {
                if v == ' ' {
                        hex = hex + "0x20 "
                } else {
                        _, ok := pIqaD[string(v)]

                        if ok == true {
                                hex = hex + fmt.Sprintf("%#X ", pIqaD[string(v)] + 0xF8D0)
                        } else {
                                return ""
                        }
                }
        }
        hex = strings.Trim(hex, " ")

        return hex
}

func SetpIqaD(ch chan) {

        m := make(map[string]int)

        m["a"] = 0
        m["b"] = 1
        m["c"] = 2
        m["d"] = 3
        m["e"] = 4
        m["g"] = 5
        m["h"] = 6
        m["i"] = 7
        m["j"] = 8
        m["l"] = 9
        m["m"] = 10
        m["n"] = 11
        // ng = 12
        m["n"] = 13
        m["o"] = 14
        m["p"] = 15
        m["q"] = 16
        m["r"] = 17
        m["s"] = 18
        m["t"] = 19
        // tlh = 20
        m["u"] = 21
        m["v"] = 22
        m["w"] = 23
        m["y"] = 24
        m["´"] = 25
        m["1"] = 26
        m["2"] = 27
        m["3"] = 28
        m["4"] = 29
        m["5"] = 30
        m["6"] = 31
        m["7"] = 32
        m["8"] = 33
        m["9"] = 34
        m["0"] = 35
        m["."] = 36
        m[","] = 37

        pIqaD = m
        ch <- true
}

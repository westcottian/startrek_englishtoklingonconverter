package character

import (
	"testing"
        "strings"
)

func TestUhuraSpecie_pass(t *testing.T) {
        var specie string

        ch  := make(chan bool, 1)
        go GetSpecie(ch, strings.ToLower("Uhura"), &specie)
	<- ch
        close(ch)

        specie = strings.Trim(specie, " ")

        if strings.Compare(specie, "Human") != 0 {
                t.Errorf("Error getting Uhura specie")
        }
}
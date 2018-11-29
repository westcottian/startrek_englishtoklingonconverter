package character

import (
    "testing"
    "strings"
)

func TestUhuraSpecie_pass(t * testing.T) {
    specie: = strings.ToLower(GetSpecie("Uhura"))

    if strings.Compare(specie, "human") != 0 {
        t.Errorf("Error getting Uhura specie")
    }
}
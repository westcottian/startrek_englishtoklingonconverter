package translate

import (
	"testing"
        "strings"
)

func TestSetpIqaD_pass(t *testing.T) {
        ch := make(chan bool, 1)
	SetpIqaD(ch)
	<- ch
        close(ch)

        if len(pIqaD) != 35 {
                t.Errorf("Error setting pIqaD")
        }
}

func TestKlingon_pass(t *testing.T) {
        hex := Klingon("Uhura")

        if strings.Compare(hex, "0XF8E5 0XF8D6 0XF8E5 0XF8E1 0XF8D0") != 0 {
                t.Errorf("Error transalting Uhura")
        }
}

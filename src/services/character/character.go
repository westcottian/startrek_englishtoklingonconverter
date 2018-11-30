package character

import (
    "fmt"
    "strings"

    "common/defines"
    "common/urls"
)

func GetSpecie(ch chan bool, name string, specie* string) {
    var aux string

    data := `name=` + name
    form := strings.NewReader(data)
    url := fmt.Sprintf("%s/search", defines.StapiRESTCharacter)

    headers := make(map[string] string)
    headers["Content-Type"] = "application/x-www-form-urlencoded"

    ret, err := urls.SendRequest("POST", url, form, headers, defines.Timeout)
    if err != nil {
        ch <-false
        return
    }

    if ret["characters"] == nil {
        ch <-false
        return
    }

    for _, v := range ret["characters"].([] interface {}) {
        for key, value := range v.(map[string] interface {}) {
            if key == "uid" {
                aux = getFullCharacter(value.(string))
                if aux != "" { * specie = aux
                    ch <-true
                    return
                }
            }
        }
    }

    ch <-false
}

func getFullCharacter(uid string) string {

    url := fmt.Sprintf("%s?uid=%s", defines.StapiRESTCharacter, uid)
    headers := make(map[string] string)
    headers["Content-Type"] = "application/x-www-form-urlencoded"
    headers["Access-Control-Allow-Origin"] = "*"

    ret,err := urls.SendRequest("GET", url, nil, nil, defines.Timeout)
    if err != nil {
        return ""
    }

    if ret["character"] == nil {
        return ""
    }

    for k,
    v := range ret["character"].(map[string] interface {}) {
        if k == "characterSpecies" {
            for _, array := range v.([] interface {}) {
                for key, value := range array.(map[string] interface {}) {
                    if key == "name" {
                        return value.(string)
                    }
                }
            }
        }
    }

        return ""
}
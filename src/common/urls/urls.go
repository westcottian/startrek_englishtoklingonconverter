package urls

import (
    "errors"
    "io"
    "time"
	"fmt"
    "encoding/json"
    "net/http"
)

func SendRequest(method, uri string, data io.Reader, headers map[string] string, timeout int)(map[string] interface {}, error) {

    var client * http.Client

    request, err := http.NewRequest(method, uri, data)
    if err != nil {
        return nil, err
    }

    for k, v := range headers {
        request.Header.Set(k, v)
    }

    client = &http.Client{Timeout: time.Duration(timeout) * time.Second}

    resp, err  := client.Do(request)
    if err != nil {
		fmt.Println("Response error:", err)
        return nil, err
    }
    defer resp.Body.Close()
	var d map[string] interface {}
	
    json.NewDecoder(resp.Body).Decode(&d)

    if resp.StatusCode < 200 || resp.StatusCode > 299 {
        return d, errors.New(resp.Status)
    }
    return d, nil

}
package urls

import (
    "fmt"
    "net/http"
    "net/http/httptest"
    "testing"
    "strconv"
)

func TestCreateDataFromService_pass(t * testing.T) {
    ts: = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Header().Set("Content-Type", "application/json")
        fmt.Fprintln(w, "{\"" + TestResponseParameter1 + "\":\"" + TestResponseValue1 + "\",\"" + TestResponseParameter2 + "\":\"" + TestResponseValue2 + "\",\"" + TestResponseParameter3 + "\":\"" + TestResponseValue3 + "\"}")
    }))
    defer ts.Close()

    ret,
    err: = CreateDataFromService(ts.URL, nil, 5)
    if err != nil {
        t.Errorf("Erro no envio/recebimento do JSON")
    }

    if ret[TestResponseParameter1] != TestResponseValue1 {
        t.Errorf("Problema no valor 1 retornado no JSON")
    }

    if ret[TestResponseParameter2] != TestResponseValue2 {
        t.Errorf("Problema no valor 2 retornado no JSON")
    }

    if ret[TestResponseParameter3] != TestResponseValue3 {
        t.Errorf("Problema no valor 3 retornado no JSON")
    }
}

func TestCreateDataFromService_failed(t * testing.T) {
    ts: = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
        w.WriteHeader(http.StatusUnauthorized)
        w.Header().Set("Content-Type", "application/json")
        fmt.Fprintln(w, "{\"" + TestResponseParameter1 + "\":\"" + TestResponseValue1 + "\",\"" + TestResponseParameter2 + "\":\"" + TestResponseValue2 + "\",\"" + TestResponseParameter3 + "\":\"" + TestResponseValue3 + "\"}")
    }))
    defer ts.Close()

    ret,
    err: = CreateDataFromService(ts.URL, nil, 5)
    i,
    _: = strconv.Atoi(err.Error()[: 3])
    if i != http.StatusUnauthorized {
        t.Errorf("Erro no envio/recebimento do JSON")
    }

    if ret[TestResponseParameter1] != TestResponseValue1 {
        t.Errorf("Problema no valor 1 retornado no JSON")
    }

    if ret[TestResponseParameter2] != TestResponseValue2 {
        t.Errorf("Problema no valor 2 retornado no JSON")
    }

    if ret[TestResponseParameter3] != TestResponseValue3 {
        t.Errorf("Problema no valor 3 retornado no JSON")
    }
}

func TestCreateDataFromServiceWithCustomHeader_pass(t * testing.T) {
    ts: = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Header().Set("Content-Type", "application/json")

        header_data: = fmt.Sprintf("%s", r.Header[TestResponseParameter1])
        fmt.Fprintln(w, "{\"" + TestResponseParameter1 + "\":\"" + header_data[1: len(header_data) - 1] + "\"}")
    }))
    defer ts.Close()

    headers: = make(map[string] string)
    headers[TestResponseParameter1] = TestResponseValue1

    ret,
    err: = CreateDataFromServiceWithCustomHeader(ts.URL, nil, headers, 5)
    if err != nil {
        t.Errorf("Erro no envio/recebimento do JSON")
    }

    if ret[TestResponseParameter1] != TestResponseValue1 {
        t.Errorf("Problema no valor 1 retornado no JSON")
    }
}

func TestCreateDataFromServiceWithCustomHeader_failed(t * testing.T) {
    ts: = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
        w.WriteHeader(http.StatusUnauthorized)
        w.Header().Set("Content-Type", "application/json")

        header_data: = fmt.Sprintf("%s", r.Header[TestResponseParameter1])
        fmt.Fprintln(w, "{\"" + TestResponseParameter1 + "\":\"" + header_data[1: len(header_data) - 1] + "\"}")
    }))
    defer ts.Close()

    headers: = make(map[string] string)
    headers[TestResponseParameter1] = TestResponseValue1

    ret,
    err: = CreateDataFromServiceWithCustomHeader(ts.URL, nil, headers, 5)
    i,
    _: = strconv.Atoi(err.Error()[: 3])
    if i != http.StatusUnauthorized {
        t.Errorf("Erro no envio/recebimento do JSON")
    }

    if ret[TestResponseParameter1] != TestResponseValue1 {
        t.Errorf("Problema no valor 1 retornado no JSON")
    }
}

func TestDeleteDataFromService_pass(t * testing.T) {
    ts: = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Header().Set("Content-Type", "application/json")
        fmt.Fprintln(w, "{\"" + TestStatusParam + "\":" + TestStatusValue + ",\"" + TestDataParam + "\":\"" + TestDataValue + "\"}")
    }))
    defer ts.Close()

    status,
    data,
    err: = DeleteDataFromService(ts.URL, 5)
    if err != nil {
        t.Errorf("Erro no processo de delecao")
    }

    i,
    _: = strconv.Atoi(TestStatusValue)
    if status != i {
        t.Errorf("Erro no retorno do status")
    }

    if data != TestDataValue {
        t.Errorf("Erro no retorno da data")
    }
}

func TestDeleteDataFromService_failed(t * testing.T) {
    ts: = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
        w.WriteHeader(http.StatusUnauthorized)
        w.Header().Set("Content-Type", "application/json")
        fmt.Fprintln(w, "{\"" + TestStatusParam + "\":" + TestStatusValue + ",\"" + TestDataParam + "\":\"" + TestDataValue + "\"}")
    }))
    defer ts.Close()

    status,
    data,
    err: = DeleteDataFromService(ts.URL, 5)
    i,
    _: = strconv.Atoi(err.Error()[: 3])
    if i != http.StatusUnauthorized {
        t.Errorf("Erro no envio/recebimento do JSON")
    }

    i,
    _ = strconv.Atoi(TestStatusValue)
    if status != i {
        t.Errorf("Erro no retorno do status")
    }

    if data != TestDataValue {
        t.Errorf("Erro no retorno da data")
    }
}

func TestGetDataFromService_pass(t * testing.T) {
    ts: = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Header().Set("Content-Type", "application/json")
        fmt.Fprintln(w, "{\"" + TestStatusParam + "\":" + TestStatusValue + ",\"" + TestDataParam + "\":\"" + TestDataValue + "\"}")
    }))
    defer ts.Close()

    status,
    data,
    err: = GetDataFromService(ts.URL, 5)
    if err != nil {
        t.Errorf("Erro no processo de delecao")
    }

    i,
    _: = strconv.Atoi(TestStatusValue)
    if status != i {
        t.Errorf("Erro no retorno do status")
    }

    if data != TestDataValue {
        t.Errorf("Erro no retorno da data")
    }
}

func TestGetDataFromService_failed(t * testing.T) {
    ts: = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
        w.WriteHeader(http.StatusUnauthorized)
        w.Header().Set("Content-Type", "application/json")
        fmt.Fprintln(w, "{\"" + TestStatusParam + "\":" + TestStatusValue + ",\"" + TestDataParam + "\":\"" + TestDataValue + "\"}")
    }))
    defer ts.Close()

    status,
    data,
    err: = GetDataFromService(ts.URL, 5)
    i,
    _: = strconv.Atoi(err.Error()[: 3])
    if i != http.StatusUnauthorized {
        t.Errorf("Erro no envio/recebimento do JSON")
    }

    i,
    _ = strconv.Atoi(TestStatusValue)
    if status != i {
        t.Errorf("Erro no retorno do status")
    }

    if data != TestDataValue {
        t.Errorf("Erro no retorno da data")
    }
}

func TestGetDataFromServiceWithCustomHeader_pass(t * testing.T) {
    ts: = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Header().Set("Content-Type", "application/json")

        header_data: = fmt.Sprintf("%s", r.Header[TestResponseParameter1])
        fmt.Fprintln(w, "{\"" + TestResponseParameter1 + "\":\"" + header_data[1: len(header_data) - 1] + "\"}")
    }))
    defer ts.Close()

    headers: = make(map[string] string)
    headers[TestResponseParameter1] = TestResponseValue1

    ret,
    err: = GetDataFromServiceWithCustomHeader(ts.URL, headers, 5)
    if err != nil {
        t.Errorf("Erro ao pegar data")
    }

    if ret[TestResponseParameter1] != TestResponseValue1 {
        t.Errorf("Problema no valor retornado no JSON")
    }
}

func TestGetDataFromServiceWithCustomHeader_failed(t * testing.T) {
    ts: = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
        w.WriteHeader(http.StatusUnauthorized)
        w.Header().Set("Content-Type", "application/json")

        header_data: = fmt.Sprintf("%s", r.Header[TestResponseParameter1])
        fmt.Fprintln(w, "{\"" + TestResponseParameter1 + "\":\"" + header_data[1: len(header_data) - 1] + "\"}")
    }))
    defer ts.Close()

    headers: = make(map[string] string)
    headers[TestResponseParameter1] = TestResponseValue1

    ret,
    err: = GetDataFromServiceWithCustomHeader(ts.URL, headers, 5)
    if err == nil {
        t.Errorf("Erro ao pegar data")
    }

    i,
    _: = strconv.Atoi(err.Error()[: 3])
    if i != http.StatusUnauthorized {
        t.Errorf("Erro no envio/recebimento do JSON")
    }

    if ret[TestResponseParameter1] != TestResponseValue1 {
        t.Errorf("Problema no valor retornado no JSON")
    }
}
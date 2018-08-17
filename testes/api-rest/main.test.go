package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	data := url.Values{}
	data.Set("valor_pi", "foo")
	req, err := http.NewRequest("POST", "http://localhost:8080/status", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	cliente := &http.Client{}
	resp, err := cliente.Do(req)
    if err != nil {
        panic(err)
    }
	defer resp.Body.Close()
	
	fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}

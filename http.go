package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"crypto/tls"
)

func main() {
	// GET HTTP/HTTPS
	tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}
	resp, err := client.Get("https://git.hkparker.com/")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	body := string(bytes)
	fmt.Println(body)
	
	// GET/POST with basic authentication
	var username string = "hayden"
    var passwd string = "password"
    req, err := http.NewRequest("GET", "http://hkparker.com", nil)
    req.SetBasicAuth(username, passwd)
    resp2, err := client.Do(req)
    if err != nil{
        fmt.Println(err)
    }
    bytes2, err := ioutil.ReadAll(resp2.Body)
    body2 := string(bytes2)
    fmt.Println(body2)
    
    //POST data
}

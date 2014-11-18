package main

//
//	This script will read in json data and
//		Send a POST or a GET to any URL		[Done]
//		Include post data optionally		?
//		Include basic auth optionally		[Done]
//		Add any custom headers				?
//
//	Based on the results, it can
//		Parse the body with any regex		?
//		Determine the response codes		[Done]

import (
	"fmt"
	"net/http"
	//"io/ioutil"
	"crypto/tls"
)

func main() {
	tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}
    
    http_verb := "GET"
    url := "http://hkparker.com"
    // need to define post data and basic auth
    
    request, err := http.NewRequest(http_verb, url, nil)
	//var username string = "hayden"
    //var passwd string = "password"
    //req.SetBasicAuth(username, passwd)
    
    response, err := client.Do(request)
    if err != nil{
        fmt.Println(err)
    }
    //bytes, err := ioutil.ReadAll(response.Body)
    //body := string(bytes)
    fmt.Println(response.StatusCode)
}

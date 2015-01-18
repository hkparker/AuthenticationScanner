package main

import (
		"fmt"
		"log"
		"flag"
		"io/ioutil"
		"net/http"
		"crypto/tls"
		"encoding/json"
		"strings"
		"strconv"
		"regexp"
	)
	
type AuthSource struct {
	Name string
	URL string
	HTTPVerb string
	Headers map[string]string
	BasicAuth int
	Body string
	ParseForm int
	DetectionMethod string
	Success string
	Failed string
}

func parse_response(site AuthSource, response *http.Response) {
	if site.DetectionMethod == "ResponseCode" {
		code := strconv.Itoa(response.StatusCode)
		if code == site.Success {
			fmt.Println("")
		} else if code == site.Failed {
			fmt.Println("")
		}
	} else if site.DetectionMethod == "Regex" {
		bytes, _ := ioutil.ReadAll(response.Body)
		match, _ := regexp.Match(site.Success, bytes)
		if match {
			fmt.Println("")
		}
		match, _ = regexp.Match(site.Failed, bytes)
		if match {
			fmt.Println("")
		}
	}
}

func test_site(site AuthSource, done chan int, client *http.Client, username string, password string) {
	fmt.Println("Checking site:", site.Name)
	request := site.build_request(username, password)
	response, err := client.Do(&request)
	if err != nil{
		fmt.Println(err)
	}
	parse_response(site, response)
	done <- 0
}

func (source AuthSource) build_request(username string, password string) http.Request {
	request, _ := http.NewRequest(source.HTTPVerb, source.URL, nil)
	for k, v := range source.Headers {
		key := strings.Replace(k, "#{USERNAME}", username, -1)
		key = strings.Replace(key, "#{PASSWORD}", password, -1)
		value := strings.Replace(v, "#{USERNAME}", username, -1)
		value = strings.Replace(value, "#{PASSWORD}", password, -1)
		request.Header.Set(key,value)
	}
	if source.BasicAuth == 1 {
		request.SetBasicAuth(username, password)
	}
	body := strings.Replace(source.Body, "#{USERNAME}", username, -1)
	body = strings.Replace(body, "#{PASSWORD}", password, -1)
	request.Body = ioutil.NopCloser(strings.NewReader(body))
	if source.ParseForm == 1 {
		request.ParseForm()
	}
	return *request
}

func authentication_source_from_file(filename string) (AuthSource, error) {
	var auth_source AuthSource
	file_contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return auth_source, err
	}
	json_err := json.Unmarshal(file_contents, &auth_source)
	return auth_source, json_err
}

func build_auth_source_slice(directory string) []AuthSource {
	var auth_sources []AuthSource
	entries, dir_err := ioutil.ReadDir(directory)
	if dir_err != nil {
		log.Fatal(dir_err)
	}
	for _, filename := range entries {
		auth_source, err := authentication_source_from_file(directory+"/"+filename.Name())
		if err != nil {
			fmt.Println("Error loading", filename.Name() ,":", err)
		} else {
			auth_sources = append(auth_sources, auth_source)
		}
	}
	return auth_sources
}

func main() {
	fmt.Println("\n\tStarting AuthenticationScanner 1.0...\n")
	var uname = flag.String("username", "", "Username to authenticate with")
	var pword = flag.String("password", "", "Password to authenticate with")
	flag.Parse()
	username := *uname
	password := *pword
	
	// Get sites to test
	auth_source_array := build_auth_source_slice("sites")
	site_count := len(auth_source_array)
	if site_count == 0 {
		log.Fatal("No usable sites loaded, exiting.")
	}
	fmt.Println("Sites loaded:", site_count)
	
	// Build HTTP Client
	tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}
    
    // Test each site
    done := make(chan int)
	for _, site := range auth_source_array {
		go test_site(site, done, client, username, password)
	}
	for i:= 0; i < site_count; i++ {
		<- done
	}
}

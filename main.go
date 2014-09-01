package main

import (
//		"net"
		"fmt"
		"io/ioutil"
		"encoding/json"
//		"os"
//		"bufio"
		"log"
//		"flag"
	)
	
type AuthSource struct {
	Name string					// String name, used for cli
	URL string					// URL to query
	HTTPVerb string				// GET or POST
	Headers string				// Any other headers
	DetectionMethod string		// Valid options: ResponseCode, Regex (have to go beneath the javascript to the api js is querying)
	LoginSuccess string			// respose code to match if response code, regex to match if regex
	LoginFailed string			// same as success but for failure.  Need both to confirm.
	// Maybe have a field for login type (username or email), then apply the sites to the crentials dynamically
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
			fmt.Println("Error loading ", filename.Name() ,": ", err)
		} else {
			auth_sources = append(auth_sources, auth_source)
		}
	}
	return auth_sources
}

func main() {
	fmt.Println("\n\tStarting AuthenticationScanner 1.0...\n")
	auth_source_array := build_auth_source_slice("sites")
	site_count := len(auth_source_array)
	if site_count == 0 {
		log.Fatal("No usable sites loaded, exiting.")
	}
	fmt.Println("Sites loaded:", site_count)
	// parse arguments for proxies and such here
	// single credentials to test?  a whole file of crenetials?
	for _, site := range auth_source_array {
		fmt.Println(site.Name)
		// each site will "go attempt_to_login(site, results_channel)".  The attemp_to_login method with handle proxies.
	}
	// a recieving function will loop site_count times and read login results from the results_channel
	// the recieving method returns a list of results, with are parsed and printed out.
}

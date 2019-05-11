package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func main() {
	fmt.Println("Starting the application...")
	// api request to sumo
	//client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.us2.sumologic.com/api/v1/users", nil)
	req.Header.Add("Authorization", "Basic "+basicAuth("", ""))
	if err != nil {
		log.Fatal(err)
	} else {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatalln("error")
		} else {
			fmt.Println(string(body))
		}
	}

	fmt.Println("Terminating the application...")
}

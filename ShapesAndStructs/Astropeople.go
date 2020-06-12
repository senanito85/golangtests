package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type people struct {
	Status string `json:"message"`
	Count int `json:"number"`
	Group []Teammember `json:"people"`
}

type Teammember  struct {
        Craft string `json:"craft"`
        Name  string `json:"name"`
}

func main() {

	url := "http://api.open-notify.org/astros.json"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-people")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	people1 := people{}
	jsonErr := json.Unmarshal(body, &people1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
     
    fmt.Println("----------------------------")
	fmt.Println("Status of request:", people1.Status)
	fmt.Println("Number of people in Space:", people1.Count)
	fmt.Println("----------------------------")

	for _, p := range people1.Group {
        fmt.Println(p.Name)
    }

}





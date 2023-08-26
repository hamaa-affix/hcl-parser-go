package main

import (
	// "context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Res struct {
	Size int    `json:"fileSizeBytes"`
	Url  string `json:"url"`
}

func main() {
	Run()
}

func Run() {
	res, err := http.Get("https://random.dog/woof.json")
	if err != nil {
		fmt.Printf("fatal get request: %v\n", err.Error())
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("fatal read body: %s\n", err.Error())
	}

	var resData Res
	if err := json.Unmarshal(data, &resData); err != nil {
		log.Fatalf("error can not parse: %s\n", err.Error())
	}

	fmt.Println(resData)
}

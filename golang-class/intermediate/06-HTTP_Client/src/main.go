package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type CatImage struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

func main() {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", "https://distribution-uat.dev.muangthai.co.th/mtl-node-red/golang-course/cat-api/list", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/json")
	q := req.URL.Query()
	q.Add("limit", "10")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("HTTP error: %s", resp.Status)
	}

	var result []CatImage
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %+v\n", result)
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"

	for {
		// generate random values for water and wind
		water := rand.Intn(100)
		wind := rand.Intn(100)

		// determine the status of water and wind
		waterStatus := ""
		windStatus := ""

		if water < 5 {
			waterStatus = "aman"
		} else if water < 9 {
			waterStatus = "siaga"
		} else {
			waterStatus = "bahaya"
		}

		if wind < 6 {
			windStatus = "aman"
		} else if wind < 16 {
			windStatus = "siaga"
		} else {
			windStatus = "bahaya"
		}

		// create the JSON payload
		data := map[string]int{
			"water": water,
			"wind":  wind,
		}

		// marshal the JSON payload
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
			return
		}

		// make the POST request
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// delete id
		var result map[string]interface{}
		json.Unmarshal([]byte(body), &result)
		delete(result, "id")
		jsonData, _ = json.Marshal(result)

		// print the JSON response
		fmt.Println(string(jsonData))

		// Print status
		fmt.Printf("status water: %s\n", waterStatus)
		fmt.Printf("status wind: %s\n", windStatus)

		// close the response body
		resp.Body.Close()

		time.Sleep(15 * time.Second)
	}
}

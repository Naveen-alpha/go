package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}

	url := "https://bankcbw.zendesk.com/api/v2/search.json?query=type:user%20external_id:004"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", "Basic Yi5uYXZlZW5rdW1hckBiYW5rY2J3Lm9yZzpDYnduYXZlZW4x")

	res, _ := client.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println("RES", string(body))

}

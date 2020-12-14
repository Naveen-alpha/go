package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}

	url := "https://bankcbw.zendesk.com/api/v2/users/create_or_update.json"

	payload := strings.NewReader("{\n    \"user\": {\n        \"external_id\": \"account_5327\",\n \"name\": \"ajith\",\n        \"email\": \"ajith+2@example.org\",\n        \"role\": \"end-user\",\n        \"phone\": \"1234567890\"\n    }\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.SetBasicAuth("b.naveenkumar@bankcbw.org", "Cbwnaveen1")

	res, _ := client.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

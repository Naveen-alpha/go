package main

import (
	"encoding/base64"
	"log"

	config "../zendesk/config"
)

func main() {
	authstring := base64.StdEncoding.EncodeToString([]byte(config.UserName + ":" + config.Password))
	log.Println(authstring)
}

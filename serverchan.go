package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	key  string
	text string
	desp string
)

func init() {
	key = os.Getenv("PLUGIN_KEY")
	text = os.Getenv("PLUGIN_TEXT")
	desp = os.Getenv("PLUGIN_DESP")
}

func main() {
	if key == "" || text == "" {
		log.Fatalln("key or text is required")
	}

	reqMsg := &url.Values{
		"text": []string{text},
		"desp": []string{desp},
	}

	reqURL := fmt.Sprintf("https://sc.ftqq.com/%s.send", key)
	_, err := http.Post(reqURL, "application/x-www-form-urlencoded", strings.NewReader(reqMsg.Encode()))
	if err != nil {
		log.Fatalln("post error: ", err)
	}
}

package main

import (
	"github.com/labstack/gommon/log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://tkstorm.cc/img/blog-red-logo.jpg")
	if err !=nil {
		panic(err)
	}
	log.Print(resp.Request)
	log.Print(resp.Header)
	log.Print(resp.TLS)
	log.Print(resp.ContentLength, resp.StatusCode, resp.Status)
}

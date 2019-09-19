package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

var (
	headImg   []byte
	indexCss  []byte
	indexHtml []byte
)

func init() {
	var err error
	if headImg, err = ioutil.ReadFile("./img/blog-red-logo.jpg"); err != nil {
		panic(err)
	}
	if indexCss, err = ioutil.ReadFile("./css/style.css"); err != nil {
		panic(err)
	}
	if indexHtml, err = ioutil.ReadFile("./html/index.html"); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", htmlHandler)
	http.HandleFunc("/img/", imgHandler)
	http.HandleFunc("/css/", cssHandler)
	log.Fatal(http.ListenAndServeTLS(":2345", "gohttp2.cert", "gohttp2.key", nil))
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	w.Write(indexCss)
}

func imgHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(headImg)
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	h2push(w, "/css/style.css", "/img/blog-red-logo.jpg")
	w.Header().Set("Content-Type", "text/html")
	w.Write(indexHtml)
}

// set http2 push
func h2push(w http.ResponseWriter, s ...string) {
	if p, ok := w.(http.Pusher); ok {
		for _, t := range s {
			p.Push(t, nil)
		}
	}
}

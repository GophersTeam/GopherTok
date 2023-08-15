package main

import (
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("/home/project/gophertok/temp/"))
	http.Handle("/gophertok/", http.StripPrefix("/gophertok/", fs))

	http.ListenAndServe(":5200", nil)
}

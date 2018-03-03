package main

import (
	"./http"
)

func main() {
	http.Request("https://httpbin.org/headers")
}

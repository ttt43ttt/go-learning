package main

import "net/http"

func main() {
	h := http.FileServer(http.Dir("c:/code/html"))
	println("listening on port:", 8888)
	http.ListenAndServe(":8888", h)
}

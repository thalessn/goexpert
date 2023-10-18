package main

import "net/http"

func main() {
	fileServer := http.FileServer(http.Dir("./public"))
	http.ListenAndServe(":8080", fileServer)
}

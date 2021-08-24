package main

import "net/http"

func ana(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/ana", ana)

}

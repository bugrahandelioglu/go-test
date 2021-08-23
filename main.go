package main

import "net/http"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("ındex page"))

	x := r.URL.Path[1:]
	data := ""

	if len(x) > 0 {
		data = "merhaba " + x + "!"
	} else {
		data = "index page"
	}
	w.Write([]byte(data))

}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}

func anaHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GİRİŞ EKRANI"))
}
func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/ana", anaHandler)

	//http.ListenAndServe(":8080", nil)
}

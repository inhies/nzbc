package main

import "net/http"

func main() {
	//http.HandleFunc("/api", server.APIhandler)
	http.HandleFunc("/api/", handler)
	http.ListenAndServe(":5001", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../testserveroutput.xml")
}

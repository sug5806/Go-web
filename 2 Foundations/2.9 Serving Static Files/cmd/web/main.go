package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Static file Serving을 위한 파일서버를 새로 만든다.
	//
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// /static 접두사를 제거한다.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))


	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

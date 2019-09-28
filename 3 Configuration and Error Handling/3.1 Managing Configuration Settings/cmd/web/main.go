package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// flag값을 읽어 세팅을한다 기본값은 4000
	// 반환값을 값 자체가 아니라 포인터이다
	addr := flag.String("addr", ":4000", "HTTP network address")
	//addr := os.Getenv("SNIPPETBOX_ADDR")
	// 항상 Parse를 해줘야함
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Static file Serving을 위한 파일서버를 새로 만든다.
	//
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// /static 접두사를 제거한다.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))


	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}

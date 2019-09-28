package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// flag값을 읽어 세팅을한다 기본값은 4000
	// 반환값을 값 자체가 아니라 포인터이다
	addr := flag.String("addr", ":4000", "HTTP network address")
	//addr := os.Getenv("SNIPPETBOX_ADDR")
	// 항상 Parse를 해줘야함
	flag.Parse()

	// 사용자 메시지 작성을 위한 로거를 만듬
	// 로그를 기록할 대상, 접두사, 추가정보(현지 날짜 및 시각)을 파이프라인(|)을 통해서 결합함
	//infoLog := log.New(os.Stdout, "INFO/t", log.Ldate|log.Ltime)
	// 기록할 대상은 stderr, 날짜, 시각, 파일이름을 포함시킴.
	// Lshortfile대신 전체 파일경로를 포함하면 Llongfile 플래그 사용가능, log.LUTC를 추가하여 UTC날짜 시간 사용가능
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)



	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Static file Serving을 위한 파일서버를 새로 만든다.
	//
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// /static 접두사를 제거한다.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr:  *addr,
		ErrorLog: errorLog,
		Handler: mux,
	}


	log.Printf("Starting server on %s", *addr) // Information message
	//err := http.ListenAndServe(*addr, mux)

	// ListenAndServe() 메소드는 작성한 http.Server 구조체를 호출한다.
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
	//log.Fatal(err) // Error message
}

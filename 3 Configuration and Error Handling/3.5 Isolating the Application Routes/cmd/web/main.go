package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// 3.3 애플리케이션 전체에 대한 종속성을 보유하도록 application 구조체 정의
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// flag값을 읽어 세팅을한다 기본값은 4000
	// 반환값을 값 자체가 아니라 포인터이다
	addr := flag.String("addr", ":4000", "HTTP network address")
	//addr := os.Getenv("SNIPPETBOX_ADDR")
	// 항상 Parse를 해줘야함
	flag.Parse()

	// 사용자 메시지 작성을 위한 로거를 만듬
	// 로그를 기록할 대상, 접두사, 추가정보(현지 날짜 및 시각)을 파이프라인(|)을 통해서 결합함
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// 기록할 대상은 stderr, 날짜, 시각, 파일이름을 포함시킴.
	// Lshortfile대신 전체 파일경로를 포함하면 Llongfile 플래그 사용가능, log.LUTC를 추가하여 UTC날짜 시간 사용가능
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// 3.3 의존성을 포함하는 새로운 구조체 인스턴스 생성
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// 3.5 routess.go 파일로 옮김
	//// 3.3 app 구조체의 메서드를 핸들러 함수로 사용하도록 경로 선언 변경
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", app.home)
	//mux.HandleFunc("/snippet", app.showSnippet)
	//mux.HandleFunc("/snippet/create", app.createSnippet)
	//
	//// Static file Serving을 위한 파일서버를 새로 만든다.
	////
	//fileServer := http.FileServer(http.Dir("./ui/static"))
	//
	//// /static 접두사를 제거한다.
	//mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr) // Information message
	//err := http.ListenAndServe(*addr, mux)

	// ListenAndServe() 메소드는 작성한 http.Server 구조체를 호출한다.
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
	//log.Fatal(err) // Error message
}

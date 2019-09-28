package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/". If it doesn't, use
	// the http.NotFound() function to send a 404 response to the client.
	// Importantly, we then return from the handler. If we don't return the handler
	// would keep executing and also write the "Hello from SnippetBox" message.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	_, _ = w.Write([]byte("Hello from Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	// Query String으로 들어온 값중 ID값을 뽑아
	// 문자열을 정수로 변환한다음 비교한다
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	//_, _ = w.Write([]byte("Display a specific snippet..."))
	// Fprintf의 첫번쨰 매개변수는 io.Writer 인데 http.ResponseWriter 객체가
	// Write 메소드를 가지고 있으므로 인자로 넘겨줄 수 있다.
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// 잘못된 요청을 보내면 POST요청만 가능하다는 메시지를 보내줌
		w.Header().Set("Allow", http.MethodPost)

		// Set 함수를 통해 Content-Tyep을 오버라이딩함
		//w.Header().Set("Content-Type", "application/json")
		//w.Write([]byte(`{"name":"Alex"}`))

		// Add, Del 함수를 통해 헤더 메시지를 추가, 삭제 할 수 있음
		w.Header().Add("asdfsaf", "xczvkljsalk")

		// 시스템이 자동으로 생성한 헤더를 삭제하고 싶다면 nil로 변경한다
		// Date를 삭제한다
		w.Header()["Date"] = nil

		// Write를 호출하기전에 WriteHeader를 호출해야 한다
		// WriteHeader는 response당 한번만 호출할 수 있으며
		// 그 이상을 호출하면 Go가 경고 메시지를 보낸다
		//w.WriteHeader(405)
		//_, _ = w.Write([]byte("Method Not Allowed"))

		http.Error(w, "Method Not Allowed", 405)
		return
	}

	_, _ = w.Write([]byte("Create a new snippet...."))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/snippet", showSnippet)
	http.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", nil)
	log.Fatal(err)
}

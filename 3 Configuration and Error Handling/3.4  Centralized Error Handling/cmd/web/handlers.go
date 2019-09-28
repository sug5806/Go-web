package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)
// 3.3 home 핸들러 함수는 이제 application에 대한 메소드이므로 해당 필드에 접근 가능
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		//http.NotFound(w, r) 3.4
		app.notFound(w) // 3.4 helper 함수 사용
		return
	}

	// html.page.html은 반드시 첫번째에 위치해야함
	files := []string{
		"./ui/html/home.page.html",
		"./ui/html/base.layout.html",
		// footer html을 포함시킨다.
		"./ui/html/footer.partial.html",
	}

	// template.ParseFiles 함수로 템플릿파일을 템플릿 셋으로 읽는다.
	// 오류가 있으면 오류를 기록하고 500에러를 보낸다
	// ParseFiles 함수는 파일을 읽고 템플릿 셋에 템플릿을 저장해라

	ts, err := template.ParseFiles(files...)
	if err != nil {
		//log.Println(err.Error())

		// 3.3 표준 로거 대신 로그메시지 작성
		app.errorLog.Println(err.Error())
		//http.Error(w, "Internal Server Error", 500) 3.4
		app.serverError(w, err) // 3.4 helper 함수 사용
		return
	}

	// Execute 메소드를 사용하여 템플릿 컨텐츠를 response body으로 작성한다.
	// 마지막 인자는 동적 데이터를 나타내며 현재는 nil로 둔다
	err = ts.Execute(w, nil)
	if err != nil{
		//log.Println(err.Error())

		// 3.3 표준 로거 대신 로그메시지 작성
		app.errorLog.Println(err.Error())
		//http.Error(w, "Internal Server Error", 500) 3.4
		app.serverError(w, err) // 3.4 helper 함수 사용
	}

	//_, _ = w.Write([]byte("Hello from Snippetbox"))
}

// 3.3 showSnippet 핸들러 함수는 이제 application에 대한 메소드이므로 해당 필드에 접근 가능
func (app *application)showSnippet(w http.ResponseWriter, r *http.Request) {
	// Query String으로 들어온 값중 ID값을 뽑아
	// 문자열을 정수로 변환한다음 비교한다
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		//http.NotFound(w, r) 3.4
		app.notFound(w) // 3.4 helper 함수 사용
		return
	}
	//_, _ = w.Write([]byte("Display a specific snippet..."))
	// Fprintf의 첫번쨰 매개변수는 io.Writer 인데 http.ResponseWriter 객체가
	// Write 메소드를 가지고 있으므로 인자로 넘겨줄 수 있다.
	_, _ = fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// 3.3 createSnippet 핸들러 함수는 이제 application에 대한 메소드이므로 해당 필드에 접근 가능
func (app *application)createSnippet(w http.ResponseWriter, r *http.Request) {
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

		//http.Error(w, "Method Not Allowed", 405) 3.4
		//app.clientError(w, http.StatusMethodNotAllowed) 3.4 helper 함수 사용
		return
	}

	_, _ = w.Write([]byte("Create a new snippet...."))
}
package main

import (
	"fmt" // 3.4
	"net/http" // 3.4
	"runtime/debug" // 3.4
)

// 3.4 serverError 도우미는 오류 메시지와 스택 추적을 error Log에 기록한다음
// 3.4 일반 500 내부 서버 오류 응답을 사용자에게 보냄
func (app *application) serverError(w http.ResponseWriter, err error){
	// 3.4 debug.Stack() 함수를 통해 현재 고루틴에 대한 스택추적을 가져와서 로그 메시지에 추가한다.
	// 3.4 스택 추적을 통해 응용 프로그램의 실행 경로를 확인할 수 있으면 오류를 디버깅할 때 도움이 될 수 있다.
	trace := fmt.Sprintf("%n%s", err.Error(), debug.Stack())
	//app.errorLog.Println(trace) 3.4
	// 파일 이름과 줄 번호를 스택 추적에서 한단계 뒤로 올림으로써 3.4
	// 오류가 어디서 발생 했는지 명확하게 알수있다 3.4
	app.errorLog.Output(3, trace)

	// 3.4 StatusText 함수를 사용하여 지정된 HTTP상태 코드의 사람이 읽기 쉬운 텍스트 표현을 자동으로 생성
	// 3.4 예를 들어 400은 Bad Request 문자열을 반환함
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// 3.4 특정 상태코드와 해당 설명을 사용자에게 보냄
func (app *application) clientError(w http.ResponseWriter, status int){
	// 3.4 StatusText 함수를 사용하여 지정된 HTTP상태 코드의 사람이 읽기 쉬운 텍스트 표현을 자동으로 생성
	// 3.4 예를 들어 400은 Bad Request 문자열을 반환함
	http.Error(w, http.StatusText(status), status)
}

// 3.4 일관성 있게 Not Found도 구현
// 3.4 404에러를 사용자에게 보냄
func (app *application) notFound(w http.ResponseWriter){
	// 3.4 StatusText 함수를 사용하여 지정된 HTTP상태 코드의 사람이 읽기 쉬운 텍스트 표현을 자동으로 생성
	// 3.4 예를 들어 400은 Bad Request 문자열을 반환함
	app.clientError(w, http.StatusNotFound)
}
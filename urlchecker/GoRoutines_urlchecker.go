package main

import (
	"fmt"
	"net/http"
)

type RequestResult struct {
	url    string
	status string
}

func main() {
	c := make(chan RequestResult) //c라는 이름의 채널을 생성.

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		go hitURL(url, c)
		//go를 앞에 붙여 GoRoutines를 활성화
		//go를 앞에 붙이면 함수가 순차 실행이 아닌 동시 실행됨.(속도 향상)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c) //urls의 길이만큼 채널 c에 메시지가 왔을 것. 그만큼 받아와서 출력
	}
}

func hitURL(url string, c chan RequestResult) {
	resp, err := http.Get(url)
	status := "OK"                            //status OK로 생성
	if err != nil || resp.StatusCode >= 400 { //err가 nil이 아님 & StatusCode가 400을 넘어가는 건 응답에 실패했다는 것
		status = "FAILED" //응답에 문제 있으면 FAILED로 변경
	}
	c <- RequestResult{url: url, status: status} //채널 c에 전송
}

package main

import (
	"fmt"
	"learngo/myretriever/mock"
	real2 "learngo/myretriever/real"
)

type Retiever interface {
	Get(url string) string
	//func Get(url string) string
}

func download(r Retiever) string {
	return r.Get("http://www.gqapwd.com")
}

func main() {
	var r Retiever
	r = mock.Retrieve{"this is gqa"}
	r = real2.Retriever{}
	fmt.Println(download(r))
}

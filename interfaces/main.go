package main

import (
	"first/interfaces/mock"
	"first/interfaces/real"
	"fmt"
)

const url = "https://www.google.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string)
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func download(r Retriever) string {
	return r.Get("https://www.google.com")
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"Content": "This is source code from google.com",
	})
	return s.Get(url)
}

type Book struct {
	name string
	title string
	author string
}

type PageTurning interface {
	TurnPage() string
}

func (book *Book) TurnPage() string {
	return "It's been done"
}
func SumPages(turning PageTurning) uint {
	return 10000
}

func main() {
	//var r Retriever
	//fmt.Printf("r Retriever: %T, %v\n", r, r)


	//r = &mock.Retriever{
	//	Contents: "This is fake retriever",
	//}
	//fmt.Printf("mock Retriever: %T, %v\n", r, r)
	//fmt.Println(download(r))

	rp := mock.Retriever{
		Contents: "default nothing",
	}
	fmt.Println(rp)
	rp.Post(url, map[string]string{"Content": "ss"})
	fmt.Println(session(&rp))
	fmt.Println(rp)

	fmt.Println("-----------------------")


	book := Book{"Flipped", "Do you remember your first love?", "American writer"}
	fmt.Println(book)

	fmt.Println(book.TurnPage())

	fmt.Println(SumPages(&book))



	//r = real.Retriever{
	//	UserAgent: "Mozilla: 5.0",
	//}
	//fmt.Printf("real Retriever %T, %v\n", r, r)
	//fmt.Println(r)

	//inspect(r)

	// Type assersion
	//realRetriever := r.(real.Retriever)
	//fmt.Println(realRetriever.UserAgent)
}

func inspect(r Retriever) {
	// switch 类型检查
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("mock.Retriever -->", v.Contents)
	case real.Retriever:
		fmt.Println("real.Retriever -->", v.UserAgent)
	}
}
package mock

import "fmt"

type Retriever struct {
	Contents string
}

func (r *Retriever) Get(url string) string {
	return r.Contents
}

func (r *Retriever) Post(url string, form map[string]string) {
	fmt.Println(form["Content"])
	if value, ok := form["name"]; ok {
		fmt.Println(value)
	} else {
		fmt.Printf("key %s not found\n", "name")
	}
	r.Contents = form["Content"]
}


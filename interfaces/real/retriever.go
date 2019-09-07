package real

type Retriever struct {
	UserAgent string
}

func (r Retriever) Get(url string) string {
	return r.UserAgent
}

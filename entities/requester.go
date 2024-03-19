package entities

import (
	"io"
	"net/http"
)

func NewRequester(dataStructure HasURL, ch chan<- string) *requester {
	return &requester{dataStructure, ch}
}

type requester struct {
	dataStructure HasURL
	ch            chan<- string
}

func (e *requester) Execute() {
	data := e.getData()
	e.ch <- string(*data)
}

func (e *requester) getData() *[]byte {
	req, err := http.Get(e.dataStructure.GetURL())
	if err != nil {
		panic(err.Error())
	}
	defer req.Body.Close()

	data, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err.Error())
	}

	return &data
}

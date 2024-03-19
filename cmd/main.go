package main

import (
	"os"
	"time"

	"github.com/juliopjr/fullcycle-desafio-multithreading/entities"
)

func main() {
	cep := os.Args[1]

	chViaCEP := make(chan string)
	chBrasilAPI := make(chan string)

	go entities.NewRequester(entities.NewViaCEP(cep), chViaCEP).Execute()
	go entities.NewRequester(entities.NewBrasilAPI(cep), chBrasilAPI).Execute()

	select {
	case data := <-chViaCEP:
		println("-----> VIA CEP")
		println(data)
	case data := <-chBrasilAPI:
		println("-----> Brasil API")
		println(data)
	case <-time.After(1 * time.Second):
		println("timeout")
	}
}

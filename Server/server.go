package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("iniciando server....")
	http.HandleFunc("/", ApiCotacao)
	http.ListenAndServe(":8080", nil)

}

func ApiCotacao(w http.ResponseWriter, r *http.Request) {
	fmt.Println("iniciando Api....")
	req, err := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro na requisição da url %v \n", err)
	}
	resp, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro na resposta da url %v \n", err)
	}

	fmt.Println("chegou aqui")

	fmt.Println(string(resp))

}

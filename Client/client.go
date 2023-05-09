package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*300)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost/8080", nil)
	if err != nil {
		log.Println("Erro ao tentar acessar o server", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Erro ao retornar a resposta do Server", err)
	}

	// defer resp.Body.Close()
	// io.Copy(os.Stdout, resp.Body)
	fmt.Println(resp)

}

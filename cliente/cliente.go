package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Nota struct {
	Nome string  `json:"nome"`
	Nota float64 `json:"nota"`
}

type Media struct {
	Nota
}

func main() {
	notas := []Nota{
		{
			Nome: "João",
			Nota: 8.0,
		},
		{
			Nome: "Bruno",
			Nota: 5.0,
		},
		{
			Nome: "Giovanna",
			Nota: 10.0,
		},
	}

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(notas); err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:8080/nota", "application/json", &body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	medias := []Media{}

	if err := json.NewDecoder(resp.Body).Decode(&medias); err != nil {
		panic(err)
	}

	fmt.Println("Medias", medias)

}

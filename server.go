package main

//Importando pacotes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Estruturação dos dados

type Nota struct {
	Nome string  `json:"nome"`
	Nota float64 `json:"nota"`
}

// Função com as notas estáticas do servidor
// Método ListenAndServe inicia um servidor HTTP com o endereço determinado

func main() {

	http.HandleFunc("/nota", calculaNota)

	fmt.Println("servidor rodando em http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}

// função para calcular nota utilizando como parametros metodos que permitem a construção da resposta HTTP e trata exceções

func calculaNota(w http.ResponseWriter, r *http.Request) {
	notas := []Nota{}
	if err := json.NewDecoder(r.Body).Decode(&notas); err != nil {
		panic(err)
	}

	maiorNota := 0.0

	for _, nota := range notas {
		if nota.Nota > maiorNota {
			maiorNota = nota.Nota
		}
	}

	for i := range notas {
		notas[i].Nota = notas[i].Nota * 10 / maiorNota
	}

	if err := json.NewEncoder(w).Encode(notas); err != nil {
		panic(err)
	}

	fmt.Printf("Recebendo conexão de %s\nNotas: %v\n", r.RemoteAddr, notas)
}

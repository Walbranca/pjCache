package src

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func IniciarAplicacao() {
	fmt.Println("teste2")
	Database = BD{}
	erro := Database.Conectar()
	if erro != nil {
		log.Fatal(erro)
		return
	}
	router := mux.NewRouter()
	router.HandleFunc("/pj/{cnpj}", HttpGetPessoaJuridica).Methods("GET")
	log.Fatal(http.ListenAndServe(":80", router))
}

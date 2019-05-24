package src

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HttpGetPessoaJuridica(w http.ResponseWriter, r *http.Request) {
	cnpj := mux.Vars(r)["cnpj"]
	pj, _ := GetPessoaJuridica(cnpj)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	erro := json.NewEncoder(w).Encode(pj)
	if erro != nil {
		log.Fatal(erro)
	}
}

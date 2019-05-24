package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

func getReceitaWS(cnpj string) (receitaws ReceitaWS, erro error) {
	url := fmt.Sprintf("https://www.ReceitaWS.com.br/v1/cnpj/%s", cnpj)
	response, erro := http.Get(url)
	if erro != nil {
		return
	}
	erro = json.NewDecoder(response.Body).Decode(&receitaws)
	if erro != nil {
		return
	}
	if receitaws.Status != "OK" {
		erro = errors.New(receitaws.Message)
	}
	return
}

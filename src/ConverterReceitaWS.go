package main

import (
	"strings"
	"time"
)

func converterReceitaWS(receitaws ReceitaWS) (pj PessoaJuridica) {
	pj.AtividadePrincipal = receitaws.AtividadePrincipal[0].Text
	pj.Bairro = receitaws.Bairro
	pj.CEP = receitaws.CEP
	pj.CNPJComMascara = receitaws.CNPJ
	pj.Cidade = receitaws.Municipio
	pj.Complemento = receitaws.Complemento
	dataAbertura, _ := time.Parse("02/01/2006", receitaws.Abertura)
	pj.DataAbertura = dataAbertura
	pj.Email = receitaws.Email
	pj.Estado = receitaws.UF
	pj.Fantasia = receitaws.Fantasia
	pj.Logradouro = receitaws.Logradouro
	pj.NaturezaJuridica = receitaws.NaturezaJuridica
	pj.Numero = receitaws.Numero
	pj.RazaoSocial = receitaws.Nome
	pj.Situacao = receitaws.Situacao
	socios := make([]Socio, 0)
	for i := 0; i < len(receitaws.QSA); i++ {
		socio := Socio{
			Nome:         receitaws.QSA[i].Nome,
			Qualificacao: receitaws.QSA[i].Qual,
		}
		socios = append(socios, socio)
	}
	pj.Socios = socios
	pj.Telefone = receitaws.Telefone
	pj.Tipo = receitaws.Tipo
	pj.CNPJSemMascara = strings.Replace(pj.CNPJComMascara, ".", "", 2)
	pj.CNPJSemMascara = strings.Replace(pj.CNPJSemMascara, "/", "", 1)
	pj.CNPJSemMascara = strings.Replace(pj.CNPJSemMascara, "-", "", 1)
	pj.UltimaAtualizacao = time.Now()
	pj.Porte = receitaws.Porte
	return
}

package main

import (
	"time"
)

type PessoaJuridica struct {
	ID                 int64     `json:"-"                   gorm:"column:id;primary_key"`
	Cache              bool      `json:"cache"               gorm:"-"`
	UltimaAtualizacao  time.Time `json:"ultima_atualizacao"  gorm:"column:ultima_atualizacao"`
	CNPJSemMascara     string    `json:"cnpj_sem_mascara"    gorm:"column:cnpj_sem_mascara"`
	CNPJComMascara     string    `json:"cnpj_com_mascara"    gorm:"column:cnpj_com_mascara"`
	RazaoSocial        string    `json:"razao_social"        gorm:"column:razao_social"`
	Fantasia           string    `json:"fantasia"            gorm:"column:fantasia"`
	Tipo               string    `json:"tipo"                gorm:"column:tipo"`
	DataAbertura       time.Time `json:"data_abertura"       gorm:"column:data_abertura"`
	AtividadePrincipal string    `json:"atividade_principal" gorm:"column:atividade_principal"`
	NaturezaJuridica   string    `json:"natureza_juridica"   gorm:"column:natureza_juridica"`
	Telefone           string    `json:"telefone"            gorm:"column:telefone"`
	Email              string    `json:"email"               gorm:"column:email"`
	Situacao           string    `json:"situacao"            gorm:"column:situacao"`
	Socios             []Socio   `json:"socios"              gorm:"foreignkey:IDPessoaJuridica;association_foreignkey:ID"`
	Logradouro         string    `json:"logradouro"          gorm:"column:logradouro"`
	Numero             string    `json:"numero"              gorm:"column:numero"`
	Complemento        string    `json:"complemento"         gorm:"column:complemento"`
	Bairro             string    `json:"bairro"              gorm:"column:bairro"`
	Cidade             string    `json:"cidade"              gorm:"column:cidade"`
	Estado             string    `json:"estado"              gorm:"column:estado"`
	CEP                string    `json:"cep"                 gorm:"column:cep"`
	Porte              string    `json:"porte"               gorm:"porte"`
}

func (PessoaJuridica) TableName() string {
	return "pjcache.pessoa_juridica"
}

type Socio struct {
	ID               int64  `json:"-" json:"-"   gorm:"Column:id;PRIMARY_KEY"`
	IDPessoaJuridica int64  `json:"-"            gorm:"Column:id_pessoa_juridica"`
	Nome             string `json:"nome"         gorm:"Column:nome"`
	Qualificacao     string `json:"qualificacao" gorm:"Column:qualificacao"`
	PaisOrigem       string `json:"pais_origem"  gorm:"Column:pais_origem"`
}

func (Socio) TableName() string {
	return "pjcache.socio"
}

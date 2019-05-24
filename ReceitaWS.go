package main

type ReceitaWS struct {
	Status                string      `json:"status"`                 // Indica a situação da requisição. Valores possíveis: OK, ERROR.
	Message               string      `json:"message"`                // Mensagem explicativa indicando erro. Válido apenas quando o campo status é ERROR.
	Billing               billing     `json:"billing"`                // Indica para a requisição como foi registrado a cobrança da consulta.
	CNPJ                  string      `json:"cnpj"`                   // CNPJ no formato 00.000.000/0000-00.
	Tipo                  string      `json:"tipo"`                   // Valores possíveis: MATRIZ, FILIAL.
	Abertura              string      `json:"abertura"`               // Data de abertura no formato dd/mm/aaaa.
	Nome                  string      `json:"nome"`                   // Razão social.
	Fantasia              string      `json:"fantasia"`               // Nome fantasia.
	AtividadePrincipal    []atividade `json:"atividade_principal"`    // Atividade principal. Um array com um elemento.
	AtividadesSecundarias []atividade `json:"atividades_secundarias"` // Atividades secundárias.
	NaturezaJuridica      string      `json:"natureza_juridica"`      // Natureza jurídica.
	Logradouro            string      `json:"logradouro"`             // Logradouro
	Numero                string      `json:"numero"`                 // Número
	Complemento           string      `json:"complemento"`            // Complemento
	CEP                   string      `json:"cep"`                    // CEP no formato 00.000-000.
	Bairro                string      `json:"bairro"`                 // Bairro
	Municipio             string      `json:"municipio"`              // Município
	UF                    string      `json:"uf"`                     // Sigla da Unidade da Federação.
	Email                 string      `json:"email"`                  // Email
	Telefone              string      `json:"telefone"`               // Telefone
	EFR                   string      `json:"efr"`                    // Ente Federativo Responsável, disponibilizado apenas para CNPJs da administração pública.
	Situacao              string      `json:"situacao"`               // Situação
	DataSituacao          string      `json:"data_situacao"`          // Data da situação no formato dd/mm/aaaa.
	MotivoSituacao        string      `json:"motivo_situacao"`        // Motivo da situação.
	SituacaoSspecial      string      `json:"situacao_especial"`      // Situação especial.
	DataSituacaoEspecial  string      `json:"data_situacao_especial"` // Data da situação especial no formato dd/mm/aaaa.
	CapitalSocial         string      `json:"capital_social"`         // Valor do capital social no formato 0.00.
	QSA                   []qsa       `json:"qsa"`                    // Quadro de Sócios e Administradores.
	Porte                 string      `json:"porte"`
}

type billing struct {
	Free     bool `json:"free"` // Indica se a requisição foi gratuita.
	Database bool `json:"BD"`   // Indica como a requisição foi resolvida: true (resolvida pelo banco de dados), false (resolvida em tempo real).
}

type atividade struct {
	Code string `json:"code"` // Código CNAE da atividade no formato 00.00-0-00.
	Text string `json:"text"` // Descrição da atividade.
}

type qsa struct {
	Nome         string `json:"nome"`           // Nome do sócio.
	Qual         string `json:"qual"`           // Qualificação do sócio.
	PaisOrigem   string `json:"pais_origem"`    // País de origem do sócio. Disponível apenas para sócios estrangeiros.
	NomeRepLegal string `json:"nome_rep_legal"` // Nome do representante legal. Disponível apenas para sócios com representantes.
	QualRepLegal string `json:"qual_rep_legal"` // Qualificação do representante legal. Disponível apenas para sócios com representantes.
}

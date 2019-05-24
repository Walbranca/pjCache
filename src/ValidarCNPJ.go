package src

const zero = uint8(48)

var pesosPrimeiroDigito = []uint8{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

var pesosSegundoDigito = []uint8{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

func ValidarCNPJ(cnpj string) bool {
	return validarCNPJAux(cnpj, pesosPrimeiroDigito) && validarCNPJAux(cnpj, pesosSegundoDigito)
}

func validarCNPJAux(cnpj string, pesos []uint8) bool {
	var i int
	var peso, soma uint8
	for i, peso = range pesos {
		numero := cnpj[i] - zero
		soma += numero * peso
	}
	numero := cnpj[i+1] - zero
	digito := 11 - (soma % 11)
	return digito == numero
}

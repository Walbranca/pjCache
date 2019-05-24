package src

func GetPessoaJuridica(cnpj string) (pj PessoaJuridica, erro error) {
	Database.GetConexao().Where("cnpj_sem_mascara = ?", cnpj).First(&pj)
	Database.GetConexao().Model(&pj).Related(&pj.Socios, "Socios")
	if pj.ID == 0 {
		receitaws, erro := getReceitaWS(cnpj)
		if erro == nil {
			pj = converterReceitaWS(receitaws)
			Database.GetConexao().Create(&pj)
			pj.Cache = false
		}
	} else {
		pj.Cache = true
	}
	return
}

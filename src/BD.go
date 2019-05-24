package src

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var Database BD

const PostgresHost = "192.168.1.142"

const PostgresPort = 50002

const PostgresUser = "golang"

const PostgresPassword = "golang"

const PostgresDatabase = "pjcache"

type BD struct {
	conexao *gorm.DB
}

func (database *BD) Conectar() (erro error) {
	linkConexao := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		PostgresUser,
		PostgresPassword,
		PostgresHost,
		PostgresPort,
		PostgresDatabase,
		"require",
	)
	fmt.Println(linkConexao)
	database.conexao, erro = gorm.Open("postgres", linkConexao)
	return
}

func (database *BD) GetConexao() *gorm.DB {
	return database.conexao
}

package config

import (
	"github.com/RiadMefti/go-api-boilerplate/types"
)

func NewConfig(DbHost, DbPort, DbUser, DbPassword, DbName, ServerAdress string) types.Config {

	return types.Config{
		DbHost: DbHost,
		DbPort: DbPort,
		DbUser: DbUser,
		DbPassword: DbPassword,
		DbName: DbName,
		ServerAdress: ServerAdress,
	}

}

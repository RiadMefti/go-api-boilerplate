package types

type User struct {
	ID                string `json:"id"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	EncryptedPassword string `json:"-"`
}

type Config struct {
	DbHost       string
	DbPort       string
	DbUser       string
	DbPassword   string
	DbName       string
	ServerAdress string
}

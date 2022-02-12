package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Host string `env:"APP_HOST" env-default:"localhost"`
	Port string `env:"APP_PORT" env-default:"8080"`
}

type DatabaseConfig struct {
	Driver      string `env:"DB_DRIVER" env-default:"mysql"`
	Host        string `env:"DB_HOST" env-default:"127.0.0.1"`
	Username    string `env:"DB_USER" env-default:"root"`
	Password    string `env:"DB_PASS" env-default:""`
	DBName      string `env:"DB_NAME" env-default:"fibo_dev"`
	Port        int    `env:"DB_PORT" env-default:"3306"`
	TablePrefix string `env:"DB_TABLE_PREFIX" env-default:"tbl_"`
	SSLMode     string `env:"DB_SSL_MODE" env-default:"disable"`
	SQLiteFile  string `env:"DB_SQLITE_FILE" env-default:"sqlite.db"`
}

type JWTConfig struct {
	Expire int64  `env:"JWT_EXPIRE" env-default:"3600"`
	Secret string `env:"JWT_SECRET" env-default:"1894cde6c936a294a478cff0a9227fd276d86df6573b51af5dc59c9064edf426"`
}

var App AppConfig

func LoadConfigFromEnv() {
	var err error
	err = godotenv.Load()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	err = cleanenv.ReadEnv(&App)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

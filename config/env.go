package config

import (
	"github.com/kelseyhightower/envconfig"
)

// インスタンスとして外部では読み込む
var AppEnv Env

func init() {
	// 環境変数の読み込み
	if err := envconfig.Process("", &AppEnv); err != nil {
		panic(err.Error())
	}
}

// AppEnv 環境変数の構造体
type Env struct {
	// 開発環境の名前
	Env string `default:"local"`
	// サーバのポート番号
	ServerPort string `default:"8000"`
	// DB関連
	DbHost     string `envconfig:"DB_HOST" default:"localhost"`
	DbPort     string `envconfig:"DB_PORT" default:"4432"`
	DbName     string `envconfig:"DB_NAME" default:"postgres"`
	DbUser     string `envconfig:"DB_USER" default:"postgres"`
	DbPassword string `envconfig:"DB_PASSWORD" default:"postgres"`
}

// // dev環境であるかどうか
func (a *Env) IsDev() bool {
	return a.Env == "development"
}

// prod環境であるかどうか
func (a *Env) IsProd() bool {
	return a.Env == "production"
}

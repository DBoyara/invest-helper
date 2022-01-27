package common

import (
	"os"
	"reflect"
)

type Settings struct {
	// App
	HttpPort string `env:"HTTP_PORT"`

	// DB
	DbHost     string `env:"DB_HOST"`
	DbPort     string `env:"DB_PORT"`
	DbName     string `env:"DB_NAME"`
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`

	// Tinkoff
	TinkoffToken        string `env:"TINKOFF_TOKEN"`
	TinkoffSandBoxToken string `env:"TINKOFF_SAND_BOX_TOKEN"`
	// Telegram
	TelegramToken  string `env:"TELEGRAM_TOKEN"`
	TelegramChatId string `env:"TELEGRAM_CHAT_ID"`
}

var globalSettings *Settings

func init() {
	// defaults
	globalSettings = &Settings{
		HttpPort:            "9000",
		DbHost:              "localhost",
		DbPort:              "5433",
		DbName:              "invest-helper",
		DbUser:              "user",
		DbPassword:          "pass",
		TinkoffToken:        "token",
		TinkoffSandBoxToken: "token",
		TelegramToken:       "token",
		TelegramChatId:      "0000",
	}

	// load from env
	st := reflect.TypeOf(*globalSettings)
	sv := reflect.ValueOf(globalSettings).Elem()

	for n := 0; n < st.NumField(); n++ {
		field := st.Field(n)
		envTag := field.Tag.Get("env")

		if envTag == "" {
			continue
		}

		if val, ok := os.LookupEnv(envTag); ok {
			sv.FieldByName(field.Name).SetString(val)
		}
	}
}

func GetSettings() *Settings {
	return globalSettings
}

package config

import (
	"flag"
	"log"

	"github.com/jinzhu/configor"
)

var Config = struct {
	Telegram struct {
		Token string `required:"true" env:"tgToken"`
	}
}{}

func init() {
	config := flag.String("file", "configs/config.yml", "configuration file")
	flag.StringVar(&Config.Telegram.Token, "tgToken", "", "telegram bot api token")
	flag.Parse()
	if err := configor.Load(&Config, *config); err != nil {
		log.Fatalln(err)
		return
	}
}

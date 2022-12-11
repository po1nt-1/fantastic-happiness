package config

import (
	"flag"
	"log"

	"github.com/jinzhu/configor"
)

var Config = struct {
	Tg struct {
		Token string `required:"true" env:"tgToken"`
	}
}{}

func init() {
	flag.StringVar(&Config.Tg.Token, "tgToken", "", "telegram bot api token")
	flag.Parse()
	if err := configor.Load(&Config, "configs/config.yml"); err != nil {
		log.Fatalln(err)
		return
	}
}

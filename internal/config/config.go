package config

import (
	"flag"
	"log"

	"github.com/jinzhu/configor"
)

var Config = struct {
	Tg struct {
		Debug bool   `required:"false" env:"tgDebug"`
		Token string `required:"true" env:"tgToken"`
	}
}{}

func init() {
	flag.BoolVar(&Config.Tg.Debug, "tgDebug", false, "telegram bot mode api")
	flag.StringVar(&Config.Tg.Token, "tgToken", "", "telegram bot api token")
	flag.Parse()
	if err := configor.Load(&Config, "configs/config.yml"); err != nil {
		log.Fatal(err)
	}
}

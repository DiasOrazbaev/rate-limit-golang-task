package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	Limit         int `env:"LIMIT"`
	BlockDuration int `env:"BLOCK"`
}

var (
	once     sync.Once
	instance *Config
)

func GetConfig() *Config {
	once.Do(func() {
		log.Println("getting config from env")
		instance = &Config{}
		if err := cleanenv.ReadEnv(instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Fatalln(help, err)
		}

		log.Println("configuration:", *instance)
	})

	return instance
}

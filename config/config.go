package config

import (
	"errors"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Link         string `toml:"link"`
	Token        string `toml:"token"`
	PhaseGroupId int64  `toml:"phaseGroupId"`
}

func LoadConfig(file string) (Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(file, &cfg)
	if err != nil {
		return Config{}, err
	}

	switch {
	case len(cfg.Link) == 0:
		return Config{}, errors.New("slug: field is empty")
	case len(cfg.Token) == 0:
		return Config{}, errors.New("token: field is empty")
	case cfg.PhaseGroupId <= 0:
		return Config{}, errors.New("phaseGroupId: field is empty")
	default:
		return cfg, nil
	}
}

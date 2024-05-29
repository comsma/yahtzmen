package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Environment string `yaml:"environment"`
	DSN         string `yaml:"dsn"`
	StripeKey   string `yaml:"stripe_key"`
	Storage     struct {
		Endpoint  string `yaml:"endpoint"`
		AccessKey string `yaml:"access_key"`
		SecretKey string `yaml:"secret_key"`
	} `yaml:"storage"`
}

func NewConfig(path string) Config {

	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("error opening config file: %v", err)
	}
	defer f.Close()

	var cfg Config

	dec := yaml.NewDecoder(f)

	if err := dec.Decode(&cfg); err != nil {
		log.Fatalf("error decoding config file: %v", err)
	}
	return cfg
}

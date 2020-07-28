package config

import (
	"github.com/joeshaw/envdecode"
)

type Config struct {
	Host         string `env:"HOST,default=0.0.0.0"`
	Port         string `env:"PORT,default=3000"`
	GinMode      string `env:"GIN_MODE,default=debug"`
	DatabaseFile string `env:"DATABASE_FILE,default=jqplay.boltdb"`
	AssetHost    string `env:"ASSET_HOST"`
	JQPath       string
	JQVer        string
}

func (c *Config) IsProd() bool {
	return c.GinMode == "release"
}

func Load() (*Config, error) {
	conf := &Config{}
	err := envdecode.Decode(conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

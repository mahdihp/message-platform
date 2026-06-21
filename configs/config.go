package configs

import (
	"fmt"
	"strings"

	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	MessagePlatform MessagePlatform `koanf:"mp"`
	Database        Database        `koanf:"mp.db"`
}

type MessagePlatform struct {
	Host                    string `koanf:"host"`
	Port                    int    `koanf:"port"`
	GracefulShutdownTimeout int    `koanf:"gracefulshutdowntimeout"`
}

type Database struct {
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	Name     string `koanf:"name"`
	Username string `koanf:"username"`
	Password string `koanf:"password"`
}

func LoadConfig() {

	k := koanf.New(".")
	err := k.Load(env.Provider("MP_", ".", func(s string) string {
		s = strings.TrimPrefix(s, "MP_")
		s = strings.ToLower(s)
		return strings.ReplaceAll(s, "_", ".")
	}), nil)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", k.All())
	var cfg Config

	if err := k.Unmarshal("", &cfg); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", cfg)
}

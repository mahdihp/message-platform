package configs

import (
	"strings"

	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	MessagePlatform MessagePlatform `koanf:"mp"`
	Database        Database        `koanf:"db"`
}

type MessagePlatform struct {
	Host                    string `koanf:"host"`
	Port                    int    `koanf:"port"`
	GracefulShutdownTimeout int    `koanf:"graceful_shutdown_timeout"`
}

type Database struct {
	Host     string `koanf:"db_host"`
	Port     int    `koanf:"db_port"`
	Name     string `koanf:"db_name"`
	Username string `koanf:"db_username"`
	Password string `koanf:"db_password"`
}

var k = koanf.New(".")

func LoadConfig() Config {

	err := k.Load(
		file.Provider(".env"),
		dotenv.Parser(),
	)

	if err != nil {
		panic(err)
	}
	nk := koanf.New(".")

	for k1, v := range k.All() {
		key := strings.ReplaceAll(k1, "MP_", "")
		key = strings.ToLower(key)
		nk.Set(key, v)

	}
	//fmt.Printf("%#v\n\n", nk.All())
	var cfg Config

	if err := nk.Unmarshal("", &cfg.MessagePlatform); err != nil {
		panic(err)
	}
	if err := nk.Unmarshal("", &cfg.Database); err != nil {
		panic(err)
	}

	//fmt.Printf("%+v\n", cfg)
	return cfg
}

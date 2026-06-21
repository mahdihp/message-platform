package configs

import (
	"fmt"
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
	GracefulShutdownTimeout int    `koanf:"gracefulshutdowntimeout"`
}

type Database struct {
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	Name     string `koanf:"name"`
	Username string `koanf:"username"`
	Password string `koanf:"password"`
}

var k = koanf.New(".")

func LoadConfig() {

	err := k.Load(
		file.Provider(".env"),
		dotenv.Parser(),
	)

	if err != nil {
		panic(err)
	}
	nk := koanf.New(".")

	for k1, v := range k.All() {

		//if strings.TrimPrefix(k1, "MP_DB_")  {
		//
		//}
		key := strings.ToLower(strings.TrimPrefix(k1, "MP_"))
		key = strings.ReplaceAll(key, "_", ".")

		nk.Set(key, v)
	}
	//fmt.Printf("%#v\n", nk.All())
	var cfg Config

	if err := nk.Unmarshal("", &cfg.MessagePlatform); err != nil {
		panic(err)
	}
	if err := nk.Unmarshal("", &cfg.Database); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", cfg)
}

package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type config struct {
	Server  server
	Mailgun mailgun
	From    contact
	To      contact
}

type server struct {
	Port int
}

type mailgun struct {
	Key    string
	Domain string
}

type contact struct {
	Name    string
	Address string
}

var v *viper.Viper
var c *config

func init() {
	v = viper.New()

	// Support for config files
	cp := os.Getenv("FEN_CONFIG_PATH")
	v.SetConfigName("config")
	v.AddConfigPath(cp)
	v.AddConfigPath(".")

	// Support for environment variables
	replacer := strings.NewReplacer(".", "_")
	v.SetEnvKeyReplacer(replacer)
	v.AutomaticEnv()

	// Map environment variables to structs
	v.BindEnv("mailgun.key", "FEN_MAILGUN_KEY")
	v.BindEnv("mailgun.domain", "FEN_MAILGUN_DOMAIN")
	v.BindEnv("from.name", "FEN_FROM_NAME")
	v.BindEnv("from.address", "FEN_FROM_ADDRESS")
	v.BindEnv("to.name", "FEN_TO_NAME")
	v.BindEnv("to.address", "FEN_TO_ADDRESS")
	v.BindEnv("server.port", "FEN_SERVER_PORT")

	// Configure default values
	v.SetDefault("server.port", 3000)

	v.ReadInConfig()
	v.WatchConfig()
}

func GetConfig() *config {
	v.Unmarshal(&c)
	return c
}

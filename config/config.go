package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type config struct {
	Server    server
	Sendgrid  sendgrid
	EmailFrom contact
	EmailTo   contact
}

type server struct {
	Port int
}

type sendgrid struct {
	Key string
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
	v.SetEnvPrefix("FEN")
	v.AutomaticEnv()

	// Map environment variables to structs
	v.BindEnv("sendgrid.key", "SENDGRID_KEY")
	v.BindEnv("emailFrom.name", "FROM_NAME")
	v.BindEnv("emailFrom.address", "FROM_ADDRESS")
	v.BindEnv("emailTo.name", "TO_NAME")
	v.BindEnv("emailTo.address", "TO_ADDRESS")
	v.BindEnv("server.port", "SERVER_PORT")

	// Configure default values
	v.SetDefault("server.port", 3000)

	v.ReadInConfig()
	v.WatchConfig()
}

func GetConfig() *config {
	v.Unmarshal(&c)
	return c
}

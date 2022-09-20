package mmclient

import (
	"fmt"
	"os"

	"github.com/tkanos/gonfig"
)

type Config struct {
	Server struct {
		HOST     string `yaml:"host"`
		PROTOCOL string `yaml:"protocol"`
		PORT     string `yaml:"port"`
	} `yaml:"server"`
	Bot struct {
		SAMPLE_NAME   string `yaml:"sample_name"`
		USER_EMAIL    string `yaml:"user_email"`
		USERNAME      string `yaml:"username"`
		USER_FIRST    string `yaml:"user_first"`
		USER_LAST     string `yaml:"user_last"`
		USER_PASSWORD string `yaml:"user_password"`
		TEAM_NAME     string `yaml:"team_name"`
		LOG_NAME      string `yaml:"log_name"`
	} `yaml:"bot"`
}

func GetConfig(params ...string) (*Config, error) {
	cfg := Config{}

	//TODO: Set default env to prod
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	// This here is only going to work so long as the bot is started in the main working directory.
	// We might want to set this to pull from some config directory in $GOROOT
	fileName := fmt.Sprintf("./%s_config.yml", env)

	err := gonfig.GetConf(fileName, &cfg)

	return &cfg, err
}

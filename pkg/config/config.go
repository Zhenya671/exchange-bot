package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	TelegramToken string
	DBPath        string `mapstructure:"db_file"`
}

func Init() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	viper.AddConfigPath("config")
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cnf Config
	if err := viper.Unmarshal(&cnf); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("db_file", &cnf.DBPath); err != nil {
		return nil, err
	}

	if err := parseEnv(&cnf); err != nil {
		return nil, err
	}

	return &cnf, nil
}

func parseEnv(cnf *Config) error {
	if err := viper.BindEnv("token"); err != nil {
		return err
	}

	cnf.TelegramToken = viper.GetString("token")
	return nil
}

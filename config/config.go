package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	TOKEN string
	AiUrl string
	AiKey string
}

func LoadConfig() (*Config, error) {
	wd, _ := os.Getwd()
	absEnv := filepath.Join(wd, ".env")
	parentAbsEnv := filepath.Join(filepath.Dir(wd), ".env")
	_ = godotenv.Load(".env", "../.env", absEnv, parentAbsEnv)

	cfg := &Config{
		TOKEN: os.Getenv("TELEGRAM_BOT_TOKEN"),
		AiUrl: os.Getenv("OPENAI_BASE_URL"),
		AiKey: os.Getenv("OPENAI_API_KEY"),
	}

	if cfg.TOKEN == "" {
		return nil, fmt.Errorf("TELEGRAM_BOT_TOKEN 环境变量未设置")
	}
	if cfg.AiUrl == "" {
		return nil, fmt.Errorf("AI_URL 环境变量未设置")
	}

	return cfg, nil
}

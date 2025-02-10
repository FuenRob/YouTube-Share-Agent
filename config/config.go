package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

type Config struct {
	YouTubeAPIKey    string
	YouTubeChannelID string
	DiscordToken     string
	ChannelID        string
	DeepSeekAPIKey   string
}

func LoadConfig() *Config {
	return &Config{
		YouTubeAPIKey:    os.Getenv("API_KEY_YOUTUBE"),
		YouTubeChannelID: os.Getenv("ID_CHANNEL_YOUTUBE"),
		DiscordToken:     os.Getenv("TOKEN_DISCORD"),
		ChannelID:        os.Getenv("ID_CHANNEL_DISCORD"),
		DeepSeekAPIKey:   os.Getenv("API_KEY_DEEPSEEK"),
	}
}

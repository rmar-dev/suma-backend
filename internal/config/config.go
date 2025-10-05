package config

import (
	"log"
	"github.com/spf13/viper"
)

type Config struct {
	Server     ServerConfig
	Database   DatabaseConfig
	JWT        JWTConfig
	Email      EmailConfig
	Tink       TinkConfig
	Encryption EncryptionConfig
}

type ServerConfig struct {
	Port        string
	Environment string
	FrontendURL string
	APIUrl      string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type JWTConfig struct {
	Secret        string
	Expiry        string
	RefreshSecret string
	RefreshExpiry string
}

type EmailConfig struct {
	From     string
	SMTPHost string
	SMTPPort int
	SMTPUser string
	SMTPPass string
}

type TinkConfig struct {
	ClientID     string
	ClientSecret string
	Environment  string
}

type EncryptionConfig struct {
	Key string
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file: %v", err)
	}

	config := &Config{
		Server: ServerConfig{
			Port:        viper.GetString("PORT"),
			Environment: viper.GetString("ENV"),
			FrontendURL: viper.GetString("FRONTEND_URL"),
			APIUrl:      viper.GetString("API_URL"),
		},
		Database: DatabaseConfig{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetString("DB_PORT"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			DBName:   viper.GetString("DB_NAME"),
			SSLMode:  viper.GetString("DB_SSLMODE"),
		},
		JWT: JWTConfig{
			Secret:        viper.GetString("JWT_SECRET"),
			Expiry:        viper.GetString("JWT_EXPIRY"),
			RefreshSecret: viper.GetString("REFRESH_SECRET"),
			RefreshExpiry: viper.GetString("REFRESH_EXPIRY"),
		},
		Email: EmailConfig{
			From:     viper.GetString("EMAIL_FROM"),
			SMTPHost: viper.GetString("SMTP_HOST"),
			SMTPPort: viper.GetInt("SMTP_PORT"),
			SMTPUser: viper.GetString("SMTP_USER"),
			SMTPPass: viper.GetString("SMTP_PASSWORD"),
		},
		Tink: TinkConfig{
			ClientID:     viper.GetString("TINK_CLIENT_ID"),
			ClientSecret: viper.GetString("TINK_CLIENT_SECRET"),
			Environment:  viper.GetString("TINK_ENVIRONMENT"),
		},
		Encryption: EncryptionConfig{
			Key: viper.GetString("ENCRYPTION_KEY"),
		},
	}

	// Set defaults
	if config.Server.Port == "" {
		config.Server.Port = "8080"
	}
	if config.Server.Environment == "" {
		config.Server.Environment = "development"
	}
	if config.Database.SSLMode == "" {
		config.Database.SSLMode = "disable"
	}

	return config, nil
}
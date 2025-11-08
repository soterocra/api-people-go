package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Config é a struct principal que guarda TUDO
type Config struct {
	Server ServerConfig `mapstructure:"server"`
	DBName string       `mapstructure:"DB_NAME"`
	DBUser string       `mapstructure:"DB_USER"`
	DBPass string       `mapstructure:"DB_PASS"`
}

// ServerConfig espelha a estrutura do config.yml
type ServerConfig struct {
	Port string `mapstructure:"port"`
}

// LoadConfig é a "Fábrica" de configuração
func LoadConfig() (Config, error) {
	var cfg Config

	// 1. Carregar .env com lib godotenv que joga os valores para as variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: .env não encontrado, lendo vars de ambiente.")
	}

	// 2. Configurar o Viper para ler o config.yml
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".") // Procurar na raiz do projeto

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("erro ao ler config.yml: %w", err)
	}

	// 3. Configurar o viper para também ler as env vars
	viper.SetEnvPrefix("")
	viper.AutomaticEnv() // Habilita leitura

	// 4. Ativa o resgate das variáveis desejadas
	viper.BindEnv("DB_NAME")
	viper.BindEnv("DB_USER")
	viper.BindEnv("DB_PASS")

	// 5. Deserializar (unmarshal) tudo
	// Viper vai juntar tudo que leu (env e yml) e preecher a struct cfg
	if err := viper.Unmarshal(&cfg); err != nil {
		return Config{}, fmt.Errorf("erro ao 'deserializar' config: %w", err)
	}

	return cfg, nil
}

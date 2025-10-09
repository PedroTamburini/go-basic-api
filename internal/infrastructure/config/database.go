package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// ConfigBancoDeDados armazena as conficurações para conexão com o banco de dados
type ConfigBancoDeDados struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	TimeZone string
}

// CarregarConfigBancoDeDados carrega as configurações do banco de dados a partir das variáveis de ambiente
// Em dev. tentarar carregar os dados de um arquivo .env
func CarregarConfigBancoDeDados() *ConfigBancoDeDados {
	godotenv.Load()
	return &ConfigBancoDeDados{
		Host:     BuscarAmbiente("DB_HOST", "localhost"),
		Port:     BuscarAmbiente("DB_PORT", "5432"),
		User:     BuscarAmbiente("DB_USER", "postgres"),
		Password: BuscarAmbiente("DB_PASSWORD", "postgres"),
		DBName:   BuscarAmbiente("DB_NAME", "go_basic_db"),
		SSLMode:  BuscarAmbiente("DB_SSLMODE", "disable"),
		TimeZone: BuscarAmbiente("DB_TIMEZONE", "Brazil/Acre"),
	}
}

// BuscarDSN retorna uma string com o DSN (Data Source Name) para conexão com o banco de dados
func (c *ConfigBancoDeDados) BuscarDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s timezone=%s", c.Host, c.User, c.Password, c.DBName, c.Port, c.SSLMode, c.TimeZone)
}

// BuscarAmbiente busca uma variável de ambiente retornando um valor padrão caso não exista
func BuscarAmbiente(chave, valorPadrao string) string {
	if valor, existe := os.LookupEnv(chave); existe {
		return valor
	}
	return valorPadrao
}

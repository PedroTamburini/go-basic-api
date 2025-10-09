package database

import (
	"fmt"
	"log"

	"github.com/PedroTamburini/go-basic-api/internal/domain"
	"github.com/PedroTamburini/go-basic-api/internal/infrastructure/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// IniciarDBPostgres inicializa a conexão com o banco de dados e realiza migrações
func IniciarDBPostgres(c *config.ConfigBancoDeDados) *gorm.DB {
	dsn := c.BuscarDSN()

	// Realiza a conexão com o banco de dados
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados PostgreSQL: %v", err)
	}
	fmt.Println("Conexão com PostgreSQL estabelecida com sucesso!")

	// Realiza a migração automática da tabela 'usuarios'
	err = db.AutoMigrate(&domain.Usuario{})
	if err != nil {
		log.Fatalf("Falha ao migrar o esquema do banco de dados: %v", err)
	}
	fmt.Println("Migrações do banco de dados realizadas com sucesso!")

	return db
}

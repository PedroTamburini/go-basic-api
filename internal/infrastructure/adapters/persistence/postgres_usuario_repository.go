package persistence

import (
	"errors"

	"github.com/PedroTamburini/go-basic-api/internal/application/ports"
	"github.com/PedroTamburini/go-basic-api/internal/domain"
	"gorm.io/gorm"
)

// PostgresUsuarioRepositorio é um adaptador que implementa a porta UsuarioRespositorio
// Utilizando o GORM para interagir com o banco de dados PostgreSQL
type PostgresUsuarioRepositorio struct {
	db *gorm.DB
}

// NovoPostgresUsuarioRepositorio cria uma nova instância de UsuarioRespositorio
func NovoPostgresUsuarioRepositorio(db *gorm.DB) ports.UsuarioRespositorio {
	return &PostgresUsuarioRepositorio{db: db}
}

// Salvar implementa a lógica de inserção de um novo usuário no banco de dados
func (r *PostgresUsuarioRepositorio) Salvar(usuario *domain.Usuario) error {
	resultado := r.db.Create(usuario)
	if resultado.Error != nil {
		// Implementar tratamento de erros expecífico mais tarde, algo como unique constraint violation para itens unique
		return resultado.Error
	}
	return nil
}

// EncontrarPorID implementa a lógica de busca de um usuário utilizando seu ID
func (r *PostgresUsuarioRepositorio) EncontrarPorID(idUsuario string) (*domain.Usuario, error) {
	var usuario domain.Usuario
	resultado := r.db.First(&usuario, "id=?", idUsuario)
	if resultado.Error != nil {
		if errors.Is(resultado.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("usuário não encontrado")
		}
		return nil, resultado.Error
	}
	return &usuario, nil
}

// EncontrarPorEmail implementa a lógica de busca de um usuário utilizando seu E-mail
func (r *PostgresUsuarioRepositorio) EncontrarPorEmail(email string) (*domain.Usuario, error) {
	var usuario domain.Usuario
	resultado := r.db.Where("email=?", email).First(&usuario)
	if resultado.Error != nil {
		if errors.Is(resultado.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("usuário não encontrado")
		}
		return nil, resultado.Error
	}
	return &usuario, nil
}

// EncontrarPorStatus implementa a lógica de busca de usuários com base em seu status
func (r *PostgresUsuarioRepositorio) EncontrarPorStatus(status string) ([]domain.Usuario, error) {
	var usuarios []domain.Usuario
	resultado := r.db.Where("status=?", status).Find(&usuarios)
	if resultado.Error != nil {
		return nil, resultado.Error
	}
	return usuarios, nil
}

// Atualizar implementa a lógica de atualização dos dados de um usuário no banco de dados
func (r *PostgresUsuarioRepositorio) Atualizar(usuario *domain.Usuario) error {
	resultado := r.db.Save(usuario) // Tanto insere quanto atualiza, com base em chave primária
	if resultado.Error != nil {
		return resultado.Error
	}
	return nil
}

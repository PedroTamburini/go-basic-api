package ports

import (
	"time"

	"github.com/PedroTamburini/go-basic-api/internal/domain"
)

// Configuração base de métodos Drivers (actor -> core)
// UsuarioServico define a porta de entrada para a lógica de negócios relacionada a usuários. Come in!
type UsuarioServico interface {
	RegistrarUsuario(nome, cpf, cargo, matricula, setor, email, telefone, sexo string, dataNascimento time.Time, senha string) (*domain.Usuario, error)
	LogarUsuario(email, senha string) (string, error)
	AprovarUsuario(idUsuario string) (*domain.Usuario, error)
	BuscarUsuariosPendentes() ([]domain.Usuario, error)
	// Mais métodos para manipulação de usuarios em breve...
}

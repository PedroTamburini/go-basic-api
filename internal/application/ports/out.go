package ports

import "github.com/PedroTamburini/go-basic-api/internal/domain"

// Configuração base de métodos Driven (core -> application)
// UsuarioRespositorio define a porta de saída para operações de persistência de usuário. Let's go outside!
type UsuarioRespositorio interface {
	Salvar(usuario *domain.Usuario) error
	Atualizar(usuario *domain.Usuario) error
	EncontrarPorID(idUsuario string) (*domain.Usuario, error)
	EncontrarPorEmail(email string) (*domain.Usuario, error)
	EncontrarPorStatus(status string) ([]domain.Usuario, error)
}

// TokenServico define a porta de saída para operações de gerenciamento de sessão
type TokenServico interface {
	GerarToken(idUsuario, cargo string) (string, error)
	ValidarToken(tokenString string) (string, string, error)
}

// HasherDeSenha define a porta de saída para operações de criptografia de senhas
type HasherDeSenha interface {
	CriarHashDeSenha(senha string) (string, error)
	VerificarHashDeSenha(senha, hash string) bool
}

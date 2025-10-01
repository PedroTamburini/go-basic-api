package services

import (
	"errors"
	"time"

	"github.com/PedroTamburini/go-basic-api/internal/application/ports"
	"github.com/PedroTamburini/go-basic-api/internal/domain"
)

// usuarioServico é a implementação da porta UsuarioServico
type usuarioServico struct {
	repo   ports.UsuarioRespositorio
	hasher ports.HasherDeSenha
	token  ports.TokenServico
}

// NovoUsuarioServico cria uma nova instância do serviço de usuário
func NovoUsuarioServico(repo ports.UsuarioRespositorio, hasher ports.HasherDeSenha, token ports.TokenServico) ports.UsuarioServico {
	return &usuarioServico{
		repo:   repo,
		hasher: hasher,
		token:  token,
	}
}

// RegistrarUsuario implementa a lógica de registro de usuário
func (s *usuarioServico) RegistrarUsuario(nome, cpf, cargo, matricula, setor, email, telefone, sexo string, dataNascimento time.Time, senha string) (*domain.Usuario, error) {
	// Definição de regras de negócios básicas
	if len(senha) < 8 {
		return nil, errors.New("a senha deve conter pelo menos 8 caracteres")
	}

	senhaHash, err := s.hasher.CriarHashDeSenha(senha)
	if err != nil {
		return nil, err
	}

	usuario := &domain.Usuario{
		Nome:           nome,
		CPF:            cpf,
		Cargo:          cargo,
		Matricula:      matricula,
		Setor:          setor,
		Email:          email,
		Telefone:       telefone,
		Sexo:           sexo,
		DataNascimento: dataNascimento.Format("2006-01-02"),
		SenhaHash:      senhaHash,
		Status:         domain.UsuarioStatusPendente,
	}

	if err := s.repo.Salvar(usuario); err != nil {
		return nil, err
	}

	return usuario, nil
}

// LogarUsuario implementa lógica de geração de token e login do usuário no sistema
func (s *usuarioServico) LogarUsuario(email, senha string) (string, error) {
	// Definição de regras de negócios básicas
	usuario, err := s.repo.EncontrarPorEmail(email)
	if err != nil {
		return "", errors.New("credenciais inválidas")
	}

	if !s.hasher.VerificarHashDeSenha(senha, usuario.SenhaHash) {
		return "", errors.New("credenciais inválidas")
	}

	token, err := s.token.GerarToken(usuario.ID, usuario.Cargo)
	if err != nil {
		return "", errors.New("falha ao gerar token de autenticação")
	}

	return token, nil
}

// AprovarUsuario implementa a lógica de aprovação de cadastro do usuário no sistema
func (s *usuarioServico) AprovarUsuario(idUsuario string) (*domain.Usuario, error) {
	// Definição de regras de negócios básicas
	usuario, err := s.repo.EncontrarPorID(idUsuario)
	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	if usuario.Status != domain.UsuarioStatusPendente {
		return nil, errors.New("o status do usuário não é pendente")
	}

	usuario.Status = domain.UsuarioStatusAprovado

	if err := s.repo.Atualizar(usuario); err != nil {
		return nil, err
	}

	return usuario, nil
}

// BuscarUsuariosPendentes implementa a lógica de buscar usuarios com status "PENDENTE"
func (s *usuarioServico) BuscarUsuariosPendentes() ([]domain.Usuario, error) {
	return s.repo.EncontrarPorStatus(domain.UsuarioStatusPendente)
}

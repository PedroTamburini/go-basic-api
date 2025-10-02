package security

import (
	"github.com/PedroTamburini/go-basic-api/internal/application/ports"
	"golang.org/x/crypto/bcrypt"
)

// BcryptHasher é um adaptador que implementa a porta HasherDeSenha usando bcrypt
type BcryptHasher struct{}

// NovoBcryptHasher cria uma nova instância de BcryptHasher
func NovoBcryptHasher() ports.HasherDeSenha {
	return &BcryptHasher{}
}

// CriarHashDeSenha implementa a lógica de gração de hash bcrypt para a senha fornecida
func (h *BcryptHasher) CriarHashDeSenha(senha string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	return string(bytes), err
}

// VerificarHashDeSenha implementa a lógica de comparação da senha em texto puro com o hash bcrypt
func (h *BcryptHasher) VerificarHashDeSenha(senha, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(senha))
	return err == nil
}

package security

import (
	"errors"
	"time"

	"github.com/PedroTamburini/go-basic-api/internal/application/ports"
	"github.com/golang-jwt/jwt"
)

// JWTTokenServico é um adaptador que implementa a  porta TokenServico para JWT
type JWTTokenServico struct {
	chaveSecreta []byte
}

// NovoJWTTokenServico cria uma nova instância para JWTTokenServico
func NovoJWTTokenServico(segredo string) ports.TokenServico {
	return &JWTTokenServico{chaveSecreta: []byte(segredo)}
}

// GerarToken implementa a lógica de criação de token JWT para o usuário com base em seu ID e cargo
func (s *JWTTokenServico) GerarToken(idUsuaio, cargo string) (string, error) {
	revindicacao := jwt.MapClaims{
		"authorized": true,
		"user_id":    idUsuaio,
		"role":       cargo,
		"exp":        time.Now().Add(time.Hour * 24).Unix(), // Token expira em 24h...
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, revindicacao)
	tokenString, err := token.SignedString(s.chaveSecreta)
	if err != nil {
		return "", errors.New("falha ao assinar o token")
	}
	return tokenString, nil
}

// ValidarToken implementa a lógica de análise e validação do token JWT
// Retorna o ID e o cargo do usuário se o token for válido
func (s *JWTTokenServico) ValidarToken(tokenString string) (string, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de assinatura inesperado")
		}
		return s.chaveSecreta, nil
	})

	if err != nil {
		return "", "", err
	}

	if revindicacao, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		idUsusario, ok := revindicacao["user_id"].(string)
		if !ok {
			return "", "", errors.New("user_id inválido no token")
		}

		role, ok := revindicacao["role"].(string)
		if !ok {
			return "", "", errors.New("role inválido no token")
		}
		return idUsusario, role, nil
	}
	return "", "", errors.New("token inválido")
}

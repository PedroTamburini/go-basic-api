package http

import (
	"net/http"
	"time"

	"github.com/PedroTamburini/go-basic-api/internal/application/ports"
	"github.com/PedroTamburini/go-basic-api/internal/infrastructure/adapters/http/dto"
	"github.com/gin-gonic/gin"
)

// UsuarioHandler é o adaptador HTTP para a porta UsuarioServico
type UsuarioHandler struct {
	usuarioServico ports.UsuarioServico
}

// NovoUsuarioHandler cria uma nova instância de UsuarioHandler
func NovoUsuarioHandler(us ports.UsuarioServico) *UsuarioHandler {
	return &UsuarioHandler{usuarioServico: us}
}

// @Summary Registrar novo usuário
// @Description Realiza o registro de um novo usuário no sistema.
// @Tags Usuarios
// @Accept json
// @Produce json
// @Param usuario body dto.RequisicaoDeRegistroDeUsuario true "Dados do usuário para registro"
// @Success 201 {object} domain.Usuario "Usuário registrado com sucesso"
// @Failure 400 {object} map[string]string "Requisição inválida"
// @Failure 500 {object} map[string]string "Erro interno do servidor"
// @Router /usuarios/registrar [post]
func (h *UsuarioHandler) RegistrarUsuario(c *gin.Context) {
	var req dto.RequisicaoDeRegistroDeUsuario
	if err := c.ShouldBindJSON(&req); err != nil {
		return
	}
	dataNascimento, err := time.Parse("2006-01-02", req.DataNascimento)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de data de nascimento inválido. Use YYYY-MM-DD."})
		return
	}
	usuario, err := h.usuarioServico.RegistrarUsuario(req.Nome, req.CPF, req.Cargo, req.Matricula, req.Setor, req.Email, req.Telefone, req.Sexo, dataNascimento, req.Senha)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, usuario)
}

// @Summary Logar usuário
// @Description Realiza o login de um usuário e retorna um token JWT.
// @Tags Autenticação
// @Accept json
// @Produce json
// @Param credenciais body dto.RequisicaoDeLogin true "Credenciais do usuário"
// @Success 200 {object} map[string]string "Token JWT"
// @Failure 400 {object} map[string]string "Requisição inválida"
// @Failure 401 {object} map[string]string "Credenciais inválidas ou usuário não aprovado"
// @Failure 500 {object} map[string]string "Erro interno do servidor"
// @Router /login [post]
func (h *UsuarioHandler) LogarUsuario(c *gin.Context) {
	var req dto.RequisicaoDeLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := h.usuarioServico.LogarUsuario(req.Email, req.Senha)
	if err != nil {
		if err.Error() == "credenciais inválidas" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha no login:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// @Summary Aprovar usuário
// @Description Aprova o cadastro de um usuário pendente. (Exige autenticação de atendente/admin)
// @Tags Usuarios
// @Accept json
// @Produce json
// @Param id path string true "ID do usuário a ser aprovado"
// @Success 200 {object} domain.Usuario "Usuário aprovado com sucesso"
// @Failure 400 {object} map[string]string "Requisição inválida"
// @Failure 404 {object} map[string]string "Usuário não encontrado"
// @Failure 500 {object} map[string]string "Erro interno do servidor"
// @Router /usuarios/{id}/aprovar [put]
func (h *UsuarioHandler) AprovarUsuario(c *gin.Context) {
	idUsuaio := c.Param("id")
	usuarioAprovado, err := h.usuarioServico.AprovarUsuario(idUsuaio)
	if err != nil {
		if err.Error() == "usuário não encontrado" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err.Error() == "o status do usuário não é pendente" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao aprovar usuário" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, usuarioAprovado)
}

// @Summary Buscar usuários pendentes
// @Description Retorna uma lista de usuários com status pendente. (Exige autenticação de atendente/admin)
// @Tags Usuarios
// @Produce json
// @Success 200 {array} domain.Usuario "Lista de usuários pendentes"
// @Failure 500 {object} map[string]string "Erro interno do servidor"
// @Router /usuarios/pendentes [get]
func (h *UsuarioHandler) BuscarUsuariosPendentes(c *gin.Context) {
	usuarios, err := h.usuarioServico.BuscarUsuariosPendentes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, usuarios)
}

package main

import (
	"fmt"
	"log"

	"github.com/PedroTamburini/go-basic-api/internal/application/services"
	"github.com/PedroTamburini/go-basic-api/internal/infrastructure/adapters/http"
	"github.com/PedroTamburini/go-basic-api/internal/infrastructure/adapters/persistence"
	"github.com/PedroTamburini/go-basic-api/internal/infrastructure/adapters/security"
	"github.com/PedroTamburini/go-basic-api/internal/infrastructure/config"
	"github.com/PedroTamburini/go-basic-api/internal/infrastructure/database"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/PedroTamburini/go-basic-api/docs"
)

func main() {
	// Buscar configurações do banco de dados e chave secreta do JWT
	ConfigBancoDeDados := config.CarregarConfigBancoDeDados()
	JWTSegredo := config.BuscarAmbiente("JWT_SECRET_KEY", "senhasupersecreta")

	// Inicializar o banco de dados
	bancoDeDados := database.IniciarDBPostgres(ConfigBancoDeDados)

	// Criar adaptadores de saída
	usuarioRepo := persistence.NovoPostgresUsuarioRepositorio(bancoDeDados)
	hasherDeSenha := security.NovoBcryptHasher()
	tokenServico := security.NovoJWTTokenServico(JWTSegredo)

	// Criar lógica de negócio.
	usuarioServico := services.NovoUsuarioServico(usuarioRepo, hasherDeSenha, tokenServico)

	// Criar adaptadores de entrada (handlers HTTP)
	usuarioHandler := http.NovoUsuarioHandler(usuarioServico)

	// Configurar roteador do Gin
	roteador := gin.Default()

	// Configurar rotas do Swagger
	roteador.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.BasePath = "/api/v1"

	// Agrupar rotas da API (v1)
	v1 := roteador.Group("/api/v1")
	{
		v1.POST("/login", usuarioHandler.LogarUsuario)
		v1.POST("/usuarios/registrar", usuarioHandler.RegistrarUsuario)
		v1.PUT("/usuarios/:id/aprovar", usuarioHandler.AprovarUsuario)
		v1.GET("/usuarios/pendentes", usuarioHandler.BuscarUsuariosPendentes)
	}

	// Iniciar o servidor Gin.
	porta := config.BuscarAmbiente("API_PORT", "8080")
	log.Printf("Go Basic API rodando na porta :%s", porta)
	if err := roteador.Run(fmt.Sprintf(":%s", porta)); err != nil {
		log.Fatalf("Erro ao iniciar o servidor Gin: %v", err)
	}
}

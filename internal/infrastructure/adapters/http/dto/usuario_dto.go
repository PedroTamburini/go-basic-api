package dto

// RequisicaoDeRegistroDeUsuario representa os dados necessários para registrar um novo usuário
type RequisicaoDeRegistroDeUsuario struct {
	Nome           string `json:"none" binding:"required"`
	CPF            string `json:"cpf" binding:"required"`
	Cargo          string `json:"cargo" binding:"required"`
	Matricula      string `json:"matricula" binding:"required"`
	Setor          string `json:"setor" binding:"required"`
	Email          string `json:"email" binding:"required"`
	Telefone       string `json:"telefone" binding:"required"`
	Sexo           string `json:"sexo" binding:"required"`
	DataNascimento string `json:"data_nascimento" binding:"required"`
	Senha          string `json:"senha" binding:"required,min=8"`
}

// RequisicaoDeLogin representa os dados necessários para o login de um usuário
type RequisicaoDeLogin struct {
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"required"`
}

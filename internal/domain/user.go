package domain

import "time"

// Entidade usuário representada na aplicação
type Usuario struct {
	ID             string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Nome           string    `json:"nome" gorm:"not null"`
	CPF            string    `json:"cpf" gorm:"unique;not null"`
	Cargo          string    `json:"cargo" gorm:"not null"`
	Matricula      string    `json:"matricula" gorm:"unique;not null"`
	Setor          string    `json:"setor" gorm:"not null"`
	Email          string    `json:"email" gorm:"not null"`
	Telefone       string    `json:"telefone"`
	Sexo           string    `json:"" gorm:"not null"`
	DataNascimento string    `json:"data_nascimento" gorm:"not null"`
	SenhaHash      string    `json:"-" gorm:"not null"`
	Status         string    `json:"status" gorm:"type:varchar(20);default:'PENDENTE'"`
	CriadoEm       time.Time `json:"criado_em" gorm:"autoCreateTime"`
	AtualizadoEm   time.Time `json:"atualizado_em" gorm:"autoCreateTime"`
}

const (
	UsuarioStatusPendente  = "PENDENTE"
	UsuarioStatusAprovado  = "APROVADO"
	UsuarioStatusRejeitado = "REJEITADO"
	UsuarioStatusInativo   = "INATIVO"
)

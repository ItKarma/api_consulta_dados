package controller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConsultCpf(c *gin.Context, db *sql.DB) {
	cpf := c.Query("cpf")

	var record struct {
		CPF             string `json:"cpf"`
		CNS             string `json:"cns"`
		Nome            string `json:"nome"`
		Nascimento      string `json:"nascimento"`
		Sexo            string `json:"sexo"`
		RacaCor         string `json:"raca_cor"`
		Falecido        string `json:"falecido"`
		DataFalecimento string `json:"data_falecimento"`
		Mae             string `json:"mae"`
		Pai             string `json:"pai"`
		Celular         string `json:"celular"`
		Telefone        string `json:"telefone"`
		Contato         string `json:"contato"`
		Email           string `json:"email"`
		Rua             string `json:"rua"`
		Numero          string `json:"numero"`
		Complemento     string `json:"complemento"`
		Bairro          string `json:"bairro"`
		Cidade          string `json:"cidade"`
		Estado          string `json:"estado"`
		CEP             string `json:"cep"`
		RG              string `json:"rg"`
		RGOrgaoEmissor  string `json:"rg_orgao_emissor"`
		RGDataEmissao   string `json:"rg_data_emissao"`
		NIS             string `json:"nis"`
	}

	query := `SELECT cpf, cns, nome, nascimento, sexo, raca_cor, falecido, data_falecimento,
		mae, pai, celular, telefone, contato, email, rua, numero, complemento,
		bairro, cidade, estado, cep, rg, rg_orgao_emissor, rg_data_emissao, nis
		FROM cadsus WHERE cpf = ?`

	err := db.QueryRow(query, cpf).Scan(
		&record.CPF, &record.CNS, &record.Nome, &record.Nascimento, &record.Sexo, &record.RacaCor,
		&record.Falecido, &record.DataFalecimento, &record.Mae, &record.Pai, &record.Celular, &record.Telefone,
		&record.Contato, &record.Email, &record.Rua, &record.Numero, &record.Complemento, &record.Bairro,
		&record.Cidade, &record.Estado, &record.CEP, &record.RG, &record.RGOrgaoEmissor, &record.RGDataEmissao, &record.NIS,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"Message": "CPF not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": "Error while querying CPF", "Error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, record)
}

func ConsultNome(c *gin.Context, db *sql.DB) {
	nome := c.Query("nome")

	var records []struct {
		CPF             string `json:"cpf"`
		CNS             string `json:"cns"`
		Nome            string `json:"nome"`
		Nascimento      string `json:"nascimento"`
		Sexo            string `json:"sexo"`
		RacaCor         string `json:"raca_cor"`
		Falecido        string `json:"falecido"`
		DataFalecimento string `json:"data_falecimento"`
		Mae             string `json:"mae"`
		Pai             string `json:"pai"`
		Celular         string `json:"celular"`
		Telefone        string `json:"telefone"`
		Contato         string `json:"contato"`
		Email           string `json:"email"`
		Rua             string `json:"rua"`
		Numero          string `json:"numero"`
		Complemento     string `json:"complemento"`
		Bairro          string `json:"bairro"`
		Cidade          string `json:"cidade"`
		Estado          string `json:"estado"`
		CEP             string `json:"cep"`
		RG              string `json:"rg"`
		RGOrgaoEmissor  string `json:"rg_orgao_emissor"`
		RGDataEmissao   string `json:"rg_data_emissao"`
		NIS             string `json:"nis"`
	}

	query := `SELECT cpf, cns, nome, nascimento, sexo, raca_cor, falecido, data_falecimento,
		mae, pai, celular, telefone, contato, email, rua, numero, complemento,
		bairro, cidade, estado, cep, rg, rg_orgao_emissor, rg_data_emissao, nis
		FROM cadsus WHERE nome = ?`

	// Executa a consulta
	rows, err := db.Query(query, nome)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Error while querying name", "Error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var record struct {
			CPF             string `json:"cpf"`
			CNS             string `json:"cns"`
			Nome            string `json:"nome"`
			Nascimento      string `json:"nascimento"`
			Sexo            string `json:"sexo"`
			RacaCor         string `json:"raca_cor"`
			Falecido        string `json:"falecido"`
			DataFalecimento string `json:"data_falecimento"`
			Mae             string `json:"mae"`
			Pai             string `json:"pai"`
			Celular         string `json:"celular"`
			Telefone        string `json:"telefone"`
			Contato         string `json:"contato"`
			Email           string `json:"email"`
			Rua             string `json:"rua"`
			Numero          string `json:"numero"`
			Complemento     string `json:"complemento"`
			Bairro          string `json:"bairro"`
			Cidade          string `json:"cidade"`
			Estado          string `json:"estado"`
			CEP             string `json:"cep"`
			RG              string `json:"rg"`
			RGOrgaoEmissor  string `json:"rg_orgao_emissor"`
			RGDataEmissao   string `json:"rg_data_emissao"`
			NIS             string `json:"nis"`
		}

		// Scaneia cada linha para a estrutura
		err := rows.Scan(
			&record.CPF, &record.CNS, &record.Nome, &record.Nascimento, &record.Sexo, &record.RacaCor,
			&record.Falecido, &record.DataFalecimento, &record.Mae, &record.Pai, &record.Celular, &record.Telefone,
			&record.Contato, &record.Email, &record.Rua, &record.Numero, &record.Complemento, &record.Bairro,
			&record.Cidade, &record.Estado, &record.CEP, &record.RG, &record.RGOrgaoEmissor, &record.RGDataEmissao, &record.NIS,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": "Error while scanning row", "Error": err.Error()})
			return
		}

		records = append(records, record)
	}

	if len(records) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Message": "Name not found"})
		return
	}

	c.JSON(http.StatusOK, records)
}

func ConsultCelular(c *gin.Context, db *sql.DB) {
	Celular := c.Query("celular")

	var records []struct {
		CPF             string `json:"cpf"`
		CNS             string `json:"cns"`
		Nome            string `json:"nome"`
		Nascimento      string `json:"nascimento"`
		Sexo            string `json:"sexo"`
		RacaCor         string `json:"raca_cor"`
		Falecido        string `json:"falecido"`
		DataFalecimento string `json:"data_falecimento"`
		Mae             string `json:"mae"`
		Pai             string `json:"pai"`
		Celular         string `json:"celular"`
		Telefone        string `json:"telefone"`
		Contato         string `json:"contato"`
		Email           string `json:"email"`
		Rua             string `json:"rua"`
		Numero          string `json:"numero"`
		Complemento     string `json:"complemento"`
		Bairro          string `json:"bairro"`
		Cidade          string `json:"cidade"`
		Estado          string `json:"estado"`
		CEP             string `json:"cep"`
		RG              string `json:"rg"`
		RGOrgaoEmissor  string `json:"rg_orgao_emissor"`
		RGDataEmissao   string `json:"rg_data_emissao"`
		NIS             string `json:"nis"`
	}

	query := `SELECT cpf, cns, nome, nascimento, sexo, raca_cor, falecido, data_falecimento,
		mae, pai, celular, telefone, contato, email, rua, numero, complemento,
		bairro, cidade, estado, cep, rg, rg_orgao_emissor, rg_data_emissao, nis
		FROM cadsus WHERE TRIM(celular) = ?`

	rows, err := db.Query(query, Celular)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Error while querying celular", "Error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var record struct {
			CPF             string `json:"cpf"`
			CNS             string `json:"cns"`
			Nome            string `json:"nome"`
			Nascimento      string `json:"nascimento"`
			Sexo            string `json:"sexo"`
			RacaCor         string `json:"raca_cor"`
			Falecido        string `json:"falecido"`
			DataFalecimento string `json:"data_falecimento"`
			Mae             string `json:"mae"`
			Pai             string `json:"pai"`
			Celular         string `json:"celular"`
			Telefone        string `json:"telefone"`
			Contato         string `json:"contato"`
			Email           string `json:"email"`
			Rua             string `json:"rua"`
			Numero          string `json:"numero"`
			Complemento     string `json:"complemento"`
			Bairro          string `json:"bairro"`
			Cidade          string `json:"cidade"`
			Estado          string `json:"estado"`
			CEP             string `json:"cep"`
			RG              string `json:"rg"`
			RGOrgaoEmissor  string `json:"rg_orgao_emissor"`
			RGDataEmissao   string `json:"rg_data_emissao"`
			NIS             string `json:"nis"`
		}

		if err := rows.Scan(
			&record.CPF, &record.CNS, &record.Nome, &record.Nascimento, &record.Sexo, &record.RacaCor,
			&record.Falecido, &record.DataFalecimento, &record.Mae, &record.Pai, &record.Celular, &record.Telefone,
			&record.Contato, &record.Email, &record.Rua, &record.Numero, &record.Complemento, &record.Bairro,
			&record.Cidade, &record.Estado, &record.CEP, &record.RG, &record.RGOrgaoEmissor, &record.RGDataEmissao, &record.NIS,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Message": "Error while scanning row", "Error": err.Error()})
			return
		}

		records = append(records, record)
	}

	if len(records) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Message": "No records found for the given celular"})
	} else {
		c.JSON(http.StatusOK, records)
	}
}

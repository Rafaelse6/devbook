package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES (?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
	SELECT DISTINCT p.*,  u.nick
	FROM publicacoes p
	JOIN usuarios u ON p.autor_id = u.id
	LEFT JOIN seguidores s ON u.id = s.usuario_id
	WHERE u.id = ? OR s.seguidor_id = ?
	ORDER BY p.criadaEm DESC;`,
		usuarioID, usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {
	linha, erro := repositorio.db.Query(`
	SELECT p.*, u.nick FROM 
	publicacoes p INNER JOIN usuarios u
	ON u.id = p.autor_id WHERE p.id = ?`,
		publicacaoID,
	)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}
	defer linha.Close()

	var publicacao modelos.Publicacao

	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao modelos.Publicacao) error {
	statement, erro := repositorio.db.Prepare("UPDATE publicacoes SET titulo = ?, conteudo = ? WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Publicacoes) Deletar(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare("DELETE FROM publicacoes WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		SELECT p.*, u.nick FROM publicacoes p
		JOIN usuarios u ON u.id = p.autor_id
		WHERE p.autor_id = ?`,
		usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (repositorio Publicacoes) Curtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare("UPDATE publicacoes SET curtidas = curtidas + 1 WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Publicacoes) Descurtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(`
		UPDATE publicacoes SET curtidas = 
		CASE 
			WHEN curtidas > 0 THEN curtidas - 1
			ELSE 0 
		END
		where id = ?
	`)
	if erro != nil {
		return erro
	}

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}

package models

import (
	"go_studies/web-store-app/db"
)

type Produto struct {
	Id          int
	Nome        string
	Descricao string
	Preco       float64
	Quantidade      int
}

func SelectProdutos() []Produto {
	db := db.ConnectDatabase()
	defer db.Close()

	selectDb, err := db.Query("SELECT * FROM produtos ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDb.Next() {
		var id, Quantidade int
		var Nome, Descricao string
		var Preco float64

		err = selectDb.Scan(&id, &Nome, &Descricao, &Preco, &Quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = Nome
		p.Descricao = Descricao
		p.Preco = Preco
		p.Quantidade = Quantidade

		produtos = append(produtos, p)
	}

	return produtos
}

func CreateProduto(Nome, Descricao string, Preco float64, Quantidade int) {
	db := db.ConnectDatabase()
	defer db.Close()

	insertDb, err := db.Prepare("INSERT INTO Produtos (Nome, Descricao, Preco, Quantidade) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertDb.Exec(Nome, Descricao, Preco, Quantidade)
}

func DeleteProduto(id string) {
	db := db.ConnectDatabase()
	defer db.Close()

	deleteDb, err := db.Prepare("DELETE FROM Produtos WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}

	deleteDb.Exec(id)
}

func EditProduto(id string) Produto {
	db := db.ConnectDatabase()
	defer db.Close()

	editDb, err := db.Query("SELECT * FROM Produtos WHERE id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	Produto := Produto{}

	for editDb.Next() {
		var id, Quantidade int
		var Nome, Descricao string
		var Preco float64

		err = editDb.Scan(&id, &Nome, &Descricao, &Preco, &Quantidade)
		if err != nil {
			panic(err.Error())
		}

		Produto.Id = id
		Produto.Nome = Nome
		Produto.Descricao = Descricao
		Produto.Preco = Preco
		Produto.Quantidade = Quantidade
	}

	return Produto
}

func UpdateProduto(Nome, Descricao string, Preco float64, id, Quantidade int) {
	db := db.ConnectDatabase()
	defer db.Close()

	updateDb, err := db.Prepare("UPDATE produtos SET Nome=$1, Descricao=$2, Preco=$3, Quantidade=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}

	updateDb.Exec(Nome, Descricao, Preco, Quantidade, id)
}

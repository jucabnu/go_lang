package controllers

import (
	"go_studies/web-store-app/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.SelectProdutos()

	templates.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		p := r.FormValue("preco")
		a := r.FormValue("quantidade")

		preco, err := strconv.ParseFloat(p, 64)
		if err != nil {
			log.Println("Erro convertendo o preço")
			panic(err.Error())
		}

		quantidade, err := strconv.Atoi(a)
		if err != nil {
			log.Println("Erro convertendo quantidade")
			panic(err.Error())
		}

		models.CreateProduto(nome, descricao, preco, quantidade)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	product_id := r.URL.Query().Get("id")
	models.DeleteProduto(product_id)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	product_id := r.URL.Query().Get("id")
	product := models.EditProduto(product_id)

	templates.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		i := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		p := r.FormValue("preco")
		a := r.FormValue("quantidade")

		id, err := strconv.Atoi(i)
		if err != nil {
			log.Println("Erro convertendo o ID para int")
			panic(err.Error())
		}

		preco, err := strconv.ParseFloat(p, 64)
		if err != nil {
			log.Println("Erro convertendo preço para float")
			panic(err.Error())
		}

		quantidade, err := strconv.Atoi(a)
		if err != nil {
			log.Println("Erro convertendo quantidade para int")
			panic(err.Error())
		}

		models.UpdateProduto(nome, descricao, preco, id, quantidade)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

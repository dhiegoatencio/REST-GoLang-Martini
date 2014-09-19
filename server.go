package main

import (
	"github.com/go-martini/martini"
	"net/http"
	"encoding/json"
)

func main() {
  	m := martini.Classic()

	m.Group("/books", func(r martini.Router) {
		r.Get   ("",     GetBooks)
	    r.Post  ("",     NewBook)
	    r.Get   ("/:id", GetBook)
	    r.Put   ("/:id", UpdateBook)
	    r.Delete("/:id", DeleteBook)
	})

  	m.Run()
}

func GetBooks(req http.ResponseWriter, res *http.Request) string {
	return "GetBooks"
}

func GetBook(req    http.ResponseWriter,
	         res    *http.Request,
	         params martini.Params) string{
	return "GetBook/" + params["id"]
}

func NewBook(req *http.Request, res http.ResponseWriter, ) string {
	var teste string
    json.NewDecoder(req.Body).Decode(&teste) 
	return "NewBook " + teste
}

func UpdateBook(req    http.ResponseWriter,
	            res    *http.Request,
	            params martini.Params) string {
	return "UpdateBook/" + params["id"]
}

func DeleteBook(req    http.ResponseWriter,
	            res    *http.Request,
	            params martini.Params) string {
	return "DeleteBook/" + params["id"]
}
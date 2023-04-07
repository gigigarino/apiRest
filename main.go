package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const port = ":8080"

type Book struct {
	ID     int //`json: "id"` estructura etiqueta
	Author string
	Title  string
	Price  int
}

var db []Book

func main() {

	b1 := Book{
		ID:     1,
		Title:  "Dune",
		Price:  1965,
		Author: "Frank Herbert",
	}

	b2 := Book{
		ID:     2,
		Title:  "Cita con Rama",
		Price:  1974,
		Author: "Arthur C. Clarke",
	}

	b3 := Book{
		ID:     3,
		Title:  "Un guijarro en el cielo",
		Price:  500,
		Author: "Isaac Asimov",
	}

	db = append(db, b1, b2, b3)

	r := gin.Default()

	//GETS
	r.GET("/", index)
	r.GET("/books", getbooks)
	r.GET("/books/:id", getbookid)

	//POST
	r.POST("/books", addBook)

	//PUT UPDATE
	r.PUT("/books/:id", updateBook)

	//delete 
	r.DELETE("/books/:id", deleteBook)

	r.Run(port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	//msj de escuchando el puerto
	log.Println("server listening to the port", port)

	//mj de error si algo falla
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalln(err)
	}

}

//w responde: respuesta del servidor al cliente
//r request: peticion del cliente al servidor

// metodo get index
func index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "bienvenido a mi increible api")
}

// metodo getbooks
func getbooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  db,
	})
}

type ResponseInfo struct {
	Error bool   `json:"error"`
	Data  string `json:"data"`
}

// metodo getbookid
func getbookid(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}

	for _, v := range db {
		if v.ID == id {
			ctx.JSON(http.StatusOK, gin.H{
				"error": false,
				"data":  v,
			})
			return
		}

	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": true,
		"data":  "book not found",
	})
}

// metodo post
func addBook(ctx *gin.Context) {
	request := ctx.Request

	var book Book
	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return

	}
	//agregar book
	db = append(db, book)
	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  db,
	})
}

// metodo put
func updateBook(ctx *gin.Context) {
	r := ctx.Request
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}
	var book Book
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  err.Error(),
		})
		return
	}
	//aca se usa el id que chiilla
	for i, v := range db {
		if v.ID == id {
			db[i] = book

		}
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data": db,
	})

}

func deleteBook(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data": err.Error(),
		})
		return
	}
	for i,v := range db {
		if v.ID == id {
			db = append(db[:i], db[i+1:]...)
		}
		}
			ctx.JSON(http.StatusOK, gin.H{
				"error": false,
				"data": db,
			})
		}
	


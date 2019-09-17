package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	melody "gopkg.in/olahol/melody.v1"
)

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "WebProyector - Iglesia Evangelica JERUSALEM",
	})
}

func control(c *gin.Context) {
	c.HTML(http.StatusOK, "control.html", gin.H{
		"title": "hola",
	})
}

type Titulos struct {
	Id     byte
	Titulo string
}

type Items struct {
	Id_item   byte
	Id_titulo byte
	Es_text   string
	En_text   string
}

var ar_clase string

func main() {
	puerto := flag.String("port", "8080", "Puerto del servidor")
	clase := flag.String("clase", "clasego", "Nombre de la clase a proyectar")

	flag.Parse()
	ar_clase = *clase

	r := gin.Default()
	m := melody.New()

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/public", "./public")

	r.GET("/", index)
	r.GET("/titulos", traertitulos)
	r.GET("/items/:item", traeritems)
	r.GET("/control", control)

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)

	})

	m.HandleDisconnect(func(s *melody.Session) {
		//fmt.Println(s)
	})

	fmt.Println(ar_clase)
	r.Run(":" + *puerto) // listen and serve on 0.0.0.0:8080

}

func traertitulos(c *gin.Context) {
	db, _ := sql.Open("sqlite3", "./dbs/"+ar_clase+".db")
	defer db.Close()
	rows, _ := db.Query("SELECT * FROM titulos")

	defer rows.Close()
	var titulos []Titulos
	for rows.Next() {
		var r Titulos
		rows.Scan(&r.Id, &r.Titulo)
		titulos = append(titulos, r)
	}
	c.JSON(http.StatusOK, titulos)
}

func traeritems(c *gin.Context) {
	item := c.Param("item")
	db, _ := sql.Open("sqlite3", "./dbs/"+ar_clase+".db")
	defer db.Close()
	rows, _ := db.Query("SELECT id_item, id_titulo, es_text, en_text  FROM items WHERE id_titulo = '" + item + "'")
	defer rows.Close()
	var items []Items
	for rows.Next() {
		var r Items
		rows.Scan(&r.Id_item, &r.Id_titulo, &r.Es_text, &r.En_text)
		items = append(items, r)
	}
	c.JSON(http.StatusOK, items)
}

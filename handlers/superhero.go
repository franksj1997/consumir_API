package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"consumir_Api/superhero"
)

func InitRoutes() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", HandlerInicio)
	r.GET("/superhero", HandlerSuperHero)
	r.Run(":8080").Error()
}

func HandlerSuperHero(c *gin.Context) {
	id := c.Query("input_search")
	log.Println("ID Recibido: ", id)
	datos := superhero.MainSuperHero(id)
	if datos.Id == "" {
		log.Println(datos)
		c.HTML(http.StatusBadRequest, "error", nil)
		return
	}
	c.HTML(http.StatusOK, "superhero", datos)
}

func HandlerInicio(c *gin.Context) {
	//c.HTML(http.StatusOK, "inicio.html", nil) este método corre solo cuando es html puro
	c.HTML(http.StatusOK, "inicio", nil) //este método corre cuando son templates
}

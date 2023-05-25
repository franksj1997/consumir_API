package handlers

import (
	"consumir_Api/superhero"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitRoutes() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.DateTime})
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/", HandlerInicio)
	r.GET("/superhero", HandlerSuperHero)
	r.Run(":8080")
}

func HandlerSuperHero(c *gin.Context) {
	id := c.Query("input_busqueda")
	log.Info().Msg("ID Recibido: " + id)
	datos := superhero.MainSuperHero(id)
	if datos.Id == "" {
		log.Warn().Msgf("No se obtuvieron datos: ", datos)
		c.HTML(http.StatusBadRequest, "error", nil)
		return
	}
	c.HTML(http.StatusOK, "superhero", datos)
}

func HandlerInicio(c *gin.Context) {
	//c.HTML(http.StatusOK, "inicio.html", nil) este método corre solo cuando es html puro
	c.HTML(http.StatusOK, "inicio", nil) //este método corre cuando son templates
}

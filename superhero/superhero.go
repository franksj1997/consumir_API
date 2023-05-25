package superhero

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"

	"consumir_Api/config"
	"consumir_Api/superhero/modelo"
)

func MainSuperHero(id string) modelo.SuperHeroModel {
	token := config.Token
	log.Info().Msgf("token: ", token)
	data, err := http.Get("https://www.superheroapi.com/api.php/" + token + "/" + id)
	if err != nil {
		log.Error().Msgf("", err)
	}
	defer data.Body.Close()

	dataResponse, alert := ioutil.ReadAll(data.Body)
	if err != nil {
		log.Error().Msgf("", alert)
	}

	var datos modelo.SuperHeroModel
	err = json.Unmarshal(dataResponse, &datos)
	if err != nil {
		log.Error().Msgf("", err)
	}
	return datos
	//fmt.Println(string(dataResponse))
}

package superhero

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"consumir_Api/superhero/modelo"
)

func MainSuperHero(id string) modelo.SuperHeroModel {
	token := "9315978835142264"
	data, err := http.Get("https://www.superheroapi.com/api.php/" + token + "/" + id)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Body.Close()

	dataResponse, alert := ioutil.ReadAll(data.Body)
	if err != nil {
		log.Fatal(alert)
	}

	var datos modelo.SuperHeroModel
	err = json.Unmarshal(dataResponse, &datos)
	if err != nil {
		log.Fatal(err)
	}
	return datos
	//fmt.Println(string(dataResponse))
}

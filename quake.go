package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Response struct {
	IDArea                 string  `json:"idArea"`
	Country                string  `json:"country"`
	LastSismicActivityDate string  `json:"lastSismicActivityDate"`
	UpdateDate             string  `json:"updateDate"`
	Owner                  string  `json:"owner"`
	Data                   []Quake `json:"data"`
}

type Quake struct {
	Googlemapref string `json:"googlemapref"`
	Degree       string `json: "degree"`
	DataUpdate   string `json: "dataUpdate"`
	MagType      string `json:"magType"`
	ObsRegion    string `json: "obsRegion"`
	Lon          string `json:"lon"`
	Source       string `json: "source"`
	Depth        int    `json:"depth"`
	TensorRef    string `json:"tensorRef"`
	Sensed       string `json:"sensed"`
	Shakemapid   string `json:"shakemapid"`
	Time         string `json:"time"`
	Lat          string `json:"lat"`
	Shakemapref  string `json:"shakemapref"`
	Local        string `json:"local"`
	Magnitud     string `json: "magnitud"`
}

const BaseURL = "http://api.ipma.pt/open-data/observation/seismic/"
const ArqAcores = 3
const ContinenteMadeira = 7

func launchInitialRequest(location int) Response {
	url := BaseURL + strconv.Itoa(location) + ".json"
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	return responseObject
}

/* should retrieve by magnitude, can receive a number returning all that are bigger or equal
or a MAX retrieving the maximum magnitude registered
*/
func retrieveByMagnitude(location int, magnitude string) {
	quakes := launchInitialRequest(location).Data
	if strings.ToUpper(magnitude) == "MAX" {
		var newQuake []Quake
		newQuake = append(newQuake, quakes[len(quakes)-1])
		print(newQuake)
	}
}

func print(quakes []Quake) {
	fmt.Print("\n     MAGNITUDE      |      DATE      |      TIME      |      LOCATION\n")
	fmt.Print("-------------------------------------------------------------------------")
	fmt.Print("\n")

	for _, quake := range quakes {
		time := strings.Split(quake.Time, "T")
		fmt.Printf("      %s               %s       %s        %s", quake.Magnitud, time[0], time[1], quake.ObsRegion)
	}

	fmt.Print("\n\n")
}

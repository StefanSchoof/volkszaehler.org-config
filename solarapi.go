package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"os"
)

type Response struct {
	Name string
	Body Body
	Head Head
}

type Body struct {
	Data Data
}

type Data struct {
	Site Site
}

type Site struct {
	P_PV    int `json:"P_PV"`
	E_Total int `json:"E_Total"`
}

type Head struct {
	Timestamp time.Time
}

func main() {
	response, err := http.Get("http://pv.fritz.box/solar_api/v1/GetPowerFlowRealtimeData.fcgi")

	if err != nil {
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		os.Exit(1)
	}
	var result Response
	json.Unmarshal(responseData, &result)
	var timestamp = result.Head.Timestamp.Unix()
	fmt.Printf("%v: Power = %v\n", timestamp, result.Body.Data.Site.P_PV)
	fmt.Printf("%v: Export = %v\n", timestamp, result.Body.Data.Site.E_Total)
}

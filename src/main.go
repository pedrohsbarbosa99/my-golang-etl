package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"my-golang-etl/src/models"
	"net/http"
)

func get_holidays(country, year string) string {
	url := fmt.Sprintf("https://date.nager.at/api/v1/Get/%s/%s", country, year)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	return sb
}

func tranform_holidays(holidays string) []models.Holiday {
	var holiday []models.Holiday
	err := json.Unmarshal([]byte(holidays), &holiday)
	if err != nil {
		fmt.Println(err)
	}
	return holiday
}

func main() {
	results := get_holidays("BR", "2022")
	holidays := tranform_holidays(results)
	for _, field := range holidays {
		fmt.Println(field.LocalName, field.Date)
	}
}

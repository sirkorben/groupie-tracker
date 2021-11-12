package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Artists []struct {
	Id             int                 `json:"id"`
	Image          string              `json:"image"`
	Name           string              `json:"name"`
	Members        []string            `json:"members"`
	CreationDate   int                 `json:"creationdate"`
	FirstAlbum     string              `json:"firstalbum"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Relation struct {
	Index []struct {
		Id             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

var data Artists
var data2 Relation

func getArtistsData() {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err2 := json.Unmarshal(body, &data)
	if err2 != nil {
		fmt.Println("--> smth wrong in getArtistData func")
		panic(err)

	}
}

func getRelationData() {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err2 := json.Unmarshal(body, &data2)
	if err2 != nil {
		fmt.Println("--> smth wrong in getRelationData func")
		panic(err)
	}

	// add .DatesLocations from Relation struct to Artist struct
	for index, value := range data2.Index {
		data[index].DatesLocations = value.DatesLocations
	}
}

func formatData() {
	for _, element := range data {
		preFormat := element.DatesLocations
		for key, value := range preFormat {
			newKey := strings.ReplaceAll(key, "-", ", ")
			newKey = strings.ReplaceAll(newKey, "_", " ")
			newKey = strings.Title(newKey)

			var newValue []string
			for _, val := range value {
				newVal := strings.ReplaceAll(val, "-", ".")
				newValue = append(newValue, newVal)
			}
			delete(element.DatesLocations, key)
			element.DatesLocations[newKey] = newValue
		}
	}
	for i, element := range data {
		newval := strings.ReplaceAll(element.FirstAlbum, "-", ".")
		data[i].FirstAlbum = newval
	}
}

func giveData() {
	getArtistsData()
	getRelationData()
	formatData()
}

package services

import (
	"bytes"
	"encoding/json"
	"github.com/abnergarcia1/SalesloftEngineeringTest/pkg/salesloft/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)
type PeopleService struct{}


func(s *PeopleService) GetPeopleFromAPI()(peopleList []models.People, err error){
	request:=models.SalesloftRequest{}
	url := "https://api.salesloft.com/v2/people.json"

	token:=os.Getenv("SALESLOFT_TOKEN")
	// Create a Bearer string by appending string access token
	bearer := "Bearer " + token

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
		return
	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)

		err = json.Unmarshal(data, &request)

		peopleList=request.Data

	}
	return
}




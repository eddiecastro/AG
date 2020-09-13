package server


import (
	"encoding/json"
	"fmt"
	"github.com/abnergarcia1/SalesloftEngineeringTest/pkg/salesloft/models"
	"net/http"
	//"strconv"

	//"github.com/abnergarcia1/SalesloftEngineeringTest/pkg/salesloft/models"
	"github.com/abnergarcia1/SalesloftEngineeringTest/pkg/salesloft/services"
	"github.com/gorilla/mux"
)

type api struct {
	router          http.Handler
	peopleService   *services.PeopleService
}

type Server interface {
	Router() http.Handler
}

func New() Server {
	a := &api{}

	r := mux.NewRouter()

	fmt.Println("Starting REST endpoints...")
	r.HandleFunc("/people", a.GetPeople).Methods(http.MethodGet)
	r.HandleFunc("/people/wordscounter", a.WordsCounter).Methods(http.MethodGet)
	r.HandleFunc("/people/deduping", a.DedupingEmails).Methods(http.MethodGet)

	fmt.Println("Running REST endpoints!")
	r.PathPrefix("/webclient/").Handler(http.StripPrefix("/webclient/",
		http.FileServer(http.Dir("../../pkg/salesloft/server/static"))))
	a.router = r

	a.peopleService=&services.PeopleService{}

	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) GetPeople(w http.ResponseWriter, r *http.Request){

	people, err := a.peopleService.GetPeopleFromAPI()
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
		json.NewEncoder(w).Encode(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(people)

}

func (a *api) WordsCounter(w http.ResponseWriter, r *http.Request){
	var wordList []models.EmailCharCounter

	peopleList, err := a.peopleService.GetPeopleFromAPI()
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
		json.NewEncoder(w).Encode(err)
		return
	}

	charcounter := make(map[string]int)
	for _, people:=range peopleList{
		for _,char:=range people.EmailAddress{
			_, ok := charcounter[string(char)]
			if !ok{
				charcounter[string(char)]=1
			}else{
				charcounter[string(char)]++
			}
		}
	}

	for word, count:=range charcounter{
		wordList = append(wordList, models.EmailCharCounter{Counter:count,Word:word})
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(wordList)
}

func (a *api) DedupingEmails(w http.ResponseWriter, r *http.Request){}

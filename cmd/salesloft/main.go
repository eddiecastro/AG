package main

import (
	"github.com/abnergarcia1/SalesloftEngineeringTest/pkg/salesloft/server"
	"log"
	"net/http"
	"github.com/gorilla/handlers"
	"os"
)

func main(){
	s := server.New()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8000",handlers.CORS(originsOk, headersOk, methodsOk)(s.Router())))
}

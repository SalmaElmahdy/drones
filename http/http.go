package http

import (
	"log"
	"net/http"

	"github.com/SalmaElmahdy/drones/http/apis"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type APIs struct {
	DroneAPIs apis.DroneAPIs
}

func StartServer(api APIs) {
	r := mux.NewRouter()

	droneSubRouter := r.PathPrefix("/drone").Subrouter()
	droneSubRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		api.DroneAPIs.Create(w, r)

	}).Methods("POST")
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8010", r))
}

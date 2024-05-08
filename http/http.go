//	@title			Drones API
//	@version		1.16.3
//	@description	This is a sample serice for managing Drones
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	API Support
//	@contact.email	soberkoder@swagger.io
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			localhost:8010
//	@BasePath		/

package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SalmaElmahdy/drones/http/apis"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type APIs struct {
	DroneAPIs      apis.DroneAPIs
	MedicationAPIs apis.MedicationAPIs
}

func StartServer(api APIs) {
	r := mux.NewRouter()

	droneSubRouter := r.PathPrefix("/drone").Subrouter()
	droneSubRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		api.DroneAPIs.Create(w, r)

	}).Methods("POST")

	droneSubRouter.HandleFunc("/{id}/medications", func(w http.ResponseWriter, r *http.Request) {
		api.DroneAPIs.GetLoadedMedications(w, r)

	}).Methods("GET")

	droneSubRouter.HandleFunc("/load", func(w http.ResponseWriter, r *http.Request) {
		api.DroneAPIs.LoadMedications(w, r)

	}).Methods("POST")

	medicationSubRouter := r.PathPrefix("/medication").Subrouter()
	medicationSubRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		api.MedicationAPIs.Create(w, r)

	}).Methods("POST")
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	fmt.Println("server running on port:8010")
	log.Fatal(http.ListenAndServe(":8010", r))
}

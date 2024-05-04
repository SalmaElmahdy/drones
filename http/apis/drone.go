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

package apis

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/SalmaElmahdy/drones/docs"
	"github.com/SalmaElmahdy/drones/usecase"
	"github.com/gorilla/mux"
)

type DroneAPIs struct {
	droneUseCase usecase.IDroneUseCase
}

func NewDroneAPIs(droneUseCase usecase.IDroneUseCase) DroneAPIs {
	return DroneAPIs{
		droneUseCase: droneUseCase,
	}
}

// @Summary		Create a new drone
// @Description	Create a new drone entity
// @Tags			Drone
// @Accept			json
// @Produce		json
// @Param			request	body		entity.DroneRequest	true	"Request of Creating Drone"
// @Success		200		{object}	entity.DroneRequest
// @Failure		400		{string}	string	"Bad Request"
// @Failure		500		{string}	string	"Internal Server Error"
// @Router			/drone [post]
func (api DroneAPIs) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("[Error]: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{ "error" : "%s"}`, err.Error())))
		return
	}

	response, err := api.droneUseCase.Create(ctx, requestByte)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{ "error" : "%s"}`, err.Error())))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (api DroneAPIs) GetLoadedMedications(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		err := errors.New("ID is required")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{ "error" : %q}`, err.Error())))
		return
	}
	ID, err := strconv.Atoi(id)
	if err != nil {
		err := errors.New("invalid ID")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{ "error" : %q}`, err.Error())))
		return
	}
	response, err := api.droneUseCase.GetLoadedMedications(ctx, uint(ID))
	if err != nil {
		log.Printf("[Error]: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{ "error" : "%s"}`, err.Error())))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

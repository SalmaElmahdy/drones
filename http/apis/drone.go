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
// @Router			/drone/ [post]
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

// @Summary		Get Loaded Medications
// @Description	checking loaded medication items for a given drone
// @Tags			Drone
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Drone ID"
// @Success		200	{object}	[]entity.MedicationRequest
// @Failure		400	{string}	string	"Bad Request"
// @Failure		500	{string}	string	"Internal Server Error"
// @Failure		404	{string}	string	"Not Found"
// @Router			/drone/{id}/medications [get]
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

// @Summary		Load Medications
// @Description	Loading a drone with medication items
// @Tags			Drone
// @Accept			json
// @Produce		json
// @Param			request	body		entity.LoadMedicationsRequest	true	"Request of load medications"
// @Success		200		{object}	entity.DroneRequest
// @Failure		400		{string}	string	"Bad Request"
// @Failure		500		{string}	string	"Internal Server Error"
// @Router			/drone/load [post]
func (api DroneAPIs) LoadMedications(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("[Error]: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{ "error" : "%s"}`, err.Error())))
		return
	}

	response, err := api.droneUseCase.LoadMedications(ctx, requestByte)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{ "error" : "%s"}`, err.Error())))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

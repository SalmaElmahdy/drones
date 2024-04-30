package apis

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/SalmaElmahdy/drones/usecase"
)

type MedicationAPIs struct {
	medicationUseCase usecase.IMedicationUseCase
}

func NewMedicationAPIs(medicationUseCase usecase.IMedicationUseCase) MedicationAPIs {
	return MedicationAPIs{
		medicationUseCase: medicationUseCase,
	}
}

func (api MedicationAPIs) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("[Error]: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{ "error" : "%s"}`, err.Error())))
		return
	}

	response, err := api.medicationUseCase.Create(ctx, requestByte)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{ "error" : "%s"}`, err.Error())))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

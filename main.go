package main

import (
	_ "github.com/SalmaElmahdy/drones/docs"
	serverHTTP "github.com/SalmaElmahdy/drones/http"
	"github.com/SalmaElmahdy/drones/http/apis"
	"github.com/SalmaElmahdy/drones/repository"
	"github.com/SalmaElmahdy/drones/repository/db"
	"github.com/SalmaElmahdy/drones/usecase"
)

func main() {

	db := db.ConnectToDatabse()
	droneRepo := repository.NewDroneRepository(db)
	droneUseCase := usecase.NewDroneUseCase(droneRepo)
	droneAPIs := apis.NewDroneAPIs(droneUseCase)

	apis := serverHTTP.APIs{
		DroneAPIs: droneAPIs,
	}

	serverHTTP.StartServer(apis)
}

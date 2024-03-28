package main

import (
	"flag"
	"log"

	_ "github.com/SalmaElmahdy/drones/docs"
	serverHTTP "github.com/SalmaElmahdy/drones/http"
	"github.com/SalmaElmahdy/drones/http/apis"
	"github.com/SalmaElmahdy/drones/repository"
	"github.com/SalmaElmahdy/drones/repository/db"
	"github.com/SalmaElmahdy/drones/seeds"
	"github.com/SalmaElmahdy/drones/usecase"
	gorm_seeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
)

var seedFlag = flag.Bool("seeds", false, "Seed the database")
var clearSeedFlag = flag.Bool("clearseeds", false, "Clear seeds from the database")

func main() {
	db := db.ConnectToDatabse()
	flag.Parse()
	if *seedFlag {
		seed(db)
		return
	}
	if *clearSeedFlag {
		clearSeed(db)
		return
	}

	droneRepo := repository.NewDroneRepository(db)
	droneUseCase := usecase.NewDroneUseCase(droneRepo)
	droneAPIs := apis.NewDroneAPIs(droneUseCase)

	apis := serverHTTP.APIs{
		DroneAPIs: droneAPIs,
	}

	serverHTTP.StartServer(apis)
}

func seed(db *gorm.DB) {
	droneSeeder := seeds.NewDroneSeeder(gorm_seeder.SeederConfiguration{Rows: 5})
	seedersStack := gorm_seeder.NewSeedersStack(db)
	seedersStack.AddSeeder(&droneSeeder)
	if err := seedersStack.Seed(); err != nil {
		log.Fatalf("Error seeding database: %v", err)
	}

	log.Println("Seeds inserted successfully")
}

func clearSeed(db *gorm.DB) {
	droneSeeder := seeds.NewDroneSeeder(gorm_seeder.SeederConfiguration{})
	if err := droneSeeder.Clear(db); err != nil {
		log.Fatalf("Error clearing seeded data: %v", err)
	}
}

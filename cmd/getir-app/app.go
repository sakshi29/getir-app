package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sakshi29/getir-app/common/config"
	"github.com/sakshi29/getir-app/internal/delivery"
	"github.com/sakshi29/getir-app/internal/model/lib"
	mongoDB "github.com/sakshi29/getir-app/internal/repository/mongo-db"
	v1UseCase "github.com/sakshi29/getir-app/internal/usecase/v1"
)

func main() {

	cfg, err := config.Init("development")
	if err != nil {
		log.Fatalf("failed to set config, err: %+v", err)
	}

	//initialse mongo DB
	mongoDBRepo, err := mongoDB.Init(cfg.MongoDB)
	if err != nil {
		log.Fatalf("failed to initialise mongo DB, err: %+v", err)
	}

	useCaseV1 := v1UseCase.InitUseCaseV1(cfg, mongoDBRepo)

	lib.InitErrorMap()

	//Initialise HTTP Handler
	delivery.Init(cfg, useCaseV1)

	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.Server.Port
	}

	log.Printf("app started on port %s", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

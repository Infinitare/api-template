package main

import (
	"api-template/db"
	"api-template/handler"
	"github.com/Infinitare/types-template/console"
	"github.com/Infinitare/types-template/metrics"
	"github.com/Infinitare/types-template/requests"
	"github.com/Infinitare/types-template/secrets"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	const service = "" // INSERT service name here
	requests.SetDomain()
	console.Setup(service)

	console.Info("Initialize", "Url", service)
	console.Info("Initialize", "Secrets")
	secrets.Load()

	console.Info("Initialize", "Datadog")
	tracerStop, profilerStop := metrics.Setup(service)

	defer tracerStop()
	defer profilerStop()

	console.Info("Initialize", "Database")
	dbClose := db.Setup()
	defer dbClose()

	port := os.Getenv("PORT")
	console.Info("Initialize", "Port", port)

	router := handler.Add()
	console.Info("Initialize", "Finished", "Listen and Serve")

	console.Error(http.ListenAndServe(":"+port, router))

}

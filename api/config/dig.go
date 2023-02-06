package config

import (
	"ohlc-data-api/api/repository"
	"ohlc-data-api/api/service"

	"go.uber.org/dig"
)

var container = dig.New()

func BuildProject() *dig.Container {

	// configuration
	container.Provide(InitializeDB)

	//Service
	container.Provide(service.NewDataService)

	//Repository
	container.Provide(repository.NewDataRepository)

	return container

}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}

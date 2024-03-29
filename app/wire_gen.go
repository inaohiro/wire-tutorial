// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/inaohiro-sandbox/wire-tutorial3/app/config"
	"github.com/inaohiro-sandbox/wire-tutorial3/app/domain/application"
	"github.com/inaohiro-sandbox/wire-tutorial3/app/infra/database"
	"github.com/inaohiro-sandbox/wire-tutorial3/app/infra/database/mysql"
	"github.com/inaohiro-sandbox/wire-tutorial3/app/infra/inmemorydb"
)

// Injectors from wire.go:

func initializeService() (application.Service, func(), error) {
	mysqlConfig := config.MysqlConfig()
	dataSource := mysql.NewDataSource(mysqlConfig)
	db, cleanup, err := database.Open(dataSource)
	if err != nil {
		return nil, nil, err
	}
	userRepository := database.NewUserRepository(db)
	service := application.NewService(userRepository)
	return service, func() {
		cleanup()
	}, nil
}

func initializeMockService() application.Service {
	inMemoryUserRepository := inmemorydb.NewUserRepository()
	service := application.NewService(inMemoryUserRepository)
	return service
}

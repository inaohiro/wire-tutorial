//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/inaohiro-sandbox/wire-tutorial3/app/config"
	"github.com/inaohiro-sandbox/wire-tutorial3/app/domain/application"
	"github.com/inaohiro-sandbox/wire-tutorial3/app/domain/domain"
	"github.com/inaohiro-sandbox/wire-tutorial3/app/infra/database"
	"github.com/inaohiro-sandbox/wire-tutorial3/app/infra/database/mysql"
	"github.com/inaohiro-sandbox/wire-tutorial3/app/infra/inmemorydb"
)

func initializeService() (application.Service, func(), error) {
	wire.Build(
		application.Set,
		mysql.Set,
		config.Set,
		database.RepositorySet,
	)
	return nil, nil, nil
}

func initializeMockService() application.Service {
	wire.Build(
		application.Set,
		inmemorydb.NewUserRepository,
		wire.Bind(new(domain.UserRepository), new(*inmemorydb.InMemoryUserRepository)),
	)
	return nil
}

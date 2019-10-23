package mysql

import (
	"github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/inaohiro-sandbox/wire-tutorial3/app/infra/database"
)

var Set = wire.NewSet(
	wire.NewSet(database.Open),
	NewDataSource,
)

func NewDataSource(cfg *mysql.Config) database.DataSource {
	return database.DataSource{
		DriverName:     "mysql",
		DataSourceName: cfg.FormatDSN(),
	}
}

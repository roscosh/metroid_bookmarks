package misc

import (
	"github.com/amacneil/dbmate/pkg/dbmate"
	"github.com/amacneil/dbmate/pkg/driver/postgres"
	"net/url"
)

func DbMigrate(dbUrl string, migrationsDir string) error {
	URL, _ := url.Parse(dbUrl)
	dbmate.RegisterDriver(postgres.NewDriver, "postgres")
	db := dbmate.New(URL)
	db.MigrationsDir = migrationsDir
	return db.Migrate()
}

package misc

import (
	"net/url"
	"os"
)
import "github.com/amacneil/dbmate/pkg/dbmate"
import "github.com/amacneil/dbmate/pkg/driver/postgres"

func DbMigrate() error {
	dbURL, _ := url.Parse(config.Db.Dsn)
	dbmate.RegisterDriver(postgres.NewDriver, "postgres")
	db := dbmate.New(dbURL)
	db.MigrationsDir = os.Getenv("DBMATE_MIGRATIONS_DIR")
	return db.Migrate()
}

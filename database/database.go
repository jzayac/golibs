package database

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Info contains the database configurations
type Info struct {
	TypeDb string
	Path   string
}

type databaseService struct {
	db        Info
	isDevelop bool
}

func (ds databaseService) Connect() (*gorm.DB, error) {
	dbinst, err := gorm.Open(ds.db.TypeDb, ds.db.Path)

	if err != nil {
		return nil, ErrSqlConnectionProblem
	}

	if ds.isDevelop {
		return dbinst.Debug(), err
	}

	return dbinst, err
}

func NewDatabaseService(d Info, dbname string, develop bool) (*databaseService, error) {
	d.Path = strings.Replace(d.Path, "::dbname::", dbname, 1)
	con, err := gorm.Open(d.TypeDb, d.Path)
	if err != nil {
		return nil, ErrSqlInitDbDriver
	}
	con.Close()

	return &databaseService{
		db:        d,
		isDevelop: develop,
	}, nil
}

var ErrSqlConnectionProblem = errors.New("Connect| sql connection problem")
var ErrSqlInitDbDriver = errors.New("Initialize| PG Drive Error")

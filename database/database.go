package database

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	// "videolib/logger"
)

var (
	db Info
	// l  = logger.SetLogger("database")
)

// Info contains the database configurations
type Info struct {
	TypeDb string
	Path   string
}

func Connect() (*gorm.DB, error) {
	con, err := gorm.Open(db.TypeDb, db.Path)
	if err != nil {
		// logger.Error.Println("Connect| sql connection problem", err)
		return nil, ErrSqlConnectionProblem
	}
	return con, err
}

// Connect to the database
func Initialize(d Info) error {

	// Connect to sqlite
	con, err := gorm.Open(d.TypeDb, d.Path)
	if err != nil {
		// logger.Error.Println("Initialize| Sqlite Drive Error", err)
		// return err
		return ErrSqlInitDbDriver
	}
	con.Close()

	// l.Debug("db connection established", "Initialize")
	db = d

	return nil
}

// ReadConfig returns the database information
func ReadConfig() Info {
	return db
}

var ErrSqlConnectionProblem = errors.New("Connect| sql connection problem")
var ErrSqlInitDbDriver = errors.New("Initialize| Sqlite Drive Error")

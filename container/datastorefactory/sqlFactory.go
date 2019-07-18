package datastorefactory

import (
	"database/sql"
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/jfeng45/servicetmpl/tools/gdbc"
	"github.com/pkg/errors"
)

// sqlFactory is receiver for Build method
type sqlFactory struct {}

// implement Build method for SQL database
func (sf *sqlFactory) Build(c container.Container, dsc *configs.DataStoreConfig) (DataStoreInterface, error) {
	logger.Log.Debug("sqlFactory")
	key := dsc.Code

	if SQLDB != key {
		errMsg := SQLDB + " in sqlFactory doesn't match key = " + key
		return nil, errors.New(errMsg)
	}
	//if it is already in container, return
	if value, found := c.Get(key); found {
		//logger.Log.Debugf("found db value %+v\n:", value)
		sdb :=value.(*sql.DB)
		sdt := gdbc.SqlDBTx{DB: sdb}
		logger.Log.Debug("found db in container for key:", key)
		return &sdt, nil
	}

	db, err := sql.Open(dsc.DriverName, dsc.UrlAddress)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	// check the connection
	err =db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	dt := gdbc.SqlDBTx{DB: db}
	//c.Put(key, &dt)
	c.Put(key, db)
	return &dt, nil

}


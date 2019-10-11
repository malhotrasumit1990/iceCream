package dataAccess

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/iceCream/src/dataSource"
)

var db *dataSource.MySQLDataSource
var DBConnectionString string

var LoadDBConfig = func(filePath string) error {

	file, _ := ioutil.ReadFile(filePath)
	dbs := dataSource.DBS{}
	err := json.Unmarshal([]byte(file), &dbs)
	if err != nil {
		log.Print(err)
		return err
	}
	conString := makeDBConnectionString(dbs)
	db, err = setupDBInstance(conString)
	if err != nil {
		return err
	}
	return nil

}

func makeDBConnectionString(dbs dataSource.DBS) string {

	userName := dbs.DBs[0].UserID + ":"
	pass := dbs.DBs[0].Password + "@"
	server := "tcp(" + dbs.DBs[0].Server + ")/"
	database := dbs.DBs[0].Database

	connectionString := userName + pass + server + database
	fmt.Println(connectionString)
	return connectionString
}

var setupDBInstance = func(connectionString string) (*dataSource.MySQLDataSource, error) {
	database, err := dataSource.GetDB(connectionString)

	if err != nil {
		return nil, err
	}
	return database, nil
}

package dataSource

type DBS struct {
	DBs []DB `json:"DBs"`
}
type DB struct {
	UserID   string `json:"USERID"`
	Password string `json:"Password"`
	Server   string `json:"Server"`
	Database string `json:"Database"`
}

func GetDB(connectionString string) (*MySQLDataSource, error) {
	return NewMySqlDataStore(connectionString)

}

package connection

type DataBase interface {
	Connect(dbUrl string) error
	InitializeTables(models ...interface{}) error
}

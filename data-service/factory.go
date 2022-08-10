package data_service

import (
	"poker-on-dices/config"
	db "poker-on-dices/data-service/database"
	"poker-on-dices/data-service/database/connection"
)

func NewDataService(config config.Config) (*db.DBDataService, error) {
	databaseURL := config.GetDataProviderURL()
	dataBaseConnection, err := connection.SetupDatabase(databaseURL)
	if err != nil {
		return nil, err
	}
	return &db.DBDataService{
		DataBase: dataBaseConnection,
	}, nil
}

package connection

import (
	"sync"
	"poker-on-dices/data-service/database/models"
)

var instance DataBase
var dbOnce sync.Once

func GetInstance() DataBase {
	dbOnce.Do(func() {
		instance = &PgDataBase{}
	})

	return instance
}

func SetupDatabase(dbUrl string) (DataBase, error) {
	var db = GetInstance()

	if err := db.Connect(dbUrl); err != nil {
		return nil, err
	}

	if err := db.InitializeTables(models.Player{}, models.Game{}, models.Round{}, models.Stake{}); err != nil {
		return nil, err
	}

	return db, nil
}

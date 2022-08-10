package database

import (
	"poker-on-dices/data-service/database/connection"
)

type DBDataService struct {
	DataBase connection.DataBase
}

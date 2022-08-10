package models

import "time"

type Game struct {
	ID        uint     `gorm:"primarykey"`
	Players   []Player `gorm:"many2many:players_per_games;foreignKey:ID;joinForeignKey:GameID;References:ID;joinReferences:PlayerID"`
	CreatedAt time.Time
	Rounds    []Round
	IsOver    bool
}

type Player struct {
	ID       uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Username string `gorm:"uniqueIndex"`
	TgID     string
}

type Round struct {
	ID           uint `gorm:"primary_key;AUTO_INCREMENT"`
	GameID       uint
	NumberOfDice uint
	Number       uint
	Stakes       []Stake `gorm:"preload:true"`
}

type Stake struct {
	ID       uint `gorm:"primary_key;AUTO_INCREMENT"`
	RoundID  uint
	PlayerID uint
	Bet      uint
	Bribe    *uint
}

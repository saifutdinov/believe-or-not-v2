package database

import (
	"saifutdinov/believe-or-not/backend/api/domain"

	"gorm.io/gorm"
)

func Migrate(dbc *gorm.DB) {

	dbc.AutoMigrate(
		&domain.Player{},
		&domain.Room{},
	)

}

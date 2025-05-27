package api

import (
	"saifutdinov/believe-or-not/backend/database"
	"saifutdinov/believe-or-not/backend/packages/dotenv"

	gamedelivery "saifutdinov/believe-or-not/backend/api/game/delivery"
	gamerepo "saifutdinov/believe-or-not/backend/api/game/repository"
	gameusecase "saifutdinov/believe-or-not/backend/api/game/usecase"

	"github.com/go-chi/chi"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initHandlers(chirouter *chi.Mux, config *dotenv.Env) error {

	database, err := initDatabase(config)
	if err != nil {
		return err
	}

	// game handlers
	gameRepo := gamerepo.NewGameRepository(database, config)
	gameUsecase := gameusecase.NewGameUsecase(gameRepo, config)
	gamedelivery.NewGameHandler(chirouter, gameUsecase, config)

	return nil
}

func initDatabase(config *dotenv.Env) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(config.SqliteConnection), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// confirm exists migrations
	database.Migrate(db)
	return db, nil
}

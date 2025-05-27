package repository

import (
	"saifutdinov/believe-or-not/backend/api/domain"
	"saifutdinov/believe-or-not/backend/packages/dotenv"

	"gorm.io/gorm"
)

type (
	GameRepository struct {
		*gorm.DB
		Config *dotenv.Env
	}
)

func NewGameRepository(dbc *gorm.DB, config *dotenv.Env) domain.GameRepository {
	return &GameRepository{
		DB:     dbc,
		Config: config,
	}
}

func (gr *GameRepository) CreateRoom(room *domain.Room) error {
	result := gr.Create(room)
	return result.Error
}

func (gr *GameRepository) ReadRoom(gc domain.GameCode) (*domain.Room, error) {
	room := new(domain.Room)
	result := gr.First(room, "game_code=?", gc)
	return room, result.Error
}

func (gr *GameRepository) UpdateRoom(room *domain.Room) error {
	result := gr.Save(room)
	return result.Error
}

func (gr *GameRepository) ReadRoomPlayers(gc domain.GameCode) ([]*domain.Player, error) {
	players := make([]*domain.Player, 0)
	result := gr.Where("game_code=?", gc).Find(&players)
	return players, result.Error
}

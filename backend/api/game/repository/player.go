package repository

import "saifutdinov/believe-or-not/backend/api/domain"

func (gr *GameRepository) CreatePlayer(player *domain.Player) error {
	result := gr.Create(player)
	return result.Error
}

func (gr *GameRepository) ReadPlayer(playerId uint) (*domain.Player, error) {
	player := new(domain.Player)
	result := gr.First(player, "id=?", playerId)
	return player, result.Error
}

func (gr *GameRepository) UpdatePlayer(player *domain.Player) error {
	result := gr.Save(player)
	return result.Error
}

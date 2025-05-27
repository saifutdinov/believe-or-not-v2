package domain

import "saifutdinov/believe-or-not/backend/packages/cards"

type (
	GameUsecase interface {
		PlayerReady(playerId uint) error
		//
		CreatePlayer(playerName string) (*Player, error)
		//
		CreateRoom(creatorId uint, gp GameParams) (*Room, error)
		//
		AuthInRoom(gc GameCode, password, name string) (*Player, error)
		//
		StartGame(gc GameCode) (map[uint][]*cards.Card, error)
	}

	GameRepository interface {
		//
		CreateRoom(room *Room) error
		//
		ReadRoom(gc GameCode) (*Room, error)
		//
		ReadRoomPlayers(gc GameCode) ([]*Player, error)
		//
		UpdateRoom(room *Room) error
		//
		CreatePlayer(player *Player) error
		//
		ReadPlayer(playerId uint) (*Player, error)
		//
		UpdatePlayer(player *Player) error
	}
)

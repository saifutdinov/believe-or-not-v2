package delivery

import (
	"saifutdinov/believe-or-not/backend/api/domain"
	"saifutdinov/believe-or-not/backend/packages/cards"
	"saifutdinov/believe-or-not/backend/packages/request"
)

type (
	CreatePlayerRequest struct {
		PlayerName string `json:"playerName"`
	}

	CreateRoomRequest struct {
		IsPrivate request.CBool `json:"isPrivate"`
		Password  string        `json:"password"`
	}

	AuthInRoomRequest struct {
		GameCode   domain.GameCode `json:"game_code"`
		PlayerName string          `json:"playerName"`
		Password   string          `json:"password,omitempty"`
	}

	StartGameRequest struct {
		GameCode domain.GameCode `json:"game_code"`
	}

	GameCardsStatus struct {
		PlayerCards map[int]*cards.Card
	}
)

package domain

import (
	"saifutdinov/believe-or-not/backend/packages/cards"

	"gorm.io/gorm"
)

const (
	StageDefault Stage = iota
	StagePreparing
	StagePaused
	StageOnline
	StageDone
)

type (
	Stage int
	// параметры для создания игры
	GameParams struct {
		IsPrivate bool
		Password  string
	}
	// игрок
	Player struct {
		gorm.Model
		Name     string        `gorm:"name;not null;"`
		GameCode GameCode      `gorm:"game_code"`
		Hand     []*cards.Card `gorm:"hand;serializer:json"`
		IsWinner bool          `gorm:"is_winner"`
		IsReady  bool          `gorm:"is_ready"`
	}
	// игровая комната
	Room struct {
		gorm.Model
		GameCode  GameCode `gorm:"game_code"`
		Players   []uint   `gorm:"players;serializer:json"`
		Order     []string `gorm:"order;serializer:json"`
		TurnIdx   int
		Stack     *cards.Stack `gorm:"stack;serializer:json"`
		Discard   []string     `gorm:"discard;serializer:json"`
		IsPrivate bool         `gorm:"is_private"`
		Password  string       `gorm:"password"`
		Stage     Stage
	}
)

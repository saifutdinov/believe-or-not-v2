package usecase

import (
	"errors"
	"saifutdinov/believe-or-not/backend/api/domain"
	"saifutdinov/believe-or-not/backend/packages/cards"
	"saifutdinov/believe-or-not/backend/packages/dotenv"
)

type (
	GameUsecase struct {
		GameRepository domain.GameRepository
		Config         *dotenv.Env
	}
)

func NewGameUsecase(
	GameRepository domain.GameRepository,
	Config *dotenv.Env,
) domain.GameUsecase {
	return &GameUsecase{
		GameRepository: GameRepository,
		Config:         Config,
	}
}

func (gu *GameUsecase) PlayerReady(playerId uint) error {
	player, err := gu.GameRepository.ReadPlayer(playerId)
	if err != nil {
		return err
	}
	player.IsReady = !player.IsReady

	return gu.GameRepository.UpdatePlayer(player)
}

func (gu *GameUsecase) CreatePlayer(playerName string) (*domain.Player, error) {
	player := &domain.Player{
		Name: playerName,
	}
	if err := gu.GameRepository.CreatePlayer(player); err != nil {
		return nil, err
	}
	return player, nil
}

func (gu *GameUsecase) CreateRoom(creatorId uint, gp domain.GameParams) (*domain.Room, error) {
	newGameCode := domain.GenerateGameCode()

	room := &domain.Room{
		GameCode:  newGameCode,
		IsPrivate: gp.IsPrivate,
		Password:  gp.Password,
		Players:   []uint{creatorId},
		Stage:     domain.StageDefault,
	}

	if err := gu.GameRepository.CreateRoom(room); err != nil {
		return nil, err
	}

	creator, err := gu.GameRepository.ReadPlayer(creatorId)
	if err != nil {
		return nil, err
	}

	creator.GameCode = newGameCode

	if err := gu.GameRepository.UpdatePlayer(creator); err != nil {
		return nil, err
	}

	return room, nil
}

func (gu *GameUsecase) AuthInRoom(gc domain.GameCode, password, playerName string) (*domain.Player, error) {
	room, err := gu.GameRepository.ReadRoom(gc)
	if err != nil {
		return nil, err
	}
	if room.IsPrivate && room.Password != password {
		return nil, errors.New("wrong password")
	}

	player := &domain.Player{
		Name:     playerName,
		GameCode: gc,
	}
	if err := gu.GameRepository.CreatePlayer(player); err != nil {
		return nil, err
	}

	room.Players = append(room.Players, player.ID)

	if err := gu.GameRepository.UpdateRoom(room); err != nil {
		return nil, err
	}

	return player, nil
}

func (gu *GameUsecase) StartGame(gc domain.GameCode) (map[uint][]*cards.Card, error) {
	room, err := gu.GameRepository.ReadRoom(gc)
	if err != nil {
		return nil, err
	}

	if len(room.Players) <= 1 {
		return nil, errors.New("not enough players in room")
	}

	players, err := gu.GameRepository.ReadRoomPlayers(gc)
	if err != nil {
		return nil, err
	}
	playerById := make(map[uint]*domain.Player)
	for _, p := range players {
		playerById[p.ID] = p
	}

	stack := cards.NewStack()

	playersHands := make(map[uint][]*cards.Card)

	playerCards := func(pcount int) int {
		switch len(room.Players) {
		case domain.Duo:
			return domain.DuoPlayersCards
		case domain.Trio:
			return domain.TrioPlayersCards
		default:
			return domain.MorePlayersCards
		}
	}

	cardsForPlay := playerCards(len(room.Players))

	for _, playerId := range room.Players {
		hand := stack.Deal(cardsForPlay)
		playersHands[playerId] = hand
		player := playerById[playerId]
		player.Hand = hand
		if err := gu.GameRepository.UpdatePlayer(player); err != nil {
			return nil, err
		}
	}

	room.Stack = stack

	if err := gu.GameRepository.UpdateRoom(room); err != nil {
		return nil, err
	}

	return playersHands, nil
}

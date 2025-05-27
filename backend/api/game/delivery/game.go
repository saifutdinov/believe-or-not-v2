package delivery

import (
	"net/http"
	"saifutdinov/believe-or-not/backend/api/domain"
	"saifutdinov/believe-or-not/backend/middlewares"
	"saifutdinov/believe-or-not/backend/packages/request"
	"saifutdinov/believe-or-not/backend/packages/response"
)

func (gh *GameHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	userID, _ := middlewares.UserIDFromContext(r.Context())

	if userID == 0 {
		response.Json(w, response.Error{Message: "user id not found."})
		return
	}

	createRequestForm := new(CreateRoomRequest)
	if err := request.Bind(r, createRequestForm); err != nil {
		response.ErrJson(w, err)
		return
	}

	room, err := gh.GameUsecase.CreateRoom(uint(userID), domain.GameParams{
		IsPrivate: bool(createRequestForm.IsPrivate),
		Password:  createRequestForm.Password,
	})
	if err != nil {
		response.ErrJson(w, err)
		return
	}

	response.Json(w, room)
	return
}

func (gh *GameHandler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	createRequestForm := new(CreatePlayerRequest)
	if err := request.Bind(r, createRequestForm); err != nil {
		response.ErrJson(w, err)
		return
	}

	if createRequestForm.PlayerName == "" {
		response.Set(w, response.Error{Status: 10001, Message: "playerName required!"})
		return
	}

	player, err := gh.GameUsecase.CreatePlayer(createRequestForm.PlayerName)
	if err != nil {
		response.ErrJson(w, err)
		return
	}

	token, err := middlewares.GenerateJWT(player.ID)

	response.Json(w, map[string]any{"token": token})
	return
}

func (gh *GameHandler) AuthInRoom(w http.ResponseWriter, r *http.Request) {
	userID, _ := middlewares.UserIDFromContext(r.Context())

	if userID == 0 {
		response.Json(w, response.Error{Message: "user id not found."})
		return
	}

	authRequest := new(AuthInRoomRequest)
	if err := request.Bind(r, authRequest); err != nil {
		response.ErrJson(w, err)
		return
	}

	player, err := gh.GameUsecase.AuthInRoom(authRequest.GameCode, authRequest.Password, authRequest.PlayerName)
	if err != nil {
		response.ErrJson(w, err)
		return
	}

	token, err := middlewares.GenerateJWT(player.ID)

	response.Json(w, map[string]any{"token": token})
	return
}

func (gh *GameHandler) PlayerReady(w http.ResponseWriter, r *http.Request) {
	userID, _ := middlewares.UserIDFromContext(r.Context())

	if userID == 0 {
		response.Json(w, response.Error{Message: "user id not found."})
		return
	}

	if err := gh.GameUsecase.PlayerReady(uint(userID)); err != nil {
		response.ErrJson(w, err)
		return
	}

	response.Json(w)
	return
}

func (gh *GameHandler) StartGame(w http.ResponseWriter, r *http.Request) {
	userID, _ := middlewares.UserIDFromContext(r.Context())

	if userID == 0 {
		response.Json(w, response.Error{Message: "user id not found."})
		return
	}

	startGameRequest := new(StartGameRequest)
	if err := request.Bind(r, startGameRequest); err != nil {
		response.ErrJson(w, err)
		return
	}

	cards, err := gh.GameUsecase.StartGame(startGameRequest.GameCode)
	if err != nil {
		response.ErrJson(w, err)
		return
	}

	response.Json(w, cards)
	return
}

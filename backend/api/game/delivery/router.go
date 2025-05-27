package delivery

import (
	"saifutdinov/believe-or-not/backend/api/domain"
	"saifutdinov/believe-or-not/backend/middlewares"
	"saifutdinov/believe-or-not/backend/packages/dotenv"

	"github.com/go-chi/chi"
)

type GameHandler struct {
	GameUsecase domain.GameUsecase
	Config      *dotenv.Env
}

func NewGameHandler(
	chirouter *chi.Mux,
	gameUsecase domain.GameUsecase,
	config *dotenv.Env,
) {
	handler := &GameHandler{
		GameUsecase: gameUsecase,
		Config:      config,
	}

	// available before auth (without JWT in header)
	chirouter.Put("/api/create-player", handler.CreatePlayer)
	chirouter.Post("/api/auth/room", handler.AuthInRoom)

	chirouter.Use(middlewares.JWTMiddleware)
	// available only auth (with JWT ih header)
	chirouter.Put("/api/create-room", handler.CreateRoom)
	chirouter.Post("/api/game/ready", handler.PlayerReady)
	chirouter.Post("/api/game/start", handler.StartGame)

}

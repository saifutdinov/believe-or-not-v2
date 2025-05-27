package api

import (
	"fmt"
	"log"
	"net/http"
	"saifutdinov/believe-or-not/backend/packages/dotenv"

	"github.com/go-chi/chi"
	extmiddw "github.com/go-chi/chi/middleware"
)

func StartListen(config *dotenv.Env) {

	// redis.InitRedis()

	// router object
	chirouter := chi.NewRouter()
	// logs
	chirouter.Use(extmiddw.Logger)
	// recovery
	chirouter.Use(extmiddw.Recoverer)
	// start
	if err := initHandlers(chirouter, config); err != nil {
		panic(err)
	}

	//
	host := fmt.Sprintf(":%s", config.BackendPort)

	log.Printf("Server started on: http://localhost%s\n", host)
	http.ListenAndServe(host, chirouter)
}

package chi_server

import (
	"archetype/app/shared/archetype/container"
	"archetype/app/shared/config"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

var Chi *chi.Mux

func init() {
	config.Installations.EnableHTTPServer = true

	container.InjectInstallation(func() error {
		Chi = chi.NewRouter()
		Chi.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("UP"))
		})
		return nil
	}, container.InjectionProps{Parallel: false, DependencyID: uuid.NewString()})

	container.InjectHTTPServer(func() error {
		fmt.Println("starting server on port :" + config.PORT.Get())
		err := http.ListenAndServe(":"+config.PORT.Get(), Chi)
		if err != nil {
			log.Error().Err(err).Msg("error initializing application server")
			return err
		}
		return nil
	}, container.InjectionProps{Parallel: false, DependencyID: uuid.NewString()})

}

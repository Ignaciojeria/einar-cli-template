package controller

import (
	einar "archetype/app/shared/archetype/chi_server"
	"archetype/app/shared/archetype/container"
	"net/http"

	"github.com/google/uuid"
)

func init() {
	container.InjectInBoundAdapter(func() error {
		einar.Chi.Put("/INSERT_YOUR_PATTERN_HERE", archetypePutController)
		return nil
	}, container.InjectionProps{
		DependencyID: uuid.NewString(),
		Parallel:     true,
	})
}

func archetypePutController(w http.ResponseWriter, r *http.Request) {
	// Write your handling process here
	w.Write([]byte("YOUR CUSTOM RESPONSE"))
}

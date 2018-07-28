package web

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Run(addr string) error {
	router := chi.NewRouter()
	router.Post("/", writeBlockChainHandler)
	router.Get("/", getBlockChainHandler)
	return http.ListenAndServe(addr, router)
}

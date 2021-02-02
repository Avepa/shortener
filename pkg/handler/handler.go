package handler

import (
	"net/http"

	"github.com/Avepa/shortener/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/a/", h.Add)
	mux.HandleFunc("/s/", h.Get)

	return mux
}

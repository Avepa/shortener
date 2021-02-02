package handler

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/Avepa/shortener/pkg"
)

type PathData struct {
	Path string `json:"path"`
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	path, err := h.services.URL.Add(url, "")
	if err != nil {
		if err == pkg.ErrorPathIsBusy {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	data := PathData{path}
	json.NewEncoder(w).Encode(data)
	return
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	path := regexp.MustCompile(`[^//]*$`).FindString(r.URL.EscapedPath())
	url, err := h.services.URL.Get(path)
	if err != nil {
		if err == pkg.ErrorDcumentNotFound {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Redirect(w, r, url, http.StatusSeeOther)
	return
}

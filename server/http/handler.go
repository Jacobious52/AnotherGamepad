package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Handler is the main handler for the server
type Handler struct {
	*httprouter.Router
}

func NewHandler() *Handler {

	h := &Handler{
		Router: httprouter.New(),
	}

	h.GET("/", h.index)
	h.GET("/status", h.status)
	h.ServeFiles("/static/*filepath", http.Dir("static"))

	return h
}

func (h *Handler) index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "static/index.html")
}

func (h *Handler) status(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("webserver running"))
}

package api

import (
	"net/http"

	"github.com/arthurvicencio/nestlink/pkg/html"
)

func Register() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		html.Render(w, "index.html", nil)
	})
}

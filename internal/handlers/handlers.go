package handlers

import (
	"ascii-converter/internal/templates"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templates.Home("ASCII Converter").Render(r.Context(), w)
}

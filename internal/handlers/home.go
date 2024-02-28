package handlers

import (
	"fmt"
	"net/http"
)

func (h *handler) home(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	data := h.app.NewTemplateData(r)
	h.app.Render(w, http.StatusOK, "home.html", data)
}

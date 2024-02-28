package handlers

import (
	"net/http"
)

const (
	pageSize = 5
)

func (h *handler) home(w http.ResponseWriter, r *http.Request) {
	data := h.app.NewTemplateData(r)
	categories, err := h.service.GetAllCategory()
	if err != nil {
		h.app.ServerError(w, err)
	}
	data.Categories = categories

	posts, err := h.service.GetAllPostPaginated(1, pageSize)

	if err != nil {
		h.app.ServerError(w, err)
	}

	data.Posts = posts
	h.app.Render(w, http.StatusOK, "home.html", data)
}

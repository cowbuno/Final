package handlers

import (
	"net/http"
	"strconv"
)

const (
	pageSize    = 2
	defaultPage = 1
)

func (h *handler) home(w http.ResponseWriter, r *http.Request) {
	currentPageStr := r.URL.Query().Get("page")
	pageNumber, err := h.service.GetPageNumber(pageSize)
	if err != nil {
		h.app.ServerError(w, err)
	}
	currentPage, err := strconv.Atoi(currentPageStr)
	if err != nil || currentPage < 1 || currentPage > pageNumber {
		currentPage = defaultPage
	}

	data := h.app.NewTemplateData(r)
	categories, err := h.service.GetAllCategory()
	if err != nil {
		h.app.ServerError(w, err)
	}
	data.Categories = categories

	posts, err := h.service.GetAllPostPaginated(currentPage, pageSize)
	if err != nil {
		h.app.ServerError(w, err)
	}

	data.Posts = posts
	data.CurrentPage = currentPage
	data.NumberOfPage = pageNumber
	h.app.Render(w, http.StatusOK, "home.html", data)
}

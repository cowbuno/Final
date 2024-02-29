package handlers

import (
	"net/http"
	"strconv"
)

const (
	pageSize    = 5
	defaultPage = 1
)

func (h *handler) home(w http.ResponseWriter, r *http.Request) {
	methodResolver(w, r, h.homeGet, h.homePost)
}

func (h *handler) homeGet(w http.ResponseWriter, r *http.Request) {
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

func (h *handler) homePost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.app.ClientError(w, http.StatusBadRequest)
		return
	}

	filterCategoriesString := r.Form["categories"]
	filterCategories, err := ConverCategories(filterCategoriesString)
	if err != nil {
		h.app.ClientError(w, http.StatusBadRequest)
	}
	posts, err := h.service.GetAllPostByCategories(filterCategories)
	if err != nil {
		h.app.ServerError(w, err)
	}
	data := h.app.NewTemplateData(r)

	categories, err := h.service.GetAllCategory()
	if err != nil {
		h.app.ServerError(w, err)
	}
	data.Categories = categories

	data.Posts = posts
	h.app.Render(w, http.StatusOK, "home.html", data)
}

func ConverCategories(CategoriesString []string) ([]int, error) {
	categories := make([]int, len(CategoriesString))
	for i, str := range CategoriesString {
		nb, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		categories[i] = nb
	}

	return categories, nil
}

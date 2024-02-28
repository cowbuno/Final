package handlers

import (
	"fmt"
	"forum/models"
	"forum/pkg/validator"
	"net/http"
	"strings"
)

func (h *handler) postCreate(w http.ResponseWriter, r *http.Request) {
	methodResolver(w, r, h.postCreateGet, h.postCreatePost)
}

func (h *handler) postCreateGet(w http.ResponseWriter, r *http.Request) {
	data := h.app.NewTemplateData(r)

	data.Form = models.PostForm{}
	categories, err := h.service.GetAllCategory()
	if err != nil {
		h.app.ServerError(w, err)
	}
	data.Categories = categories
	h.app.Render(w, http.StatusOK, "create.html", data)
}

func (h *handler) postCreatePost(w http.ResponseWriter, r *http.Request) {
	form := models.PostForm{
		Title:            r.FormValue("title"),
		Content:          r.FormValue("content"),
		CategoriesString: r.Form["categories"],
	}

	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
	form.CheckField(validator.NotSelected(form.CategoriesString), "categories", "This field cannot be selected")

	if !form.Valid() {
		data := h.app.NewTemplateData(r)
		data.Form = form
		categories, err := h.service.GetAllCategory()
		if err != nil {
			h.app.ServerError(w, err)
		}
		data.Categories = categories
		h.app.Render(w, http.StatusUnprocessableEntity, "create.html", data)
		return
	}

	
}

func (h *handler) postView(w http.ResponseWriter, r *http.Request) {
	id, _ := strings.CutPrefix(r.URL.Path, "/post/")
	fmt.Fprintf(w, "View Post %s", id)
}

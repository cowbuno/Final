package models

import (
	"forum/pkg/validator"
	"time"
)

type Post struct {
	PostID    int
	UserID    int
	Title     string
	Content   string
	ImageName string
	Created   time.Time
	Like      int
	Dislike   int
	Comment   *[]Comment
	Category  map[int]string
}

type Comment struct {
	CommentId      int
	PostID         int
	CreatedUserID  int
	Content        string
	CreatedTime    time.Time
	LikeCounter    string
	DislikeCounter string
}

type PostForm struct {
	Title               string   `form:"title"`
	Content             string   `form:"content"`
	Categories            []int    `form:"category"`
	CategoriesString      []string `form:"category"`
	validator.Validator `form:"-"`
}

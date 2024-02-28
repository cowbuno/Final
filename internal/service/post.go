package service

import (
	"fmt"
	"forum/models"
)

func (s *service) CreatePost(title, content, token string, categories []int) (int, error) {
	userID, err := s.repo.GetUserIDByToken(token)
	if err != nil {
		return 0, err
	}

	postID, err := s.repo.CreatePost(userID, title, content, "Nan")
	if err != nil {
		return 0, err
	}

	if err = s.repo.AddCategoryToPost(postID, AddCategory(categories)); err != nil {
		return 0, err
	}
	return postID, err
}

func (s *service) GetPostByID(id int) (*models.Post, error) {
	post, err := s.repo.GetPostByID(id)
	if err != nil {
		return nil, err
	}

	categories, err := s.repo.GetCategoriesByPostID(id)
	if err != nil {
		return nil, err
	}
	post.Categories = categories
	return post, nil
}

func (s *service) GetAllPostPaginated(curentPage, pageSize int) (*[]models.Post, error) {
	posts, err := s.repo.GetAllPostPaginated(curentPage, pageSize)
	if err != nil {
		return nil, err
	}

	fmt.Println(posts)
	for _, post := range *posts {
		categories, err := s.repo.GetCategoriesByPostID(post.PostID)
		if err != nil {
			return nil, err
		}
		post.Categories = categories
	}
	return posts, nil
}

func (s *service) GetPageNumber(pageSize int) (int, error) {
	return 0, nil
}

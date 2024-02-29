package repo

import (
	"forum/internal/repo/sqlite"
	"forum/models"
)

type UserRepo interface {
	CreateUser(models.User) error
	GetUserByID(int) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
	UpdateUserByID(string) (*models.User, error)
	Authenticate(email, password string) (int, error)
}

type SessionRepo interface {
	GetUserIDByToken(string) (int, error)
	CreateSession(*models.Session) error
	DeleteSessionByUserID(int) error
	DeleteSessionByToken(string) error
}

type PostRepo interface {
	CreatePost(userID int, title, content, imageName string) (int, error)
	GetPostByID(int) (*models.Post, error)
	GetCategoriesByPostID(int) (map[int]string, error)
	// GetAllPost() (*models.Post, error)
	// UpdatePost(string, *models.Post) error
	//AddLikeAndDislike(bool, string, string) error
	// DeleteLikeAndDislike(int, int) error
	GetAllPostByUserID(int) (*[]models.Post, error)
	GetAllPostByCategories(categories []int) (*[]models.Post, error)
	GetPageNumber(pageSize int) (int, error)
	GetAllPostPaginated(page int, pageSize int) (*[]models.Post, error)
}

type CategoryRepo interface {
	AddCategoryToPost(int, []int) error
	GetALLCategory() ([]string, error)
	// CreateCategory(string) error
}

// type CommentRepo interface {
// 	CreateComment(*models.Comment) error
// 	GetAllCommentByPostID(string) (*[]models.Post, error)
// 	GetAllCommentByUserID(string) (*[]models.Post, error)
// 	AddLikeAndDislike(bool, string, string) error
// }

type RepoI interface {
	UserRepo
	SessionRepo
	PostRepo
	CategoryRepo
	// CommentRepo
}

func New(storagePath string) (RepoI, error) {
	return sqlite.NewDB(storagePath)
}

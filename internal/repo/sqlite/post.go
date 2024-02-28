package sqlite

import (
	"forum/models"
	"strings"
)

func (s *Sqlite) CreatePost(post *models.Post) error {
	const query = `INSERT INTO Post (user_id, title, content, image_name) VALUES (?, ?, ?, ?)`
	_, err := s.db.Exec(query, post.UserID, post.Title, post.Content, post.ImageName)

	return err
}

func (s *Sqlite) GetAllPost() ([]models.Post, error) {
	const query = `SELECT post_id, user_id, title, content, created, like, dislike, image_name FROM Post`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.PostID, &post.UserID, &post.Title, &post.Content, &post.Created, &post.Like, &post.Dislike, &post.ImageName)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (s *Sqlite) AddLikeAndDislike(isLike bool, userID, postID string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	// Check if the user has already liked/disliked the post
	var isExists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM Post_User_Like WHERE user_id = ? AND post_id = ?)`
	err = tx.QueryRow(checkQuery, userID, postID).Scan(&isExists)
	if err != nil {
		tx.Rollback()
		return err
	}

	if !isExists {
		// Insert like/dislike
		insertQuery := `INSERT INTO Post_User_Like (user_id, post_id, is_like) VALUES (?, ?, ?)`
		_, err := tx.Exec(insertQuery, userID, postID, isLike)
		if err != nil {
			tx.Rollback()
			return err
		}

		// Update Post like/dislike count
		updateQuery := ""
		if isLike {
			updateQuery = `UPDATE Post SET like = like + 1 WHERE post_id = ?`
		} else {
			updateQuery = `UPDATE Post SET dislike = dislike + 1 WHERE post_id = ?`
		}
		_, err = tx.Exec(updateQuery, postID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (s *Sqlite) DeleteLikeAndDislike(userID, postID int) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	//is user liked or disliked
	var isLike bool
	checkQuery := `SELECT is_like FROM Post_User_Like WHERE user_id = ? AND post_id = ?`
	err = tx.QueryRow(checkQuery, userID, postID).Scan(&isLike)
	if err != nil {
		tx.Rollback()
		return err
	}

	// delete the like/dislike
	deleteQuery := `DELETE FROM Post_User_Like WHERE user_id = ? AND post_id = ?`
	_, err = tx.Exec(deleteQuery, userID, postID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// decrement the like or dislike
	updateQuery := ""
	if isLike {
		updateQuery = `UPDATE Post SET like = like - 1 WHERE post_id = ? AND like > 0`
	} else {
		updateQuery = `UPDATE Post SET dislike = dislike - 1 WHERE post_id = ? AND dislike > 0`
	}
	_, err = tx.Exec(updateQuery, postID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (s *Sqlite) GetAllPostByUserID(userID int) (*[]models.Post, error) {
	const query = `SELECT post_id, user_id, title, content, created, like, dislike, image_name FROM Post WHERE user_id=?`
	rows, err := s.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.PostID, &post.UserID, &post.Title, &post.Content, &post.Created, &post.Like, &post.Dislike, &post.ImageName)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return &posts, nil
}

func (s *Sqlite) GetAllPostByCategories(categoryIDs []int) ([]*models.Post, error) {
	query := `SELECT p.post_id, p.user_id, p.title, p.content, p.created, p.like, p.dislike, p.image_name
              FROM Post AS p
              INNER JOIN Post_Category AS pc ON p.post_id = pc.post_id
              WHERE pc.category_id IN (?` + strings.Repeat(",?", len(categoryIDs)-1) + `)
              GROUP BY p.post_id`

	args := make([]interface{}, len(categoryIDs))
	for i, id := range categoryIDs {
		args[i] = id
	}

	// Execute the query
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.PostID, &post.UserID, &post.Title, &post.Content, &post.Created, &post.Like, &post.Dislike, &post.ImageName); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	return posts, nil
}

func (s *Sqlite) GetAllPostPaginated(page int, pageSize int) (*[]models.Post, error) {
	offset := (page - 1) * pageSize
	query := `SELECT post_id, user_id, title, content, created, like, dislike, image_name FROM Post LIMIT ? OFFSET ?`

	rows, err := s.db.Query(query, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.PostID, &post.UserID, &post.Title, &post.Content, &post.Created, &post.Like, &post.Dislike, &post.ImageName); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return &posts, nil
}

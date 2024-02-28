package sqlite

import "fmt"

func AddCategoryToPost(int, []int) error {
	return nil
}

func (s *Sqlite) GetALLCategory() ([]string, error) {
	op := "sqlite.GetAllCategory"
	stmt := `SELECT name FROM category ORDER BY id ASC`

	rows, err := s.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []string
	for rows.Next() {
		var categoryName string
		err := rows.Scan(&categoryName)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		categories = append(categories, categoryName)
	}

	return categories, nil
}

func CreateCategory(string) error {
	return nil
}

func (s *Sqlite) GetCategoryByPostID(postID int) (map[int]string, error) {
	query := `SELECT 
	category_id 
	category.name as name
	FROM 
	Post_Category 
	INNER JOIN Category ON Post_Category.category_id = Category.category_id
	WHERE post_id=? `

	rows, err := s.db.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var category map[int]string
	for rows.Next() {
		var categoryID int
		var categoryName string
		err := rows.Scan(&categoryID, &categoryName)
		if err != nil {
			return nil, err
		}
		category[categoryID] = categoryName
	}

	return category, nil
}

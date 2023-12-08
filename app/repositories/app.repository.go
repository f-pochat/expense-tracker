package repositories

import "expense-track/app/db"

func GetCategories(userID string) ([]string, error) {
	var categories []string

	sql := "SELECT category_name FROM categories WHERE user_id = $1"
	rows, err := db.DB.Query(sql, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var categoryName string
		err := rows.Scan(&categoryName)
		if err != nil {
			return nil, err
		}
		categories = append(categories, categoryName)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

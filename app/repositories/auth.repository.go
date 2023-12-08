package repositories

import (
	"expense-track/app/db"
	"expense-track/app/models/dto"
	"expense-track/app/models/entities"
)

func GetUser(username string) (entities.User, error) {
	var user entities.User
	sql := "SELECT id, email, username, password FROM users WHERE username = $1 LIMIT 1"
	err := db.DB.QueryRow(sql, username).Scan(&user.ID, &user.Email, &user.Username, &user.Password)

	return user, err
}

func CreateUser(user dto.RegisterDTO) (entities.User, error) {
	var newUser entities.User
	sql := "INSERT INTO users (email, username, password) VALUES ($1, $2, $3) RETURNING id"
	err := db.DB.QueryRow(sql, user.Email, user.Username, user.Password).Scan(&newUser.ID)

	return newUser, err
}

package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

func (u users) CreateUser(user models.User) (uint64, error) {
	statement, err := u.db.Prepare("insert into users (name ,nick ,email,password) values(?,?,?,?)")
	if err != nil {
		return 0, nil
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return uint64(userID), nil
}

func (u users) GetUser(searchTerm string) ([]models.User, error) {
	searchTerm = fmt.Sprintf("%%%s%%", searchTerm)

	rows, err := u.db.Query(
		"SELECT id, name, nick, email, createDate FROM users WHERE name LIKE ? or nick LIKE ?",
		searchTerm, searchTerm,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreateDate,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

func (u users) GetUserByID(ID uint64) (models.User, error) {
	rows, err := u.db.Query(
		"SELECT id, name, nick, email, createDate FROM users WHERE id = ?",
		ID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreateDate,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

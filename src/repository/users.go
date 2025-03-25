package repository

import (
	"api/src/models"
	"database/sql"
	"errors"
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
	fmt.Println("erro", err)
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

func (u users) UpdateUser(ID uint64, user models.User) error {
	statement, err := u.db.Prepare(
		"UPDATE users set name = ?, nick = ?, email = ? WHERE id = ?",
	)

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(
		user.Name,
		user.Nick,
		user.Email,
		ID,
	); err != nil {
		return err
	}

	return nil
}

func (u users) DeleteUser(ID uint64) error {
	statement, err := u.db.Prepare("DELETE from users WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (u users) GetUserByEmail(email string) (models.User, error) {
	row := u.db.QueryRow("SELECT id, password FROM users WHERE email = ?", email)

	var user models.User
	err := row.Scan(&user.ID, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, errors.New("user not found")
		}
		return models.User{}, err
	}

	return user, nil
}

func (u users) FollowUser(userID, followerID uint64) error {
	statement, err := u.db.Prepare("insert ignore into followers (user_id, follower_id) values (?, ?)")

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

func (u users) UnfollowUser(userID, followerID uint64) error {
	statement, err := u.db.Prepare("delete from followers where user_id = ? and follower_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

func (u users) GetFollowers(userID uint64) ([]models.User, error) {
	rows, err := u.db.Query(`
	select u.id, u.name, u.nick, u.email, u.createDate
	from users u inner join followers s on u.id = s.follower_id where s.user_id = ?
	`, userID)
	if err != nil {
		return nil, err
	}

	var followers []models.User
	for rows.Next() {
		var follower models.User

		if err = rows.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nick,
			&follower.Email,
			&follower.CreateDate,
		); err != nil {
			return nil, err
		}

		followers = append(followers, follower)
	}

	return followers, nil
}

func (u users) GetFollowing(userID uint64) ([]models.User, error) {
	rows, err := u.db.Query(`
	SELECT u.id, u.name, u.nick, u.email, u.createDate
	FROM users u INNER JOIN followers s on u.id = s.user_id where s.follower_id = ?
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []models.User
	for rows.Next() {
		var follower models.User

		if err = rows.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nick,
			&follower.Email,
			&follower.CreateDate,
		); err != nil {
			return nil, err
		}

		followers = append(followers, follower)
	}

	return followers, nil
}

func (u users) GetPassword(userID uint64) (string, error) {
	row, err := u.db.Query("select password from users where id = ?", userID)
	if err != nil {
		return "", err
	}
	defer row.Close()

	var user models.User

	if row.Next() {
		if err = row.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

func (u users) UpdatePassword(userID uint64, password string) error {
	statement, err := u.db.Prepare("UPDATE users SET password = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(password, userID); err != nil {
		return err
	}

	return nil
}

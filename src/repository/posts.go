package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Posts struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

func (p Posts) CreatePost(post models.Post) (uint64, error) {
	statement, err := p.db.Prepare("insert into posts (title, content,author_id) values (?,?,?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil
}

func (p Posts) GetPosts(userID uint64) ([]models.Post, error) {
	rows, err := p.db.Query(`
	select distinct p.*, u.nick from posts p
	inner join users u on u.id = p.author_id
	inner join followers s on p.author_id = s.user_id 
	where u.id = ? or s.follower_id= ?
`, userID, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreateDate,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (p Posts) GetPostById(postID uint64) (models.Post, error) {
	row, err := p.db.Query(`
	select p.*, u.nick from
	posts p inner join users u
	on u.id = p.author_id where p.id = ?
	`, postID)

	if err != nil {
		return models.Post{}, err
	}
	defer row.Close()
	fmt.Println("row", row)
	var post models.Post

	if row.Next() {
		if err = row.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreateDate,
			&post.AuthorNick,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil

}

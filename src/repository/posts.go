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
	order by 1 desc
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

func (p Posts) UpdatePost(postID uint64, post models.Post) error {
	statement, err := p.db.Prepare("update posts set title = ?, content = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(post.Title, post.Content, postID); err != nil {
		return err
	}

	return nil
}

func (p Posts) DeletePost(postID uint64) error {
	statement, err := p.db.Prepare("delete from posts where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil

}

func (p Posts) GetPostByUser(userID uint64) ([]models.Post, error) {
	rows, err := p.db.Query(`
	select p.*, u.nick from posts p
	join users u on u.id = p.author_id
	where p.author_id = ?
	`, userID)
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

func (p Posts) LikeInPost(postID uint64) error {
	statement, err := p.db.Prepare("update posts set likes = likes + 1 where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}

func (p Posts) UnlikeInPost(postID uint64) error {
	statement, err := p.db.Prepare(`
	update posts set likes = 
	CASE WHEN likes > 0 THEN likes - 1
	ELSE 0 END
	where id = ?
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}

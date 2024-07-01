package model

import (
	"database/sql"
)

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	// UserID    sql.NullInt64 `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type postModel struct {
	conn *sql.DB
}

func NewPostModel(db *sql.DB) *postModel {
	return &postModel{conn: db}
}

func (p *postModel) FindAll() ([]Post, error) {
	rows, err := p.conn.Query("SELECT id, title, content, created_at FROM posts ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	posts := []Post{}
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (p *postModel) FindById(id int) (Post, error) {
	row := p.conn.QueryRow("SELECT id, title, content, created_at FROM posts WHERE id = ?", id)
	var post Post
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt)
	if err != nil {
		return Post{}, err
	}
	return post, nil
}

func (p *postModel) Create(user_id int, data Post) (Post, error) {
	stmt, err := p.conn.Prepare("INSERT INTO posts(title, content, user_id) VALUES(?, ?, ?)")
	if err != nil {
		return Post{}, err
	}
	_, err = stmt.Exec(data.Title, data.Content, user_id)
	if err != nil {
		return Post{}, err
	}
	return data, nil
}

func (p *postModel) Update(id int, data Post) (Post, error) {
	stmt, err := p.conn.Prepare("UPDATE posts SET title = ?, content = ?, user_id = ? WHERE id = ?")
	if err != nil {
		return Post{}, err
	}
	_, err = stmt.Exec(data.Title, data.Content, id)
	if err != nil {
		return Post{}, err
	}
	return data, nil
}

func (p *postModel) Delete(id int) error {
	stmt, err := p.conn.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (p *postModel) FindByUserId(id int) ([]Post, error) {
	rows, err := p.conn.Query("SELECT * FROM posts WHERE user_id = ?", id)
	if err != nil {
		return nil, err
	}
	posts := []Post{}
	for rows.Next() {
		var post Post

		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (p *postModel) Search(query string) ([]Post, error) {
	rows, err := p.conn.Query("SELECT id, title, content, created_at FROM posts WHERE title LIKE ? OR content LIKE ?", "%"+query+"%", "%"+query+"%")
	if err != nil {
		return nil, err
	}
	posts := []Post{}
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

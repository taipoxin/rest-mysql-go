package models

import (
	"database/sql"
	"fmt"
	"log"
)

// Post entity of table `posts`
type Post struct {
	ID    int64
	Title string
}

func (p *Post) String() string {
	return fmt.Sprintf("{id: %v, title: %v}", p.ID, p.Title)
}

// AllPosts return all `posts` from db
func (db *DbHelper) AllPosts() ([]*Post, error) {
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := make([]*Post, 0)
	for rows.Next() {
		post := &Post{}
		err := rows.Scan(&post.ID, &post.Title)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

// GetPost return post by id from db
func (db *DbHelper) GetPost(id int64) (*Post, error) {
	rows, err := db.Query("SELECT * FROM posts WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	post := &Post{}
	for rows.Next() {
		err := rows.Scan(&post.ID, &post.Title)
		if err != nil {
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return post, nil
}

// curl -X POST -d "{\"title\": \"that\"}" http://localhost:3000/post

// AddPost add new post with title
func (db *DbHelper) AddPost(title string) error {
	result, err := db.Exec("INSERT INTO  posts (title) VALUES (?)", title)
	_, _, err = logResultSet(result)
	if err != nil {
		return err
	}
	return nil
}

// curl -X PUT -d "{\"title\": \"this\", \"id\" : 5}" http://localhost:3000/post

// UpdatePost update post with new title by id
func (db *DbHelper) UpdatePost(id int64, title string) (bool, error) {
	result, err := db.Exec("UPDATE posts SET title = ? WHERE id = ?", title, id)
	if err != nil {
		return false, err
	}
	_, ra, err := logResultSet(result)
	if err != nil {
		return false, err
	}
	if ra != 0 {
		return true, nil
	}
	return false, nil
}

// curl -X DELETE  http://localhost:3000/post?id=3

// DeletePost delete post by id
func (db *DbHelper) DeletePost(id int64) (bool, error) {
	result, err := db.Exec("DELETE from posts WHERE id = ?", id)
	if err != nil {
		return false, err
	}
	_, ra, err := logResultSet(result)
	if err != nil {
		return false, err
	}
	if ra != 0 {
		return true, nil
	}
	return false, nil
}

func logResultSet(result sql.Result) (lastInserted int64, rowsAffected int64, err error) {
	li, err := result.LastInsertId()
	if err != nil {
		return 0, 0, err
	}
	ra, err := result.RowsAffected()
	if err != nil {
		return 0, 0, err
	}
	log.Println("exec on post:{ last id: ", li, " rows affected: ", ra, " }")
	return li, ra, nil
}

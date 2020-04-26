package models

import (
	"fmt"
)

// Post entity of table `posts`
type Post struct {
	ID    int64
	Title string
}

func (p *Post) String() string {
	return fmt.Sprintf("id: %v, title: %v", p.ID, p.Title)
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

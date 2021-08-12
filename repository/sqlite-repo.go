package repository

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/seonicklaus/rest-api-go/entity"
)

type sqliteRepo struct{}

func NewSQLiteRepository() PostRepository {
	os.Remove("./posts.db")

	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sqlStmt := `
	CREATE TABLE POSTS (id INTEGER NOT NULL PRIMARY KEY, title TEXT, txt TEXT);
	delete from post;
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s", err, sqlStmt)
	}
	return &sqliteRepo{}
}

func (*sqliteRepo) Save(post *entity.Post) (*entity.Post, error) {
	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	stmt, err := tx.Prepare("INSERT INTO posts (id, title, txt) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(post.ID, post.Title, post.Text)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	tx.Commit()
	return post, nil
}

func (*sqliteRepo) FindAll() ([]entity.Post, error) {
	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	rows, err := db.Query("SELECT id, title, txt FROM posts")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer rows.Close()
	var posts []entity.Post

	for rows.Next() {
		var id int64
		var title string
		var text string

		rows.Scan(&id, &title, &text)

		post := entity.Post{ID: id, Title: title, Text: text}
		posts = append(posts, post)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return posts, nil
}

func (*sqliteRepo) Delete(post *entity.Post) error {
	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return err
	}

	stmt, err := tx.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(post.ID)
	if err != nil {
		log.Fatal(err)
		return err
	}

	tx.Commit()
	return nil
}

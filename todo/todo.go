package todo

import (
	"time"
	"todox/db"
)

func init() {
	db.ConnectPostgres()
}

func Add(content string) {
	_, err := db.DB.Exec("INSERT INTO todos (content, created_at, done) VALUES ($1, $2, $3)", content, time.Now(), false)
	if err != nil {
		panic(err)
	}
}

func List() []*Todo {
	rows, err := db.DB.Query("SELECT id, content, created_at, done FROM todos")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var todos []*Todo
	for rows.Next() {
		t := &Todo{}
		err := rows.Scan(&t.ID, &t.Content, &t.createdAt, &t.Done)
		if err != nil {
			panic(err)
		}
		todos = append(todos, t)
	}

	return todos
}

func Done(index int) error {
	_, err := db.DB.Exec("UPDATE todos SET done = true WHERE id = $1", index)
	if err != nil {
		return err
	}
	return nil
}

func Undone(index int) error {
	_, err := db.DB.Exec("UPDATE todos SET done = false WHERE id = $1", index)
	if err != nil {
		return err
	}
	return nil
}

func Del(index int) error {
	_, err := db.DB.Exec("DELETE FROM todos WHERE id = $1", index)
	if err != nil {
		return err
	}

	return nil
}

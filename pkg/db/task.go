package db

import (
	"context"
	"fmt"
)

type Task struct {
	ID      int64  `json:"id"` // при сериализации нужно будет перевести в строку
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

const addTaskDB = `INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)`

func AddTask(ctx context.Context, task *Task) (int64, error) {
	query := addTaskDB
	res, err := db.ExecContext(ctx, query, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return 0, fmt.Errorf("insert task: %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("last insert: %w", err)
	}
	return id, nil
}

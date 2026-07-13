package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type Task struct {
	ID      int64  `json:"id,string"` // при сериализации нужно будет перевести в строку
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

// SQL запросы
const addTaskDB = `
INSERT INTO scheduler (date, title, comment, repeat)
 VALUES (?, ?, ?, ?)
 `

const listTasks = `
SELECT id, date, title, comment, repeat 
FROM scheduler ORDER BY date LIMIT ?
`

const searchTasks = `SELECT id, date, title, comment, repeat FROM scheduler
WHERE title LIKE ? OR comment LIKE ? ORDER BY date LIMIT ?`

const getTaskDB = `SELECT id, date, title, comment, repeat FROM scheduler
	WHERE id = ?`

const updateTaskDB = `UPDATE scheduler 
	SET date = ?, title = ?, comment = ?, repeat = ? WHERE id = ?`

const deleteTaskDB = `DELETE FROM scheduler WHERE id = ?`

// Функция для создания задачи
func AddTask(ctx context.Context, task *Task) (int64, error) {
	res, err := db.ExecContext(ctx, addTaskDB, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return 0, fmt.Errorf("insert task: %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("last insert: %w", err)
	}
	return id, nil
}

// Функция получения задачи
func Tasks(ctx context.Context, limit int, search string) ([]*Task, error) {
	var rows *sql.Rows
	var err error

	if search != "" {
		pattern := "%" + search + "%"

		rows, err = db.QueryContext(
			ctx,
			searchTasks,
			pattern,
			pattern,
			limit,
		)
		if err != nil {
			return nil, err
		}
	} else {
		rows, err = db.QueryContext(
			ctx,
			listTasks,
			limit,
		)
		if err != nil {
			return nil, err
		}
	}

	tasks := make([]*Task, 0)

	defer rows.Close()

	for rows.Next() {
		task := &Task{}

		err := rows.Scan(
			&task.ID,
			&task.Date,
			&task.Title,
			&task.Comment,
			&task.Repeat,
		)

		if err != nil {
			return tasks, err
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return tasks, err
	}

	return tasks, nil
}

func GetTask(ctx context.Context, id string) (*Task, error) {
	task := &Task{}

	err := db.QueryRowContext(ctx, getTaskDB, id).Scan(
		&task.ID,
		&task.Date,
		&task.Title,
		&task.Comment,
		&task.Repeat,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("Задача не найдена")
	}

	if err != nil {
		return nil, fmt.Errorf("get task: %w", err)
	}

	return task, nil
}

func UpdateTask(ctx context.Context, task *Task) error {
	res, err := db.ExecContext(
		ctx,
		updateTaskDB,
		task.Date,
		task.Title,
		task.Comment,
		task.Repeat,
		task.ID,
	)

	if err != nil {
		return fmt.Errorf("update task: %w", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}

	if count == 0 {
		return fmt.Errorf("задача не найдена")
	}

	return nil
}

func DeleteTask(ctx context.Context, id int64) error {

	res, err := db.ExecContext(ctx, deleteTaskDB, id)
	if err != nil {
		return fmt.Errorf("delete task: %w", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}

	if count == 0 {
		return fmt.Errorf("задача не найдена")
	}

	return nil
}

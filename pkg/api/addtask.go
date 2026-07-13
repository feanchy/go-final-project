package api

import (
	"encoding/json"
	"net/http"
	"time"

	"go1f/pkg/db"
)

func addTask(w http.ResponseWriter, r *http.Request) error {
	var task db.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		return err
	}
	if task.Title == "" {
		return err
	}

	err = checkDate(&task)
	if err != nil {
		return err
	}

	id, err := db.AddTask(r.Context(), &task)
	if err != nil {
		return err
	}

	return nil
}

// Проверка даты на корректность
func checkDate(task *db.Task) error {
	now := time.Now()

	if task.Date == "" {
		task.Date = now.Format(dateFormat)
	}

	t, err := time.Parse(dateFormat, task.Date)
	if err != nil {
		return err
	}

	var next string
	if task.Repeat != "" {
		next, err = NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}
	}
	_ = next

	return nil
}

// Хэндлер для добавления задачи
func addTaskHandler(w http.ResponseWriter, r *http.Request) {

}

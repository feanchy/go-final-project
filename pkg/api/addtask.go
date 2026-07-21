package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"go1f/pkg/db"
)

func addTask(w http.ResponseWriter, r *http.Request) (int64, error) {
	var task db.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		return 0, err
	}
	if task.Title == "" {
		return 0, errors.New("title is empty")
	}

	err = checkDate(&task)
	if err != nil {
		return 0, err
	}

	id, err := db.AddTask(r.Context(), &task)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func writeJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Проверка даты на корректность
func checkDate(task *db.Task) error {
	now := time.Now()

	if task.Date == "" || task.Date == "today" {
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

	if afterNow(now, t) {
		if task.Repeat == "" {
			task.Date = now.Format(dateFormat)
		} else {
			task.Date = next
		}
	}

	return nil
}

// Хэндлер для добавления задачи
func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	id, err := addTask(w, r)
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, map[string]string{"id": strconv.FormatInt(id, 10)})
}

package api

import (
	"go1f/pkg/db"
	"net/http"
	"strconv"
	"time"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	if id == "" {
		writeJSON(w, map[string]string{
			"error": "Не указан идентификатор",
		})
		return
	}

	task, err := db.GetTask(r.Context(), id)
	if err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	taskID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		writeJSON(w, map[string]string{
			"error": "Неккоректный id",
		})
		return
	}
	if task.Repeat == "" {

		err := db.DeleteTask(r.Context(), taskID)
		if err != nil {
			writeJSON(w, map[string]string{
				"error": err.Error(),
			})
			return
		}
	} else {
		taskDate, err := time.Parse(dateFormat, task.Date)
		if err != nil {
			writeJSON(w, map[string]string{"error": err.Error()})
			return
		}

		date, err := NextDate(taskDate, task.Date, task.Repeat)
		if err != nil {
			writeJSON(w, map[string]string{
				"error": err.Error(),
			})
			return
		}
		err = db.UpdateDate(r.Context(), date, taskID)
		if err != nil {
			writeJSON(w, map[string]string{
				"error": err.Error(),
			})
			return
		}

	}

	writeJSON(w, map[string]string{})
}

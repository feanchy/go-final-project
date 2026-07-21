package api

import (
	"encoding/json"
	"go1f/pkg/db"
	"net/http"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if task.ID == 0 {
		writeJSON(w, map[string]string{
			"error": "не указан id",
		})
		return
	}

	if task.Title == "" {
		writeJSON(w, map[string]string{
			"error": "поле title пустое",
		})
		return
	}

	err = checkDate(&task)
	if err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}
	err = db.UpdateTask(r.Context(), &task)
	if err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	writeJSON(w, map[string]string{})
}

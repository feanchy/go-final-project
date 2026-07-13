package api

import (
	"net/http"
	"strconv"

	"go1f/pkg/db"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		writeJSON(w, map[string]string{
			"error": "Не указан идентификатор",
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

	err = db.DeleteTask(r.Context(), taskID)

	if err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	writeJSON(w, map[string]string{})
}

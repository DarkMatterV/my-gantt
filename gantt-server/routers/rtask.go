package routers

import (
	"gantt-server/httputils"
	"gantt-server/models"
	"net/http"
)

func HandleGetTaskList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		task := models.Task{}
		l, err := task.List()
		if err != nil {
			_, _ = httputils.ResponseJson(w, httputils.CodeInterval, "get task list failed", nil)
			return
		}
		_, _ = httputils.ResponseJson(w, httputils.CodeOK, "success", l)
	}
}

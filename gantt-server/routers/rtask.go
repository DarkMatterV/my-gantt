package routers

import (
	"encoding/json"
	"gantt-server/httputils"
	"gantt-server/models"
	"net/http"
)

func HandleGetTaskList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		task := models.Task{}
		l, err := task.ListTree()
		if err != nil {
			panic(err)
			_, _ = httputils.ResponseJson(w, httputils.CodeErrInterval, "get task list failed", nil)
			return
		}
		_, _ = httputils.ResponseJson(w, httputils.CodeOK, "success", l)
	}
}

func HandleAddTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.ReqAddTask{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			_, _ = httputils.ResponseJson(w, httputils.CodeErrParams, "add task param failed", nil)
			return
		}
		task := models.Task{}
		if err := task.Add(req); err != nil {
			_, _ = httputils.ResponseJson(w, httputils.CodeErrInterval, "add task failed", nil)
		}
		_, _ = httputils.ResponseJson(w, httputils.CodeOK, "success", nil)
	}
}

func HandleUpdateTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.ReqUpdateTask{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			_, _ = httputils.ResponseJson(w, httputils.CodeErrParams, "update task param failed", nil)
			return
		}
		d := models.Task{ID: req.ID}
		info, exist, err := d.GetByID()
		if err != nil {
			_, _ = httputils.ResponseJson(w, httputils.CodeErrInterval, "update task failed", nil)
			return
		}
		if !exist {
			_, _ = httputils.ResponseJson(w, httputils.CodeErrInterval, "task id not exist", nil)
			return
		}
		if err = info.Update(req); err != nil {
			_, _ = httputils.ResponseJson(w, httputils.CodeErrInterval, "update task failed", nil)
			return
		}
		_, _ = httputils.ResponseJson(w, httputils.CodeOK, "success", nil)
	}
}

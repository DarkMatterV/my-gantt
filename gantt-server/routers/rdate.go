package routers

import (
	"encoding/json"
	"gantt-server/httputils"
	"gantt-server/models"
	"net/http"
)

func HandleGetDateList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		date := models.Date{}
		l, err := date.List()
		if err != nil {
			_, _ = httputils.ResponseJson(w, httputils.CodeErrInterval, "get date list failed", nil)
			return
		}
		_, _ = httputils.ResponseJson(w, httputils.CodeOK, "success", l)
	}
}

func HandleAddDate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.ReqAddDate{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			_, _ = httputils.ResponseJson(w, httputils.CodeErrParams, "add date param failed", nil)
			return
		}
		d := models.Date{}
		if err := d.Add(req); err != nil {
			_, _ = httputils.ResponseJson(w, httputils.CodeErrInterval, "add date failed", nil)
		}
		_, _ = httputils.ResponseJson(w, httputils.CodeOK, "success", nil)
	}
}

func HandleUpdateDate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.ReqUpdateDate{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			_, _ = httputils.ResponseJson(w, httputils.CodeErrParams, "update date param failed", nil)
			return
		}
		d := models.Date{ID: req.ID}
		info, exist, err := d.GetByID()
		if err != nil {
			_, _ = httputils.ResponseJson(w, httputils.CodeErrInterval, "update date failed", nil)
			return
		}
		if !exist {
			_, _ = httputils.ResponseJson(w, httputils.CodeErrInterval, "date id not exist", nil)
			return
		}
		if err = info.Update(req); err != nil {
			_, _ = httputils.ResponseJson(w, httputils.CodeErrInterval, "update date failed", nil)
			return
		}
		_, _ = httputils.ResponseJson(w, httputils.CodeOK, "success", nil)
	}
}

package routers

import (
	"gantt-server/httputils"
	"gantt-server/models"
	"net/http"
)

func HandleGetDateList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		date := models.Date{}
		l, err := date.List()
		if err != nil {
			_, _ = httputils.ResponseJson(w, httputils.CodeInterval, "get date list failed", nil)
			return
		}
		_, _ = httputils.ResponseJson(w, httputils.CodeOK, "success", l)
	}
}

package routers

import (
	"fmt"
	"my-gantt/models"
	"net/http"
)

func HandleGetTaskList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		task := models.Task{}
		l, err := task.List()
		if err != nil {
		}
		fmt.Println(len(l))
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Write([]byte("hello"))
	}
}

func GenerateData(data interface{}) {

}
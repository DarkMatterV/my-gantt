package main

import (
	"gantt-server/models"
	"gantt-server/routers"
	"gantt-server/settings"
	"net/http"
)

func init() {
	if err := settings.InitSetting("./config/config.json"); err != nil {
		panic(err)
	}

	if err := models.NewEngine(settings.Setting.DB.Addr); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/date/list", routers.HandleGetDateList())
	http.HandleFunc("/date/add", routers.HandleAddDate())
	http.HandleFunc("/date/update", routers.HandleUpdateDate())

	http.HandleFunc("/task/list", routers.HandleGetTaskList())
	http.HandleFunc("/task/add", routers.HandleAddTask())
	http.HandleFunc("/task/update", routers.HandleUpdateTask())

	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		panic(err)
	}
}

package models

type Task struct {
	ID     int64                        `xorm:"'id' pk autoincr" json:"id" form:"id"`
	No     string                       `xorm:"'no' notnull unique" json:"no" form:"no"`
	Title  string                       `xorm:"'title' notnull longtext" json:"title" form:"title"`
	Record map[string]map[string]string `xorm:"'record' longtext" json:"record" form:"record"`
	PID    int64                        `xorm:"'pid'" json:"pid" form:"pid"`
}

func (t *Task) Add() error {
	_, err := engine.Table("task").Insert(t)
	return err
}

func (t *Task) GetByID(id int64) (info Task, exist bool, err error) {
	exist, err = engine.Table("task").ID(id).Get()
	return
}

func (t *Task) List() (l []Task, err error) {
	tList := make([]Task, 0)
	if err = engine.Table("task").Find(&tList); err != nil {
		return l, err
	}
	return l, nil
}

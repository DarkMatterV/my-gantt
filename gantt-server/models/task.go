package models

type Task struct {
	ID     int64                        `xorm:"'id' pk autoincr" json:"id" form:"id"`
	No     string                       `xorm:"'no' notnull unique" json:"no" form:"no"`
	Title  string                       `xorm:"'title' notnull longtext" json:"title" form:"title"`
	Record map[string]map[string]string `xorm:"'record' longtext" json:"record" form:"record"`
	PID    int64                        `xorm:"'pid'" json:"pid" form:"pid"`
}

type ReqAddTask struct {
	No     string                       `json:"no" form:"no"`
	Title  string                       `json:"title" form:"title"`
	Record map[string]map[string]string `json:"record" form:"record"`
}

type ReqUpdateTask struct {
	ID     int64                        `json:"id" form:"id"`
	No     string                       `json:"no" form:"no"`
	Title  string                       `json:"title" form:"title"`
	Record map[string]map[string]string `json:"record" form:"record"`
}

func (t *Task) Add(req ReqAddTask) error {
	n := Task{
		No:     req.No,
		Title:  req.Title,
		Record: req.Record,
	}
	_, err := engine.Table("task").Insert(&n)
	return err
}

func (t *Task) Update(req ReqUpdateTask) error {
	t.No = req.No
	t.Title = req.Title
	t.Record = req.Record
	_, err := engine.Table("task").
		Where("id=?", t.ID).Cols("no", "title", "record").Update(t)
	return err
}

func (t *Task) GetByID() (info Task, exist bool, err error) {
	exist, err = engine.Table("task").ID(t.ID).Get(&info)
	return
}

func (t *Task) List() (l []Task, err error) {
	err = engine.Table("task").Find(&l)
	return l, err
}

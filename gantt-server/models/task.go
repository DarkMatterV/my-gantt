package models

type Task struct {
	ID     int64                        `xorm:"'id' pk autoincr" json:"id" form:"id"`
	No     string                       `xorm:"'no' notnull unique" json:"no" form:"no"`
	Title  string                       `xorm:"'title' notnull longtext" json:"title" form:"title"`
	Status string                       `xorm:"'status'" json:"status" form:"status"`
	Record map[string]map[string]string `xorm:"'record' longtext" json:"record" form:"record"`
	PID    int64                        `xorm:"'pid' default 0" json:"pid" form:"pid"`
}

type ReqAddTask struct {
	No     string                       `json:"no" form:"no"`
	Title  string                       `json:"title" form:"title"`
	Status string                       `json:"status" form:"status"`
	Record map[string]map[string]string `json:"record" form:"record"`
	PID    int64                        `json:"pid" form:"pid"`
}

type ReqUpdateTask struct {
	ID     int64                        `json:"id" form:"id"`
	No     string                       `json:"no" form:"no"`
	Title  string                       `json:"title" form:"title"`
	Status string                       `json:"status" form:"status"`
	Record map[string]map[string]string `json:"record" form:"record"`
}

type RespTask struct {
	ID       int64                        `xorm:"'id'" json:"id" form:"id"`
	No       string                       `xorm:"'no'" json:"no" form:"no"`
	Title    string                       `xorm:"'title'" json:"title" form:"title"`
	Status   string                       `xorm:"'status'" json:"status" form:"status"`
	Record   map[string]map[string]string `xorm:"'record'" json:"record" form:"record"`
	PID      int64                        `xorm:"'pid'" json:"pid" form:"pid"`
	Children []*RespTask                  `xorm:"-" json:"children" form:"children"`
}

func (t *Task) Add(req ReqAddTask) error {
	n := Task{
		No:     req.No,
		Title:  req.Title,
		Status: req.Status,
		Record: req.Record,
		PID:    req.PID,
	}
	_, err := engine.Table("task").Insert(&n)
	return err
}

func (t *Task) Update(req ReqUpdateTask) error {
	t.No = req.No
	t.Title = req.Title
	t.Status = req.Status
	t.Record = req.Record
	_, err := engine.Table("task").
		Where("id=?", t.ID).Cols("no", "title", "status", "record").Update(t)
	return err
}

func (t *Task) GetByID() (info Task, exist bool, err error) {
	exist, err = engine.Table("task").ID(t.ID).Get(&info)
	return
}

func (t *Task) List() (l []RespTask, err error) {
	err = engine.Table("task").Find(&l)
	return l, err
}

func (t *Task) ListTree() (l []RespTask, err error) {
	tl, err := t.List()
	if err != nil {
		return nil, err
	}
	pids := make(map[int64][]*RespTask)
	for i, task := range tl {
		if _, ok := pids[task.PID]; !ok {
			pids[task.PID] = make([]*RespTask, 0)
		}
		pids[task.PID] = append(pids[task.PID], &tl[i])
	}
	for _, task := range tl {
		if task.PID == 0 {
			task.Children = pids[task.ID]
			l = append(l, task)
		}
	}
	return l, nil
}

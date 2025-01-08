package models

type Date struct {
	ID      int64  `xorm:"'id' pk autoincr" json:"id" form:"id"`
	Date    string `xorm:"'date' notnull unique" json:"date" form:"date"`
	Week    string `xorm:"'week'" json:"week" form:"week"`
	Comment string `xorm:"'comment' longtext" json:"comment" form:"comment"`
}

type ReqAddDate struct {
	Date    string `json:"date" form:"date"`
	Week    string `json:"week" form:"week"`
	Comment string `json:"comment" form:"comment"`
}

type ReqUpdateDate struct {
	ID      int64  `json:"id" form:"id"`
	Date    string `json:"date" form:"date"`
	Week    string `json:"week" form:"week"`
	Comment string `json:"comment" form:"comment"`
}

func (d *Date) Add(req ReqAddDate) error {
	n := Date{
		Date:    req.Date,
		Week:    req.Week,
		Comment: req.Comment,
	}
	_, err := engine.Table("date").Insert(&n)
	return err
}

func (d *Date) Update(req ReqUpdateDate) error {
	d.Date = req.Date
	d.Week = req.Week
	d.Comment = req.Comment
	_, err := engine.Table("date").
		Where("id=?", d.ID).Cols("date", "week", "comment").Update(d)
	return err
}

func (d *Date) GetByID() (info Date, exist bool, err error) {
	exist, err = engine.Table("date").Where("id=?", d.ID).Get(&info)
	return
}

func (d *Date) List() (dList []Date, err error) {
	err = engine.Table("date").Find(&dList)
	return
}

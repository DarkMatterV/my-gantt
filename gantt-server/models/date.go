package models

type Date struct {
	ID      int64  `xorm:"'id' pk autoincr" json:"id" form:"id"`
	Date    string `xorm:"'date' notnull unique" json:"date" form:"date"`
	Week    string `xorm:"'week'" json:"week" form:"week"`
	Comment string `xorm:"'comment' longtext" json:"comment" form:"comment"`
}

func (d *Date) Add() error {
	_, err := engine.Table("date").Insert(d)
	return err
}

func (d *Date) List() (err error) {
	dList := make([]Date, 0)
	if err = engine.Table("date").Find(&dList); err != nil {
		return err
	}
	return nil
}

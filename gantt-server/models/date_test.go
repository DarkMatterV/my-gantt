package models

func init() {
	if err := NewEngine("../data/task.db"); err != nil {
		panic(err)
	}
}

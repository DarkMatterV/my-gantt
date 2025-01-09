package models

import (
	"fmt"
	"testing"
)

func init() {
	NewEngine("../data/task.db")
}

func TestTask_List(t *testing.T) {
	task := Task{}
	l, err := task.List()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(l)
}

func TestTask_ListTree(t *testing.T) {
	task := Task{}
	l, err := task.ListTree()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(l)
}

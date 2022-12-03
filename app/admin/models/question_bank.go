package models

import "go-admin/common/models"

type QuestionBank struct {
	Id       int    `json:"Id" gorm:"primaryKey;autoIncrement;"` //题目Id
	Title    string `json:"Title" gorm:"size:255"`               //题目
	Option_a string `json:"Option_a" gorm:"size:255;"`           //选项A
	Option_b string `json:"Option_b"  gorm:"size:255;"`          //选项B
	Option_c string `json:"Option_c"  gorm:"size:255;"`          //选项C
	Option_d string `json:"Option_d"  gorm:"size:255;"`          //选项D
	Answer   int    `json:"Answer" gorm:"size:4;"`               //答案
	Detail   string `json:"Detail" gorm:"autoIncrement;"`        //答案详情
	Subject  int    `json:"Subject" gorm:"size:4;"`              //科目
	Type     int    `json:"Type" gorm:"size:4;"`                 //类别
	models.ModelTime
}

func (QuestionBank) TableName() string {
	return "question_bank"
}

func (e *QuestionBank) GetId() interface{} {
	return e.Id
}

package dto

import "go-admin/app/admin/models"

type QuestionBankGetQuestionReq struct {
	Id       int    `form:"id" search:"type:exact;column:id;table:question_bank" comment:"id"`              //题目Id
	Title    string `form:"title" search:"type:exact;column:title;table:question_bank" comment:"题目"`        //题目
	Option_a string `form:"option_a" search:"type:exact;column:option_a;table:question_bank" comment:"选项A"` //选项A
	Option_b string `form:"option_b" search:"type:exact;column:option_b;table:question_bank" comment:"选项B"` //选项B
	Option_c string `form:"option_c" search:"type:exact;column:option_c;table:question_bank" comment:"选项C"` //选项C
	Option_d string `form:"option_d" search:"type:exact;column:option_d;table:question_bank" comment:"选项D"` //选项D
	Answer   int    `form:"answer" search:"type:exact;column:answer;table:question_bank" comment:"答案"`      //答案
	Detail   string `form:"detail" search:"type:exact;column:detail;table:question_bank" comment:"答案详情"`    //答案详情
	Subject  int    `form:"subject" search:"type:exact;column:subject;table:question_bank" comment:"科目"`    //科目
	Type     int    `form:"type" search:"type:exact;column:type;table:question_bank" comment:"类别"`          //类别
}

type QuestionBankInsertReq struct {
	Id       int    `uri:"id" comment:"编码"`                       //题目Id
	Title    string `json:"title" comment:"题目" vd:"len($)>0"`     //题目
	Option_a string `json:"option_a" comment:"选项A" vd:"len($)>0"` //选项A
	Option_b string `json:"option_b" comment:"选项B" vd:"len($)>0"` //选项B
	Option_c string `json:"option_c" comment:"选项C" vd:"len($)>0"` //选项C
	Option_d string `json:"option_d" comment:"选项D" vd:"len($)>0"` //选项D
	Answer   int    `json:"answer" comment:"答案" vd:"$>0"`         //答案
	Detail   string `json:"detail" comment:"答案详情" vd:"len($)>0"`  //答案详情
	Subject  int    `json:"subject" comment:"科目" vd:"$>0"`        //科目
	Type     int    `json:"type" comment:"类别" vd:"$>0"`           //类别
}

func (s *QuestionBankInsertReq) Generate(model *models.QuestionBank) {
	if s.Id != 0 {
		model.Id = s.Id
	}
	model.Title = s.Title
	model.Option_a = s.Option_a
	model.Option_b = s.Option_b
	model.Option_c = s.Option_c
	model.Option_d = s.Option_d
	model.Answer = s.Answer
	model.Detail = s.Detail
	model.Subject = s.Subject
	model.Type = s.Type
}

type QuestionBankUpdateReq struct {
	Id       int    `uri:"id" comment:"编码"`                       //题目Id
	Title    string `json:"title" comment:"题目" vd:"len($)>0"`     //题目
	Option_a string `json:"option_a" comment:"选项A" vd:"len($)>0"` //选项A
	Option_b string `json:"option_b" comment:"选项B" vd:"len($)>0"` //选项B
	Option_c string `json:"option_c" comment:"选项C" vd:"len($)>0"` //选项C
	Option_d string `json:"option_d" comment:"选项D" vd:"len($)>0"` //选项D
	Answer   int    `json:"answer" comment:"答案" vd:"$>0"`         //答案
	Detail   string `json:"detail" comment:"答案详情" vd:"len($)>0"`  //答案详情
	Subject  int    `json:"subject" comment:"科目" vd:"$>0"`        //科目
	Type     int    `json:"type" comment:"类别" vd:"$>0"`
}

func (s *QuestionBankUpdateReq) Generate(model *models.QuestionBank) {
	if s.Id != 0 {
		model.Id = s.Id
	}
	model.Title = s.Title
	model.Option_a = s.Option_a
	model.Option_b = s.Option_b
	model.Option_c = s.Option_c
	model.Option_d = s.Option_d
	model.Answer = s.Answer
	model.Detail = s.Detail
	model.Subject = s.Subject
	model.Type = s.Type
}

func (m *QuestionBankGetQuestionReq) GetNeedSearch() interface{} {
	return *m
}

func (s *QuestionBankGetReq) GetId() interface{} {
	return s.Id
}

func (s *QuestionBankInsertReq) GetId() interface{} {
	return s.Id
}

func (s *QuestionBankUpdateReq) GetId() interface{} {
	return s.Id
}

type QuestionBankDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *QuestionBankDeleteReq) GetId() interface{} {
	return s.Ids
}

type QuestionBankGetReq struct {
	Id int `uri:"id"`
}

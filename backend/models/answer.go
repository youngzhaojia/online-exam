package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Answer struct {
	AnswerId    int       `gorm:"primaryKey;autoIncrement" json:"answerId"`
	ExamId      int       `json:"examId"`
	TeacherId   int       `json:"teacherId"`
	StudentName string    `json:"studentName"`
	AnswerList  string    `json:"answerList"`
	CreateTime  int64     `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
}

type AnswerInfo struct {
	AnswerId     int    `json:"answerId"`
	ExamId       int    `json:"examId"`
	ExamName     string `json:"examName"`
	QuestionList string `json:"quesionList"`
	TeacherId    int    `json:"teacherId"`
	StudentName  string `json:"studentName"`
	AnswerList   string `json:"answerList"`
	CreateTime   int64  `json:"createTime"`
}

func GetAnswerList(pageNum int, pageSize int, params interface{}) ([]AnswerInfo, error) {
	var (
		answerInfos []AnswerInfo
		err         error
	)

	if pageNum >= 0 && pageSize > 0 {
		err = db.Table("t_answer").Select("t_answer.*, t_exam.name as exam_name").Joins("left join t_exam on t_answer.exam_id = t_exam.exam_id").Where(params).Limit(pageSize).Offset(pageNum).Order("Create_time desc").Find(&answerInfos).Error
	}
	if err != nil {
		return nil, err
	}
	return answerInfos, err
}

func GetAnswerTotal(params interface{}) (count int) {
	db.Model(&Answer{}).Where(params).Count(&count)
	return
}

func GetAnswerDetail(answerId int) (*AnswerInfo, error) {
	var answerInfo AnswerInfo

	err := db.Table("t_answer").Select("t_answer.*, t_exam.name as exam_name, t_exam.question_list").Joins("left join t_exam on t_answer.exam_id = t_exam.exam_id").Where("answer_id = ?", answerId).First(&answerInfo).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &answerInfo, nil
}

func AddAnswer(params map[string]interface{}) error {
	answer := Answer{
		ExamId:      params["examId"].(int),
		TeacherId:   params["teacherId"].(int),
		StudentName: params["studentName"].(string),
		AnswerList:  params["answerList"].(string),
		CreateTime:  time.Now().Unix(),
		UpdateTime:  time.Now(),
	}

	if err := db.Create(&answer).Error; err != nil {
		return err
	}
	return nil
}

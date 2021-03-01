package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Answer struct {
	AnswerId   int       `gorm:"primaryKey;autoIncrement" json:"answerId"`
	UserId     int       `json:"userId"`
	ExamId     int       `json:"examId"`
	TeacherId  int       `json:"teacherId"`
	AnswerList string    `json:"answerList"`
	CreateTime int64     `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

type AnswerInfo struct {
	AnswerId   int    `json:"answerId"`
	UserId     int    `json:"userId"`
	UserName   string `json:"userName"`
	ExamId     int    `json:"examId"`
	ExamName   string `json:"examName"`
	TeacherId  int    `json:"teacherId"`
	AnswerList string `json:"answerList"`
	CreateTime int64  `json:"createTime"`
}

func GetAnswerList(pageNum int, pageSize int, params interface{}) ([]AnswerInfo, error) {
	var (
		answerInfos []AnswerInfo
		err         error
	)

	if pageNum >= 0 && pageSize > 0 {
		err = db.Table("t_answer").Select("t_answer.*, t_user.name as user_name, t_exam.name as exam_name").Joins("left join t_user on t_answer.user_id = t_user.user_id").Joins("left join t_exam on t_answer.exam_id = t_exam.exam_id").Where(params).Limit(pageSize).Offset(pageNum).Order("Create_time desc").Find(&answerInfos).Error
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

func GetAnswerDetail(answerId int) (*Answer, error) {
	var answer Answer
	err := db.Where("answer_id = ?", answerId).First(&answer).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &answer, nil
}

func AddAnswer(params map[string]interface{}) error {
	answer := Answer{
		UserId:     params["userId"].(int),
		AnswerList: params["answerList"].(string),
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now(),
	}

	if err := db.Create(&answer).Error; err != nil {
		return err
	}
	return nil
}

func EditAnswer(answerId int, data interface{}) error {
	if err := db.Model(&Answer{}).Where("answer_id = ?", answerId).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAnswer(answerId int) error {
	err := db.Where("answer_id = ?", answerId).Delete(&Answer{}).Error
	if err != nil {
		return err
	}
	return nil
}

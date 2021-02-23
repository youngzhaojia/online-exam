package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Answer struct {
	AnswerId   int       `gorm:"primaryKey;autoIncrement" json:"answerId"`
	UserId     int       `json:"userId"`
	AnswerList string    `json:"answerList"`
	CreateTime int64     `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

func GetAnswerList(pageNum int, pageSize int, params interface{}) ([]Answer, error) {
	var (
		answers []Answer
		err     error
	)

	if pageNum >= 0 && pageSize > 0 {
		err = db.Where(params).Limit(pageSize).Offset(pageNum).Find(&answers).Error
	} else {
		err = db.Where(params).Find(&answers).Error
	}
	if err != nil {
		return nil, err
	}
	return answers, err
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

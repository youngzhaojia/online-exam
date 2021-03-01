package models

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/unknwon/com"
)

type Question struct {
	QuestionId   int       `gorm:"primaryKey;autoIncrement" json:"questionId"`
	UserId       int       `json:"userId"`
	QuestionType int       `json:"questionType"` // 分类: 1选择题2问答题
	Title        string    `json:"title"`        // 题目
	Remark       string    `json:"remark"`       // 备注
	Option       string    `json:"option"`       // 选择题选项
	CreateTime   int64     `json:"createTime"`
	UpdateTime   time.Time `json:"updateTime"`
}

func GetQuestionList(pageNum int, pageSize int, params interface{}) ([]Question, error) {
	var (
		questions []Question
		err       error
	)

	if pageNum >= 0 && pageSize > 0 {
		err = db.Where(params).Limit(pageSize).Offset(pageNum).Order("Create_time desc").Find(&questions).Error
	} else {
		err = db.Where(params).Find(&questions).Error
	}
	if err != nil {
		return nil, err
	}
	return questions, err
}

// id 逗号隔开
func GetQuestionListByIds(ids string) ([]Question, error) {
	var (
		questions []Question
		err       error
	)
	idsStrArr := strings.Split(ids, ",")
	length := len(idsStrArr)
	idsIntArr := make([]int, length)
	for key, value := range idsStrArr {
		idsIntArr[key] = com.StrTo(value).MustInt()
	}

	err = db.Where("question_id IN (?)", idsIntArr).Order("Create_time desc").Find(&questions).Error
	if err != nil {
		return nil, err
	}
	return questions, err
}

func GetQuestionTotal(params interface{}) (count int) {
	db.Model(&Question{}).Where(params).Count(&count)
	return
}

func GetQuestionDetail(questionId int) (*Question, error) {
	var question Question
	err := db.Where("question_id = ?", questionId).First(&question).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &question, nil
}

func AddQuestion(params map[string]interface{}) error {
	question := Question{
		UserId:       params["userId"].(int),
		QuestionType: params["questionType"].(int),
		Title:        params["title"].(string),
		Remark:       params["remark"].(string),
		Option:       params["option"].(string),
		CreateTime:   time.Now().Unix(),
		UpdateTime:   time.Now(),
	}
	if err := db.Create(&question).Error; err != nil {
		return err
	}
	return nil
}

func EditQuestion(questionId int, data interface{}) error {
	if err := db.Model(&Question{}).Where("question_id = ?", questionId).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func DeleteQuestion(questionId int) error {
	err := db.Where("question_id = ?", questionId).Delete(&Question{}).Error
	if err != nil {
		return err
	}
	return nil
}

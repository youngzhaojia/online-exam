package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Exam struct {
	ExamId       int       `gorm:"primaryKey;autoIncrement" json:"examId"`
	UserId       int       `json:"userId"`
	Name         string    `json:"name"`
	Time         int       `json:"time"`
	QuestionList string    `json:"questionList"`
	CreateTime   int64     `json:"createTime"`
	UpdateTime   time.Time `json:"updateTime"`
}

func GetExamList(pageNum int, pageSize int, params interface{}) ([]Exam, error) {
	var (
		exams []Exam
		err   error
	)

	if pageNum >= 0 && pageSize > 0 {
		err = db.Where(params).Limit(pageSize).Offset(pageNum).Order("Create_time desc").Find(&exams).Error
	} else {
		err = db.Where(params).Find(&exams).Error
	}
	if err != nil {
		return nil, err
	}
	return exams, err
}

func GetExamTotal(params interface{}) (count int) {
	db.Model(&Exam{}).Where(params).Count(&count)
	return
}

func GetExamDetail(examId int) (*Exam, error) {
	var exam Exam
	err := db.Where("exam_id = ?", examId).First(&exam).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &exam, nil
}

func AddExam(params map[string]interface{}) error {
	exam := Exam{
		UserId:       params["userId"].(int),
		Name:         params["name"].(string),
		Time:         params["time"].(int),
		QuestionList: params["questionList"].(string),
		CreateTime:   time.Now().Unix(),
		UpdateTime:   time.Now(),
	}

	if err := db.Create(&exam).Error; err != nil {
		return err
	}
	return nil
}

func EditExam(examId int, data interface{}) error {
	if err := db.Model(&Exam{}).Where("exam_id = ?", examId).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func DeleteExam(examId int) error {
	err := db.Where("exam_id = ?", examId).Delete(&Exam{}).Error
	if err != nil {
		return err
	}
	return nil
}

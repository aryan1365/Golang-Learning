
package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"www.github.com/Aryan1365/day3/Config"
)

type Student struct {
	gorm.Model
	FirstName string    `gorm:"not null" json:"firstName"`
	LastName  string    `json:"lastName"`
	DOB       time.Time `json:"dob"`
	Address   string    `json:"address"`
	Marks     []Mark    `gorm:"constraint:OnDelete:CASCADE;foreignKey:StudentId"`
}


type Mark struct {
	gorm.Model
	Subject   string `json:"subject"`
	Marks     int    `json:"marks"`
	StudentId uint
}

func GetStudents(student *[]Student) error {
	if err := Config.DB.Preload("Marks").Find(student).Error; err != nil {
		return err
	}
	return nil
}

func CreateStudent(student *Student) error {
	if err := Config.DB.Create(student).Error; err != nil {
		return err
	}
	return nil
}

func GetUserById(student *Student, id string) error {
	if err := Config.DB.Preload("Marks").Where("id = ?", id).First(student).Error; err != nil {
		return err
	}
	return nil
}
func UpdateStudent(student *Student) error {
	Config.DB.Save(student)
	return nil
}


func DeleteStudent(student *Student, id string) error {
	Config.DB.Where("id = ?", id).Delete(student)
	return nil
}

func CreateMark(mark *Mark) error {
	if err := Config.DB.Create(mark).Error; err != nil {
		return err
	}
	return nil
}

func GetMarksInfo(marks *Mark, markId string) error {
	if err := Config.DB.Where("id = ?", markId).First(marks).Error; err != nil {
		return err
	}
	return nil
}

func UpdateMarks(mark *Mark) error {
	Config.DB.Save(mark)
	return nil
}


func DeleteMark(markId string) error {
	result := Config.DB.Unscoped().Where("id = ?", markId).Delete(&Mark{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("mark not found")
	}
	return nil
}
